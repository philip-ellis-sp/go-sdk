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

// ScheduledSearchCreateRequest struct for ScheduledSearchCreateRequest
type ScheduledSearchCreateRequest struct {
	// The name of the scheduled search. 
	Name *string `json:"name,omitempty"`
	// The description of the scheduled search. 
	Description *string `json:"description,omitempty"`
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

type _ScheduledSearchCreateRequest ScheduledSearchCreateRequest

// NewScheduledSearchCreateRequest instantiates a new ScheduledSearchCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledSearchCreateRequest(savedSearchId string, schedule Schedule1, recipients []TypedReference) *ScheduledSearchCreateRequest {
	this := ScheduledSearchCreateRequest{}
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

// NewScheduledSearchCreateRequestWithDefaults instantiates a new ScheduledSearchCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledSearchCreateRequestWithDefaults() *ScheduledSearchCreateRequest {
	this := ScheduledSearchCreateRequest{}
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScheduledSearchCreateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScheduledSearchCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScheduledSearchCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScheduledSearchCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSavedSearchId returns the SavedSearchId field value
func (o *ScheduledSearchCreateRequest) GetSavedSearchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SavedSearchId
}

// GetSavedSearchIdOk returns a tuple with the SavedSearchId field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetSavedSearchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavedSearchId, true
}

// SetSavedSearchId sets field value
func (o *ScheduledSearchCreateRequest) SetSavedSearchId(v string) {
	o.SavedSearchId = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledSearchCreateRequest) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledSearchCreateRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *ScheduledSearchCreateRequest) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *ScheduledSearchCreateRequest) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *ScheduledSearchCreateRequest) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledSearchCreateRequest) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledSearchCreateRequest) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *ScheduledSearchCreateRequest) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *ScheduledSearchCreateRequest) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *ScheduledSearchCreateRequest) UnsetModified() {
	o.Modified.Unset()
}

// GetSchedule returns the Schedule field value
func (o *ScheduledSearchCreateRequest) GetSchedule() Schedule1 {
	if o == nil {
		var ret Schedule1
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetScheduleOk() (*Schedule1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ScheduledSearchCreateRequest) SetSchedule(v Schedule1) {
	o.Schedule = v
}

// GetRecipients returns the Recipients field value
func (o *ScheduledSearchCreateRequest) GetRecipients() []TypedReference {
	if o == nil {
		var ret []TypedReference
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetRecipientsOk() ([]TypedReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ScheduledSearchCreateRequest) SetRecipients(v []TypedReference) {
	o.Recipients = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScheduledSearchCreateRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScheduledSearchCreateRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEmailEmptyResults returns the EmailEmptyResults field value if set, zero value otherwise.
func (o *ScheduledSearchCreateRequest) GetEmailEmptyResults() bool {
	if o == nil || o.EmailEmptyResults == nil {
		var ret bool
		return ret
	}
	return *o.EmailEmptyResults
}

// GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetEmailEmptyResultsOk() (*bool, bool) {
	if o == nil || o.EmailEmptyResults == nil {
		return nil, false
	}
	return o.EmailEmptyResults, true
}

// HasEmailEmptyResults returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasEmailEmptyResults() bool {
	if o != nil && o.EmailEmptyResults != nil {
		return true
	}

	return false
}

// SetEmailEmptyResults gets a reference to the given bool and assigns it to the EmailEmptyResults field.
func (o *ScheduledSearchCreateRequest) SetEmailEmptyResults(v bool) {
	o.EmailEmptyResults = &v
}

// GetDisplayQueryDetails returns the DisplayQueryDetails field value if set, zero value otherwise.
func (o *ScheduledSearchCreateRequest) GetDisplayQueryDetails() bool {
	if o == nil || o.DisplayQueryDetails == nil {
		var ret bool
		return ret
	}
	return *o.DisplayQueryDetails
}

// GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearchCreateRequest) GetDisplayQueryDetailsOk() (*bool, bool) {
	if o == nil || o.DisplayQueryDetails == nil {
		return nil, false
	}
	return o.DisplayQueryDetails, true
}

// HasDisplayQueryDetails returns a boolean if a field has been set.
func (o *ScheduledSearchCreateRequest) HasDisplayQueryDetails() bool {
	if o != nil && o.DisplayQueryDetails != nil {
		return true
	}

	return false
}

// SetDisplayQueryDetails gets a reference to the given bool and assigns it to the DisplayQueryDetails field.
func (o *ScheduledSearchCreateRequest) SetDisplayQueryDetails(v bool) {
	o.DisplayQueryDetails = &v
}

func (o ScheduledSearchCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
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

func (o *ScheduledSearchCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varScheduledSearchCreateRequest := _ScheduledSearchCreateRequest{}

	if err = json.Unmarshal(bytes, &varScheduledSearchCreateRequest); err == nil {
		*o = ScheduledSearchCreateRequest(varScheduledSearchCreateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
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

type NullableScheduledSearchCreateRequest struct {
	value *ScheduledSearchCreateRequest
	isSet bool
}

func (v NullableScheduledSearchCreateRequest) Get() *ScheduledSearchCreateRequest {
	return v.value
}

func (v *NullableScheduledSearchCreateRequest) Set(val *ScheduledSearchCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledSearchCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledSearchCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledSearchCreateRequest(val *ScheduledSearchCreateRequest) *NullableScheduledSearchCreateRequest {
	return &NullableScheduledSearchCreateRequest{value: val, isSet: true}
}

func (v NullableScheduledSearchCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledSearchCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


