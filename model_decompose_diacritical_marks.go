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

// DecomposeDiacriticalMarks struct for DecomposeDiacriticalMarks
type DecomposeDiacriticalMarks struct {
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DecomposeDiacriticalMarks DecomposeDiacriticalMarks

// NewDecomposeDiacriticalMarks instantiates a new DecomposeDiacriticalMarks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecomposeDiacriticalMarks() *DecomposeDiacriticalMarks {
	this := DecomposeDiacriticalMarks{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewDecomposeDiacriticalMarksWithDefaults instantiates a new DecomposeDiacriticalMarks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecomposeDiacriticalMarksWithDefaults() *DecomposeDiacriticalMarks {
	this := DecomposeDiacriticalMarks{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *DecomposeDiacriticalMarks) GetRequiresPeriodicRefresh() bool {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecomposeDiacriticalMarks) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *DecomposeDiacriticalMarks) HasRequiresPeriodicRefresh() bool {
	if o != nil && o.RequiresPeriodicRefresh != nil {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *DecomposeDiacriticalMarks) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *DecomposeDiacriticalMarks) GetInput() map[string]interface{} {
	if o == nil || o.Input == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecomposeDiacriticalMarks) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *DecomposeDiacriticalMarks) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *DecomposeDiacriticalMarks) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o DecomposeDiacriticalMarks) MarshalJSON() ([]byte, error) {
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

func (o *DecomposeDiacriticalMarks) UnmarshalJSON(bytes []byte) (err error) {
	varDecomposeDiacriticalMarks := _DecomposeDiacriticalMarks{}

	if err = json.Unmarshal(bytes, &varDecomposeDiacriticalMarks); err == nil {
		*o = DecomposeDiacriticalMarks(varDecomposeDiacriticalMarks)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDecomposeDiacriticalMarks struct {
	value *DecomposeDiacriticalMarks
	isSet bool
}

func (v NullableDecomposeDiacriticalMarks) Get() *DecomposeDiacriticalMarks {
	return v.value
}

func (v *NullableDecomposeDiacriticalMarks) Set(val *DecomposeDiacriticalMarks) {
	v.value = val
	v.isSet = true
}

func (v NullableDecomposeDiacriticalMarks) IsSet() bool {
	return v.isSet
}

func (v *NullableDecomposeDiacriticalMarks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecomposeDiacriticalMarks(val *DecomposeDiacriticalMarks) *NullableDecomposeDiacriticalMarks {
	return &NullableDecomposeDiacriticalMarks{value: val, isSet: true}
}

func (v NullableDecomposeDiacriticalMarks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecomposeDiacriticalMarks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


