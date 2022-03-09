//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

const (
	module = "radclient"
	version = "v0.0.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{	
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// CertificateObjectPropertiesEncoding - Encoding format. Default utf-8
type CertificateObjectPropertiesEncoding string

const (
	CertificateObjectPropertiesEncodingBase64 CertificateObjectPropertiesEncoding = "base64"
	CertificateObjectPropertiesEncodingHex CertificateObjectPropertiesEncoding = "hex"
	CertificateObjectPropertiesEncodingUTF8 CertificateObjectPropertiesEncoding = "utf-8"
)

// PossibleCertificateObjectPropertiesEncodingValues returns the possible values for the CertificateObjectPropertiesEncoding const type.
func PossibleCertificateObjectPropertiesEncodingValues() []CertificateObjectPropertiesEncoding {
	return []CertificateObjectPropertiesEncoding{	
		CertificateObjectPropertiesEncodingBase64,
		CertificateObjectPropertiesEncodingHex,
		CertificateObjectPropertiesEncodingUTF8,
	}
}

// ToPtr returns a *CertificateObjectPropertiesEncoding pointing to the current value.
func (c CertificateObjectPropertiesEncoding) ToPtr() *CertificateObjectPropertiesEncoding {
	return &c
}

// CertificateObjectPropertiesFormat - Certificate format. Default pem
type CertificateObjectPropertiesFormat string

const (
	CertificateObjectPropertiesFormatPem CertificateObjectPropertiesFormat = "pem"
	CertificateObjectPropertiesFormatPfx CertificateObjectPropertiesFormat = "pfx"
)

// PossibleCertificateObjectPropertiesFormatValues returns the possible values for the CertificateObjectPropertiesFormat const type.
func PossibleCertificateObjectPropertiesFormatValues() []CertificateObjectPropertiesFormat {
	return []CertificateObjectPropertiesFormat{	
		CertificateObjectPropertiesFormatPem,
		CertificateObjectPropertiesFormatPfx,
	}
}

// ToPtr returns a *CertificateObjectPropertiesFormat pointing to the current value.
func (c CertificateObjectPropertiesFormat) ToPtr() *CertificateObjectPropertiesFormat {
	return &c
}

// CertificateObjectPropertiesValue - Certificate object to be downloaded - the certificate itself, private key or public key of the certificate
type CertificateObjectPropertiesValue string

const (
	CertificateObjectPropertiesValueCertificate CertificateObjectPropertiesValue = "certificate"
	CertificateObjectPropertiesValuePrivatekey CertificateObjectPropertiesValue = "privatekey"
	CertificateObjectPropertiesValuePublickey CertificateObjectPropertiesValue = "publickey"
)

// PossibleCertificateObjectPropertiesValueValues returns the possible values for the CertificateObjectPropertiesValue const type.
func PossibleCertificateObjectPropertiesValueValues() []CertificateObjectPropertiesValue {
	return []CertificateObjectPropertiesValue{	
		CertificateObjectPropertiesValueCertificate,
		CertificateObjectPropertiesValuePrivatekey,
		CertificateObjectPropertiesValuePublickey,
	}
}

// ToPtr returns a *CertificateObjectPropertiesValue pointing to the current value.
func (c CertificateObjectPropertiesValue) ToPtr() *CertificateObjectPropertiesValue {
	return &c
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{	
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// ToPtr returns a *CheckNameAvailabilityReason pointing to the current value.
func (c CheckNameAvailabilityReason) ToPtr() *CheckNameAvailabilityReason {
	return &c
}

// ConnectionKind - The kind of connection
type ConnectionKind string

const (
	ConnectionKindAzure ConnectionKind = "azure"
	ConnectionKindAzureComKeyVault ConnectionKind = "azure.com/KeyVault"
	ConnectionKindAzureComServiceBusQueue ConnectionKind = "azure.com/ServiceBusQueue"
	ConnectionKindDaprIoInvokeHTTP ConnectionKind = "dapr.io/InvokeHttp"
	ConnectionKindDaprIoPubSubTopic ConnectionKind = "dapr.io/PubSubTopic"
	ConnectionKindDaprIoSecretStore ConnectionKind = "dapr.io/SecretStore"
	ConnectionKindDaprIoStateStore ConnectionKind = "dapr.io/StateStore"
	ConnectionKindGrpc ConnectionKind = "Grpc"
	ConnectionKindHTTP ConnectionKind = "Http"
	ConnectionKindMicrosoftComSQL ConnectionKind = "microsoft.com/SQL"
	ConnectionKindMongoComMongoDB ConnectionKind = "mongo.com/MongoDB"
	ConnectionKindRabbitmqComMessageQueue ConnectionKind = "rabbitmq.com/MessageQueue"
	ConnectionKindRedislabsComRedis ConnectionKind = "redislabs.com/Redis"
)

// PossibleConnectionKindValues returns the possible values for the ConnectionKind const type.
func PossibleConnectionKindValues() []ConnectionKind {
	return []ConnectionKind{	
		ConnectionKindAzure,
		ConnectionKindAzureComKeyVault,
		ConnectionKindAzureComServiceBusQueue,
		ConnectionKindDaprIoInvokeHTTP,
		ConnectionKindDaprIoPubSubTopic,
		ConnectionKindDaprIoSecretStore,
		ConnectionKindDaprIoStateStore,
		ConnectionKindGrpc,
		ConnectionKindHTTP,
		ConnectionKindMicrosoftComSQL,
		ConnectionKindMongoComMongoDB,
		ConnectionKindRabbitmqComMessageQueue,
		ConnectionKindRedislabsComRedis,
	}
}

// ToPtr returns a *ConnectionKind pointing to the current value.
func (c ConnectionKind) ToPtr() *ConnectionKind {
	return &c
}

// ContainerPortProtocol - Protocol in use by the port
type ContainerPortProtocol string

const (
	ContainerPortProtocolTCP ContainerPortProtocol = "TCP"
	ContainerPortProtocolUDP ContainerPortProtocol = "UDP"
)

// PossibleContainerPortProtocolValues returns the possible values for the ContainerPortProtocol const type.
func PossibleContainerPortProtocolValues() []ContainerPortProtocol {
	return []ContainerPortProtocol{	
		ContainerPortProtocolTCP,
		ContainerPortProtocolUDP,
	}
}

// ToPtr returns a *ContainerPortProtocol pointing to the current value.
func (c ContainerPortProtocol) ToPtr() *ContainerPortProtocol {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication CreatedByType = "Application"
	CreatedByTypeKey CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{	
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DaprSecretStorePropertiesKind - Radius kind for Dapr Secret Store
type DaprSecretStorePropertiesKind string

const (
	DaprSecretStorePropertiesKindGeneric DaprSecretStorePropertiesKind = "generic"
)

// PossibleDaprSecretStorePropertiesKindValues returns the possible values for the DaprSecretStorePropertiesKind const type.
func PossibleDaprSecretStorePropertiesKindValues() []DaprSecretStorePropertiesKind {
	return []DaprSecretStorePropertiesKind{	
		DaprSecretStorePropertiesKindGeneric,
	}
}

// ToPtr returns a *DaprSecretStorePropertiesKind pointing to the current value.
func (c DaprSecretStorePropertiesKind) ToPtr() *DaprSecretStorePropertiesKind {
	return &c
}

// DaprSidecarTraitProtocol - Specifies the Dapr app-protocol to use for the resource.
type DaprSidecarTraitProtocol string

const (
	DaprSidecarTraitProtocolGrpc DaprSidecarTraitProtocol = "grpc"
	DaprSidecarTraitProtocolHTTP DaprSidecarTraitProtocol = "http"
)

// PossibleDaprSidecarTraitProtocolValues returns the possible values for the DaprSidecarTraitProtocol const type.
func PossibleDaprSidecarTraitProtocolValues() []DaprSidecarTraitProtocol {
	return []DaprSidecarTraitProtocol{	
		DaprSidecarTraitProtocolGrpc,
		DaprSidecarTraitProtocolHTTP,
	}
}

// ToPtr returns a *DaprSidecarTraitProtocol pointing to the current value.
func (c DaprSidecarTraitProtocol) ToPtr() *DaprSidecarTraitProtocol {
	return &c
}

// EncryptionStatus - Indicates whether or not the encryption is enabled for container registry.
type EncryptionStatus string

const (
	EncryptionStatusDisabled EncryptionStatus = "disabled"
	EncryptionStatusEnabled EncryptionStatus = "enabled"
)

// PossibleEncryptionStatusValues returns the possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{	
		EncryptionStatusDisabled,
		EncryptionStatusEnabled,
	}
}

// ToPtr returns a *EncryptionStatus pointing to the current value.
func (c EncryptionStatus) ToPtr() *EncryptionStatus {
	return &c
}

// EphemeralVolumeManagedStore - Backing store for the ephemeral volume
type EphemeralVolumeManagedStore string

const (
	EphemeralVolumeManagedStoreDisk EphemeralVolumeManagedStore = "disk"
	EphemeralVolumeManagedStoreMemory EphemeralVolumeManagedStore = "memory"
)

// PossibleEphemeralVolumeManagedStoreValues returns the possible values for the EphemeralVolumeManagedStore const type.
func PossibleEphemeralVolumeManagedStoreValues() []EphemeralVolumeManagedStore {
	return []EphemeralVolumeManagedStore{	
		EphemeralVolumeManagedStoreDisk,
		EphemeralVolumeManagedStoreMemory,
	}
}

// ToPtr returns a *EphemeralVolumeManagedStore pointing to the current value.
func (c EphemeralVolumeManagedStore) ToPtr() *EphemeralVolumeManagedStore {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default value is "user,system"
type Origin string

const (
	OriginSystem Origin = "system"
	OriginUser Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{	
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// PersistentVolumeRbac - Container read/write access to the volume
type PersistentVolumeRbac string

const (
	PersistentVolumeRbacRead PersistentVolumeRbac = "read"
	PersistentVolumeRbacWrite PersistentVolumeRbac = "write"
)

// PossiblePersistentVolumeRbacValues returns the possible values for the PersistentVolumeRbac const type.
func PossiblePersistentVolumeRbacValues() []PersistentVolumeRbac {
	return []PersistentVolumeRbac{	
		PersistentVolumeRbacRead,
		PersistentVolumeRbacWrite,
	}
}

// ToPtr returns a *PersistentVolumeRbac pointing to the current value.
func (c PersistentVolumeRbac) ToPtr() *PersistentVolumeRbac {
	return &c
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
type SKUTier string

const (
	SKUTierFree SKUTier = "Free"
	SKUTierBasic SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{	
		SKUTierFree,
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// SecretObjectPropertiesEncoding - Encoding format. Default utf-8
type SecretObjectPropertiesEncoding string

const (
	SecretObjectPropertiesEncodingBase64 SecretObjectPropertiesEncoding = "base64"
	SecretObjectPropertiesEncodingHex SecretObjectPropertiesEncoding = "hex"
	SecretObjectPropertiesEncodingUTF8 SecretObjectPropertiesEncoding = "utf-8"
)

// PossibleSecretObjectPropertiesEncodingValues returns the possible values for the SecretObjectPropertiesEncoding const type.
func PossibleSecretObjectPropertiesEncodingValues() []SecretObjectPropertiesEncoding {
	return []SecretObjectPropertiesEncoding{	
		SecretObjectPropertiesEncodingBase64,
		SecretObjectPropertiesEncodingHex,
		SecretObjectPropertiesEncodingUTF8,
	}
}

// ToPtr returns a *SecretObjectPropertiesEncoding pointing to the current value.
func (c SecretObjectPropertiesEncoding) ToPtr() *SecretObjectPropertiesEncoding {
	return &c
}

