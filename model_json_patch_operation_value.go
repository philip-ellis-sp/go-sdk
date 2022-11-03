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

// JsonPatchOperationValue The value to be used for the operation, required for \"add\" and \"replace\" operations
type JsonPatchOperationValue struct {
	int32 *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *JsonPatchOperationValue) UnmarshalJSON(data []byte) error {
	var err error

	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32);
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(JsonPatchOperationValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *JsonPatchOperationValue) MarshalJSON() ([]byte, error) {

	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableJsonPatchOperationValue struct {
	value *JsonPatchOperationValue
	isSet bool
}

func (v NullableJsonPatchOperationValue) Get() *JsonPatchOperationValue {
	return v.value
}

func (v *NullableJsonPatchOperationValue) Set(val *JsonPatchOperationValue) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchOperationValue) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchOperationValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchOperationValue(val *JsonPatchOperationValue) *NullableJsonPatchOperationValue {
	return &NullableJsonPatchOperationValue{value: val, isSet: true}
}

func (v NullableJsonPatchOperationValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchOperationValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


