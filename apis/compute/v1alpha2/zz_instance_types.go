/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessConfigObservation struct {
}

type AccessConfigParameters struct {

	// The IP address that is be 1:1 mapped to the instance's network ip.
	// +kubebuilder:validation:Optional
	NATIP *string `json:"natIp,omitempty" tf:"nat_ip,omitempty"`

	// The networking tier used for configuring this instance. One of PREMIUM or STANDARD.
	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// The DNS domain name for the public PTR record.
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type AdvancedMachineFeaturesObservation struct {
}

type AdvancedMachineFeaturesParameters struct {

	// Whether to enable nested virtualization or not.
	// +kubebuilder:validation:Optional
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty" tf:"enable_nested_virtualization,omitempty"`

	// The number of threads per physical core. To disable simultaneous multithreading (SMT) set this to 1. If unset, the maximum number of threads supported per core by the underlying processor is assumed.
	// +kubebuilder:validation:Optional
	ThreadsPerCore *float64 `json:"threadsPerCore,omitempty" tf:"threads_per_core,omitempty"`
}

type AliasIPRangeObservation struct {
}

type AliasIPRangeParameters struct {

	// The IP CIDR range represented by this alias IP range.
	// +kubebuilder:validation:Required
	IPCidrRange *string `json:"ipCidrRange" tf:"ip_cidr_range,omitempty"`

	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.
	// +kubebuilder:validation:Optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name,omitempty"`
}

type AttachedDiskObservation struct {
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
}

type AttachedDiskParameters struct {

	// Name with which the attached disk is accessible under /dev/disk/by-id/
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRawSecretRef *v1.SecretKeySelector `json:"diskEncryptionKeyRawSecretRef,omitempty" tf:"-"`

	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name or self_link of the disk attached to this instance.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`
}

type BootDiskObservation struct {
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256,omitempty"`
}

type BootDiskParameters struct {

	// Whether the disk will be auto-deleted when the instance is deleted.
	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// Name with which attached disk will be accessible under /dev/disk/by-id/
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	DiskEncryptionKeyRawSecretRef *v1.SecretKeySelector `json:"diskEncryptionKeyRawSecretRef,omitempty" tf:"-"`

	// Parameters with which a disk was created alongside the instance.
	// +kubebuilder:validation:Optional
	InitializeParams []InitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +kubebuilder:validation:Optional
	KMSKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link,omitempty"`

	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name or self_link of the disk attached to this instance.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type ConfidentialInstanceConfigObservation struct {
}

type ConfidentialInstanceConfigParameters struct {

	// Defines whether the instance should have confidential compute enabled.
	// +kubebuilder:validation:Required
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute" tf:"enable_confidential_compute,omitempty"`
}

type GuestAcceleratorObservation struct {
}

type GuestAcceleratorParameters struct {

	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type IPv6AccessConfigObservation struct {
	ExternalIPv6 *string `json:"externalIpv6,omitempty" tf:"external_ipv6,omitempty"`

	ExternalIPv6PrefixLength *string `json:"externalIpv6PrefixLength,omitempty" tf:"external_ipv6_prefix_length,omitempty"`
}

type IPv6AccessConfigParameters struct {

	// The service-level to be provided for IPv6 traffic when the subnet has an external subnet. Only PREMIUM tier is valid for IPv6
	// +kubebuilder:validation:Required
	NetworkTier *string `json:"networkTier" tf:"network_tier,omitempty"`

	// The domain name to be used when creating DNSv6 records for the external IPv6 ranges.
	// +kubebuilder:validation:Optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name,omitempty"`
}

type InitializeParamsObservation struct {
}

type InitializeParamsParameters struct {

	// The image from which this disk was initialised.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// A set of key/value label pairs assigned to the disk.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The size of the image in gigabytes.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The Google Compute Engine disk type. One of pd-standard, pd-ssd or pd-balanced.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceObservation struct {
	CPUPlatform *string `json:"cpuPlatform,omitempty" tf:"cpu_platform,omitempty"`

	CurrentStatus *string `json:"currentStatus,omitempty" tf:"current_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint,omitempty"`

	MetadataFingerprint *string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	TagsFingerprint *string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint,omitempty"`
}

type InstanceParameters struct {

	// Controls for advanced machine-related behavior features.
	// +kubebuilder:validation:Optional
	AdvancedMachineFeatures []AdvancedMachineFeaturesParameters `json:"advancedMachineFeatures,omitempty" tf:"advanced_machine_features,omitempty"`

	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	// +kubebuilder:validation:Optional
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`

	// List of disks attached to the instance
	// +kubebuilder:validation:Optional
	AttachedDisk []AttachedDiskParameters `json:"attachedDisk,omitempty" tf:"attached_disk,omitempty"`

	// The boot disk for the instance.
	// +kubebuilder:validation:Required
	BootDisk []BootDiskParameters `json:"bootDisk" tf:"boot_disk,omitempty"`

	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	// +kubebuilder:validation:Optional
	CanIPForward *bool `json:"canIpForward,omitempty" tf:"can_ip_forward,omitempty"`

	// The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.
	// +kubebuilder:validation:Optional
	ConfidentialInstanceConfig []ConfidentialInstanceConfigParameters `json:"confidentialInstanceConfig,omitempty" tf:"confidential_instance_config,omitempty"`

	// Whether deletion protection is enabled on this instance.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A brief description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	// +kubebuilder:validation:Optional
	DesiredStatus *string `json:"desiredStatus,omitempty" tf:"desired_status,omitempty"`

	// Whether the instance has virtual displays enabled.
	// +kubebuilder:validation:Optional
	EnableDisplay *bool `json:"enableDisplay,omitempty" tf:"enable_display,omitempty"`

	// List of the type and count of accelerator cards attached to the instance.
	// +kubebuilder:validation:Optional
	GuestAccelerator []GuestAcceleratorParameters `json:"guestAccelerator,omitempty" tf:"guest_accelerator,omitempty"`

	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// A set of key/value label pairs assigned to the instance.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The machine type to create.
	// +kubebuilder:validation:Required
	MachineType *string `json:"machineType" tf:"machine_type,omitempty"`

	// Metadata key/value pairs made available within the instance.
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata startup scripts made available within the instance.
	// +kubebuilder:validation:Optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script,omitempty"`

	// The minimum CPU platform specified for the VM instance.
	// +kubebuilder:validation:Optional
	MinCPUPlatform *string `json:"minCpuPlatform,omitempty" tf:"min_cpu_platform,omitempty"`

	// The networks attached to the instance.
	// +kubebuilder:validation:Required
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the reservations that this instance can consume from.
	// +kubebuilder:validation:Optional
	ReservationAffinity []ReservationAffinityParameters `json:"reservationAffinity,omitempty" tf:"reservation_affinity,omitempty"`

	// A list of short names or self_links of resource policies to attach to the instance. Currently a max of 1 resource policy is supported.
	// +kubebuilder:validation:Optional
	ResourcePolicies []*string `json:"resourcePolicies,omitempty" tf:"resource_policies,omitempty"`

	// The scheduling strategy being used by the instance.
	// +kubebuilder:validation:Optional
	Scheduling []SchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// The scratch disks attached to the instance.
	// +kubebuilder:validation:Optional
	ScratchDisk []ScratchDiskParameters `json:"scratchDisk,omitempty" tf:"scratch_disk,omitempty"`

	// The service account to attach to the instance.
	// +kubebuilder:validation:Optional
	ServiceAccount []ServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// The shielded vm config being used by the instance.
	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []ShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// The list of tags attached to the instance.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NetworkInterfaceObservation struct {
	IPv6AccessType *string `json:"ipv6AccessType,omitempty" tf:"ipv6_access_type,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type NetworkInterfaceParameters struct {

	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	// +kubebuilder:validation:Optional
	AccessConfig []AccessConfigParameters `json:"accessConfig,omitempty" tf:"access_config,omitempty"`

	// An array of alias IP ranges for this network interface.
	// +kubebuilder:validation:Optional
	AliasIPRange []AliasIPRangeParameters `json:"aliasIpRange,omitempty" tf:"alias_ip_range,omitempty"`

	// An array of IPv6 access configurations for this interface. Currently, only one IPv6 access config, DIRECT_IPV6, is supported. If there is no ipv6AccessConfig specified, then this instance will have no external IPv6 Internet access.
	// +kubebuilder:validation:Optional
	IPv6AccessConfig []IPv6AccessConfigParameters `json:"ipv6AccessConfig,omitempty" tf:"ipv6_access_config,omitempty"`

	// The name or self_link of the network attached to this interface.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The private IP address assigned to the instance.
	// +kubebuilder:validation:Optional
	NetworkIP *string `json:"networkIp,omitempty" tf:"network_ip,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET
	// +kubebuilder:validation:Optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type,omitempty"`

	// The stack type for this network interface to identify whether the IPv6 feature is enabled or not. If not specified, IPV4_ONLY will be used.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty" tf:"stack_type,omitempty"`

	// The name or self_link of the subnetwork attached to this interface.
	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The project in which the subnetwork belongs.
	// +kubebuilder:validation:Optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty" tf:"subnetwork_project,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

type NodeAffinitiesObservation struct {
}

type NodeAffinitiesParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ReservationAffinityObservation struct {
}

type ReservationAffinityParameters struct {

	// Specifies the label selector for the reservation to use.
	// +kubebuilder:validation:Optional
	SpecificReservation []SpecificReservationParameters `json:"specificReservation,omitempty" tf:"specific_reservation,omitempty"`

	// The type of reservation from which this instance can consume resources.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SchedulingObservation struct {
}

type SchedulingParameters struct {

	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	// +kubebuilder:validation:Optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty" tf:"automatic_restart,omitempty"`

	// +kubebuilder:validation:Optional
	MinNodeCpus *float64 `json:"minNodeCpus,omitempty" tf:"min_node_cpus,omitempty"`

	// Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.
	// +kubebuilder:validation:Optional
	NodeAffinities []NodeAffinitiesParameters `json:"nodeAffinities,omitempty" tf:"node_affinities,omitempty"`

	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,
	// +kubebuilder:validation:Optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance,omitempty"`

	// Whether the instance is preemptible.
	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type ScratchDiskObservation struct {
}

type ScratchDiskParameters struct {

	// The disk interface used for attaching this disk. One of SCSI or NVME.
	// +kubebuilder:validation:Required
	Interface *string `json:"interface" tf:"interface,omitempty"`
}

type ServiceAccountObservation struct {
}

type ServiceAccountParameters struct {

	// The service account e-mail address.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// A list of service scopes.
	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`
}

type ShieldedInstanceConfigObservation struct {
}

type ShieldedInstanceConfigParameters struct {

	// Whether integrity monitoring is enabled for the instance.
	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// Whether secure boot is enabled for the instance.
	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`

	// Whether the instance uses vTPM.
	// +kubebuilder:validation:Optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm,omitempty"`
}

type SpecificReservationObservation struct {
}

type SpecificReservationParameters struct {

	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Corresponds to the label values of a reservation resource.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
