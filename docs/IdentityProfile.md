# IdentityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DtoType**](DtoType.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the object to which this reference applies | [optional] 
**Name** | Pointer to **string** | Human-readable display name of the object to which this reference applies | [optional] 
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

### NewIdentityProfile

`func NewIdentityProfile() *IdentityProfile`

NewIdentityProfile instantiates a new IdentityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProfileWithDefaults

`func NewIdentityProfileWithDefaults() *IdentityProfile`

NewIdentityProfileWithDefaults instantiates a new IdentityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdentityProfile) GetType() DtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProfile) GetTypeOk() (*DtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProfile) SetType(v DtoType)`

SetType sets Type field to given value.

### HasType

`func (o *IdentityProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *IdentityProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IdentityProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *IdentityProfile) GetOwner() BaseReferenceDto`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IdentityProfile) GetOwnerOk() (*BaseReferenceDto, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IdentityProfile) SetOwner(v BaseReferenceDto)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IdentityProfile) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPriority

`func (o *IdentityProfile) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IdentityProfile) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IdentityProfile) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IdentityProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAuthoritativeSource

`func (o *IdentityProfile) GetAuthoritativeSource() IdentityProfileAllOfAuthoritativeSource`

GetAuthoritativeSource returns the AuthoritativeSource field if non-nil, zero value otherwise.

### GetAuthoritativeSourceOk

`func (o *IdentityProfile) GetAuthoritativeSourceOk() (*IdentityProfileAllOfAuthoritativeSource, bool)`

GetAuthoritativeSourceOk returns a tuple with the AuthoritativeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoritativeSource

`func (o *IdentityProfile) SetAuthoritativeSource(v IdentityProfileAllOfAuthoritativeSource)`

SetAuthoritativeSource sets AuthoritativeSource field to given value.

### HasAuthoritativeSource

`func (o *IdentityProfile) HasAuthoritativeSource() bool`

HasAuthoritativeSource returns a boolean if a field has been set.

### GetIdentityRefreshRequired

`func (o *IdentityProfile) GetIdentityRefreshRequired() bool`

GetIdentityRefreshRequired returns the IdentityRefreshRequired field if non-nil, zero value otherwise.

### GetIdentityRefreshRequiredOk

`func (o *IdentityProfile) GetIdentityRefreshRequiredOk() (*bool, bool)`

GetIdentityRefreshRequiredOk returns a tuple with the IdentityRefreshRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRefreshRequired

`func (o *IdentityProfile) SetIdentityRefreshRequired(v bool)`

SetIdentityRefreshRequired sets IdentityRefreshRequired field to given value.

### HasIdentityRefreshRequired

`func (o *IdentityProfile) HasIdentityRefreshRequired() bool`

HasIdentityRefreshRequired returns a boolean if a field has been set.

### GetIdentityCount

`func (o *IdentityProfile) GetIdentityCount() int32`

GetIdentityCount returns the IdentityCount field if non-nil, zero value otherwise.

### GetIdentityCountOk

`func (o *IdentityProfile) GetIdentityCountOk() (*int32, bool)`

GetIdentityCountOk returns a tuple with the IdentityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCount

`func (o *IdentityProfile) SetIdentityCount(v int32)`

SetIdentityCount sets IdentityCount field to given value.

### HasIdentityCount

`func (o *IdentityProfile) HasIdentityCount() bool`

HasIdentityCount returns a boolean if a field has been set.

### GetIdentityAttributeConfig

`func (o *IdentityProfile) GetIdentityAttributeConfig() IdentityAttributeConfig`

GetIdentityAttributeConfig returns the IdentityAttributeConfig field if non-nil, zero value otherwise.

### GetIdentityAttributeConfigOk

`func (o *IdentityProfile) GetIdentityAttributeConfigOk() (*IdentityAttributeConfig, bool)`

GetIdentityAttributeConfigOk returns a tuple with the IdentityAttributeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAttributeConfig

`func (o *IdentityProfile) SetIdentityAttributeConfig(v IdentityAttributeConfig)`

SetIdentityAttributeConfig sets IdentityAttributeConfig field to given value.

### HasIdentityAttributeConfig

`func (o *IdentityProfile) HasIdentityAttributeConfig() bool`

HasIdentityAttributeConfig returns a boolean if a field has been set.

### GetIdentityExceptionReportReference

`func (o *IdentityProfile) GetIdentityExceptionReportReference() IdentityExceptionReportReference`

GetIdentityExceptionReportReference returns the IdentityExceptionReportReference field if non-nil, zero value otherwise.

### GetIdentityExceptionReportReferenceOk

`func (o *IdentityProfile) GetIdentityExceptionReportReferenceOk() (*IdentityExceptionReportReference, bool)`

GetIdentityExceptionReportReferenceOk returns a tuple with the IdentityExceptionReportReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityExceptionReportReference

`func (o *IdentityProfile) SetIdentityExceptionReportReference(v IdentityExceptionReportReference)`

SetIdentityExceptionReportReference sets IdentityExceptionReportReference field to given value.

### HasIdentityExceptionReportReference

`func (o *IdentityProfile) HasIdentityExceptionReportReference() bool`

HasIdentityExceptionReportReference returns a boolean if a field has been set.

### GetHasTimeBasedAttr

`func (o *IdentityProfile) GetHasTimeBasedAttr() bool`

GetHasTimeBasedAttr returns the HasTimeBasedAttr field if non-nil, zero value otherwise.

### GetHasTimeBasedAttrOk

`func (o *IdentityProfile) GetHasTimeBasedAttrOk() (*bool, bool)`

GetHasTimeBasedAttrOk returns a tuple with the HasTimeBasedAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTimeBasedAttr

`func (o *IdentityProfile) SetHasTimeBasedAttr(v bool)`

SetHasTimeBasedAttr sets HasTimeBasedAttr field to given value.

### HasHasTimeBasedAttr

`func (o *IdentityProfile) HasHasTimeBasedAttr() bool`

HasHasTimeBasedAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


