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

// NonEmployeeSourceLiteWithSchemaAttributesAllOf struct for NonEmployeeSourceLiteWithSchemaAttributesAllOf
type NonEmployeeSourceLiteWithSchemaAttributesAllOf struct {
	// List of schema attributes associated with this non-employee source.
	SchemaAttributes []NonEmployeeSchemaAttribute `json:"schemaAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceLiteWithSchemaAttributesAllOf NonEmployeeSourceLiteWithSchemaAttributesAllOf

// NewNonEmployeeSourceLiteWithSchemaAttributesAllOf instantiates a new NonEmployeeSourceLiteWithSchemaAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceLiteWithSchemaAttributesAllOf() *NonEmployeeSourceLiteWithSchemaAttributesAllOf {
	this := NonEmployeeSourceLiteWithSchemaAttributesAllOf{}
	return &this
}

// NewNonEmployeeSourceLiteWithSchemaAttributesAllOfWithDefaults instantiates a new NonEmployeeSourceLiteWithSchemaAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceLiteWithSchemaAttributesAllOfWithDefaults() *NonEmployeeSourceLiteWithSchemaAttributesAllOf {
	this := NonEmployeeSourceLiteWithSchemaAttributesAllOf{}
	return &this
}

// GetSchemaAttributes returns the SchemaAttributes field value if set, zero value otherwise.
func (o *NonEmployeeSourceLiteWithSchemaAttributesAllOf) GetSchemaAttributes() []NonEmployeeSchemaAttribute {
	if o == nil || o.SchemaAttributes == nil {
		var ret []NonEmployeeSchemaAttribute
		return ret
	}
	return o.SchemaAttributes
}

// GetSchemaAttributesOk returns a tuple with the SchemaAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributesAllOf) GetSchemaAttributesOk() ([]NonEmployeeSchemaAttribute, bool) {
	if o == nil || o.SchemaAttributes == nil {
		return nil, false
	}
	return o.SchemaAttributes, true
}

// HasSchemaAttributes returns a boolean if a field has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributesAllOf) HasSchemaAttributes() bool {
	if o != nil && o.SchemaAttributes != nil {
		return true
	}

	return false
}

// SetSchemaAttributes gets a reference to the given []NonEmployeeSchemaAttribute and assigns it to the SchemaAttributes field.
func (o *NonEmployeeSourceLiteWithSchemaAttributesAllOf) SetSchemaAttributes(v []NonEmployeeSchemaAttribute) {
	o.SchemaAttributes = v
}

func (o NonEmployeeSourceLiteWithSchemaAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SchemaAttributes != nil {
		toSerialize["schemaAttributes"] = o.SchemaAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeSourceLiteWithSchemaAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeSourceLiteWithSchemaAttributesAllOf := _NonEmployeeSourceLiteWithSchemaAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varNonEmployeeSourceLiteWithSchemaAttributesAllOf); err == nil {
		*o = NonEmployeeSourceLiteWithSchemaAttributesAllOf(varNonEmployeeSourceLiteWithSchemaAttributesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "schemaAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf struct {
	value *NonEmployeeSourceLiteWithSchemaAttributesAllOf
	isSet bool
}

func (v NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf) Get() *NonEmployeeSourceLiteWithSchemaAttributesAllOf {
	return v.value
}

func (v *NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf) Set(val *NonEmployeeSourceLiteWithSchemaAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceLiteWithSchemaAttributesAllOf(val *NonEmployeeSourceLiteWithSchemaAttributesAllOf) *NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf {
	return &NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceLiteWithSchemaAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


