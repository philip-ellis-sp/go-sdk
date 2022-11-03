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

// QueryType The type of query to use.  By default, the `SAILPOINT` query type is used, which requires the `query` object to be defined in the request body. To use the `queryDsl` or `typeAheadQuery` objects in the request, you must set the type to `DSL` or `TYPEAHEAD` accordingly. Additional values may be added in the future without notice.
type QueryType string

// List of QueryType
const (
	DSLQueryType = "DSL"
	SAILPOINTQueryType = "SAILPOINT"
	TYPEAHEADQueryType = "TYPEAHEAD"
)

// All allowed values of QueryType enum
var AllowedQueryTypeEnumValues = []QueryType{
	"DSL",
	"SAILPOINT",
	"TYPEAHEAD",
}

func (v *QueryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryType(value)
	for _, existing := range AllowedQueryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryType", value)
}

// NewQueryTypeFromValue returns a pointer to a valid QueryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryTypeFromValue(v string) (*QueryType, error) {
	ev := QueryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryType: valid values are %v", v, AllowedQueryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryType) IsValid() bool {
	for _, existing := range AllowedQueryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryType value
func (v QueryType) Ptr() *QueryType {
	return &v
}

type NullableQueryType struct {
	value *QueryType
	isSet bool
}

func (v NullableQueryType) Get() *QueryType {
	return v.value
}

func (v *NullableQueryType) Set(val *QueryType) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryType) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryType(val *QueryType) *NullableQueryType {
	return &NullableQueryType{value: val, isSet: true}
}

func (v NullableQueryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

