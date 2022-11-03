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

// DateMath struct for DateMath
type DateMath struct {
	// A string value of the date and time components to operation on, along with the math operations to execute. 
	Expression string `json:"expression"`
	// A boolean value to indicate whether the transform should round up or down when a rounding `/` operation is defined in the expression.    If not provided, the transform will default to `false`   `true` indicates the transform should round up (i.e., truncate the fractional date/time component indicated and then add one unit of that component)   `false` indicates the transform should round down (i.e., truncate the fractional date/time component indicated) 
	RoundUp *bool `json:"roundUp,omitempty"`
	// A value that indicates whether the transform logic should be re-evaluated every evening as part of the identity refresh process
	RequiresPeriodicRefresh *bool `json:"requiresPeriodicRefresh,omitempty"`
	// This is an optional attribute that can explicitly define the input data which will be fed into the transform logic. If input is not provided, the transform will take its input from the source and attribute combination configured via the UI.
	Input map[string]interface{} `json:"input,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DateMath DateMath

// NewDateMath instantiates a new DateMath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateMath(expression string) *DateMath {
	this := DateMath{}
	this.Expression = expression
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// NewDateMathWithDefaults instantiates a new DateMath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateMathWithDefaults() *DateMath {
	this := DateMath{}
	var requiresPeriodicRefresh bool = false
	this.RequiresPeriodicRefresh = &requiresPeriodicRefresh
	return &this
}

// GetExpression returns the Expression field value
func (o *DateMath) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *DateMath) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *DateMath) SetExpression(v string) {
	o.Expression = v
}

// GetRoundUp returns the RoundUp field value if set, zero value otherwise.
func (o *DateMath) GetRoundUp() bool {
	if o == nil || o.RoundUp == nil {
		var ret bool
		return ret
	}
	return *o.RoundUp
}

// GetRoundUpOk returns a tuple with the RoundUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateMath) GetRoundUpOk() (*bool, bool) {
	if o == nil || o.RoundUp == nil {
		return nil, false
	}
	return o.RoundUp, true
}

// HasRoundUp returns a boolean if a field has been set.
func (o *DateMath) HasRoundUp() bool {
	if o != nil && o.RoundUp != nil {
		return true
	}

	return false
}

// SetRoundUp gets a reference to the given bool and assigns it to the RoundUp field.
func (o *DateMath) SetRoundUp(v bool) {
	o.RoundUp = &v
}

// GetRequiresPeriodicRefresh returns the RequiresPeriodicRefresh field value if set, zero value otherwise.
func (o *DateMath) GetRequiresPeriodicRefresh() bool {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		var ret bool
		return ret
	}
	return *o.RequiresPeriodicRefresh
}

// GetRequiresPeriodicRefreshOk returns a tuple with the RequiresPeriodicRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateMath) GetRequiresPeriodicRefreshOk() (*bool, bool) {
	if o == nil || o.RequiresPeriodicRefresh == nil {
		return nil, false
	}
	return o.RequiresPeriodicRefresh, true
}

// HasRequiresPeriodicRefresh returns a boolean if a field has been set.
func (o *DateMath) HasRequiresPeriodicRefresh() bool {
	if o != nil && o.RequiresPeriodicRefresh != nil {
		return true
	}

	return false
}

// SetRequiresPeriodicRefresh gets a reference to the given bool and assigns it to the RequiresPeriodicRefresh field.
func (o *DateMath) SetRequiresPeriodicRefresh(v bool) {
	o.RequiresPeriodicRefresh = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *DateMath) GetInput() map[string]interface{} {
	if o == nil || o.Input == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateMath) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *DateMath) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *DateMath) SetInput(v map[string]interface{}) {
	o.Input = v
}

func (o DateMath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if o.RoundUp != nil {
		toSerialize["roundUp"] = o.RoundUp
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

func (o *DateMath) UnmarshalJSON(bytes []byte) (err error) {
	varDateMath := _DateMath{}

	if err = json.Unmarshal(bytes, &varDateMath); err == nil {
		*o = DateMath(varDateMath)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "roundUp")
		delete(additionalProperties, "requiresPeriodicRefresh")
		delete(additionalProperties, "input")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDateMath struct {
	value *DateMath
	isSet bool
}

func (v NullableDateMath) Get() *DateMath {
	return v.value
}

func (v *NullableDateMath) Set(val *DateMath) {
	v.value = val
	v.isSet = true
}

func (v NullableDateMath) IsSet() bool {
	return v.isSet
}

func (v *NullableDateMath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateMath(val *DateMath) *NullableDateMath {
	return &NullableDateMath{value: val, isSet: true}
}

func (v NullableDateMath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateMath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

