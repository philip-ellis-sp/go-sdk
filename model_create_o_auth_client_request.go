/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// CreateOAuthClientRequest struct for CreateOAuthClientRequest
type CreateOAuthClientRequest struct {
	// The name of the business the API Client should belong to
	BusinessName *string `json:"businessName,omitempty"`
	// The homepage URL associated with the owner of the API Client
	HomepageUrl *string `json:"homepageUrl,omitempty"`
	// A human-readable name for the API Client
	Name string `json:"name"`
	// A description of the API Client
	Description string `json:"description"`
	// The number of seconds an access token generated for this API Client is valid for
	AccessTokenValiditySeconds int32 `json:"accessTokenValiditySeconds"`
	// The number of seconds a refresh token generated for this API Client is valid for
	RefreshTokenValiditySeconds *int32 `json:"refreshTokenValiditySeconds,omitempty"`
	// A list of the approved redirect URIs. Provide one or more URIs when assigning the AUTHORIZATION_CODE grant type to a new OAuth Client.
	RedirectUris []string `json:"redirectUris,omitempty"`
	// A list of OAuth 2.0 grant types this API Client can be used with
	GrantTypes []GrantType `json:"grantTypes"`
	AccessType AccessType `json:"accessType"`
	Type *ClientType `json:"type,omitempty"`
	// An indicator of whether the API Client can be used for requests internal within the product.
	Internal *bool `json:"internal,omitempty"`
	// An indicator of whether the API Client is enabled for use
	Enabled bool `json:"enabled"`
	// An indicator of whether the API Client supports strong authentication
	StrongAuthSupported *bool `json:"strongAuthSupported,omitempty"`
	// An indicator of whether the API Client supports the serialization of SAML claims when used with the authorization_code flow
	ClaimsSupported *bool `json:"claimsSupported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOAuthClientRequest CreateOAuthClientRequest

// NewCreateOAuthClientRequest instantiates a new CreateOAuthClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOAuthClientRequest(name string, description string, accessTokenValiditySeconds int32, grantTypes []GrantType, accessType AccessType, enabled bool) *CreateOAuthClientRequest {
	this := CreateOAuthClientRequest{}
	this.Name = name
	this.Description = description
	this.AccessTokenValiditySeconds = accessTokenValiditySeconds
	this.GrantTypes = grantTypes
	this.AccessType = accessType
	this.Enabled = enabled
	return &this
}

// NewCreateOAuthClientRequestWithDefaults instantiates a new CreateOAuthClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOAuthClientRequestWithDefaults() *CreateOAuthClientRequest {
	this := CreateOAuthClientRequest{}
	return &this
}

// GetBusinessName returns the BusinessName field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetBusinessName() string {
	if o == nil || o.BusinessName == nil {
		var ret string
		return ret
	}
	return *o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetBusinessNameOk() (*string, bool) {
	if o == nil || o.BusinessName == nil {
		return nil, false
	}
	return o.BusinessName, true
}

// HasBusinessName returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasBusinessName() bool {
	if o != nil && o.BusinessName != nil {
		return true
	}

	return false
}

// SetBusinessName gets a reference to the given string and assigns it to the BusinessName field.
func (o *CreateOAuthClientRequest) SetBusinessName(v string) {
	o.BusinessName = &v
}

// GetHomepageUrl returns the HomepageUrl field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetHomepageUrl() string {
	if o == nil || o.HomepageUrl == nil {
		var ret string
		return ret
	}
	return *o.HomepageUrl
}

// GetHomepageUrlOk returns a tuple with the HomepageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetHomepageUrlOk() (*string, bool) {
	if o == nil || o.HomepageUrl == nil {
		return nil, false
	}
	return o.HomepageUrl, true
}

// HasHomepageUrl returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasHomepageUrl() bool {
	if o != nil && o.HomepageUrl != nil {
		return true
	}

	return false
}

// SetHomepageUrl gets a reference to the given string and assigns it to the HomepageUrl field.
func (o *CreateOAuthClientRequest) SetHomepageUrl(v string) {
	o.HomepageUrl = &v
}

// GetName returns the Name field value
func (o *CreateOAuthClientRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOAuthClientRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateOAuthClientRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateOAuthClientRequest) SetDescription(v string) {
	o.Description = v
}

// GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field value
func (o *CreateOAuthClientRequest) GetAccessTokenValiditySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccessTokenValiditySeconds
}

// GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetAccessTokenValiditySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTokenValiditySeconds, true
}

// SetAccessTokenValiditySeconds sets field value
func (o *CreateOAuthClientRequest) SetAccessTokenValiditySeconds(v int32) {
	o.AccessTokenValiditySeconds = v
}

// GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetRefreshTokenValiditySeconds() int32 {
	if o == nil || o.RefreshTokenValiditySeconds == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenValiditySeconds
}

// GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetRefreshTokenValiditySecondsOk() (*int32, bool) {
	if o == nil || o.RefreshTokenValiditySeconds == nil {
		return nil, false
	}
	return o.RefreshTokenValiditySeconds, true
}

// HasRefreshTokenValiditySeconds returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasRefreshTokenValiditySeconds() bool {
	if o != nil && o.RefreshTokenValiditySeconds != nil {
		return true
	}

	return false
}

// SetRefreshTokenValiditySeconds gets a reference to the given int32 and assigns it to the RefreshTokenValiditySeconds field.
func (o *CreateOAuthClientRequest) SetRefreshTokenValiditySeconds(v int32) {
	o.RefreshTokenValiditySeconds = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *CreateOAuthClientRequest) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetGrantTypes returns the GrantTypes field value
func (o *CreateOAuthClientRequest) GetGrantTypes() []GrantType {
	if o == nil {
		var ret []GrantType
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetGrantTypesOk() ([]GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *CreateOAuthClientRequest) SetGrantTypes(v []GrantType) {
	o.GrantTypes = v
}

// GetAccessType returns the AccessType field value
func (o *CreateOAuthClientRequest) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *CreateOAuthClientRequest) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetType() ClientType {
	if o == nil || o.Type == nil {
		var ret ClientType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetTypeOk() (*ClientType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ClientType and assigns it to the Type field.
func (o *CreateOAuthClientRequest) SetType(v ClientType) {
	o.Type = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *CreateOAuthClientRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateOAuthClientRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateOAuthClientRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetStrongAuthSupported returns the StrongAuthSupported field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetStrongAuthSupported() bool {
	if o == nil || o.StrongAuthSupported == nil {
		var ret bool
		return ret
	}
	return *o.StrongAuthSupported
}

// GetStrongAuthSupportedOk returns a tuple with the StrongAuthSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetStrongAuthSupportedOk() (*bool, bool) {
	if o == nil || o.StrongAuthSupported == nil {
		return nil, false
	}
	return o.StrongAuthSupported, true
}

// HasStrongAuthSupported returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasStrongAuthSupported() bool {
	if o != nil && o.StrongAuthSupported != nil {
		return true
	}

	return false
}

// SetStrongAuthSupported gets a reference to the given bool and assigns it to the StrongAuthSupported field.
func (o *CreateOAuthClientRequest) SetStrongAuthSupported(v bool) {
	o.StrongAuthSupported = &v
}

// GetClaimsSupported returns the ClaimsSupported field value if set, zero value otherwise.
func (o *CreateOAuthClientRequest) GetClaimsSupported() bool {
	if o == nil || o.ClaimsSupported == nil {
		var ret bool
		return ret
	}
	return *o.ClaimsSupported
}

// GetClaimsSupportedOk returns a tuple with the ClaimsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOAuthClientRequest) GetClaimsSupportedOk() (*bool, bool) {
	if o == nil || o.ClaimsSupported == nil {
		return nil, false
	}
	return o.ClaimsSupported, true
}

// HasClaimsSupported returns a boolean if a field has been set.
func (o *CreateOAuthClientRequest) HasClaimsSupported() bool {
	if o != nil && o.ClaimsSupported != nil {
		return true
	}

	return false
}

// SetClaimsSupported gets a reference to the given bool and assigns it to the ClaimsSupported field.
func (o *CreateOAuthClientRequest) SetClaimsSupported(v bool) {
	o.ClaimsSupported = &v
}

func (o CreateOAuthClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusinessName != nil {
		toSerialize["businessName"] = o.BusinessName
	}
	if o.HomepageUrl != nil {
		toSerialize["homepageUrl"] = o.HomepageUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["accessTokenValiditySeconds"] = o.AccessTokenValiditySeconds
	}
	if o.RefreshTokenValiditySeconds != nil {
		toSerialize["refreshTokenValiditySeconds"] = o.RefreshTokenValiditySeconds
	}
	if o.RedirectUris != nil {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if true {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.StrongAuthSupported != nil {
		toSerialize["strongAuthSupported"] = o.StrongAuthSupported
	}
	if o.ClaimsSupported != nil {
		toSerialize["claimsSupported"] = o.ClaimsSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateOAuthClientRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateOAuthClientRequest := _CreateOAuthClientRequest{}

	if err = json.Unmarshal(bytes, &varCreateOAuthClientRequest); err == nil {
		*o = CreateOAuthClientRequest(varCreateOAuthClientRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "businessName")
		delete(additionalProperties, "homepageUrl")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "accessTokenValiditySeconds")
		delete(additionalProperties, "refreshTokenValiditySeconds")
		delete(additionalProperties, "redirectUris")
		delete(additionalProperties, "grantTypes")
		delete(additionalProperties, "accessType")
		delete(additionalProperties, "type")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "strongAuthSupported")
		delete(additionalProperties, "claimsSupported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOAuthClientRequest struct {
	value *CreateOAuthClientRequest
	isSet bool
}

func (v NullableCreateOAuthClientRequest) Get() *CreateOAuthClientRequest {
	return v.value
}

func (v *NullableCreateOAuthClientRequest) Set(val *CreateOAuthClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOAuthClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOAuthClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOAuthClientRequest(val *CreateOAuthClientRequest) *NullableCreateOAuthClientRequest {
	return &NullableCreateOAuthClientRequest{value: val, isSet: true}
}

func (v NullableCreateOAuthClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOAuthClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

