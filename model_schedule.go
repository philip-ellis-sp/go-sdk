/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// Schedule struct for Schedule
type Schedule struct {
	// The ID of the saved search that will be executed. 
	SavedSearchId string `json:"savedSearchId"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	Schedule Schedule1 `json:"schedule"`
	// The email recipients. 
	Recipients []TypedReference `json:"recipients"`
	// Indicates if the scheduled search is enabled. 
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates if email generation should not be suppressed if search returns no results. 
	EmailEmptyResults *bool `json:"emailEmptyResults,omitempty"`
	// Indicates if the generated email should include the query and search results preview (which could include PII). 
	DisplayQueryDetails *bool `json:"displayQueryDetails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Schedule Schedule

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule(savedSearchId string, schedule Schedule1, recipients []TypedReference) *Schedule {
	this := Schedule{}
	this.SavedSearchId = savedSearchId
	this.Schedule = schedule
	this.Recipients = recipients
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// GetSavedSearchId returns the SavedSearchId field value
func (o *Schedule) GetSavedSearchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SavedSearchId
}

// GetSavedSearchIdOk returns a tuple with the SavedSearchId field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetSavedSearchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavedSearchId, true
}

// SetSavedSearchId sets field value
func (o *Schedule) SetSavedSearchId(v string) {
	o.SavedSearchId = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Schedule) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Schedule) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Schedule) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Schedule) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Schedule) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Schedule) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *Schedule) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *Schedule) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *Schedule) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *Schedule) UnsetModified() {
	o.Modified.Unset()
}

// GetSchedule returns the Schedule field value
func (o *Schedule) GetSchedule() Schedule1 {
	if o == nil {
		var ret Schedule1
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetScheduleOk() (*Schedule1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *Schedule) SetSchedule(v Schedule1) {
	o.Schedule = v
}

// GetRecipients returns the Recipients field value
func (o *Schedule) GetRecipients() []TypedReference {
	if o == nil {
		var ret []TypedReference
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetRecipientsOk() ([]TypedReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *Schedule) SetRecipients(v []TypedReference) {
	o.Recipients = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Schedule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Schedule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Schedule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEmailEmptyResults returns the EmailEmptyResults field value if set, zero value otherwise.
func (o *Schedule) GetEmailEmptyResults() bool {
	if o == nil || o.EmailEmptyResults == nil {
		var ret bool
		return ret
	}
	return *o.EmailEmptyResults
}

// GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetEmailEmptyResultsOk() (*bool, bool) {
	if o == nil || o.EmailEmptyResults == nil {
		return nil, false
	}
	return o.EmailEmptyResults, true
}

// HasEmailEmptyResults returns a boolean if a field has been set.
func (o *Schedule) HasEmailEmptyResults() bool {
	if o != nil && o.EmailEmptyResults != nil {
		return true
	}

	return false
}

// SetEmailEmptyResults gets a reference to the given bool and assigns it to the EmailEmptyResults field.
func (o *Schedule) SetEmailEmptyResults(v bool) {
	o.EmailEmptyResults = &v
}

// GetDisplayQueryDetails returns the DisplayQueryDetails field value if set, zero value otherwise.
func (o *Schedule) GetDisplayQueryDetails() bool {
	if o == nil || o.DisplayQueryDetails == nil {
		var ret bool
		return ret
	}
	return *o.DisplayQueryDetails
}

// GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetDisplayQueryDetailsOk() (*bool, bool) {
	if o == nil || o.DisplayQueryDetails == nil {
		return nil, false
	}
	return o.DisplayQueryDetails, true
}

// HasDisplayQueryDetails returns a boolean if a field has been set.
func (o *Schedule) HasDisplayQueryDetails() bool {
	if o != nil && o.DisplayQueryDetails != nil {
		return true
	}

	return false
}

// SetDisplayQueryDetails gets a reference to the given bool and assigns it to the DisplayQueryDetails field.
func (o *Schedule) SetDisplayQueryDetails(v bool) {
	o.DisplayQueryDetails = &v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["savedSearchId"] = o.SavedSearchId
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EmailEmptyResults != nil {
		toSerialize["emailEmptyResults"] = o.EmailEmptyResults
	}
	if o.DisplayQueryDetails != nil {
		toSerialize["displayQueryDetails"] = o.DisplayQueryDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Schedule) UnmarshalJSON(bytes []byte) (err error) {
	varSchedule := _Schedule{}

	if err = json.Unmarshal(bytes, &varSchedule); err == nil {
		*o = Schedule(varSchedule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "savedSearchId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "emailEmptyResults")
		delete(additionalProperties, "displayQueryDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


