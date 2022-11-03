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

// ApprovalStatusDto struct for ApprovalStatusDto
type ApprovalStatusDto struct {
	// True if the request for this item was forwarded from one owner to another.
	Forwarded *bool `json:"forwarded,omitempty"`
	OriginalOwner *BaseReferenceDto `json:"originalOwner,omitempty"`
	CurrentOwner *BaseReferenceDto `json:"currentOwner,omitempty"`
	ReviewedBy *BaseReferenceDto `json:"reviewedBy,omitempty"`
	// Time at which item was modified.
	Modified *time.Time `json:"modified,omitempty"`
	Status *ManualWorkItemState `json:"status,omitempty"`
	Scheme *ApprovalScheme `json:"scheme,omitempty"`
	// If the request failed, includes any error messages that were generated.
	ErrorMessages []ErrorMessageDto `json:"errorMessages,omitempty"`
	// Comment, if any, provided by the approver.
	Comment *string `json:"comment,omitempty"`
	// The date the role or access profile is no longer assigned to the specified identity.
	RemoveDate *time.Time `json:"removeDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApprovalStatusDto ApprovalStatusDto

// NewApprovalStatusDto instantiates a new ApprovalStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalStatusDto() *ApprovalStatusDto {
	this := ApprovalStatusDto{}
	return &this
}

// NewApprovalStatusDtoWithDefaults instantiates a new ApprovalStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalStatusDtoWithDefaults() *ApprovalStatusDto {
	this := ApprovalStatusDto{}
	return &this
}

// GetForwarded returns the Forwarded field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetForwarded() bool {
	if o == nil || o.Forwarded == nil {
		var ret bool
		return ret
	}
	return *o.Forwarded
}

// GetForwardedOk returns a tuple with the Forwarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetForwardedOk() (*bool, bool) {
	if o == nil || o.Forwarded == nil {
		return nil, false
	}
	return o.Forwarded, true
}

// HasForwarded returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasForwarded() bool {
	if o != nil && o.Forwarded != nil {
		return true
	}

	return false
}

// SetForwarded gets a reference to the given bool and assigns it to the Forwarded field.
func (o *ApprovalStatusDto) SetForwarded(v bool) {
	o.Forwarded = &v
}

// GetOriginalOwner returns the OriginalOwner field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetOriginalOwner() BaseReferenceDto {
	if o == nil || o.OriginalOwner == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.OriginalOwner
}

// GetOriginalOwnerOk returns a tuple with the OriginalOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetOriginalOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil || o.OriginalOwner == nil {
		return nil, false
	}
	return o.OriginalOwner, true
}

// HasOriginalOwner returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasOriginalOwner() bool {
	if o != nil && o.OriginalOwner != nil {
		return true
	}

	return false
}

// SetOriginalOwner gets a reference to the given BaseReferenceDto and assigns it to the OriginalOwner field.
func (o *ApprovalStatusDto) SetOriginalOwner(v BaseReferenceDto) {
	o.OriginalOwner = &v
}

// GetCurrentOwner returns the CurrentOwner field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetCurrentOwner() BaseReferenceDto {
	if o == nil || o.CurrentOwner == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.CurrentOwner
}

// GetCurrentOwnerOk returns a tuple with the CurrentOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetCurrentOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil || o.CurrentOwner == nil {
		return nil, false
	}
	return o.CurrentOwner, true
}

// HasCurrentOwner returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasCurrentOwner() bool {
	if o != nil && o.CurrentOwner != nil {
		return true
	}

	return false
}

// SetCurrentOwner gets a reference to the given BaseReferenceDto and assigns it to the CurrentOwner field.
func (o *ApprovalStatusDto) SetCurrentOwner(v BaseReferenceDto) {
	o.CurrentOwner = &v
}

// GetReviewedBy returns the ReviewedBy field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetReviewedBy() BaseReferenceDto {
	if o == nil || o.ReviewedBy == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.ReviewedBy
}

// GetReviewedByOk returns a tuple with the ReviewedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetReviewedByOk() (*BaseReferenceDto, bool) {
	if o == nil || o.ReviewedBy == nil {
		return nil, false
	}
	return o.ReviewedBy, true
}

// HasReviewedBy returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasReviewedBy() bool {
	if o != nil && o.ReviewedBy != nil {
		return true
	}

	return false
}

// SetReviewedBy gets a reference to the given BaseReferenceDto and assigns it to the ReviewedBy field.
func (o *ApprovalStatusDto) SetReviewedBy(v BaseReferenceDto) {
	o.ReviewedBy = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *ApprovalStatusDto) SetModified(v time.Time) {
	o.Modified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetStatus() ManualWorkItemState {
	if o == nil || o.Status == nil {
		var ret ManualWorkItemState
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetStatusOk() (*ManualWorkItemState, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ManualWorkItemState and assigns it to the Status field.
func (o *ApprovalStatusDto) SetStatus(v ManualWorkItemState) {
	o.Status = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetScheme() ApprovalScheme {
	if o == nil || o.Scheme == nil {
		var ret ApprovalScheme
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetSchemeOk() (*ApprovalScheme, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given ApprovalScheme and assigns it to the Scheme field.
func (o *ApprovalStatusDto) SetScheme(v ApprovalScheme) {
	o.Scheme = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetErrorMessages() []ErrorMessageDto {
	if o == nil || o.ErrorMessages == nil {
		var ret []ErrorMessageDto
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetErrorMessagesOk() ([]ErrorMessageDto, bool) {
	if o == nil || o.ErrorMessages == nil {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasErrorMessages() bool {
	if o != nil && o.ErrorMessages != nil {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []ErrorMessageDto and assigns it to the ErrorMessages field.
func (o *ApprovalStatusDto) SetErrorMessages(v []ErrorMessageDto) {
	o.ErrorMessages = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApprovalStatusDto) SetComment(v string) {
	o.Comment = &v
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise.
func (o *ApprovalStatusDto) GetRemoveDate() time.Time {
	if o == nil || o.RemoveDate == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoveDate
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalStatusDto) GetRemoveDateOk() (*time.Time, bool) {
	if o == nil || o.RemoveDate == nil {
		return nil, false
	}
	return o.RemoveDate, true
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *ApprovalStatusDto) HasRemoveDate() bool {
	if o != nil && o.RemoveDate != nil {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given time.Time and assigns it to the RemoveDate field.
func (o *ApprovalStatusDto) SetRemoveDate(v time.Time) {
	o.RemoveDate = &v
}

func (o ApprovalStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Forwarded != nil {
		toSerialize["forwarded"] = o.Forwarded
	}
	if o.OriginalOwner != nil {
		toSerialize["originalOwner"] = o.OriginalOwner
	}
	if o.CurrentOwner != nil {
		toSerialize["currentOwner"] = o.CurrentOwner
	}
	if o.ReviewedBy != nil {
		toSerialize["reviewedBy"] = o.ReviewedBy
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if o.ErrorMessages != nil {
		toSerialize["errorMessages"] = o.ErrorMessages
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.RemoveDate != nil {
		toSerialize["removeDate"] = o.RemoveDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApprovalStatusDto) UnmarshalJSON(bytes []byte) (err error) {
	varApprovalStatusDto := _ApprovalStatusDto{}

	if err = json.Unmarshal(bytes, &varApprovalStatusDto); err == nil {
		*o = ApprovalStatusDto(varApprovalStatusDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "forwarded")
		delete(additionalProperties, "originalOwner")
		delete(additionalProperties, "currentOwner")
		delete(additionalProperties, "reviewedBy")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "status")
		delete(additionalProperties, "scheme")
		delete(additionalProperties, "errorMessages")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "removeDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApprovalStatusDto struct {
	value *ApprovalStatusDto
	isSet bool
}

func (v NullableApprovalStatusDto) Get() *ApprovalStatusDto {
	return v.value
}

func (v *NullableApprovalStatusDto) Set(val *ApprovalStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalStatusDto(val *ApprovalStatusDto) *NullableApprovalStatusDto {
	return &NullableApprovalStatusDto{value: val, isSet: true}
}

func (v NullableApprovalStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


