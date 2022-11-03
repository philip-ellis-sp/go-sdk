/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"fmt"
)

// NonEmployeeSchemaAttributeType Enum representing the type of data a schema attribute accepts.
type NonEmployeeSchemaAttributeType string

// List of NonEmployeeSchemaAttributeType
const (
	TEXTNonEmployeeSchemaAttributeType = "TEXT"
	DATENonEmployeeSchemaAttributeType = "DATE"
	IDENTITYNonEmployeeSchemaAttributeType = "IDENTITY"
)

// All allowed values of NonEmployeeSchemaAttributeType enum
var AllowedNonEmployeeSchemaAttributeTypeEnumValues = []NonEmployeeSchemaAttributeType{
	"TEXT",
	"DATE",
	"IDENTITY",
}

func (v *NonEmployeeSchemaAttributeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NonEmployeeSchemaAttributeType(value)
	for _, existing := range AllowedNonEmployeeSchemaAttributeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NonEmployeeSchemaAttributeType", value)
}

// NewNonEmployeeSchemaAttributeTypeFromValue returns a pointer to a valid NonEmployeeSchemaAttributeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNonEmployeeSchemaAttributeTypeFromValue(v string) (*NonEmployeeSchemaAttributeType, error) {
	ev := NonEmployeeSchemaAttributeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NonEmployeeSchemaAttributeType: valid values are %v", v, AllowedNonEmployeeSchemaAttributeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NonEmployeeSchemaAttributeType) IsValid() bool {
	for _, existing := range AllowedNonEmployeeSchemaAttributeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NonEmployeeSchemaAttributeType value
func (v NonEmployeeSchemaAttributeType) Ptr() *NonEmployeeSchemaAttributeType {
	return &v
}

type NullableNonEmployeeSchemaAttributeType struct {
	value *NonEmployeeSchemaAttributeType
	isSet bool
}

func (v NullableNonEmployeeSchemaAttributeType) Get() *NonEmployeeSchemaAttributeType {
	return v.value
}

func (v *NullableNonEmployeeSchemaAttributeType) Set(val *NonEmployeeSchemaAttributeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSchemaAttributeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSchemaAttributeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSchemaAttributeType(val *NonEmployeeSchemaAttributeType) *NullableNonEmployeeSchemaAttributeType {
	return &NullableNonEmployeeSchemaAttributeType{value: val, isSet: true}
}

func (v NullableNonEmployeeSchemaAttributeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSchemaAttributeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

