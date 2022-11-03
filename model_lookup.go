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

// Lookup struct for Lookup
type Lookup struct {
	// This is a JSON object of key-value pairs. The key is the string that will attempt to be matched to the input, and the value is the output string that should be returned if the key is matched   >**Note** the use of the optional default key value here; if none of the three countries in the above example match the input string, the transform will return \"Unknown Region\" for the attribute that is mapped to this transform. 
	Table map[string]interface{} `json:"table"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Lookup Lookup

// NewLookup instantiates a new Lookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookup(table map[string]interface{}) *Lookup {
	this := Lookup{}
	this.Table = table
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewLookupWithDefaults instantiates a new Lookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupWithDefaults() *Lookup {
	this := Lookup{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetTable returns the Table field value
func (o *Lookup) GetTable() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *Lookup) GetTableOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Table, true
}

// SetTable sets field value
func (o *Lookup) SetTable(v map[string]interface{}) {
	o.Table = v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *Lookup) GetRequiresPeriodicRefresh() bool {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lookup) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *Lookup) HasRequiresPeriodicRefresh() bool {
	if o != nil && o.RequiresPeriodicRefresh != nil {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *Lookup) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *Lookup) GetInput() map[string]interface{} {
	if o == nil || o.Input == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lookup) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *Lookup) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *Lookup) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o Lookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["table"] = o.Table
	}
	if o.RequiresPeriodicRefresh != nil {
		toSerialize["requiresPeriodicRefresh"] = o.RequiresPeriodicRefresh
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Lookup) UnmarshalJSON(bytes []byte) (err error) {
	varLookup := _Lookup{}

	if err = json.Unmarshal(bytes, &varLookup); err == nil {
		*o = Lookup(varLookup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "table")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLookup struct {
	value *Lookup
	isSet bool
}

func (v NullableLookup) Get() *Lookup {
	return v.value
}

func (v *NullableLookup) Set(val *Lookup) {
	v.value = val
	v.isSet = true
}

func (v NullableLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookup(val *Lookup) *NullableLookup {
	return &NullableLookup{value: val, isSet: true}
}

func (v NullableLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


