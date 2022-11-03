/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// NonEmployeeApprovalItemDetailAllOf struct for NonEmployeeApprovalItemDetailAllOf
type NonEmployeeApprovalItemDetailAllOf struct {
	NonEmployeeRequest *NonEmployeeRequestWithoutApprovalItem `json:"nonEmployeeRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeApprovalItemDetailAllOf NonEmployeeApprovalItemDetailAllOf

// NewNonEmployeeApprovalItemDetailAllOf instantiates a new NonEmployeeApprovalItemDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeApprovalItemDetailAllOf() *NonEmployeeApprovalItemDetailAllOf {
	this := NonEmployeeApprovalItemDetailAllOf{}
	return &this
}

// NewNonEmployeeApprovalItemDetailAllOfWithDefaults instantiates a new NonEmployeeApprovalItemDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeApprovalItemDetailAllOfWithDefaults() *NonEmployeeApprovalItemDetailAllOf {
	this := NonEmployeeApprovalItemDetailAllOf{}
	return &this
}

// GetNonEmployeeRequest returns the NonEmployeeRequest field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetailAllOf) GetNonEmployeeRequest() NonEmployeeRequestWithoutApprovalItem {
	if o == nil || o.NonEmployeeRequest == nil {
		var ret NonEmployeeRequestWithoutApprovalItem
		return ret
	}
	return *o.NonEmployeeRequest
}

// GetNonEmployeeRequestOk returns a tuple with the NonEmployeeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetailAllOf) GetNonEmployeeRequestOk() (*NonEmployeeRequestWithoutApprovalItem, bool) {
	if o == nil || o.NonEmployeeRequest == nil {
		return nil, false
	}
	return o.NonEmployeeRequest, true
}

// HasNonEmployeeRequest returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetailAllOf) HasNonEmployeeRequest() bool {
	if o != nil && o.NonEmployeeRequest != nil {
		return true
	}

	return false
}

// SetNonEmployeeRequest gets a reference to the given NonEmployeeRequestWithoutApprovalItem and assigns it to the NonEmployeeRequest field.
func (o *NonEmployeeApprovalItemDetailAllOf) SetNonEmployeeRequest(v NonEmployeeRequestWithoutApprovalItem) {
	o.NonEmployeeRequest = &v
}

func (o NonEmployeeApprovalItemDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NonEmployeeRequest != nil {
		toSerialize["nonEmployeeRequest"] = o.NonEmployeeRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeApprovalItemDetailAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeApprovalItemDetailAllOf := _NonEmployeeApprovalItemDetailAllOf{}

	if err = json.Unmarshal(bytes, &varNonEmployeeApprovalItemDetailAllOf); err == nil {
		*o = NonEmployeeApprovalItemDetailAllOf(varNonEmployeeApprovalItemDetailAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nonEmployeeRequest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeApprovalItemDetailAllOf struct {
	value *NonEmployeeApprovalItemDetailAllOf
	isSet bool
}

func (v NullableNonEmployeeApprovalItemDetailAllOf) Get() *NonEmployeeApprovalItemDetailAllOf {
	return v.value
}

func (v *NullableNonEmployeeApprovalItemDetailAllOf) Set(val *NonEmployeeApprovalItemDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeApprovalItemDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeApprovalItemDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeApprovalItemDetailAllOf(val *NonEmployeeApprovalItemDetailAllOf) *NullableNonEmployeeApprovalItemDetailAllOf {
	return &NullableNonEmployeeApprovalItemDetailAllOf{value: val, isSet: true}
}

func (v NullableNonEmployeeApprovalItemDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeApprovalItemDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


