# CreateOAuthClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessName** | Pointer to **string** | The name of the business the API Client should belong to | [optional] 
**HomepageUrl** | Pointer to **string** | The homepage URL associated with the owner of the API Client | [optional] 
**Name** | **string** | A human-readable name for the API Client | 
**Description** | **string** | A description of the API Client | 
**AccessTokenValiditySeconds** | **int32** | The number of seconds an access token generated for this API Client is valid for | 
**RefreshTokenValiditySeconds** | Pointer to **int32** | The number of seconds a refresh token generated for this API Client is valid for | [optional] 
**RedirectUris** | Pointer to **[]string** | A list of the approved redirect URIs. Provide one or more URIs when assigning the AUTHORIZATION_CODE grant type to a new OAuth Client. | [optional] 
**GrantTypes** | [**[]GrantType**](GrantType.md) | A list of OAuth 2.0 grant types this API Client can be used with | 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**Type** | Pointer to [**ClientType**](ClientType.md) |  | [optional] 
**Internal** | Pointer to **bool** | An indicator of whether the API Client can be used for requests internal within the product. | [optional] 
**Enabled** | **bool** | An indicator of whether the API Client is enabled for use | 
**StrongAuthSupported** | Pointer to **bool** | An indicator of whether the API Client supports strong authentication | [optional] 
**ClaimsSupported** | Pointer to **bool** | An indicator of whether the API Client supports the serialization of SAML claims when used with the authorization_code flow | [optional] 

## Methods

### NewCreateOAuthClientRequest

`func NewCreateOAuthClientRequest(name string, description string, accessTokenValiditySeconds int32, grantTypes []GrantType, accessType AccessType, enabled bool, ) *CreateOAuthClientRequest`

NewCreateOAuthClientRequest instantiates a new CreateOAuthClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOAuthClientRequestWithDefaults

`func NewCreateOAuthClientRequestWithDefaults() *CreateOAuthClientRequest`

NewCreateOAuthClientRequestWithDefaults instantiates a new CreateOAuthClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessName

`func (o *CreateOAuthClientRequest) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *CreateOAuthClientRequest) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *CreateOAuthClientRequest) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *CreateOAuthClientRequest) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetHomepageUrl

`func (o *CreateOAuthClientRequest) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *CreateOAuthClientRequest) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *CreateOAuthClientRequest) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *CreateOAuthClientRequest) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### GetName

`func (o *CreateOAuthClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOAuthClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOAuthClientRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOAuthClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOAuthClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOAuthClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAccessTokenValiditySeconds

`func (o *CreateOAuthClientRequest) GetAccessTokenValiditySeconds() int32`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *CreateOAuthClientRequest) GetAccessTokenValiditySecondsOk() (*int32, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *CreateOAuthClientRequest) SetAccessTokenValiditySeconds(v int32)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.


### GetRefreshTokenValiditySeconds

`func (o *CreateOAuthClientRequest) GetRefreshTokenValiditySeconds() int32`

GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field if non-nil, zero value otherwise.

### GetRefreshTokenValiditySecondsOk

`func (o *CreateOAuthClientRequest) GetRefreshTokenValiditySecondsOk() (*int32, bool)`

GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValiditySeconds

`func (o *CreateOAuthClientRequest) SetRefreshTokenValiditySeconds(v int32)`

SetRefreshTokenValiditySeconds sets RefreshTokenValiditySeconds field to given value.

### HasRefreshTokenValiditySeconds

`func (o *CreateOAuthClientRequest) HasRefreshTokenValiditySeconds() bool`

HasRefreshTokenValiditySeconds returns a boolean if a field has been set.

### GetRedirectUris

`func (o *CreateOAuthClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *CreateOAuthClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *CreateOAuthClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *CreateOAuthClientRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetGrantTypes

`func (o *CreateOAuthClientRequest) GetGrantTypes() []GrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *CreateOAuthClientRequest) GetGrantTypesOk() (*[]GrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *CreateOAuthClientRequest) SetGrantTypes(v []GrantType)`

SetGrantTypes sets GrantTypes field to given value.


### GetAccessType

`func (o *CreateOAuthClientRequest) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreateOAuthClientRequest) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreateOAuthClientRequest) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetType

`func (o *CreateOAuthClientRequest) GetType() ClientType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOAuthClientRequest) GetTypeOk() (*ClientType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOAuthClientRequest) SetType(v ClientType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateOAuthClientRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInternal

`func (o *CreateOAuthClientRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CreateOAuthClientRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CreateOAuthClientRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *CreateOAuthClientRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateOAuthClientRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateOAuthClientRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateOAuthClientRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetStrongAuthSupported

`func (o *CreateOAuthClientRequest) GetStrongAuthSupported() bool`

GetStrongAuthSupported returns the StrongAuthSupported field if non-nil, zero value otherwise.

### GetStrongAuthSupportedOk

`func (o *CreateOAuthClientRequest) GetStrongAuthSupportedOk() (*bool, bool)`

GetStrongAuthSupportedOk returns a tuple with the StrongAuthSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongAuthSupported

`func (o *CreateOAuthClientRequest) SetStrongAuthSupported(v bool)`

SetStrongAuthSupported sets StrongAuthSupported field to given value.

### HasStrongAuthSupported

`func (o *CreateOAuthClientRequest) HasStrongAuthSupported() bool`

HasStrongAuthSupported returns a boolean if a field has been set.

### GetClaimsSupported

`func (o *CreateOAuthClientRequest) GetClaimsSupported() bool`

GetClaimsSupported returns the ClaimsSupported field if non-nil, zero value otherwise.

### GetClaimsSupportedOk

`func (o *CreateOAuthClientRequest) GetClaimsSupportedOk() (*bool, bool)`

GetClaimsSupportedOk returns a tuple with the ClaimsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsSupported

`func (o *CreateOAuthClientRequest) SetClaimsSupported(v bool)`

SetClaimsSupported sets ClaimsSupported field to given value.

### HasClaimsSupported

`func (o *CreateOAuthClientRequest) HasClaimsSupported() bool`

HasClaimsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


