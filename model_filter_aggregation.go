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

// FilterAggregation An additional filter to constrain the results of the search query.
type FilterAggregation struct {
	// The name of the filter aggregate to be included in the result.
	Name string `json:"name"`
	Type *FilterType1 `json:"type,omitempty"`
	// The search field to apply the filter to.  Prefix the field name with '@' to reference a nested object. 
	Field string `json:"field"`
	// The value to filter on.
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FilterAggregation FilterAggregation

// NewFilterAggregation instantiates a new FilterAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterAggregation(name string, field string, value string) *FilterAggregation {
	this := FilterAggregation{}
	this.Name = name
	this.Field = field
	this.Value = value
	return &this
}

// NewFilterAggregationWithDefaults instantiates a new FilterAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterAggregationWithDefaults() *FilterAggregation {
	this := FilterAggregation{}
	return &this
}

// GetName returns the Name field value
func (o *FilterAggregation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilterAggregation) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FilterAggregation) GetType() FilterType1 {
	if o == nil || o.Type == nil {
		var ret FilterType1
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetTypeOk() (*FilterType1, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FilterAggregation) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given FilterType1 and assigns it to the Type field.
func (o *FilterAggregation) SetType(v FilterType1) {
	o.Type = &v
}

// GetField returns the Field field value
func (o *FilterAggregation) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *FilterAggregation) SetField(v string) {
	o.Field = v
}

// GetValue returns the Value field value
func (o *FilterAggregation) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilterAggregation) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilterAggregation) SetValue(v string) {
	o.Value = v
}

func (o FilterAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FilterAggregation) UnmarshalJSON(bytes []byte) (err error) {
	varFilterAggregation := _FilterAggregation{}

	if err = json.Unmarshal(bytes, &varFilterAggregation); err == nil {
		*o = FilterAggregation(varFilterAggregation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "field")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilterAggregation struct {
	value *FilterAggregation
	isSet bool
}

func (v NullableFilterAggregation) Get() *FilterAggregation {
	return v.value
}

func (v *NullableFilterAggregation) Set(val *FilterAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterAggregation(val *FilterAggregation) *NullableFilterAggregation {
	return &NullableFilterAggregation{value: val, isSet: true}
}

func (v NullableFilterAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


