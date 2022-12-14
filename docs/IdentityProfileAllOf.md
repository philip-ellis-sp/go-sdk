# IdentityProfileAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Identity Profile. | [optional] 
**Owner** | Pointer to [**BaseReferenceDto**](BaseReferenceDto.md) |  | [optional] 
**Priority** | Pointer to **int64** | The priority for an Identity Profile. | [optional] 
**AuthoritativeSource** | Pointer to [**IdentityProfileAllOfAuthoritativeSource**](IdentityProfileAllOfAuthoritativeSource.md) |  | [optional] 
**IdentityRefreshRequired** | Pointer to **bool** | True if a identity refresh is needed. Typically triggered when a change on the source has been made. | [optional] 
**IdentityCount** | Pointer to **int32** | The number of identities that belong to the Identity Profile. | [optional] 
**IdentityAttributeConfig** | Pointer to [**IdentityAttributeConfig**](IdentityAttributeConfig.md) |  | [optional] 
**IdentityExceptionReportReference** | Pointer to [**IdentityExceptionReportReference**](IdentityExceptionReportReference.md) |  | [optional] 
**HasTimeBasedAttr** | Pointer to **bool** | Indicates the value of requiresPeriodicRefresh attribute for the Identity Profile. | [optional] 

## Methods

### NewIdentityProfileAllOf

`func NewIdentityProfileAllOf() *IdentityProfileAllOf`

NewIdentityProfileAllOf instantiates a new IdentityProfileAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProfileAllOfWithDefaults

`func NewIdentityProfileAllOfWithDefaults() *IdentityProfileAllOf`

NewIdentityProfileAllOfWithDefaults instantiates a new IdentityProfileAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdentityProfileAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProfileAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProfileAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProfileAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *IdentityProfileAllOf) GetOwner() BaseReferenceDto`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IdentityProfileAllOf) GetOwnerOk() (*BaseReferenceDto, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IdentityProfileAllOf) SetOwner(v BaseReferenceDto)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IdentityProfileAllOf) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPriority

`func (o *IdentityProfileAllOf) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IdentityProfileAllOf) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IdentityProfileAllOf) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IdentityProfileAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAuthoritativeSource

`func (o *IdentityProfileAllOf) GetAuthoritativeSource() IdentityProfileAllOfAuthoritativeSource`

GetAuthoritativeSource returns the AuthoritativeSource field if non-nil, zero value otherwise.

### GetAuthoritativeSourceOk

`func (o *IdentityProfileAllOf) GetAuthoritativeSourceOk() (*IdentityProfileAllOfAuthoritativeSource, bool)`

GetAuthoritativeSourceOk returns a tuple with the AuthoritativeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritativeSource

`func (o *IdentityProfileAllOf) SetAuthoritativeSource(v IdentityProfileAllOfAuthoritativeSource)`

SetAuthoritativeSource sets AuthoritativeSource field to given value.

### HasAuthoritativeSource

`func (o *IdentityProfileAllOf) HasAuthoritativeSource() bool`

HasAuthoritativeSource returns a boolean if a field has been set.

### GetIdentityRefreshRequired

`func (o *IdentityProfileAllOf) GetIdentityRefreshRequired() bool`

GetIdentityRefreshRequired returns the IdentityRefreshRequired field if non-nil, zero value otherwise.

### GetIdentityRefreshRequiredOk

`func (o *IdentityProfileAllOf) GetIdentityRefreshRequiredOk() (*bool, bool)`

GetIdentityRefreshRequiredOk returns a tuple with the IdentityRefreshRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRefreshRequired

`func (o *IdentityProfileAllOf) SetIdentityRefreshRequired(v bool)`

SetIdentityRefreshRequired sets IdentityRefreshRequired field to given value.

### HasIdentityRefreshRequired

`func (o *IdentityProfileAllOf) HasIdentityRefreshRequired() bool`

HasIdentityRefreshRequired returns a boolean if a field has been set.

### GetIdentityCount

`func (o *IdentityProfileAllOf) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *IdentityProfileAllOf) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *IdentityProfileAllOf) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *IdentityProfileAllOf) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetIdentityAttributeConfig

`func (o *IdentityProfileAllOf) GetIdentityAttributeConfig() IdentityAttributeConfig`

GetIdentityAttributeConfig returns the IdentityAttributeConfig field if non-nil, zero value otherwise.

### GetIdentityAttributeConfigOk

`func (o *IdentityProfileAllOf) GetIdentityAttributeConfigOk() (*IdentityAttributeConfig, bool)`

GetIdentityAttributeConfigOk returns a tuple with the IdentityAttributeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAttributeConfig

`func (o *IdentityProfileAllOf) SetIdentityAttributeConfig(v IdentityAttributeConfig)`

SetIdentityAttributeConfig sets IdentityAttributeConfig field to given value.

### HasIdentityAttributeConfig

`func (o *IdentityProfileAllOf) HasIdentityAttributeConfig() bool`

HasIdentityAttributeConfig returns a boolean if a field has been set.

### GetIdentityExceptionReportReference

`func (o *IdentityProfileAllOf) GetIdentityExceptionReportReference() IdentityExceptionReportReference`

GetIdentityExceptionReportReference returns the IdentityExceptionReportReference field if non-nil, zero value otherwise.

### GetIdentityExceptionReportReferenceOk

`func (o *IdentityProfileAllOf) GetIdentityExceptionReportReferenceOk() (*IdentityExceptionReportReference, bool)`

GetIdentityExceptionReportReferenceOk returns a tuple with the IdentityExceptionReportReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityExceptionReportReference

`func (o *IdentityProfileAllOf) SetIdentityExceptionReportReference(v IdentityExceptionReportReference)`

SetIdentityExceptionReportReference sets IdentityExceptionReportReference field to given value.

### HasIdentityExceptionReportReference

`func (o *IdentityProfileAllOf) HasIdentityExceptionReportReference() bool`

HasIdentityExceptionReportReference returns a boolean if a field has been set.

### GetHasTimeBasedAttr

`func (o *IdentityProfileAllOf) GetHasTimeBasedAttr() bool`

GetHasTimeBasedAttr returns the HasTimeBasedAttr field if non-nil, zero value otherwise.

### GetHasTimeBasedAttrOk

`func (o *IdentityProfileAllOf) GetHasTimeBasedAttrOk() (*bool, bool)`

GetHasTimeBasedAttrOk returns a tuple with the HasTimeBasedAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTimeBasedAttr

`func (o *IdentityProfileAllOf) SetHasTimeBasedAttr(v bool)`

SetHasTimeBasedAttr sets HasTimeBasedAttr field to given value.

### HasHasTimeBasedAttr

`func (o *IdentityProfileAllOf) HasHasTimeBasedAttr() bool`

HasHasTimeBasedAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


