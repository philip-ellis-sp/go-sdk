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

// PasswordStatus struct for PasswordStatus
type PasswordStatus struct {
	// The password change request ID
	RequestId NullableString `json:"requestId,omitempty"`
	// Password change state
	State *string `json:"state,omitempty"`
	// The errors during the password change request
	Errors []string `json:"errors,omitempty"`
	// List of source IDs in the password change request
	SourceIds []string `json:"sourceIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordStatus PasswordStatus

// NewPasswordStatus instantiates a new PasswordStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordStatus() *PasswordStatus {
	this := PasswordStatus{}
	return &this
}

// NewPasswordStatusWithDefaults instantiates a new PasswordStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordStatusWithDefaults() *PasswordStatus {
	this := PasswordStatus{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordStatus) GetRequestId() string {
	if o == nil || o.RequestId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordStatus) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *PasswordStatus) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *PasswordStatus) SetRequestId(v string) {
	o.RequestId.Set(&v)
}
// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *PasswordStatus) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *PasswordStatus) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PasswordStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PasswordStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PasswordStatus) SetState(v string) {
	o.State = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PasswordStatus) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStatus) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PasswordStatus) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *PasswordStatus) SetErrors(v []string) {
	o.Errors = v
}

// GetSourceIds returns the SourceIds field value if set, zero value otherwise.
func (o *PasswordStatus) GetSourceIds() []string {
	if o == nil || o.SourceIds == nil {
		var ret []string
		return ret
	}
	return o.SourceIds
}

// GetSourceIdsOk returns a tuple with the SourceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStatus) GetSourceIdsOk() ([]string, bool) {
	if o == nil || o.SourceIds == nil {
		return nil, false
	}
	return o.SourceIds, true
}

// HasSourceIds returns a boolean if a field has been set.
func (o *PasswordStatus) HasSourceIds() bool {
	if o != nil && o.SourceIds != nil {
		return true
	}

	return false
}

// SetSourceIds gets a reference to the given []string and assigns it to the SourceIds field.
func (o *PasswordStatus) SetSourceIds(v []string) {
	o.SourceIds = v
}

func (o PasswordStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId.IsSet() {
		toSerialize["requestId"] = o.RequestId.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.SourceIds != nil {
		toSerialize["sourceIds"] = o.SourceIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordStatus) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordStatus := _PasswordStatus{}

	if err = json.Unmarshal(bytes, &varPasswordStatus); err == nil {
		*o = PasswordStatus(varPasswordStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "state")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "sourceIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordStatus struct {
	value *PasswordStatus
	isSet bool
}

func (v NullablePasswordStatus) Get() *PasswordStatus {
	return v.value
}

func (v *NullablePasswordStatus) Set(val *PasswordStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordStatus(val *PasswordStatus) *NullablePasswordStatus {
	return &NullablePasswordStatus{value: val, isSet: true}
}

func (v NullablePasswordStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


