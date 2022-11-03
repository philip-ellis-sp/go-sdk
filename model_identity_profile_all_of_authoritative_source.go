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

// IdentityProfileAllOfAuthoritativeSource struct for IdentityProfileAllOfAuthoritativeSource
type IdentityProfileAllOfAuthoritativeSource struct {
	// Type of the object to which this reference applies
	Type *string `json:"type,omitempty"`
	// ID of the object to which this reference applies
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the object to which this reference applies
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProfileAllOfAuthoritativeSource IdentityProfileAllOfAuthoritativeSource

// NewIdentityProfileAllOfAuthoritativeSource instantiates a new IdentityProfileAllOfAuthoritativeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProfileAllOfAuthoritativeSource() *IdentityProfileAllOfAuthoritativeSource {
	this := IdentityProfileAllOfAuthoritativeSource{}
	return &this
}

// NewIdentityProfileAllOfAuthoritativeSourceWithDefaults instantiates a new IdentityProfileAllOfAuthoritativeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProfileAllOfAuthoritativeSourceWithDefaults() *IdentityProfileAllOfAuthoritativeSource {
	this := IdentityProfileAllOfAuthoritativeSource{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityProfileAllOfAuthoritativeSource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOfAuthoritativeSource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityProfileAllOfAuthoritativeSource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityProfileAllOfAuthoritativeSource) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProfileAllOfAuthoritativeSource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOfAuthoritativeSource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProfileAllOfAuthoritativeSource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProfileAllOfAuthoritativeSource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityProfileAllOfAuthoritativeSource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProfileAllOfAuthoritativeSource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityProfileAllOfAuthoritativeSource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityProfileAllOfAuthoritativeSource) SetName(v string) {
	o.Name = &v
}

func (o IdentityProfileAllOfAuthoritativeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProfileAllOfAuthoritativeSource) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProfileAllOfAuthoritativeSource := _IdentityProfileAllOfAuthoritativeSource{}

	if err = json.Unmarshal(bytes, &varIdentityProfileAllOfAuthoritativeSource); err == nil {
		*o = IdentityProfileAllOfAuthoritativeSource(varIdentityProfileAllOfAuthoritativeSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityProfileAllOfAuthoritativeSource struct {
	value *IdentityProfileAllOfAuthoritativeSource
	isSet bool
}

func (v NullableIdentityProfileAllOfAuthoritativeSource) Get() *IdentityProfileAllOfAuthoritativeSource {
	return v.value
}

func (v *NullableIdentityProfileAllOfAuthoritativeSource) Set(val *IdentityProfileAllOfAuthoritativeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProfileAllOfAuthoritativeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProfileAllOfAuthoritativeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProfileAllOfAuthoritativeSource(val *IdentityProfileAllOfAuthoritativeSource) *NullableIdentityProfileAllOfAuthoritativeSource {
	return &NullableIdentityProfileAllOfAuthoritativeSource{value: val, isSet: true}
}

func (v NullableIdentityProfileAllOfAuthoritativeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProfileAllOfAuthoritativeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


