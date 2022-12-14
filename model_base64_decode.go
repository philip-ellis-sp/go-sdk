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

// Base64Decode struct for Base64Decode
type Base64Decode struct {
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Base64Decode Base64Decode

// NewBase64Decode instantiates a new Base64Decode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBase64Decode() *Base64Decode {
	this := Base64Decode{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewBase64DecodeWithDefaults instantiates a new Base64Decode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBase64DecodeWithDefaults() *Base64Decode {
	this := Base64Decode{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *Base64Decode) GetRequiresPeriodicRefresh() bool {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base64Decode) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *Base64Decode) HasRequiresPeriodicRefresh() bool {
	if o != nil && o.RequiresPeriodicRefresh != nil {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *Base64Decode) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *Base64Decode) GetInput() map[string]interface{} {
	if o == nil || o.Input == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base64Decode) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *Base64Decode) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *Base64Decode) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o Base64Decode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequiresPeriodicRefresh != nil {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Base64Decode) UnmarshalJSON(bytes []byte) (err error) {
	varBase64Decode := _Base64Decode{}

	if err = json.Unmarshal(bytes, &varBase64Decode); err == nil {
		*o = Base64Decode(varBase64Decode)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBase64Decode struct {
	value *Base64Decode
	isSet bool
}

func (v NullableBase64Decode) Get() *Base64Decode {
	return v.value
}

func (v *NullableBase64Decode) Set(val *Base64Decode) {
	v.value = val
	v.isSet = true
}

func (v NullableBase64Decode) IsSet() bool {
	return v.isSet
}

func (v *NullableBase64Decode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBase64Decode(val *Base64Decode) *NullableBase64Decode {
	return &NullableBase64Decode{value: val, isSet: true}
}

func (v NullableBase64Decode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBase64Decode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


