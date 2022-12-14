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

// CreatePersonalAccessTokenRequest Object for specifying the name of a personal access token to create
type CreatePersonalAccessTokenRequest struct {
	// The name of the personal access token (PAT) to be created. Cannot be the same as another PAT owned by the user for whom this PAT is being created.
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreatePersonalAccessTokenRequest CreatePersonalAccessTokenRequest

// NewCreatePersonalAccessTokenRequest instantiates a new CreatePersonalAccessTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePersonalAccessTokenRequest(name string) *CreatePersonalAccessTokenRequest {
	this := CreatePersonalAccessTokenRequest{}
	this.Name = name
	return &this
}

// NewCreatePersonalAccessTokenRequestWithDefaults instantiates a new CreatePersonalAccessTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePersonalAccessTokenRequestWithDefaults() *CreatePersonalAccessTokenRequest {
	this := CreatePersonalAccessTokenRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePersonalAccessTokenRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePersonalAccessTokenRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePersonalAccessTokenRequest) SetName(v string) {
	o.Name = v
}

func (o CreatePersonalAccessTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreatePersonalAccessTokenRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreatePersonalAccessTokenRequest := _CreatePersonalAccessTokenRequest{}

	if err = json.Unmarshal(bytes, &varCreatePersonalAccessTokenRequest); err == nil {
		*o = CreatePersonalAccessTokenRequest(varCreatePersonalAccessTokenRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePersonalAccessTokenRequest struct {
	value *CreatePersonalAccessTokenRequest
	isSet bool
}

func (v NullableCreatePersonalAccessTokenRequest) Get() *CreatePersonalAccessTokenRequest {
	return v.value
}

func (v *NullableCreatePersonalAccessTokenRequest) Set(val *CreatePersonalAccessTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePersonalAccessTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePersonalAccessTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePersonalAccessTokenRequest(val *CreatePersonalAccessTokenRequest) *NullableCreatePersonalAccessTokenRequest {
	return &NullableCreatePersonalAccessTokenRequest{value: val, isSet: true}
}

func (v NullableCreatePersonalAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePersonalAccessTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


