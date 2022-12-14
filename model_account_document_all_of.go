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

// AccountDocumentAllOf struct for AccountDocumentAllOf
type AccountDocumentAllOf struct {
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// a map or dictionary of key/value pairs
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Identity *DisplayReference `json:"identity,omitempty"`
	Access []Entitlement1 `json:"access,omitempty"`
	// The number of entitlements assigned to the account
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// Indicates if the account is not correlated to an identity
	Uncorrelated *bool `json:"uncorrelated,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountDocumentAllOf AccountDocumentAllOf

// NewAccountDocumentAllOf instantiates a new AccountDocumentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDocumentAllOf() *AccountDocumentAllOf {
	this := AccountDocumentAllOf{}
	return &this
}

// NewAccountDocumentAllOfWithDefaults instantiates a new AccountDocumentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDocumentAllOfWithDefaults() *AccountDocumentAllOf {
	this := AccountDocumentAllOf{}
	return &this
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountDocumentAllOf) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountDocumentAllOf) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *AccountDocumentAllOf) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *AccountDocumentAllOf) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *AccountDocumentAllOf) UnsetModified() {
	o.Modified.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AccountDocumentAllOf) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *AccountDocumentAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AccountDocumentAllOf) GetIdentity() DisplayReference {
	if o == nil || o.Identity == nil {
		var ret DisplayReference
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentAllOf) GetIdentityOk() (*DisplayReference, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given DisplayReference and assigns it to the Identity field.
func (o *AccountDocumentAllOf) SetIdentity(v DisplayReference) {
	o.Identity = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *AccountDocumentAllOf) GetAccess() []Entitlement1 {
	if o == nil || o.Access == nil {
		var ret []Entitlement1
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentAllOf) GetAccessOk() ([]Entitlement1, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []Entitlement1 and assigns it to the Access field.
func (o *AccountDocumentAllOf) SetAccess(v []Entitlement1) {
	o.Access = v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *AccountDocumentAllOf) GetEntitlementCount() int32 {
	if o == nil || o.EntitlementCount == nil {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentAllOf) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || o.EntitlementCount == nil {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasEntitlementCount() bool {
	if o != nil && o.EntitlementCount != nil {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *AccountDocumentAllOf) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetUncorrelated returns the Uncorrelated field value if set, zero value otherwise.
func (o *AccountDocumentAllOf) GetUncorrelated() bool {
	if o == nil || o.Uncorrelated == nil {
		var ret bool
		return ret
	}
	return *o.Uncorrelated
}

// GetUncorrelatedOk returns a tuple with the Uncorrelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentAllOf) GetUncorrelatedOk() (*bool, bool) {
	if o == nil || o.Uncorrelated == nil {
		return nil, false
	}
	return o.Uncorrelated, true
}

// HasUncorrelated returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasUncorrelated() bool {
	if o != nil && o.Uncorrelated != nil {
		return true
	}

	return false
}

// SetUncorrelated gets a reference to the given bool and assigns it to the Uncorrelated field.
func (o *AccountDocumentAllOf) SetUncorrelated(v bool) {
	o.Uncorrelated = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AccountDocumentAllOf) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDocumentAllOf) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AccountDocumentAllOf) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AccountDocumentAllOf) SetTags(v []string) {
	o.Tags = v
}

func (o AccountDocumentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.EntitlementCount != nil {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if o.Uncorrelated != nil {
		toSerialize["uncorrelated"] = o.Uncorrelated
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountDocumentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccountDocumentAllOf := _AccountDocumentAllOf{}

	if err = json.Unmarshal(bytes, &varAccountDocumentAllOf); err == nil {
		*o = AccountDocumentAllOf(varAccountDocumentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "modified")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "access")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "uncorrelated")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountDocumentAllOf struct {
	value *AccountDocumentAllOf
	isSet bool
}

func (v NullableAccountDocumentAllOf) Get() *AccountDocumentAllOf {
	return v.value
}

func (v *NullableAccountDocumentAllOf) Set(val *AccountDocumentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDocumentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDocumentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDocumentAllOf(val *AccountDocumentAllOf) *NullableAccountDocumentAllOf {
	return &NullableAccountDocumentAllOf{value: val, isSet: true}
}

func (v NullableAccountDocumentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDocumentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


