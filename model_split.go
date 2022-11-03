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

// Split struct for Split
type Split struct {
	// This can be either a single character or a regex expression, and is used by the transform to identify the break point between two substrings in the incoming data
	Delimiter string `json:"delimiter"`
	// An integer value for the desired array element after the incoming data has been split into a list; the array is a 0-based object, so the first array element would be index 0, the second element would be index 1, etc.
	Index string `json:"index"`
	// A boolean (true/false) value which indicates whether an exception should be thrown and returned as an output when an index is out of bounds with the resultant array (i.e., the provided index value is larger than the size of the array)   `true` - The transform should return \"IndexOutOfBoundsException\"   `false` - The transform should return null   If not provided, the transform will default to false and return a null 
	Throws *bool `json:"throws,omitempty"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Split Split

// NewSplit instantiates a new Split object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplit(delimiter string, index string) *Split {
	this := Split{}
	this.Delimiter = delimiter
	this.Index = index
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewSplitWithDefaults instantiates a new Split object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitWithDefaults() *Split {
	this := Split{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetDelimiter returns the Delimiter field value
func (o *Split) GetDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value
// and a boolean to check if the value has been set.
func (o *Split) GetDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delimiter, true
}

// SetDelimiter sets field value
func (o *Split) SetDelimiter(v string) {
	o.Delimiter = v
}

// GetIndex returns the Index field value
func (o *Split) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *Split) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *Split) SetIndex(v string) {
	o.Index = v
}

// GetThrows returns the Throws field value if set, zero value otherwise.
func (o *Split) GetThrows() bool {
	if o == nil || o.Throws == nil {
		var ret bool
		return ret
	}
	return *o.Throws
}

// GetThrowsOk returns a tuple with the Throws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetThrowsOk() (*bool, bool) {
	if o == nil || o.Throws == nil {
		return nil, false
	}
	return o.Throws, true
}

// HasThrows returns a boolean if a field has been set.
func (o *Split) HasThrows() bool {
	if o != nil && o.Throws != nil {
		return true
	}

	return false
}

// SetThrows gets a reference to the given bool and assigns it to the Throws field.
func (o *Split) SetThrows(v bool) {
	o.Throws = &v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *Split) GetRequiresPeriodicRefresh() bool {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *Split) HasRequiresPeriodicRefresh() bool {
	if o != nil && o.RequiresPeriodicRefresh != nil {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *Split) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *Split) GetInput() map[string]interface{} {
	if o == nil || o.Input == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Split) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *Split) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *Split) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o Split) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["delimiter"] = o.Delimiter
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if o.Throws != nil {
		toSerialize["throws"] = o.Throws
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

func (o *Split) UnmarshalJSON(bytes []byte) (err error) {
	varSplit := _Split{}

	if err = json.Unmarshal(bytes, &varSplit); err == nil {
		*o = Split(varSplit)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delimiter")
		delete(additionalProperties, "index")
		delete(additionalProperties, "throws")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSplit struct {
	value *Split
	isSet bool
}

func (v NullableSplit) Get() *Split {
	return v.value
}

func (v *NullableSplit) Set(val *Split) {
	v.value = val
	v.isSet = true
}

func (v NullableSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplit(val *Split) *NullableSplit {
	return &NullableSplit{value: val, isSet: true}
}

func (v NullableSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

