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

// Source1AllOf struct for Source1AllOf
type Source1AllOf struct {
	// the type of source returned
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Source1AllOf Source1AllOf

// NewSource1AllOf instantiates a new Source1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource1AllOf() *Source1AllOf {
	this := Source1AllOf{}
	return &this
}

// NewSource1AllOfWithDefaults instantiates a new Source1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSource1AllOfWithDefaults() *Source1AllOf {
	this := Source1AllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Source1AllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source1AllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Source1AllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Source1AllOf) SetType(v string) {
	o.Type = &v
}

func (o Source1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Source1AllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSource1AllOf := _Source1AllOf{}

	if err = json.Unmarshal(bytes, &varSource1AllOf); err == nil {
		*o = Source1AllOf(varSource1AllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSource1AllOf struct {
	value *Source1AllOf
	isSet bool
}

func (v NullableSource1AllOf) Get() *Source1AllOf {
	return v.value
}

func (v *NullableSource1AllOf) Set(val *Source1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSource1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSource1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource1AllOf(val *Source1AllOf) *NullableSource1AllOf {
	return &NullableSource1AllOf{value: val, isSet: true}
}

func (v NullableSource1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


