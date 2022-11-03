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

// NonEmployeeBulkUploadStatus struct for NonEmployeeBulkUploadStatus
type NonEmployeeBulkUploadStatus struct {
	// Returns the following values indicating the progress or result of the bulk upload job. \"PENDING\" means the job is queued and waiting to be processed. \"IN_PROGRESS\" means the job is currently being processed. \"COMPLETED\" means the job has been completed without any errors. \"ERROR\" means the job failed to process with errors. null means job has been submitted to the source. 
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeBulkUploadStatus NonEmployeeBulkUploadStatus

// NewNonEmployeeBulkUploadStatus instantiates a new NonEmployeeBulkUploadStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeBulkUploadStatus() *NonEmployeeBulkUploadStatus {
	this := NonEmployeeBulkUploadStatus{}
	return &this
}

// NewNonEmployeeBulkUploadStatusWithDefaults instantiates a new NonEmployeeBulkUploadStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeBulkUploadStatusWithDefaults() *NonEmployeeBulkUploadStatus {
	this := NonEmployeeBulkUploadStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NonEmployeeBulkUploadStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeBulkUploadStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NonEmployeeBulkUploadStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NonEmployeeBulkUploadStatus) SetStatus(v string) {
	o.Status = &v
}

func (o NonEmployeeBulkUploadStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeBulkUploadStatus) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeBulkUploadStatus := _NonEmployeeBulkUploadStatus{}

	if err = json.Unmarshal(bytes, &varNonEmployeeBulkUploadStatus); err == nil {
		*o = NonEmployeeBulkUploadStatus(varNonEmployeeBulkUploadStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeBulkUploadStatus struct {
	value *NonEmployeeBulkUploadStatus
	isSet bool
}

func (v NullableNonEmployeeBulkUploadStatus) Get() *NonEmployeeBulkUploadStatus {
	return v.value
}

func (v *NullableNonEmployeeBulkUploadStatus) Set(val *NonEmployeeBulkUploadStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeBulkUploadStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeBulkUploadStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeBulkUploadStatus(val *NonEmployeeBulkUploadStatus) *NullableNonEmployeeBulkUploadStatus {
	return &NullableNonEmployeeBulkUploadStatus{value: val, isSet: true}
}

func (v NullableNonEmployeeBulkUploadStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeBulkUploadStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

