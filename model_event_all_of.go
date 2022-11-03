/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// EventAllOf struct for EventAllOf
type EventAllOf struct {
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	// The action that was performed
	Action *string `json:"action,omitempty"`
	// The type of event
	Type *string `json:"type,omitempty"`
	Actor *NameType `json:"actor,omitempty"`
	Target *NameType `json:"target,omitempty"`
	Stack *string `json:"stack,omitempty"`
	TrackingNumber *string `json:"trackingNumber,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Details *string `json:"details,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Objects []string `json:"objects,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Status *string `json:"status,omitempty"`
	TechnicalName *string `json:"technicalName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventAllOf EventAllOf

// NewEventAllOf instantiates a new EventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventAllOf() *EventAllOf {
	this := EventAllOf{}
	return &this
}

// NewEventAllOfWithDefaults instantiates a new EventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventAllOfWithDefaults() *EventAllOf {
	this := EventAllOf{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventAllOf) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *EventAllOf) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *EventAllOf) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *EventAllOf) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *EventAllOf) UnsetCreated() {
	o.Created.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventAllOf) GetSynced() time.Time {
	if o == nil || o.Synced.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventAllOf) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *EventAllOf) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *EventAllOf) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *EventAllOf) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *EventAllOf) UnsetSynced() {
	o.Synced.Unset()
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EventAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EventAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EventAllOf) SetAction(v string) {
	o.Action = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventAllOf) SetType(v string) {
	o.Type = &v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *EventAllOf) GetActor() NameType {
	if o == nil || o.Actor == nil {
		var ret NameType
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetActorOk() (*NameType, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *EventAllOf) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given NameType and assigns it to the Actor field.
func (o *EventAllOf) SetActor(v NameType) {
	o.Actor = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *EventAllOf) GetTarget() NameType {
	if o == nil || o.Target == nil {
		var ret NameType
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetTargetOk() (*NameType, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *EventAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NameType and assigns it to the Target field.
func (o *EventAllOf) SetTarget(v NameType) {
	o.Target = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *EventAllOf) GetStack() string {
	if o == nil || o.Stack == nil {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetStackOk() (*string, bool) {
	if o == nil || o.Stack == nil {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *EventAllOf) HasStack() bool {
	if o != nil && o.Stack != nil {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *EventAllOf) SetStack(v string) {
	o.Stack = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *EventAllOf) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetTrackingNumberOk() (*string, bool) {
	if o == nil || o.TrackingNumber == nil {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *EventAllOf) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber != nil {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *EventAllOf) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *EventAllOf) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *EventAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *EventAllOf) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *EventAllOf) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *EventAllOf) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *EventAllOf) SetDetails(v string) {
	o.Details = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EventAllOf) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EventAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *EventAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *EventAllOf) GetObjects() []string {
	if o == nil || o.Objects == nil {
		var ret []string
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetObjectsOk() ([]string, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *EventAllOf) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []string and assigns it to the Objects field.
func (o *EventAllOf) SetObjects(v []string) {
	o.Objects = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *EventAllOf) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *EventAllOf) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *EventAllOf) SetOperation(v string) {
	o.Operation = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTechnicalName returns the TechnicalName field value if set, zero value otherwise.
func (o *EventAllOf) GetTechnicalName() string {
	if o == nil || o.TechnicalName == nil {
		var ret string
		return ret
	}
	return *o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventAllOf) GetTechnicalNameOk() (*string, bool) {
	if o == nil || o.TechnicalName == nil {
		return nil, false
	}
	return o.TechnicalName, true
}

// HasTechnicalName returns a boolean if a field has been set.
func (o *EventAllOf) HasTechnicalName() bool {
	if o != nil && o.TechnicalName != nil {
		return true
	}

	return false
}

// SetTechnicalName gets a reference to the given string and assigns it to the TechnicalName field.
func (o *EventAllOf) SetTechnicalName(v string) {
	o.TechnicalName = &v
}

func (o EventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
	}
	if o.TrackingNumber != nil {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TechnicalName != nil {
		toSerialize["technicalName"] = o.TechnicalName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEventAllOf := _EventAllOf{}

	if err = json.Unmarshal(bytes, &varEventAllOf); err == nil {
		*o = EventAllOf(varEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "action")
		delete(additionalProperties, "type")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "target")
		delete(additionalProperties, "stack")
		delete(additionalProperties, "trackingNumber")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "details")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "objects")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "status")
		delete(additionalProperties, "technicalName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventAllOf struct {
	value *EventAllOf
	isSet bool
}

func (v NullableEventAllOf) Get() *EventAllOf {
	return v.value
}

func (v *NullableEventAllOf) Set(val *EventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAllOf(val *EventAllOf) *NullableEventAllOf {
	return &NullableEventAllOf{value: val, isSet: true}
}

func (v NullableEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


