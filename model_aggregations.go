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

// Aggregations struct for Aggregations
type Aggregations struct {
	Nested *NestedAggregation `json:"nested,omitempty"`
	Metric *MetricAggregation `json:"metric,omitempty"`
	Filter *FilterAggregation `json:"filter,omitempty"`
	Bucket *BucketAggregation `json:"bucket,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Aggregations Aggregations

// NewAggregations instantiates a new Aggregations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregations() *Aggregations {
	this := Aggregations{}
	return &this
}

// NewAggregationsWithDefaults instantiates a new Aggregations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationsWithDefaults() *Aggregations {
	this := Aggregations{}
	return &this
}

// GetNested returns the Nested field value if set, zero value otherwise.
func (o *Aggregations) GetNested() NestedAggregation {
	if o == nil || o.Nested == nil {
		var ret NestedAggregation
		return ret
	}
	return *o.Nested
}

// GetNestedOk returns a tuple with the Nested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregations) GetNestedOk() (*NestedAggregation, bool) {
	if o == nil || o.Nested == nil {
		return nil, false
	}
	return o.Nested, true
}

// HasNested returns a boolean if a field has been set.
func (o *Aggregations) HasNested() bool {
	if o != nil && o.Nested != nil {
		return true
	}

	return false
}

// SetNested gets a reference to the given NestedAggregation and assigns it to the Nested field.
func (o *Aggregations) SetNested(v NestedAggregation) {
	o.Nested = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *Aggregations) GetMetric() MetricAggregation {
	if o == nil || o.Metric == nil {
		var ret MetricAggregation
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregations) GetMetricOk() (*MetricAggregation, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *Aggregations) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given MetricAggregation and assigns it to the Metric field.
func (o *Aggregations) SetMetric(v MetricAggregation) {
	o.Metric = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *Aggregations) GetFilter() FilterAggregation {
	if o == nil || o.Filter == nil {
		var ret FilterAggregation
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregations) GetFilterOk() (*FilterAggregation, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *Aggregations) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given FilterAggregation and assigns it to the Filter field.
func (o *Aggregations) SetFilter(v FilterAggregation) {
	o.Filter = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *Aggregations) GetBucket() BucketAggregation {
	if o == nil || o.Bucket == nil {
		var ret BucketAggregation
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregations) GetBucketOk() (*BucketAggregation, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *Aggregations) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given BucketAggregation and assigns it to the Bucket field.
func (o *Aggregations) SetBucket(v BucketAggregation) {
	o.Bucket = &v
}

func (o Aggregations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nested != nil {
		toSerialize["nested"] = o.Nested
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Aggregations) UnmarshalJSON(bytes []byte) (err error) {
	varAggregations := _Aggregations{}

	if err = json.Unmarshal(bytes, &varAggregations); err == nil {
		*o = Aggregations(varAggregations)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nested")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "bucket")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregations struct {
	value *Aggregations
	isSet bool
}

func (v NullableAggregations) Get() *Aggregations {
	return v.value
}

func (v *NullableAggregations) Set(val *Aggregations) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregations) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregations(val *Aggregations) *NullableAggregations {
	return &NullableAggregations{value: val, isSet: true}
}

func (v NullableAggregations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


