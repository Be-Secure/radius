// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package kubernetes

// Commonly-used and Radius-Specific labels for Kubernetes
const (
	LabelRadiusApplication = "radius.dev/application"
	LabelRadiusComponent   = "radius.dev/component"
	LabelRadiusRevision    = "radius.dev/revision"

	LabelPartOf            = "app.kubernetes.io/part-of"
	LabelName              = "app.kubernetes.io/name"
	LabelManagedBy         = "app.kubernetes.io/managed-by"
	LabelManagedByRadiusRP = "radius-rp"

	FieldManager = "radius-rp"
)

// NOTE: the difference between descriptive labels and selector labels
//
// descriptive labels:
// - intended for humans and human troubleshooting
// - we include both our own metadata and kubernetes *recommended* labels
// - we might (in the future) include non-determinisitic details because these are informational
//
// selector labels:
// - intended for programmatic matching (selecting a pod for logging)
// - no value in redundancy between our own metadata and *recommended* labels, simpler is better
// - MUST remain deterministic
// - MUST be a subset of descriptive labels
//
// In general, descriptive labels should be applied all to Kubernetes objects, and selector labels should be used
// when programmatically querying those objects.

// MakeDescriptiveLabels returns a map of the descriptive labels for a Kubernetes resource associated with a component.
// The descriptive labels are a superset of the selector labels.
func MakeDescriptiveLabels(application string, component string) map[string]string {
	return map[string]string{
		LabelRadiusApplication: application,
		LabelRadiusComponent:   component,
		LabelName:              component,
		LabelPartOf:            application,
		LabelManagedBy:         LabelManagedByRadiusRP,
	}
}

// MakeSelectorLablels returns a map of labels suitable for a Kubernetes selector to identify a labeled Radius-managed
// Kubernetes object.
func MakeSelectorLabels(application string, component string) map[string]string {
	return map[string]string{
		LabelRadiusApplication: application,
		LabelRadiusComponent:   component,
	}
}