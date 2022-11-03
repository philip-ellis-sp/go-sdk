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

// MetricType Enum representing the currently supported metric aggregation types. Additional values may be added in the future without notice.
type MetricType string

// List of MetricType
const (
	COUNTMetricType = "COUNT"
	UNIQUE_COUNTMetricType = "UNIQUE_COUNT"
	AVGMetricType = "AVG"
	SUMMetricType = "SUM"
	MEDIANMetricType = "MEDIAN"
	MINMetricType = "MIN"
	MAXMetricType = "MAX"
)

// All allowed values of MetricType enum
var AllowedMetricTypeEnumValues = []MetricType{
	"COUNT",
	"UNIQUE_COUNT",
	"AVG",
	"SUM",
	"MEDIAN",
	"MIN",
	"MAX",
}

func (v *MetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricType(value)
	for _, existing := range AllowedMetricTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricType", value)
}

// NewMetricTypeFromValue returns a pointer to a valid MetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricTypeFromValue(v string) (*MetricType, error) {
	ev := MetricType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricType: valid values are %v", v, AllowedMetricTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricType) IsValid() bool {
	for _, existing := range AllowedMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricType value
func (v MetricType) Ptr() *MetricType {
	return &v
}

type NullableMetricType struct {
	value *MetricType
	isSet bool
}

func (v NullableMetricType) Get() *MetricType {
	return v.value
}

func (v *NullableMetricType) Set(val *MetricType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricType(val *MetricType) *NullableMetricType {
	return &NullableMetricType{value: val, isSet: true}
}

func (v NullableMetricType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
