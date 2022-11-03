# CreatePersonalAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the personal access token (PAT) to be created. Cannot be the same as another PAT owned by the user for whom this PAT is being created. | 

## Methods

### NewCreatePersonalAccessTokenRequest

`func NewCreatePersonalAccessTokenRequest(name string, ) *CreatePersonalAccessTokenRequest`

NewCreatePersonalAccessTokenRequest instantiates a new CreatePersonalAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePersonalAccessTokenRequestWithDefaults

`func NewCreatePersonalAccessTokenRequestWithDefaults() *CreatePersonalAccessTokenRequest`

NewCreatePersonalAccessTokenRequestWithDefaults instantiates a new CreatePersonalAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePersonalAccessTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePersonalAccessTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePersonalAccessTokenRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


