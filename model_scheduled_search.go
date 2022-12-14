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

// ScheduledSearch struct for ScheduledSearch
type ScheduledSearch struct {
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
	// The scheduled search ID. 
	Id *string `json:"id,omitempty"`
	Owner *TypedReference `json:"owner,omitempty"`
	// The ID of the scheduled search owner
	OwnerId *string `json:"ownerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledSearch ScheduledSearch

// NewScheduledSearch instantiates a new ScheduledSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledSearch(savedSearchId string, schedule Schedule1, recipients []TypedReference) *ScheduledSearch {
	this := ScheduledSearch{}
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

// NewScheduledSearchWithDefaults instantiates a new ScheduledSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledSearchWithDefaults() *ScheduledSearch {
	this := ScheduledSearch{}
	var enabled bool = false
	this.Enabled = &enabled
	var emailEmptyResults bool = false
	this.EmailEmptyResults = &emailEmptyResults
	var displayQueryDetails bool = false
	this.DisplayQueryDetails = &displayQueryDetails
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScheduledSearch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScheduledSearch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScheduledSearch) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScheduledSearch) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScheduledSearch) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScheduledSearch) SetDescription(v string) {
	o.Description = &v
}

// GetSavedSearchId returns the SavedSearchId field value
func (o *ScheduledSearch) GetSavedSearchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SavedSearchId
}

// GetSavedSearchIdOk returns a tuple with the SavedSearchId field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetSavedSearchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SavedSearchId, true
}

// SetSavedSearchId sets field value
func (o *ScheduledSearch) SetSavedSearchId(v string) {
	o.SavedSearchId = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledSearch) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledSearch) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *ScheduledSearch) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *ScheduledSearch) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *ScheduledSearch) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *ScheduledSearch) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledSearch) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledSearch) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *ScheduledSearch) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *ScheduledSearch) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *ScheduledSearch) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *ScheduledSearch) UnsetModified() {
	o.Modified.Unset()
}

// GetSchedule returns the Schedule field value
func (o *ScheduledSearch) GetSchedule() Schedule1 {
	if o == nil {
		var ret Schedule1
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetScheduleOk() (*Schedule1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *ScheduledSearch) SetSchedule(v Schedule1) {
	o.Schedule = v
}

// GetRecipients returns the Recipients field value
func (o *ScheduledSearch) GetRecipients() []TypedReference {
	if o == nil {
		var ret []TypedReference
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetRecipientsOk() ([]TypedReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ScheduledSearch) SetRecipients(v []TypedReference) {
	o.Recipients = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScheduledSearch) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScheduledSearch) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScheduledSearch) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEmailEmptyResults returns the EmailEmptyResults field value if set, zero value otherwise.
func (o *ScheduledSearch) GetEmailEmptyResults() bool {
	if o == nil || o.EmailEmptyResults == nil {
		var ret bool
		return ret
	}
	return *o.EmailEmptyResults
}

// GetEmailEmptyResultsOk returns a tuple with the EmailEmptyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetEmailEmptyResultsOk() (*bool, bool) {
	if o == nil || o.EmailEmptyResults == nil {
		return nil, false
	}
	return o.EmailEmptyResults, true
}

// HasEmailEmptyResults returns a boolean if a field has been set.
func (o *ScheduledSearch) HasEmailEmptyResults() bool {
	if o != nil && o.EmailEmptyResults != nil {
		return true
	}

	return false
}

// SetEmailEmptyResults gets a reference to the given bool and assigns it to the EmailEmptyResults field.
func (o *ScheduledSearch) SetEmailEmptyResults(v bool) {
	o.EmailEmptyResults = &v
}

// GetDisplayQueryDetails returns the DisplayQueryDetails field value if set, zero value otherwise.
func (o *ScheduledSearch) GetDisplayQueryDetails() bool {
	if o == nil || o.DisplayQueryDetails == nil {
		var ret bool
		return ret
	}
	return *o.DisplayQueryDetails
}

// GetDisplayQueryDetailsOk returns a tuple with the DisplayQueryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetDisplayQueryDetailsOk() (*bool, bool) {
	if o == nil || o.DisplayQueryDetails == nil {
		return nil, false
	}
	return o.DisplayQueryDetails, true
}

// HasDisplayQueryDetails returns a boolean if a field has been set.
func (o *ScheduledSearch) HasDisplayQueryDetails() bool {
	if o != nil && o.DisplayQueryDetails != nil {
		return true
	}

	return false
}

// SetDisplayQueryDetails gets a reference to the given bool and assigns it to the DisplayQueryDetails field.
func (o *ScheduledSearch) SetDisplayQueryDetails(v bool) {
	o.DisplayQueryDetails = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduledSearch) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduledSearch) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduledSearch) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ScheduledSearch) GetOwner() TypedReference {
	if o == nil || o.Owner == nil {
		var ret TypedReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetOwnerOk() (*TypedReference, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ScheduledSearch) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given TypedReference and assigns it to the Owner field.
func (o *ScheduledSearch) SetOwner(v TypedReference) {
	o.Owner = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ScheduledSearch) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledSearch) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ScheduledSearch) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ScheduledSearch) SetOwnerId(v string) {
	o.OwnerId = &v
}

func (o ScheduledSearch) MarshalJSON() ([]byte, error) {
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScheduledSearch) UnmarshalJSON(bytes []byte) (err error) {
	varScheduledSearch := _ScheduledSearch{}

	if err = json.Unmarshal(bytes, &varScheduledSearch); err == nil {
		*o = ScheduledSearch(varScheduledSearch)
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
		delete(additionalProperties, "id")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "ownerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledSearch struct {
	value *ScheduledSearch
	isSet bool
}

func (v NullableScheduledSearch) Get() *ScheduledSearch {
	return v.value
}

func (v *NullableScheduledSearch) Set(val *ScheduledSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledSearch(val *ScheduledSearch) *NullableScheduledSearch {
	return &NullableScheduledSearch{value: val, isSet: true}
}

func (v NullableScheduledSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


