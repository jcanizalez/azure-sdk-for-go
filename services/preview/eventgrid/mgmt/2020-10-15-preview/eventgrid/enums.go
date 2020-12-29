package eventgrid

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DomainProvisioningState enumerates the values for domain provisioning state.
type DomainProvisioningState string

const (
	// Canceled ...
	Canceled DomainProvisioningState = "Canceled"
	// Creating ...
	Creating DomainProvisioningState = "Creating"
	// Deleting ...
	Deleting DomainProvisioningState = "Deleting"
	// Failed ...
	Failed DomainProvisioningState = "Failed"
	// Succeeded ...
	Succeeded DomainProvisioningState = "Succeeded"
	// Updating ...
	Updating DomainProvisioningState = "Updating"
)

// PossibleDomainProvisioningStateValues returns an array of possible values for the DomainProvisioningState const type.
func PossibleDomainProvisioningStateValues() []DomainProvisioningState {
	return []DomainProvisioningState{Canceled, Creating, Deleting, Failed, Succeeded, Updating}
}

// DomainTopicProvisioningState enumerates the values for domain topic provisioning state.
type DomainTopicProvisioningState string

const (
	// DomainTopicProvisioningStateCanceled ...
	DomainTopicProvisioningStateCanceled DomainTopicProvisioningState = "Canceled"
	// DomainTopicProvisioningStateCreating ...
	DomainTopicProvisioningStateCreating DomainTopicProvisioningState = "Creating"
	// DomainTopicProvisioningStateDeleting ...
	DomainTopicProvisioningStateDeleting DomainTopicProvisioningState = "Deleting"
	// DomainTopicProvisioningStateFailed ...
	DomainTopicProvisioningStateFailed DomainTopicProvisioningState = "Failed"
	// DomainTopicProvisioningStateSucceeded ...
	DomainTopicProvisioningStateSucceeded DomainTopicProvisioningState = "Succeeded"
	// DomainTopicProvisioningStateUpdating ...
	DomainTopicProvisioningStateUpdating DomainTopicProvisioningState = "Updating"
)

// PossibleDomainTopicProvisioningStateValues returns an array of possible values for the DomainTopicProvisioningState const type.
func PossibleDomainTopicProvisioningStateValues() []DomainTopicProvisioningState {
	return []DomainTopicProvisioningState{DomainTopicProvisioningStateCanceled, DomainTopicProvisioningStateCreating, DomainTopicProvisioningStateDeleting, DomainTopicProvisioningStateFailed, DomainTopicProvisioningStateSucceeded, DomainTopicProvisioningStateUpdating}
}

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// EndpointTypeAzureFunction ...
	EndpointTypeAzureFunction EndpointType = "AzureFunction"
	// EndpointTypeEventHub ...
	EndpointTypeEventHub EndpointType = "EventHub"
	// EndpointTypeEventSubscriptionDestination ...
	EndpointTypeEventSubscriptionDestination EndpointType = "EventSubscriptionDestination"
	// EndpointTypeHybridConnection ...
	EndpointTypeHybridConnection EndpointType = "HybridConnection"
	// EndpointTypeServiceBusQueue ...
	EndpointTypeServiceBusQueue EndpointType = "ServiceBusQueue"
	// EndpointTypeServiceBusTopic ...
	EndpointTypeServiceBusTopic EndpointType = "ServiceBusTopic"
	// EndpointTypeStorageQueue ...
	EndpointTypeStorageQueue EndpointType = "StorageQueue"
	// EndpointTypeWebHook ...
	EndpointTypeWebHook EndpointType = "WebHook"
)

// PossibleEndpointTypeValues returns an array of possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{EndpointTypeAzureFunction, EndpointTypeEventHub, EndpointTypeEventSubscriptionDestination, EndpointTypeHybridConnection, EndpointTypeServiceBusQueue, EndpointTypeServiceBusTopic, EndpointTypeStorageQueue, EndpointTypeWebHook}
}

// EndpointTypeBasicDeadLetterDestination enumerates the values for endpoint type basic dead letter
// destination.
type EndpointTypeBasicDeadLetterDestination string

const (
	// EndpointTypeDeadLetterDestination ...
	EndpointTypeDeadLetterDestination EndpointTypeBasicDeadLetterDestination = "DeadLetterDestination"
	// EndpointTypeStorageBlob ...
	EndpointTypeStorageBlob EndpointTypeBasicDeadLetterDestination = "StorageBlob"
)

// PossibleEndpointTypeBasicDeadLetterDestinationValues returns an array of possible values for the EndpointTypeBasicDeadLetterDestination const type.
func PossibleEndpointTypeBasicDeadLetterDestinationValues() []EndpointTypeBasicDeadLetterDestination {
	return []EndpointTypeBasicDeadLetterDestination{EndpointTypeDeadLetterDestination, EndpointTypeStorageBlob}
}

// EventChannelProvisioningState enumerates the values for event channel provisioning state.
type EventChannelProvisioningState string

const (
	// EventChannelProvisioningStateCanceled ...
	EventChannelProvisioningStateCanceled EventChannelProvisioningState = "Canceled"
	// EventChannelProvisioningStateCreating ...
	EventChannelProvisioningStateCreating EventChannelProvisioningState = "Creating"
	// EventChannelProvisioningStateDeleting ...
	EventChannelProvisioningStateDeleting EventChannelProvisioningState = "Deleting"
	// EventChannelProvisioningStateFailed ...
	EventChannelProvisioningStateFailed EventChannelProvisioningState = "Failed"
	// EventChannelProvisioningStateSucceeded ...
	EventChannelProvisioningStateSucceeded EventChannelProvisioningState = "Succeeded"
	// EventChannelProvisioningStateUpdating ...
	EventChannelProvisioningStateUpdating EventChannelProvisioningState = "Updating"
)

// PossibleEventChannelProvisioningStateValues returns an array of possible values for the EventChannelProvisioningState const type.
func PossibleEventChannelProvisioningStateValues() []EventChannelProvisioningState {
	return []EventChannelProvisioningState{EventChannelProvisioningStateCanceled, EventChannelProvisioningStateCreating, EventChannelProvisioningStateDeleting, EventChannelProvisioningStateFailed, EventChannelProvisioningStateSucceeded, EventChannelProvisioningStateUpdating}
}


// ManagedIdentityType enumerates the values for managed identity type.
type ManagedIdentityType string

const (
	// ManagedIdentityTypeNone No managed identity.
	ManagedIdentityTypeNone ManagedIdentityType = "None"
	// ManagedIdentityTypeSystemAssigned A system-assigned managed identity.
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
)

// PossibleManagedIdentityTypeValues returns an array of possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{ManagedIdentityTypeNone, ManagedIdentityTypeSystemAssigned}
}

// EventDeliverySchema enumerates the values for event delivery schema.
type EventDeliverySchema string

const (
	// CloudEventSchemaV10 ...
	CloudEventSchemaV10 EventDeliverySchema = "CloudEventSchemaV1_0"
	// CustomInputSchema ...
	CustomInputSchema EventDeliverySchema = "CustomInputSchema"
	// EventGridSchema ...
	EventGridSchema EventDeliverySchema = "EventGridSchema"
)

// PossibleEventDeliverySchemaValues returns an array of possible values for the EventDeliverySchema const type.
func PossibleEventDeliverySchemaValues() []EventDeliverySchema {
	return []EventDeliverySchema{CloudEventSchemaV10, CustomInputSchema, EventGridSchema}
}

// EventSubscriptionIdentityType enumerates the values for event subscription identity type.
type EventSubscriptionIdentityType string

const (
	// SystemAssigned ...
	SystemAssigned EventSubscriptionIdentityType = "SystemAssigned"
	// UserAssigned ...
	UserAssigned EventSubscriptionIdentityType = "UserAssigned"
)

// PossibleEventSubscriptionIdentityTypeValues returns an array of possible values for the EventSubscriptionIdentityType const type.
func PossibleEventSubscriptionIdentityTypeValues() []EventSubscriptionIdentityType {
	return []EventSubscriptionIdentityType{SystemAssigned, UserAssigned}
}

// EventSubscriptionProvisioningState enumerates the values for event subscription provisioning state.
type EventSubscriptionProvisioningState string

const (
	// EventSubscriptionProvisioningStateAwaitingManualAction ...
	EventSubscriptionProvisioningStateAwaitingManualAction EventSubscriptionProvisioningState = "AwaitingManualAction"
	// EventSubscriptionProvisioningStateCanceled ...
	EventSubscriptionProvisioningStateCanceled EventSubscriptionProvisioningState = "Canceled"
	// EventSubscriptionProvisioningStateCreating ...
	EventSubscriptionProvisioningStateCreating EventSubscriptionProvisioningState = "Creating"
	// EventSubscriptionProvisioningStateDeleting ...
	EventSubscriptionProvisioningStateDeleting EventSubscriptionProvisioningState = "Deleting"
	// EventSubscriptionProvisioningStateFailed ...
	EventSubscriptionProvisioningStateFailed EventSubscriptionProvisioningState = "Failed"
	// EventSubscriptionProvisioningStateSucceeded ...
	EventSubscriptionProvisioningStateSucceeded EventSubscriptionProvisioningState = "Succeeded"
	// EventSubscriptionProvisioningStateUpdating ...
	EventSubscriptionProvisioningStateUpdating EventSubscriptionProvisioningState = "Updating"
)

// PossibleEventSubscriptionProvisioningStateValues returns an array of possible values for the EventSubscriptionProvisioningState const type.
func PossibleEventSubscriptionProvisioningStateValues() []EventSubscriptionProvisioningState {
	return []EventSubscriptionProvisioningState{EventSubscriptionProvisioningStateAwaitingManualAction, EventSubscriptionProvisioningStateCanceled, EventSubscriptionProvisioningStateCreating, EventSubscriptionProvisioningStateDeleting, EventSubscriptionProvisioningStateFailed, EventSubscriptionProvisioningStateSucceeded, EventSubscriptionProvisioningStateUpdating}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// IdentityTypeNone ...
	IdentityTypeNone IdentityType = "None"
	// IdentityTypeSystemAssigned ...
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	// IdentityTypeSystemAssignedUserAssigned ...
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned, UserAssigned"
	// IdentityTypeUserAssigned ...
	IdentityTypeUserAssigned IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{IdentityTypeNone, IdentityTypeSystemAssigned, IdentityTypeSystemAssignedUserAssigned, IdentityTypeUserAssigned}
}

// InputSchema enumerates the values for input schema.
type InputSchema string

const (
	// InputSchemaCloudEventSchemaV10 ...
	InputSchemaCloudEventSchemaV10 InputSchema = "CloudEventSchemaV1_0"
	// InputSchemaCustomEventSchema ...
	InputSchemaCustomEventSchema InputSchema = "CustomEventSchema"
	// InputSchemaEventGridSchema ...
	InputSchemaEventGridSchema InputSchema = "EventGridSchema"
)

// PossibleInputSchemaValues returns an array of possible values for the InputSchema const type.
func PossibleInputSchemaValues() []InputSchema {
	return []InputSchema{InputSchemaCloudEventSchemaV10, InputSchemaCustomEventSchema, InputSchemaEventGridSchema}
}

// InputSchemaMappingType enumerates the values for input schema mapping type.
type InputSchemaMappingType string

const (
	// InputSchemaMappingTypeInputSchemaMapping ...
	InputSchemaMappingTypeInputSchemaMapping InputSchemaMappingType = "InputSchemaMapping"
	// InputSchemaMappingTypeJSON ...
	InputSchemaMappingTypeJSON InputSchemaMappingType = "Json"
)

// PossibleInputSchemaMappingTypeValues returns an array of possible values for the InputSchemaMappingType const type.
func PossibleInputSchemaMappingTypeValues() []InputSchemaMappingType {
	return []InputSchemaMappingType{InputSchemaMappingTypeInputSchemaMapping, InputSchemaMappingTypeJSON}
}

// IPActionType enumerates the values for ip action type.
type IPActionType string

const (
	// Allow ...
	Allow IPActionType = "Allow"
)

// PossibleIPActionTypeValues returns an array of possible values for the IPActionType const type.
func PossibleIPActionTypeValues() []IPActionType {
	return []IPActionType{Allow}
}

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// OperatorTypeAdvancedFilter ...
	OperatorTypeAdvancedFilter OperatorType = "AdvancedFilter"
	// OperatorTypeBoolEquals ...
	OperatorTypeBoolEquals OperatorType = "BoolEquals"
	// OperatorTypeNumberGreaterThan ...
	OperatorTypeNumberGreaterThan OperatorType = "NumberGreaterThan"
	// OperatorTypeNumberGreaterThanOrEquals ...
	OperatorTypeNumberGreaterThanOrEquals OperatorType = "NumberGreaterThanOrEquals"
	// OperatorTypeNumberIn ...
	OperatorTypeNumberIn OperatorType = "NumberIn"
	// OperatorTypeNumberLessThan ...
	OperatorTypeNumberLessThan OperatorType = "NumberLessThan"
	// OperatorTypeNumberLessThanOrEquals ...
	OperatorTypeNumberLessThanOrEquals OperatorType = "NumberLessThanOrEquals"
	// OperatorTypeNumberNotIn ...
	OperatorTypeNumberNotIn OperatorType = "NumberNotIn"
	// OperatorTypeStringBeginsWith ...
	OperatorTypeStringBeginsWith OperatorType = "StringBeginsWith"
	// OperatorTypeStringContains ...
	OperatorTypeStringContains OperatorType = "StringContains"
	// OperatorTypeStringEndsWith ...
	OperatorTypeStringEndsWith OperatorType = "StringEndsWith"
	// OperatorTypeStringIn ...
	OperatorTypeStringIn OperatorType = "StringIn"
	// OperatorTypeStringNotIn ...
	OperatorTypeStringNotIn OperatorType = "StringNotIn"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{OperatorTypeAdvancedFilter, OperatorTypeBoolEquals, OperatorTypeNumberGreaterThan, OperatorTypeNumberGreaterThanOrEquals, OperatorTypeNumberIn, OperatorTypeNumberLessThan, OperatorTypeNumberLessThanOrEquals, OperatorTypeNumberNotIn, OperatorTypeStringBeginsWith, OperatorTypeStringContains, OperatorTypeStringEndsWith, OperatorTypeStringIn, OperatorTypeStringNotIn}
}

// PartnerNamespaceProvisioningState enumerates the values for partner namespace provisioning state.
type PartnerNamespaceProvisioningState string

const (
	// PartnerNamespaceProvisioningStateCanceled ...
	PartnerNamespaceProvisioningStateCanceled PartnerNamespaceProvisioningState = "Canceled"
	// PartnerNamespaceProvisioningStateCreating ...
	PartnerNamespaceProvisioningStateCreating PartnerNamespaceProvisioningState = "Creating"
	// PartnerNamespaceProvisioningStateDeleting ...
	PartnerNamespaceProvisioningStateDeleting PartnerNamespaceProvisioningState = "Deleting"
	// PartnerNamespaceProvisioningStateFailed ...
	PartnerNamespaceProvisioningStateFailed PartnerNamespaceProvisioningState = "Failed"
	// PartnerNamespaceProvisioningStateSucceeded ...
	PartnerNamespaceProvisioningStateSucceeded PartnerNamespaceProvisioningState = "Succeeded"
	// PartnerNamespaceProvisioningStateUpdating ...
	PartnerNamespaceProvisioningStateUpdating PartnerNamespaceProvisioningState = "Updating"
)

// PossiblePartnerNamespaceProvisioningStateValues returns an array of possible values for the PartnerNamespaceProvisioningState const type.
func PossiblePartnerNamespaceProvisioningStateValues() []PartnerNamespaceProvisioningState {
	return []PartnerNamespaceProvisioningState{PartnerNamespaceProvisioningStateCanceled, PartnerNamespaceProvisioningStateCreating, PartnerNamespaceProvisioningStateDeleting, PartnerNamespaceProvisioningStateFailed, PartnerNamespaceProvisioningStateSucceeded, PartnerNamespaceProvisioningStateUpdating}
}

// PartnerRegistrationProvisioningState enumerates the values for partner registration provisioning state.
type PartnerRegistrationProvisioningState string

const (
	// PartnerRegistrationProvisioningStateCanceled ...
	PartnerRegistrationProvisioningStateCanceled PartnerRegistrationProvisioningState = "Canceled"
	// PartnerRegistrationProvisioningStateCreating ...
	PartnerRegistrationProvisioningStateCreating PartnerRegistrationProvisioningState = "Creating"
	// PartnerRegistrationProvisioningStateDeleting ...
	PartnerRegistrationProvisioningStateDeleting PartnerRegistrationProvisioningState = "Deleting"
	// PartnerRegistrationProvisioningStateFailed ...
	PartnerRegistrationProvisioningStateFailed PartnerRegistrationProvisioningState = "Failed"
	// PartnerRegistrationProvisioningStateSucceeded ...
	PartnerRegistrationProvisioningStateSucceeded PartnerRegistrationProvisioningState = "Succeeded"
	// PartnerRegistrationProvisioningStateUpdating ...
	PartnerRegistrationProvisioningStateUpdating PartnerRegistrationProvisioningState = "Updating"
)

// PossiblePartnerRegistrationProvisioningStateValues returns an array of possible values for the PartnerRegistrationProvisioningState const type.
func PossiblePartnerRegistrationProvisioningStateValues() []PartnerRegistrationProvisioningState {
	return []PartnerRegistrationProvisioningState{PartnerRegistrationProvisioningStateCanceled, PartnerRegistrationProvisioningStateCreating, PartnerRegistrationProvisioningStateDeleting, PartnerRegistrationProvisioningStateFailed, PartnerRegistrationProvisioningStateSucceeded, PartnerRegistrationProvisioningStateUpdating}
}

// PartnerRegistrationVisibilityState enumerates the values for partner registration visibility state.
type PartnerRegistrationVisibilityState string

const (
	// GenerallyAvailable ...
	GenerallyAvailable PartnerRegistrationVisibilityState = "GenerallyAvailable"
	// Hidden ...
	Hidden PartnerRegistrationVisibilityState = "Hidden"
	// PublicPreview ...
	PublicPreview PartnerRegistrationVisibilityState = "PublicPreview"
)

// PossiblePartnerRegistrationVisibilityStateValues returns an array of possible values for the PartnerRegistrationVisibilityState const type.
func PossiblePartnerRegistrationVisibilityStateValues() []PartnerRegistrationVisibilityState {
	return []PartnerRegistrationVisibilityState{GenerallyAvailable, Hidden, PublicPreview}
}

// PartnerTopicActivationState enumerates the values for partner topic activation state.
type PartnerTopicActivationState string

const (
	// Activated ...
	Activated PartnerTopicActivationState = "Activated"
	// Deactivated ...
	Deactivated PartnerTopicActivationState = "Deactivated"
	// NeverActivated ...
	NeverActivated PartnerTopicActivationState = "NeverActivated"
)

// PossiblePartnerTopicActivationStateValues returns an array of possible values for the PartnerTopicActivationState const type.
func PossiblePartnerTopicActivationStateValues() []PartnerTopicActivationState {
	return []PartnerTopicActivationState{Activated, Deactivated, NeverActivated}
}

// PartnerTopicProvisioningState enumerates the values for partner topic provisioning state.
type PartnerTopicProvisioningState string

const (
	// PartnerTopicProvisioningStateCanceled ...
	PartnerTopicProvisioningStateCanceled PartnerTopicProvisioningState = "Canceled"
	// PartnerTopicProvisioningStateCreating ...
	PartnerTopicProvisioningStateCreating PartnerTopicProvisioningState = "Creating"
	// PartnerTopicProvisioningStateDeleting ...
	PartnerTopicProvisioningStateDeleting PartnerTopicProvisioningState = "Deleting"
	// PartnerTopicProvisioningStateFailed ...
	PartnerTopicProvisioningStateFailed PartnerTopicProvisioningState = "Failed"
	// PartnerTopicProvisioningStateSucceeded ...
	PartnerTopicProvisioningStateSucceeded PartnerTopicProvisioningState = "Succeeded"
	// PartnerTopicProvisioningStateUpdating ...
	PartnerTopicProvisioningStateUpdating PartnerTopicProvisioningState = "Updating"
)

// PossiblePartnerTopicProvisioningStateValues returns an array of possible values for the PartnerTopicProvisioningState const type.
func PossiblePartnerTopicProvisioningStateValues() []PartnerTopicProvisioningState {
	return []PartnerTopicProvisioningState{PartnerTopicProvisioningStateCanceled, PartnerTopicProvisioningStateCreating, PartnerTopicProvisioningStateDeleting, PartnerTopicProvisioningStateFailed, PartnerTopicProvisioningStateSucceeded, PartnerTopicProvisioningStateUpdating}
}

// PartnerTopicReadinessState enumerates the values for partner topic readiness state.
type PartnerTopicReadinessState string

const (
	// ActivatedByUser ...
	ActivatedByUser PartnerTopicReadinessState = "ActivatedByUser"
	// DeactivatedByUser ...
	DeactivatedByUser PartnerTopicReadinessState = "DeactivatedByUser"
	// DeletedByUser ...
	DeletedByUser PartnerTopicReadinessState = "DeletedByUser"
	// NotActivatedByUserYet ...
	NotActivatedByUserYet PartnerTopicReadinessState = "NotActivatedByUserYet"
)

// PossiblePartnerTopicReadinessStateValues returns an array of possible values for the PartnerTopicReadinessState const type.
func PossiblePartnerTopicReadinessStateValues() []PartnerTopicReadinessState {
	return []PartnerTopicReadinessState{ActivatedByUser, DeactivatedByUser, DeletedByUser, NotActivatedByUserYet}
}

// PartnerTopicTypeAuthorizationState enumerates the values for partner topic type authorization state.
type PartnerTopicTypeAuthorizationState string

const (
	// Authorized ...
	Authorized PartnerTopicTypeAuthorizationState = "Authorized"
	// NotApplicable ...
	NotApplicable PartnerTopicTypeAuthorizationState = "NotApplicable"
	// NotAuthorized ...
	NotAuthorized PartnerTopicTypeAuthorizationState = "NotAuthorized"
)

// PossiblePartnerTopicTypeAuthorizationStateValues returns an array of possible values for the PartnerTopicTypeAuthorizationState const type.
func PossiblePartnerTopicTypeAuthorizationStateValues() []PartnerTopicTypeAuthorizationState {
	return []PartnerTopicTypeAuthorizationState{Authorized, NotApplicable, NotAuthorized}
}

// PersistedConnectionStatus enumerates the values for persisted connection status.
type PersistedConnectionStatus string

const (
	// Approved ...
	Approved PersistedConnectionStatus = "Approved"
	// Disconnected ...
	Disconnected PersistedConnectionStatus = "Disconnected"
	// Pending ...
	Pending PersistedConnectionStatus = "Pending"
	// Rejected ...
	Rejected PersistedConnectionStatus = "Rejected"
)

// PossiblePersistedConnectionStatusValues returns an array of possible values for the PersistedConnectionStatus const type.
func PossiblePersistedConnectionStatusValues() []PersistedConnectionStatus {
	return []PersistedConnectionStatus{Approved, Disconnected, Pending, Rejected}
}

// PublicNetworkAccess enumerates the values for public network access.
type PublicNetworkAccess string

const (
	// Disabled ...
	Disabled PublicNetworkAccess = "Disabled"
	// Enabled ...
	Enabled PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns an array of possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{Disabled, Enabled}
}

// ResourceProvisioningState enumerates the values for resource provisioning state.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateCanceled ...
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateCreating ...
	ResourceProvisioningStateCreating ResourceProvisioningState = "Creating"
	// ResourceProvisioningStateDeleting ...
	ResourceProvisioningStateDeleting ResourceProvisioningState = "Deleting"
	// ResourceProvisioningStateFailed ...
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateSucceeded ...
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
	// ResourceProvisioningStateUpdating ...
	ResourceProvisioningStateUpdating ResourceProvisioningState = "Updating"
)

// PossibleResourceProvisioningStateValues returns an array of possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{ResourceProvisioningStateCanceled, ResourceProvisioningStateCreating, ResourceProvisioningStateDeleting, ResourceProvisioningStateFailed, ResourceProvisioningStateSucceeded, ResourceProvisioningStateUpdating}
}

// ResourceRegionType enumerates the values for resource region type.
type ResourceRegionType string

const (
	// GlobalResource ...
	GlobalResource ResourceRegionType = "GlobalResource"
	// RegionalResource ...
	RegionalResource ResourceRegionType = "RegionalResource"
)

// PossibleResourceRegionTypeValues returns an array of possible values for the ResourceRegionType const type.
func PossibleResourceRegionTypeValues() []ResourceRegionType {
	return []ResourceRegionType{GlobalResource, RegionalResource}
}

// Sku enumerates the values for sku.
type Sku string

const (
	// Basic ...
	Basic Sku = "Basic"
	// Premium ...
	Premium Sku = "Premium"
)

// PossibleSkuValues returns an array of possible values for the Sku const type.
func PossibleSkuValues() []Sku {
	return []Sku{Basic, Premium}
}

// TopicProvisioningState enumerates the values for topic provisioning state.
type TopicProvisioningState string

const (
	// TopicProvisioningStateCanceled ...
	TopicProvisioningStateCanceled TopicProvisioningState = "Canceled"
	// TopicProvisioningStateCreating ...
	TopicProvisioningStateCreating TopicProvisioningState = "Creating"
	// TopicProvisioningStateDeleting ...
	TopicProvisioningStateDeleting TopicProvisioningState = "Deleting"
	// TopicProvisioningStateFailed ...
	TopicProvisioningStateFailed TopicProvisioningState = "Failed"
	// TopicProvisioningStateSucceeded ...
	TopicProvisioningStateSucceeded TopicProvisioningState = "Succeeded"
	// TopicProvisioningStateUpdating ...
	TopicProvisioningStateUpdating TopicProvisioningState = "Updating"
)

// PossibleTopicProvisioningStateValues returns an array of possible values for the TopicProvisioningState const type.
func PossibleTopicProvisioningStateValues() []TopicProvisioningState {
	return []TopicProvisioningState{TopicProvisioningStateCanceled, TopicProvisioningStateCreating, TopicProvisioningStateDeleting, TopicProvisioningStateFailed, TopicProvisioningStateSucceeded, TopicProvisioningStateUpdating}
}

// TopicTypeProvisioningState enumerates the values for topic type provisioning state.
type TopicTypeProvisioningState string

const (
	// TopicTypeProvisioningStateCanceled ...
	TopicTypeProvisioningStateCanceled TopicTypeProvisioningState = "Canceled"
	// TopicTypeProvisioningStateCreating ...
	TopicTypeProvisioningStateCreating TopicTypeProvisioningState = "Creating"
	// TopicTypeProvisioningStateDeleting ...
	TopicTypeProvisioningStateDeleting TopicTypeProvisioningState = "Deleting"
	// TopicTypeProvisioningStateFailed ...
	TopicTypeProvisioningStateFailed TopicTypeProvisioningState = "Failed"
	// TopicTypeProvisioningStateSucceeded ...
	TopicTypeProvisioningStateSucceeded TopicTypeProvisioningState = "Succeeded"
	// TopicTypeProvisioningStateUpdating ...
	TopicTypeProvisioningStateUpdating TopicTypeProvisioningState = "Updating"
)

// PossibleTopicTypeProvisioningStateValues returns an array of possible values for the TopicTypeProvisioningState const type.
func PossibleTopicTypeProvisioningStateValues() []TopicTypeProvisioningState {
	return []TopicTypeProvisioningState{TopicTypeProvisioningStateCanceled, TopicTypeProvisioningStateCreating, TopicTypeProvisioningStateDeleting, TopicTypeProvisioningStateFailed, TopicTypeProvisioningStateSucceeded, TopicTypeProvisioningStateUpdating}
}
