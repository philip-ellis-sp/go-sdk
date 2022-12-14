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

// Identity Identity
type Identity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type DocumentType `json:"_type"`
	// The display name of the identity
	DisplayName *string `json:"displayName,omitempty"`
	// The first name of the identity
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the identity
	LastName *string `json:"lastName,omitempty"`
	// The identity's primary email address
	Email *string `json:"email,omitempty"`
	// A date-time in ISO-8601 format
	Created NullableTime `json:"created,omitempty"`
	// A date-time in ISO-8601 format
	Modified NullableTime `json:"modified,omitempty"`
	// A date-time in ISO-8601 format
	Synced NullableTime `json:"synced,omitempty"`
	// The phone number of the identity
	Phone *string `json:"phone,omitempty"`
	// Indicates if the identity is inactive
	Inactive *bool `json:"inactive,omitempty"`
	Protected *bool `json:"protected,omitempty"`
	// The identity's status in SailPoint
	Status *string `json:"status,omitempty"`
	EmployeeNumber *string `json:"employeeNumber,omitempty"`
	Manager *DisplayReference `json:"manager,omitempty"`
	// Indicates if this identity is a manager of other identities
	IsManager *bool `json:"isManager,omitempty"`
	IdentityProfile *Reference1 `json:"identityProfile,omitempty"`
	Source *Reference1 `json:"source,omitempty"`
	// a map or dictionary of key/value pairs
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	ProcessingState NullableString `json:"processingState,omitempty"`
	ProcessingDetails *ProcessingDetails `json:"processingDetails,omitempty"`
	// List of accounts associated with the identity
	Accounts []BaseAccount `json:"accounts,omitempty"`
	// Number of accounts associated with the identity
	AccountCount *int32 `json:"accountCount,omitempty"`
	// The list of applications the identity has access to
	Apps []App `json:"apps,omitempty"`
	// The number of applications the identity has access to
	AppCount *int32 `json:"appCount,omitempty"`
	// The list of access items assigned to the identity
	Access []Access1 `json:"access,omitempty"`
	// The number of access items assigned to the identity
	AccessCount *int32 `json:"accessCount,omitempty"`
	// The number of access profiles assigned to the identity
	AccessProfileCount *int32 `json:"accessProfileCount,omitempty"`
	// The number of entitlements assigned to the identity
	EntitlementCount *int32 `json:"entitlementCount,omitempty"`
	// The number of roles assigned to the identity
	RoleCount *int32 `json:"roleCount,omitempty"`
	Owns *Owns `json:"owns,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Identity Identity

// NewIdentity instantiates a new Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity(id string, name string, type_ DocumentType) *Identity {
	this := Identity{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewIdentityWithDefaults instantiates a new Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithDefaults() *Identity {
	this := Identity{}
	return &this
}

// GetId returns the Id field value
func (o *Identity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Identity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Identity) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Identity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Identity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Identity) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Identity) GetType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Identity) GetTypeOk() (*DocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Identity) SetType(v DocumentType) {
	o.Type = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Identity) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Identity) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Identity) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Identity) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Identity) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Identity) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Identity) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Identity) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Identity) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Identity) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Identity) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Identity) SetEmail(v string) {
	o.Email = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Identity) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Identity) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Identity) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Identity) UnsetCreated() {
	o.Created.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *Identity) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *Identity) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *Identity) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *Identity) UnsetModified() {
	o.Modified.Unset()
}

// GetSynced returns the Synced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetSynced() time.Time {
	if o == nil || o.Synced.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Synced.Get()
}

// GetSyncedOk returns a tuple with the Synced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetSyncedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Synced.Get(), o.Synced.IsSet()
}

// HasSynced returns a boolean if a field has been set.
func (o *Identity) HasSynced() bool {
	if o != nil && o.Synced.IsSet() {
		return true
	}

	return false
}

// SetSynced gets a reference to the given NullableTime and assigns it to the Synced field.
func (o *Identity) SetSynced(v time.Time) {
	o.Synced.Set(&v)
}
// SetSyncedNil sets the value for Synced to be an explicit nil
func (o *Identity) SetSyncedNil() {
	o.Synced.Set(nil)
}

// UnsetSynced ensures that no value is present for Synced, not even an explicit nil
func (o *Identity) UnsetSynced() {
	o.Synced.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Identity) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Identity) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Identity) SetPhone(v string) {
	o.Phone = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *Identity) GetInactive() bool {
	if o == nil || o.Inactive == nil {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetInactiveOk() (*bool, bool) {
	if o == nil || o.Inactive == nil {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *Identity) HasInactive() bool {
	if o != nil && o.Inactive != nil {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *Identity) SetInactive(v bool) {
	o.Inactive = &v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *Identity) GetProtected() bool {
	if o == nil || o.Protected == nil {
		var ret bool
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetProtectedOk() (*bool, bool) {
	if o == nil || o.Protected == nil {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *Identity) HasProtected() bool {
	if o != nil && o.Protected != nil {
		return true
	}

	return false
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *Identity) SetProtected(v bool) {
	o.Protected = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Identity) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Identity) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Identity) SetStatus(v string) {
	o.Status = &v
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise.
func (o *Identity) GetEmployeeNumber() string {
	if o == nil || o.EmployeeNumber == nil {
		var ret string
		return ret
	}
	return *o.EmployeeNumber
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetEmployeeNumberOk() (*string, bool) {
	if o == nil || o.EmployeeNumber == nil {
		return nil, false
	}
	return o.EmployeeNumber, true
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *Identity) HasEmployeeNumber() bool {
	if o != nil && o.EmployeeNumber != nil {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given string and assigns it to the EmployeeNumber field.
func (o *Identity) SetEmployeeNumber(v string) {
	o.EmployeeNumber = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *Identity) GetManager() DisplayReference {
	if o == nil || o.Manager == nil {
		var ret DisplayReference
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetManagerOk() (*DisplayReference, bool) {
	if o == nil || o.Manager == nil {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *Identity) HasManager() bool {
	if o != nil && o.Manager != nil {
		return true
	}

	return false
}

// SetManager gets a reference to the given DisplayReference and assigns it to the Manager field.
func (o *Identity) SetManager(v DisplayReference) {
	o.Manager = &v
}

// GetIsManager returns the IsManager field value if set, zero value otherwise.
func (o *Identity) GetIsManager() bool {
	if o == nil || o.IsManager == nil {
		var ret bool
		return ret
	}
	return *o.IsManager
}

// GetIsManagerOk returns a tuple with the IsManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetIsManagerOk() (*bool, bool) {
	if o == nil || o.IsManager == nil {
		return nil, false
	}
	return o.IsManager, true
}

// HasIsManager returns a boolean if a field has been set.
func (o *Identity) HasIsManager() bool {
	if o != nil && o.IsManager != nil {
		return true
	}

	return false
}

// SetIsManager gets a reference to the given bool and assigns it to the IsManager field.
func (o *Identity) SetIsManager(v bool) {
	o.IsManager = &v
}

// GetIdentityProfile returns the IdentityProfile field value if set, zero value otherwise.
func (o *Identity) GetIdentityProfile() Reference1 {
	if o == nil || o.IdentityProfile == nil {
		var ret Reference1
		return ret
	}
	return *o.IdentityProfile
}

// GetIdentityProfileOk returns a tuple with the IdentityProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetIdentityProfileOk() (*Reference1, bool) {
	if o == nil || o.IdentityProfile == nil {
		return nil, false
	}
	return o.IdentityProfile, true
}

// HasIdentityProfile returns a boolean if a field has been set.
func (o *Identity) HasIdentityProfile() bool {
	if o != nil && o.IdentityProfile != nil {
		return true
	}

	return false
}

// SetIdentityProfile gets a reference to the given Reference1 and assigns it to the IdentityProfile field.
func (o *Identity) SetIdentityProfile(v Reference1) {
	o.IdentityProfile = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Identity) GetSource() Reference1 {
	if o == nil || o.Source == nil {
		var ret Reference1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetSourceOk() (*Reference1, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Identity) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Reference1 and assigns it to the Source field.
func (o *Identity) SetSource(v Reference1) {
	o.Source = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Identity) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Identity) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Identity) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetProcessingState returns the ProcessingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetProcessingState() string {
	if o == nil || o.ProcessingState.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProcessingState.Get()
}

// GetProcessingStateOk returns a tuple with the ProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetProcessingStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessingState.Get(), o.ProcessingState.IsSet()
}

// HasProcessingState returns a boolean if a field has been set.
func (o *Identity) HasProcessingState() bool {
	if o != nil && o.ProcessingState.IsSet() {
		return true
	}

	return false
}

// SetProcessingState gets a reference to the given NullableString and assigns it to the ProcessingState field.
func (o *Identity) SetProcessingState(v string) {
	o.ProcessingState.Set(&v)
}
// SetProcessingStateNil sets the value for ProcessingState to be an explicit nil
func (o *Identity) SetProcessingStateNil() {
	o.ProcessingState.Set(nil)
}

// UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
func (o *Identity) UnsetProcessingState() {
	o.ProcessingState.Unset()
}

// GetProcessingDetails returns the ProcessingDetails field value if set, zero value otherwise.
func (o *Identity) GetProcessingDetails() ProcessingDetails {
	if o == nil || o.ProcessingDetails == nil {
		var ret ProcessingDetails
		return ret
	}
	return *o.ProcessingDetails
}

// GetProcessingDetailsOk returns a tuple with the ProcessingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetProcessingDetailsOk() (*ProcessingDetails, bool) {
	if o == nil || o.ProcessingDetails == nil {
		return nil, false
	}
	return o.ProcessingDetails, true
}

// HasProcessingDetails returns a boolean if a field has been set.
func (o *Identity) HasProcessingDetails() bool {
	if o != nil && o.ProcessingDetails != nil {
		return true
	}

	return false
}

// SetProcessingDetails gets a reference to the given ProcessingDetails and assigns it to the ProcessingDetails field.
func (o *Identity) SetProcessingDetails(v ProcessingDetails) {
	o.ProcessingDetails = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Identity) GetAccounts() []BaseAccount {
	if o == nil || o.Accounts == nil {
		var ret []BaseAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAccountsOk() ([]BaseAccount, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Identity) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []BaseAccount and assigns it to the Accounts field.
func (o *Identity) SetAccounts(v []BaseAccount) {
	o.Accounts = v
}

// GetAccountCount returns the AccountCount field value if set, zero value otherwise.
func (o *Identity) GetAccountCount() int32 {
	if o == nil || o.AccountCount == nil {
		var ret int32
		return ret
	}
	return *o.AccountCount
}

// GetAccountCountOk returns a tuple with the AccountCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAccountCountOk() (*int32, bool) {
	if o == nil || o.AccountCount == nil {
		return nil, false
	}
	return o.AccountCount, true
}

// HasAccountCount returns a boolean if a field has been set.
func (o *Identity) HasAccountCount() bool {
	if o != nil && o.AccountCount != nil {
		return true
	}

	return false
}

// SetAccountCount gets a reference to the given int32 and assigns it to the AccountCount field.
func (o *Identity) SetAccountCount(v int32) {
	o.AccountCount = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *Identity) GetApps() []App {
	if o == nil || o.Apps == nil {
		var ret []App
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAppsOk() ([]App, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *Identity) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []App and assigns it to the Apps field.
func (o *Identity) SetApps(v []App) {
	o.Apps = v
}

// GetAppCount returns the AppCount field value if set, zero value otherwise.
func (o *Identity) GetAppCount() int32 {
	if o == nil || o.AppCount == nil {
		var ret int32
		return ret
	}
	return *o.AppCount
}

// GetAppCountOk returns a tuple with the AppCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAppCountOk() (*int32, bool) {
	if o == nil || o.AppCount == nil {
		return nil, false
	}
	return o.AppCount, true
}

// HasAppCount returns a boolean if a field has been set.
func (o *Identity) HasAppCount() bool {
	if o != nil && o.AppCount != nil {
		return true
	}

	return false
}

// SetAppCount gets a reference to the given int32 and assigns it to the AppCount field.
func (o *Identity) SetAppCount(v int32) {
	o.AppCount = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *Identity) GetAccess() []Access1 {
	if o == nil || o.Access == nil {
		var ret []Access1
		return ret
	}
	return o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAccessOk() ([]Access1, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *Identity) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given []Access1 and assigns it to the Access field.
func (o *Identity) SetAccess(v []Access1) {
	o.Access = v
}

// GetAccessCount returns the AccessCount field value if set, zero value otherwise.
func (o *Identity) GetAccessCount() int32 {
	if o == nil || o.AccessCount == nil {
		var ret int32
		return ret
	}
	return *o.AccessCount
}

// GetAccessCountOk returns a tuple with the AccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAccessCountOk() (*int32, bool) {
	if o == nil || o.AccessCount == nil {
		return nil, false
	}
	return o.AccessCount, true
}

// HasAccessCount returns a boolean if a field has been set.
func (o *Identity) HasAccessCount() bool {
	if o != nil && o.AccessCount != nil {
		return true
	}

	return false
}

// SetAccessCount gets a reference to the given int32 and assigns it to the AccessCount field.
func (o *Identity) SetAccessCount(v int32) {
	o.AccessCount = &v
}

// GetAccessProfileCount returns the AccessProfileCount field value if set, zero value otherwise.
func (o *Identity) GetAccessProfileCount() int32 {
	if o == nil || o.AccessProfileCount == nil {
		var ret int32
		return ret
	}
	return *o.AccessProfileCount
}

// GetAccessProfileCountOk returns a tuple with the AccessProfileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetAccessProfileCountOk() (*int32, bool) {
	if o == nil || o.AccessProfileCount == nil {
		return nil, false
	}
	return o.AccessProfileCount, true
}

// HasAccessProfileCount returns a boolean if a field has been set.
func (o *Identity) HasAccessProfileCount() bool {
	if o != nil && o.AccessProfileCount != nil {
		return true
	}

	return false
}

// SetAccessProfileCount gets a reference to the given int32 and assigns it to the AccessProfileCount field.
func (o *Identity) SetAccessProfileCount(v int32) {
	o.AccessProfileCount = &v
}

// GetEntitlementCount returns the EntitlementCount field value if set, zero value otherwise.
func (o *Identity) GetEntitlementCount() int32 {
	if o == nil || o.EntitlementCount == nil {
		var ret int32
		return ret
	}
	return *o.EntitlementCount
}

// GetEntitlementCountOk returns a tuple with the EntitlementCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetEntitlementCountOk() (*int32, bool) {
	if o == nil || o.EntitlementCount == nil {
		return nil, false
	}
	return o.EntitlementCount, true
}

// HasEntitlementCount returns a boolean if a field has been set.
func (o *Identity) HasEntitlementCount() bool {
	if o != nil && o.EntitlementCount != nil {
		return true
	}

	return false
}

// SetEntitlementCount gets a reference to the given int32 and assigns it to the EntitlementCount field.
func (o *Identity) SetEntitlementCount(v int32) {
	o.EntitlementCount = &v
}

// GetRoleCount returns the RoleCount field value if set, zero value otherwise.
func (o *Identity) GetRoleCount() int32 {
	if o == nil || o.RoleCount == nil {
		var ret int32
		return ret
	}
	return *o.RoleCount
}

// GetRoleCountOk returns a tuple with the RoleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetRoleCountOk() (*int32, bool) {
	if o == nil || o.RoleCount == nil {
		return nil, false
	}
	return o.RoleCount, true
}

// HasRoleCount returns a boolean if a field has been set.
func (o *Identity) HasRoleCount() bool {
	if o != nil && o.RoleCount != nil {
		return true
	}

	return false
}

// SetRoleCount gets a reference to the given int32 and assigns it to the RoleCount field.
func (o *Identity) SetRoleCount(v int32) {
	o.RoleCount = &v
}

// GetOwns returns the Owns field value if set, zero value otherwise.
func (o *Identity) GetOwns() Owns {
	if o == nil || o.Owns == nil {
		var ret Owns
		return ret
	}
	return *o.Owns
}

// GetOwnsOk returns a tuple with the Owns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetOwnsOk() (*Owns, bool) {
	if o == nil || o.Owns == nil {
		return nil, false
	}
	return o.Owns, true
}

// HasOwns returns a boolean if a field has been set.
func (o *Identity) HasOwns() bool {
	if o != nil && o.Owns != nil {
		return true
	}

	return false
}

// SetOwns gets a reference to the given Owns and assigns it to the Owns field.
func (o *Identity) SetOwns(v Owns) {
	o.Owns = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Identity) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Identity) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Identity) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Identity) SetTags(v []string) {
	o.Tags = v
}

func (o Identity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["_type"] = o.Type
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Synced.IsSet() {
		toSerialize["synced"] = o.Synced.Get()
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Inactive != nil {
		toSerialize["inactive"] = o.Inactive
	}
	if o.Protected != nil {
		toSerialize["protected"] = o.Protected
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.EmployeeNumber != nil {
		toSerialize["employeeNumber"] = o.EmployeeNumber
	}
	if o.Manager != nil {
		toSerialize["manager"] = o.Manager
	}
	if o.IsManager != nil {
		toSerialize["isManager"] = o.IsManager
	}
	if o.IdentityProfile != nil {
		toSerialize["identityProfile"] = o.IdentityProfile
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ProcessingState.IsSet() {
		toSerialize["processingState"] = o.ProcessingState.Get()
	}
	if o.ProcessingDetails != nil {
		toSerialize["processingDetails"] = o.ProcessingDetails
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	if o.AccountCount != nil {
		toSerialize["accountCount"] = o.AccountCount
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.AppCount != nil {
		toSerialize["appCount"] = o.AppCount
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.AccessCount != nil {
		toSerialize["accessCount"] = o.AccessCount
	}
	if o.AccessProfileCount != nil {
		toSerialize["accessProfileCount"] = o.AccessProfileCount
	}
	if o.EntitlementCount != nil {
		toSerialize["entitlementCount"] = o.EntitlementCount
	}
	if o.RoleCount != nil {
		toSerialize["roleCount"] = o.RoleCount
	}
	if o.Owns != nil {
		toSerialize["owns"] = o.Owns
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Identity) UnmarshalJSON(bytes []byte) (err error) {
	varIdentity := _Identity{}

	if err = json.Unmarshal(bytes, &varIdentity); err == nil {
		*o = Identity(varIdentity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_type")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "synced")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "inactive")
		delete(additionalProperties, "protected")
		delete(additionalProperties, "status")
		delete(additionalProperties, "employeeNumber")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "isManager")
		delete(additionalProperties, "identityProfile")
		delete(additionalProperties, "source")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "processingState")
		delete(additionalProperties, "processingDetails")
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "accountCount")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "appCount")
		delete(additionalProperties, "access")
		delete(additionalProperties, "accessCount")
		delete(additionalProperties, "accessProfileCount")
		delete(additionalProperties, "entitlementCount")
		delete(additionalProperties, "roleCount")
		delete(additionalProperties, "owns")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentity struct {
	value *Identity
	isSet bool
}

func (v NullableIdentity) Get() *Identity {
	return v.value
}

func (v *NullableIdentity) Set(val *Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity(val *Identity) *NullableIdentity {
	return &NullableIdentity{value: val, isSet: true}
}

func (v NullableIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


