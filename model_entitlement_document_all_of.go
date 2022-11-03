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

// EntitlementDocumentAllOf struct for EntitlementDocumentAllOf
type EntitlementDocumentAllOf struct {
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	// The display name of the entitlement
	DisplayName *string `json:"displayName,omitempty"`
	Source *Reference1 `json:"source,omitempty"`
	Privileged *bool `json:"privileged,omitempty"`
	IdentityCount *int32 `json:"identityCount,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDocumentAllOf EntitlementDocumentAllOf

// NewEntitlementDocumentAllOf instantiates a new EntitlementDocumentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDocumentAllOf() *EntitlementDocumentAllOf {
	this := EntitlementDocumentAllOf{}
	return &this
}

// NewEntitlementDocumentAllOfWithDefaults instantiates a new EntitlementDocumentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDocumentAllOfWithDefaults() *EntitlementDocumentAllOf {
	this := EntitlementDocumentAllOf{}
	return &this
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocumentAllOf) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocumentAllOf) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *EntitlementDocumentAllOf) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *EntitlementDocumentAllOf) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *EntitlementDocumentAllOf) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementDocumentAllOf) GetSynced() time.Time {
	if o == nil || o.Synced.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementDocumentAllOf) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *EntitlementDocumentAllOf) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *EntitlementDocumentAllOf) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *EntitlementDocumentAllOf) UnsetSynced() {
	o.Synced.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOf) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EntitlementDocumentAllOf) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOf) GetSource() Reference1 {
	if o == nil || o.Source == nil {
		var ret Reference1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOf) GetSourceOk() (*Reference1, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference1 and assigns it to the Source field.
func (o *EntitlementDocumentAllOf) SetSource(v Reference1) {
	o.Source = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOf) GetPrivileged() bool {
	if o == nil || o.Privileged == nil {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOf) GetPrivilegedOk() (*bool, bool) {
	if o == nil || o.Privileged == nil {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasPrivileged() bool {
	if o != nil && o.Privileged != nil {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *EntitlementDocumentAllOf) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetIdentityCount returns the IdentityCount field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOf) GetIdentityCount() int32 {
	if o == nil || o.IdentityCount == nil {
		var ret int32
		return ret
	}
	return *o.IdentityCount
}

// GetIdentityCountOk returns a tuple with the IdentityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOf) GetIdentityCountOk() (*int32, bool) {
	if o == nil || o.IdentityCount == nil {
		return nil, false
	}
	return o.IdentityCount, true
}

// HasIdentityCount returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasIdentityCount() bool {
	if o != nil && o.IdentityCount != nil {
		return true
	}

	return false
}

// SetIdentityCount gets a reference to the given int32 and assigns it to the IdentityCount field.
func (o *EntitlementDocumentAllOf) SetIdentityCount(v int32) {
	o.IdentityCount = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EntitlementDocumentAllOf) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDocumentAllOf) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EntitlementDocumentAllOf) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EntitlementDocumentAllOf) SetTags(v []string) {
	o.Tags = v
}

func (o EntitlementDocumentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Privileged != nil {
		toSerialize["privileged"] = o.Privileged
	}
	if o.IdentityCount != nil {
		toSerialize["identityCount"] = o.IdentityCount
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementDocumentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementDocumentAllOf := _EntitlementDocumentAllOf{}

	if err = json.Unmarshal(bytes, &varEntitlementDocumentAllOf); err == nil {
		*o = EntitlementDocumentAllOf(varEntitlementDocumentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "source")
		delete(additionalProperties, "privileged")
		delete(additionalProperties, "identityCount")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDocumentAllOf struct {
	value *EntitlementDocumentAllOf
	isSet bool
}

func (v NullableEntitlementDocumentAllOf) Get() *EntitlementDocumentAllOf {
	return v.value
}

func (v *NullableEntitlementDocumentAllOf) Set(val *EntitlementDocumentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDocumentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDocumentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDocumentAllOf(val *EntitlementDocumentAllOf) *NullableEntitlementDocumentAllOf {
	return &NullableEntitlementDocumentAllOf{value: val, isSet: true}
}

func (v NullableEntitlementDocumentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDocumentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

