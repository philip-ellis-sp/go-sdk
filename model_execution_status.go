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

// ExecutionStatus The current state of execution.
type ExecutionStatus string

// List of ExecutionStatus
const (
	EXECUTINGExecutionStatus = "EXECUTING"
	VERIFYINGExecutionStatus = "VERIFYING"
	TERMINATEDExecutionStatus = "TERMINATED"
	COMPLETEDExecutionStatus = "COMPLETED"
)

// All allowed values of ExecutionStatus enum
var AllowedExecutionStatusEnumValues = []ExecutionStatus{
	"EXECUTING",
	"VERIFYING",
	"TERMINATED",
	"COMPLETED",
}

func (v *ExecutionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExecutionStatus(value)
	for _, existing := range AllowedExecutionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExecutionStatus", value)
}

// NewExecutionStatusFromValue returns a pointer to a valid ExecutionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExecutionStatusFromValue(v string) (*ExecutionStatus, error) {
	ev := ExecutionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExecutionStatus: valid values are %v", v, AllowedExecutionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExecutionStatus) IsValid() bool {
	for _, existing := range AllowedExecutionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExecutionStatus value
func (v ExecutionStatus) Ptr() *ExecutionStatus {
	return &v
}

type NullableExecutionStatus struct {
	value *ExecutionStatus
	isSet bool
}

func (v NullableExecutionStatus) Get() *ExecutionStatus {
	return v.value
}

func (v *NullableExecutionStatus) Set(val *ExecutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionStatus(val *ExecutionStatus) *NullableExecutionStatus {
	return &NullableExecutionStatus{value: val, isSet: true}
}

func (v NullableExecutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
