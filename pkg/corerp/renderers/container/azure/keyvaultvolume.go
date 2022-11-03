// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package azure

import (
	"context"
	"errors"

	"github.com/project-radius/radius/pkg/corerp/datamodel"
	"github.com/project-radius/radius/pkg/corerp/handlers"
	"github.com/project-radius/radius/pkg/corerp/renderers"
	"github.com/project-radius/radius/pkg/kubernetes"
	"github.com/project-radius/radius/pkg/resourcekinds"
	"github.com/project-radius/radius/pkg/rp"
	"github.com/project-radius/radius/pkg/rp/outputresource"
	"github.com/project-radius/radius/pkg/ucp/resources"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	csiv1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1"
)

var (
	errInvalidKeyVaultResourceID = errors.New("failed to parse KeyVault ResourceID. Unable to create secret provider class")
	errUnsupportedIdentityKind   = errors.New("unsupported identity kind")
)

// MakeKeyVaultVolumeSpec builds CSI volume spec for Azure Keyvault.
func MakeKeyVaultVolumeSpec(volumeName string, mountPath, spcName string) (corev1.Volume, corev1.VolumeMount, error) {
	// Make Volume Spec which uses the SecretProvider created above
	volumeSpec := corev1.Volume{
		Name: volumeName,
		VolumeSource: corev1.VolumeSource{
			CSI: &corev1.CSIVolumeSource{
				Driver: "secrets-store.csi.k8s.io",
				// We will support only Read operations
				ReadOnly: to.Ptr(true),
				VolumeAttributes: map[string]string{
					"secretProviderClass": spcName,
				},
			},
		},
	}

	// Make Volume mount spec
	volumeMountSpec := corev1.VolumeMount{
		Name:      volumeName,
		MountPath: mountPath,
		// We will support only reads to the secret store volume
		ReadOnly: true,
	}

	return volumeSpec, volumeMountSpec, nil
}

// TransformSecretProviderClass mutates Kubernetes SecretProviderClass type resource.
func TransformSecretProviderClass(ctx context.Context, options *handlers.PutOptions) error {
	spc, ok := options.Resource.Resource.(*csiv1.SecretProviderClass)
	if !ok {
		return errors.New("cannot transform service account")
	}

	// Update the clientID and tenantID only for azure workload identity.
	if spc.Annotations != nil && spc.Annotations[kubernetes.AnnotationIdentityType] == string(rp.AzureIdentityWorkload) {
		clientID, tenantID, err := extractIdentityInfo(options)
		if err != nil {
			return err
		}

		spc.Spec.Parameters["clientID"] = clientID
		spc.Spec.Parameters["tenantID"] = tenantID
	}

	return nil
}

// MakeKeyVaultSecretProviderClass builds SecretProviderClass CR for keyvault secrets.
func MakeKeyVaultSecretProviderClass(appName, name string, res *datamodel.VolumeResource, objSpec string, envOpt *renderers.EnvironmentOptions) (*outputresource.OutputResource, error) {
	prop := res.Properties.AzureKeyVault

	kvResourceID, err := resources.ParseResource(prop.Resource)
	if err != nil {
		return nil, errInvalidKeyVaultResourceID
	}

	params := map[string]string{
		"usePodIdentity": "false",
		"keyvaultName":   kvResourceID.Name(),
		"objects":        objSpec,
	}

	switch envOpt.Identity.Kind {
	case rp.AzureIdentitySystemAssigned:
		// https://azure.github.io/secrets-store-csi-driver-provider-azure/docs/configurations/identity-access-modes/system-assigned-msi-mode/
		params["useVMManagedIdentity"] = "true"
		// clientID must be empty for system assigned managed identity
		params["clientID"] = ""
		// tenantID is a fake id to bypass crd validation because CSI doesn't require a tenant ID for System/User assigned managed identity.
		params["tenantID"] = "placeholder"

	case rp.AzureIdentityWorkload:
		break

	default:
		return nil, errUnsupportedIdentityKind
	}

	secretProvider := &csiv1.SecretProviderClass{
		TypeMeta: metav1.TypeMeta{
			Kind:       "SecretProviderClass",
			APIVersion: "secrets-store.csi.x-k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: envOpt.Namespace,
			Labels:    kubernetes.MakeDescriptiveLabels(appName, res.Name, res.Type),
			Annotations: map[string]string{
				kubernetes.AnnotationIdentityType: string(envOpt.Identity.Kind),
			},
		},
		Spec: csiv1.SecretProviderClassSpec{
			Provider:   "azure",
			Parameters: params,
		},
	}

	or := outputresource.NewKubernetesOutputResource(
		resourcekinds.SecretProviderClass,
		outputresource.LocalIDSecretProviderClass,
		secretProvider,
		secretProvider.ObjectMeta)

	or.Dependencies = []outputresource.Dependency{
		{
			LocalID: outputresource.LocalIDServiceAccount,
		},
	}

	return &or, nil

}