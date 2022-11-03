/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"os"
)

// UploadSourceAccountsSchemaRequest struct for UploadSourceAccountsSchemaRequest
type UploadSourceAccountsSchemaRequest struct {
	File **os.File `json:"file,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UploadSourceAccountsSchemaRequest UploadSourceAccountsSchemaRequest

// NewUploadSourceAccountsSchemaRequest instantiates a new UploadSourceAccountsSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadSourceAccountsSchemaRequest() *UploadSourceAccountsSchemaRequest {
	this := UploadSourceAccountsSchemaRequest{}
	return &this
}

// NewUploadSourceAccountsSchemaRequestWithDefaults instantiates a new UploadSourceAccountsSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadSourceAccountsSchemaRequestWithDefaults() *UploadSourceAccountsSchemaRequest {
	this := UploadSourceAccountsSchemaRequest{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *UploadSourceAccountsSchemaRequest) GetFile() *os.File {
	if o == nil || o.File == nil {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadSourceAccountsSchemaRequest) GetFileOk() (**os.File, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *UploadSourceAccountsSchemaRequest) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *UploadSourceAccountsSchemaRequest) SetFile(v *os.File) {
	o.File = &v
}

func (o UploadSourceAccountsSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.File != nil {
		toSerialize["file"] = o.File
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UploadSourceAccountsSchemaRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUploadSourceAccountsSchemaRequest := _UploadSourceAccountsSchemaRequest{}

	if err = json.Unmarshal(bytes, &varUploadSourceAccountsSchemaRequest); err == nil {
		*o = UploadSourceAccountsSchemaRequest(varUploadSourceAccountsSchemaRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadSourceAccountsSchemaRequest struct {
	value *UploadSourceAccountsSchemaRequest
	isSet bool
}

func (v NullableUploadSourceAccountsSchemaRequest) Get() *UploadSourceAccountsSchemaRequest {
	return v.value
}

func (v *NullableUploadSourceAccountsSchemaRequest) Set(val *UploadSourceAccountsSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadSourceAccountsSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadSourceAccountsSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadSourceAccountsSchemaRequest(val *UploadSourceAccountsSchemaRequest) *NullableUploadSourceAccountsSchemaRequest {
	return &NullableUploadSourceAccountsSchemaRequest{value: val, isSet: true}
}

func (v NullableUploadSourceAccountsSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadSourceAccountsSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

