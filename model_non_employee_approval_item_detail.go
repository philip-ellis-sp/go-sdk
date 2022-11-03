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

// NonEmployeeApprovalItemDetail struct for NonEmployeeApprovalItemDetail
type NonEmployeeApprovalItemDetail struct {
	// Non-Employee approval item id
	Id *string `json:"id,omitempty"`
	Approver *NonEmployeeIdentityReferenceWithId `json:"approver,omitempty"`
	// Requested identity account name
	AccountName *string `json:"accountName,omitempty"`
	ApprovalStatus *ApprovalStatus `json:"approvalStatus,omitempty"`
	// Approval order
	ApprovalOrder *float32 `json:"approvalOrder,omitempty"`
	// comment of approver
	Comment *string `json:"comment,omitempty"`
	// When the request was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// When the request was created.
	Created *time.Time `json:"created,omitempty"`
	NonEmployeeRequest *NonEmployeeRequestWithoutApprovalItem `json:"nonEmployeeRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeApprovalItemDetail NonEmployeeApprovalItemDetail

// NewNonEmployeeApprovalItemDetail instantiates a new NonEmployeeApprovalItemDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeApprovalItemDetail() *NonEmployeeApprovalItemDetail {
	this := NonEmployeeApprovalItemDetail{}
	return &this
}

// NewNonEmployeeApprovalItemDetailWithDefaults instantiates a new NonEmployeeApprovalItemDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeApprovalItemDetailWithDefaults() *NonEmployeeApprovalItemDetail {
	this := NonEmployeeApprovalItemDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeApprovalItemDetail) SetId(v string) {
	o.Id = &v
}

// GetApprover returns the Approver field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetApprover() NonEmployeeIdentityReferenceWithId {
	if o == nil || o.Approver == nil {
		var ret NonEmployeeIdentityReferenceWithId
		return ret
	}
	return *o.Approver
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetApproverOk() (*NonEmployeeIdentityReferenceWithId, bool) {
	if o == nil || o.Approver == nil {
		return nil, false
	}
	return o.Approver, true
}

// HasApprover returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasApprover() bool {
	if o != nil && o.Approver != nil {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NonEmployeeIdentityReferenceWithId and assigns it to the Approver field.
func (o *NonEmployeeApprovalItemDetail) SetApprover(v NonEmployeeIdentityReferenceWithId) {
	o.Approver = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *NonEmployeeApprovalItemDetail) SetAccountName(v string) {
	o.AccountName = &v
}

// GetApprovalStatus returns the ApprovalStatus field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetApprovalStatus() ApprovalStatus {
	if o == nil || o.ApprovalStatus == nil {
		var ret ApprovalStatus
		return ret
	}
	return *o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetApprovalStatusOk() (*ApprovalStatus, bool) {
	if o == nil || o.ApprovalStatus == nil {
		return nil, false
	}
	return o.ApprovalStatus, true
}

// HasApprovalStatus returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasApprovalStatus() bool {
	if o != nil && o.ApprovalStatus != nil {
		return true
	}

	return false
}

// SetApprovalStatus gets a reference to the given ApprovalStatus and assigns it to the ApprovalStatus field.
func (o *NonEmployeeApprovalItemDetail) SetApprovalStatus(v ApprovalStatus) {
	o.ApprovalStatus = &v
}

// GetApprovalOrder returns the ApprovalOrder field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetApprovalOrder() float32 {
	if o == nil || o.ApprovalOrder == nil {
		var ret float32
		return ret
	}
	return *o.ApprovalOrder
}

// GetApprovalOrderOk returns a tuple with the ApprovalOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetApprovalOrderOk() (*float32, bool) {
	if o == nil || o.ApprovalOrder == nil {
		return nil, false
	}
	return o.ApprovalOrder, true
}

// HasApprovalOrder returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasApprovalOrder() bool {
	if o != nil && o.ApprovalOrder != nil {
		return true
	}

	return false
}

// SetApprovalOrder gets a reference to the given float32 and assigns it to the ApprovalOrder field.
func (o *NonEmployeeApprovalItemDetail) SetApprovalOrder(v float32) {
	o.ApprovalOrder = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NonEmployeeApprovalItemDetail) SetComment(v string) {
	o.Comment = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *NonEmployeeApprovalItemDetail) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NonEmployeeApprovalItemDetail) SetCreated(v time.Time) {
	o.Created = &v
}

// GetNonEmployeeRequest returns the NonEmployeeRequest field value if set, zero value otherwise.
func (o *NonEmployeeApprovalItemDetail) GetNonEmployeeRequest() NonEmployeeRequestWithoutApprovalItem {
	if o == nil || o.NonEmployeeRequest == nil {
		var ret NonEmployeeRequestWithoutApprovalItem
		return ret
	}
	return *o.NonEmployeeRequest
}

// GetNonEmployeeRequestOk returns a tuple with the NonEmployeeRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeApprovalItemDetail) GetNonEmployeeRequestOk() (*NonEmployeeRequestWithoutApprovalItem, bool) {
	if o == nil || o.NonEmployeeRequest == nil {
		return nil, false
	}
	return o.NonEmployeeRequest, true
}

// HasNonEmployeeRequest returns a boolean if a field has been set.
func (o *NonEmployeeApprovalItemDetail) HasNonEmployeeRequest() bool {
	if o != nil && o.NonEmployeeRequest != nil {
		return true
	}

	return false
}

// SetNonEmployeeRequest gets a reference to the given NonEmployeeRequestWithoutApprovalItem and assigns it to the NonEmployeeRequest field.
func (o *NonEmployeeApprovalItemDetail) SetNonEmployeeRequest(v NonEmployeeRequestWithoutApprovalItem) {
	o.NonEmployeeRequest = &v
}

func (o NonEmployeeApprovalItemDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Approver != nil {
		toSerialize["approver"] = o.Approver
	}
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	if o.ApprovalStatus != nil {
		toSerialize["approvalStatus"] = o.ApprovalStatus
	}
	if o.ApprovalOrder != nil {
		toSerialize["approvalOrder"] = o.ApprovalOrder
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.NonEmployeeRequest != nil {
		toSerialize["nonEmployeeRequest"] = o.NonEmployeeRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeApprovalItemDetail) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeApprovalItemDetail := _NonEmployeeApprovalItemDetail{}

	if err = json.Unmarshal(bytes, &varNonEmployeeApprovalItemDetail); err == nil {
		*o = NonEmployeeApprovalItemDetail(varNonEmployeeApprovalItemDetail)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "approver")
		delete(additionalProperties, "accountName")
		delete(additionalProperties, "approvalStatus")
		delete(additionalProperties, "approvalOrder")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "created")
		delete(additionalProperties, "nonEmployeeRequest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeApprovalItemDetail struct {
	value *NonEmployeeApprovalItemDetail
	isSet bool
}

func (v NullableNonEmployeeApprovalItemDetail) Get() *NonEmployeeApprovalItemDetail {
	return v.value
}

func (v *NullableNonEmployeeApprovalItemDetail) Set(val *NonEmployeeApprovalItemDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeApprovalItemDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeApprovalItemDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeApprovalItemDetail(val *NonEmployeeApprovalItemDetail) *NullableNonEmployeeApprovalItemDetail {
	return &NullableNonEmployeeApprovalItemDetail{value: val, isSet: true}
}

func (v NullableNonEmployeeApprovalItemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeApprovalItemDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


