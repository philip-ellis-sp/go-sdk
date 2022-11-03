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

// Source struct for Source
type Source struct {
	// the id of the Source
	Id *string `json:"id,omitempty"`
	// Human-readable description of the source
	Description *string `json:"description,omitempty"`
	Owner *BaseReferenceDto `json:"owner,omitempty"`
	Cluster *BaseReferenceDto `json:"cluster,omitempty"`
	AccountCorrelationConfig *BaseReferenceDto `json:"accountCorrelationConfig,omitempty"`
	AccountCorrelationRule *BaseReferenceDto `json:"accountCorrelationRule,omitempty"`
	ManagerCorrelationMapping *ManagerCorrelationMapping `json:"managerCorrelationMapping,omitempty"`
	ManagerCorrelationRule *BaseReferenceDto `json:"managerCorrelationRule,omitempty"`
	BeforeProvisioningRule *BaseReferenceDto `json:"beforeProvisioningRule,omitempty"`
	// List of references to Schema objects
	Schemas []BaseReferenceDto `json:"schemas,omitempty"`
	// List of references to the associated PasswordPolicy objects.
	PasswordPolicies []BaseReferenceDto `json:"passwordPolicies,omitempty"`
	// Optional features that can be supported by a source.
	Features []SourceFeature `json:"features,omitempty"`
	// Specifies the type of system being managed e.g. Active Directory, Workday, etc..
	Type *string `json:"type,omitempty"`
	// Connector script name.
	Connector *string `json:"connector,omitempty"`
	// The fully qualified name of the Java class that implements the connector interface.
	ConnectorClass *string `json:"connectorClass,omitempty"`
	// Connector specific configuration; will differ from type to type.
	ConnectorAttributes map[string]interface{} `json:"connectorAttributes,omitempty"`
	// Number from 0 to 100 that specifies when to skip the delete phase.
	DeleteThreshold *int32 `json:"deleteThreshold,omitempty"`
	// When true indicates the source is referenced by an IdentityProfile.
	Authoritative *bool `json:"authoritative,omitempty"`
	ManagementWorkgroup *BaseReferenceDto `json:"managementWorkgroup,omitempty"`
	// When true indicates a healthy source
	Healthy *bool `json:"healthy,omitempty"`
	// A status identifier, giving specific information on why a source is healthy or not
	Status *string `json:"status,omitempty"`
	// Timestamp showing when a source health check was last performed
	Since *string `json:"since,omitempty"`
	// The id of connector
	ConnectorId *string `json:"connectorId,omitempty"`
	// The name of the connector that was chosen on source creation
	ConnectorName *string `json:"connectorName,omitempty"`
	// The type of connection (direct or file)
	ConnectionType *string `json:"connectionType,omitempty"`
	// The connector implementstion id
	ConnectorImplementstionId *string `json:"connectorImplementstionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Source Source

// NewSource instantiates a new Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource() *Source {
	this := Source{}
	return &this
}

// NewSourceWithDefaults instantiates a new Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceWithDefaults() *Source {
	this := Source{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Source) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Source) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Source) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Source) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Source) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Source) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Source) GetOwner() BaseReferenceDto {
	if o == nil || o.Owner == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetOwnerOk() (*BaseReferenceDto, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Source) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BaseReferenceDto and assigns it to the Owner field.
func (o *Source) SetOwner(v BaseReferenceDto) {
	o.Owner = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *Source) GetCluster() BaseReferenceDto {
	if o == nil || o.Cluster == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetClusterOk() (*BaseReferenceDto, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *Source) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given BaseReferenceDto and assigns it to the Cluster field.
func (o *Source) SetCluster(v BaseReferenceDto) {
	o.Cluster = &v
}

// GetAccountCorrelationConfig returns the AccountCorrelationConfig field value if set, zero value otherwise.
func (o *Source) GetAccountCorrelationConfig() BaseReferenceDto {
	if o == nil || o.AccountCorrelationConfig == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.AccountCorrelationConfig
}

// GetAccountCorrelationConfigOk returns a tuple with the AccountCorrelationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetAccountCorrelationConfigOk() (*BaseReferenceDto, bool) {
	if o == nil || o.AccountCorrelationConfig == nil {
		return nil, false
	}
	return o.AccountCorrelationConfig, true
}

// HasAccountCorrelationConfig returns a boolean if a field has been set.
func (o *Source) HasAccountCorrelationConfig() bool {
	if o != nil && o.AccountCorrelationConfig != nil {
		return true
	}

	return false
}

// SetAccountCorrelationConfig gets a reference to the given BaseReferenceDto and assigns it to the AccountCorrelationConfig field.
func (o *Source) SetAccountCorrelationConfig(v BaseReferenceDto) {
	o.AccountCorrelationConfig = &v
}

// GetAccountCorrelationRule returns the AccountCorrelationRule field value if set, zero value otherwise.
func (o *Source) GetAccountCorrelationRule() BaseReferenceDto {
	if o == nil || o.AccountCorrelationRule == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.AccountCorrelationRule
}

// GetAccountCorrelationRuleOk returns a tuple with the AccountCorrelationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetAccountCorrelationRuleOk() (*BaseReferenceDto, bool) {
	if o == nil || o.AccountCorrelationRule == nil {
		return nil, false
	}
	return o.AccountCorrelationRule, true
}

// HasAccountCorrelationRule returns a boolean if a field has been set.
func (o *Source) HasAccountCorrelationRule() bool {
	if o != nil && o.AccountCorrelationRule != nil {
		return true
	}

	return false
}

// SetAccountCorrelationRule gets a reference to the given BaseReferenceDto and assigns it to the AccountCorrelationRule field.
func (o *Source) SetAccountCorrelationRule(v BaseReferenceDto) {
	o.AccountCorrelationRule = &v
}

// GetManagerCorrelationMapping returns the ManagerCorrelationMapping field value if set, zero value otherwise.
func (o *Source) GetManagerCorrelationMapping() ManagerCorrelationMapping {
	if o == nil || o.ManagerCorrelationMapping == nil {
		var ret ManagerCorrelationMapping
		return ret
	}
	return *o.ManagerCorrelationMapping
}

// GetManagerCorrelationMappingOk returns a tuple with the ManagerCorrelationMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetManagerCorrelationMappingOk() (*ManagerCorrelationMapping, bool) {
	if o == nil || o.ManagerCorrelationMapping == nil {
		return nil, false
	}
	return o.ManagerCorrelationMapping, true
}

// HasManagerCorrelationMapping returns a boolean if a field has been set.
func (o *Source) HasManagerCorrelationMapping() bool {
	if o != nil && o.ManagerCorrelationMapping != nil {
		return true
	}

	return false
}

// SetManagerCorrelationMapping gets a reference to the given ManagerCorrelationMapping and assigns it to the ManagerCorrelationMapping field.
func (o *Source) SetManagerCorrelationMapping(v ManagerCorrelationMapping) {
	o.ManagerCorrelationMapping = &v
}

// GetManagerCorrelationRule returns the ManagerCorrelationRule field value if set, zero value otherwise.
func (o *Source) GetManagerCorrelationRule() BaseReferenceDto {
	if o == nil || o.ManagerCorrelationRule == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.ManagerCorrelationRule
}

// GetManagerCorrelationRuleOk returns a tuple with the ManagerCorrelationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetManagerCorrelationRuleOk() (*BaseReferenceDto, bool) {
	if o == nil || o.ManagerCorrelationRule == nil {
		return nil, false
	}
	return o.ManagerCorrelationRule, true
}

// HasManagerCorrelationRule returns a boolean if a field has been set.
func (o *Source) HasManagerCorrelationRule() bool {
	if o != nil && o.ManagerCorrelationRule != nil {
		return true
	}

	return false
}

// SetManagerCorrelationRule gets a reference to the given BaseReferenceDto and assigns it to the ManagerCorrelationRule field.
func (o *Source) SetManagerCorrelationRule(v BaseReferenceDto) {
	o.ManagerCorrelationRule = &v
}

// GetBeforeProvisioningRule returns the BeforeProvisioningRule field value if set, zero value otherwise.
func (o *Source) GetBeforeProvisioningRule() BaseReferenceDto {
	if o == nil || o.BeforeProvisioningRule == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.BeforeProvisioningRule
}

// GetBeforeProvisioningRuleOk returns a tuple with the BeforeProvisioningRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetBeforeProvisioningRuleOk() (*BaseReferenceDto, bool) {
	if o == nil || o.BeforeProvisioningRule == nil {
		return nil, false
	}
	return o.BeforeProvisioningRule, true
}

// HasBeforeProvisioningRule returns a boolean if a field has been set.
func (o *Source) HasBeforeProvisioningRule() bool {
	if o != nil && o.BeforeProvisioningRule != nil {
		return true
	}

	return false
}

// SetBeforeProvisioningRule gets a reference to the given BaseReferenceDto and assigns it to the BeforeProvisioningRule field.
func (o *Source) SetBeforeProvisioningRule(v BaseReferenceDto) {
	o.BeforeProvisioningRule = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *Source) GetSchemas() []BaseReferenceDto {
	if o == nil || o.Schemas == nil {
		var ret []BaseReferenceDto
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetSchemasOk() ([]BaseReferenceDto, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *Source) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []BaseReferenceDto and assigns it to the Schemas field.
func (o *Source) SetSchemas(v []BaseReferenceDto) {
	o.Schemas = v
}

// GetPasswordPolicies returns the PasswordPolicies field value if set, zero value otherwise.
func (o *Source) GetPasswordPolicies() []BaseReferenceDto {
	if o == nil || o.PasswordPolicies == nil {
		var ret []BaseReferenceDto
		return ret
	}
	return o.PasswordPolicies
}

// GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetPasswordPoliciesOk() ([]BaseReferenceDto, bool) {
	if o == nil || o.PasswordPolicies == nil {
		return nil, false
	}
	return o.PasswordPolicies, true
}

// HasPasswordPolicies returns a boolean if a field has been set.
func (o *Source) HasPasswordPolicies() bool {
	if o != nil && o.PasswordPolicies != nil {
		return true
	}

	return false
}

// SetPasswordPolicies gets a reference to the given []BaseReferenceDto and assigns it to the PasswordPolicies field.
func (o *Source) SetPasswordPolicies(v []BaseReferenceDto) {
	o.PasswordPolicies = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Source) GetFeatures() []SourceFeature {
	if o == nil || o.Features == nil {
		var ret []SourceFeature
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetFeaturesOk() ([]SourceFeature, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Source) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []SourceFeature and assigns it to the Features field.
func (o *Source) SetFeatures(v []SourceFeature) {
	o.Features = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Source) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Source) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Source) SetType(v string) {
	o.Type = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *Source) GetConnector() string {
	if o == nil || o.Connector == nil {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectorOk() (*string, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *Source) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *Source) SetConnector(v string) {
	o.Connector = &v
}

// GetConnectorClass returns the ConnectorClass field value if set, zero value otherwise.
func (o *Source) GetConnectorClass() string {
	if o == nil || o.ConnectorClass == nil {
		var ret string
		return ret
	}
	return *o.ConnectorClass
}

// GetConnectorClassOk returns a tuple with the ConnectorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectorClassOk() (*string, bool) {
	if o == nil || o.ConnectorClass == nil {
		return nil, false
	}
	return o.ConnectorClass, true
}

// HasConnectorClass returns a boolean if a field has been set.
func (o *Source) HasConnectorClass() bool {
	if o != nil && o.ConnectorClass != nil {
		return true
	}

	return false
}

// SetConnectorClass gets a reference to the given string and assigns it to the ConnectorClass field.
func (o *Source) SetConnectorClass(v string) {
	o.ConnectorClass = &v
}

// GetConnectorAttributes returns the ConnectorAttributes field value if set, zero value otherwise.
func (o *Source) GetConnectorAttributes() map[string]interface{} {
	if o == nil || o.ConnectorAttributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ConnectorAttributes
}

// GetConnectorAttributesOk returns a tuple with the ConnectorAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectorAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.ConnectorAttributes == nil {
		return nil, false
	}
	return o.ConnectorAttributes, true
}

// HasConnectorAttributes returns a boolean if a field has been set.
func (o *Source) HasConnectorAttributes() bool {
	if o != nil && o.ConnectorAttributes != nil {
		return true
	}

	return false
}

// SetConnectorAttributes gets a reference to the given map[string]interface{} and assigns it to the ConnectorAttributes field.
func (o *Source) SetConnectorAttributes(v map[string]interface{}) {
	o.ConnectorAttributes = v
}

// GetDeleteThreshold returns the DeleteThreshold field value if set, zero value otherwise.
func (o *Source) GetDeleteThreshold() int32 {
	if o == nil || o.DeleteThreshold == nil {
		var ret int32
		return ret
	}
	return *o.DeleteThreshold
}

// GetDeleteThresholdOk returns a tuple with the DeleteThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetDeleteThresholdOk() (*int32, bool) {
	if o == nil || o.DeleteThreshold == nil {
		return nil, false
	}
	return o.DeleteThreshold, true
}

// HasDeleteThreshold returns a boolean if a field has been set.
func (o *Source) HasDeleteThreshold() bool {
	if o != nil && o.DeleteThreshold != nil {
		return true
	}

	return false
}

// SetDeleteThreshold gets a reference to the given int32 and assigns it to the DeleteThreshold field.
func (o *Source) SetDeleteThreshold(v int32) {
	o.DeleteThreshold = &v
}

// GetAuthoritative returns the Authoritative field value if set, zero value otherwise.
func (o *Source) GetAuthoritative() bool {
	if o == nil || o.Authoritative == nil {
		var ret bool
		return ret
	}
	return *o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetAuthoritativeOk() (*bool, bool) {
	if o == nil || o.Authoritative == nil {
		return nil, false
	}
	return o.Authoritative, true
}

// HasAuthoritative returns a boolean if a field has been set.
func (o *Source) HasAuthoritative() bool {
	if o != nil && o.Authoritative != nil {
		return true
	}

	return false
}

// SetAuthoritative gets a reference to the given bool and assigns it to the Authoritative field.
func (o *Source) SetAuthoritative(v bool) {
	o.Authoritative = &v
}

// GetManagementWorkgroup returns the ManagementWorkgroup field value if set, zero value otherwise.
func (o *Source) GetManagementWorkgroup() BaseReferenceDto {
	if o == nil || o.ManagementWorkgroup == nil {
		var ret BaseReferenceDto
		return ret
	}
	return *o.ManagementWorkgroup
}

// GetManagementWorkgroupOk returns a tuple with the ManagementWorkgroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetManagementWorkgroupOk() (*BaseReferenceDto, bool) {
	if o == nil || o.ManagementWorkgroup == nil {
		return nil, false
	}
	return o.ManagementWorkgroup, true
}

// HasManagementWorkgroup returns a boolean if a field has been set.
func (o *Source) HasManagementWorkgroup() bool {
	if o != nil && o.ManagementWorkgroup != nil {
		return true
	}

	return false
}

// SetManagementWorkgroup gets a reference to the given BaseReferenceDto and assigns it to the ManagementWorkgroup field.
func (o *Source) SetManagementWorkgroup(v BaseReferenceDto) {
	o.ManagementWorkgroup = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *Source) GetHealthy() bool {
	if o == nil || o.Healthy == nil {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetHealthyOk() (*bool, bool) {
	if o == nil || o.Healthy == nil {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *Source) HasHealthy() bool {
	if o != nil && o.Healthy != nil {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *Source) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Source) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Source) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Source) SetStatus(v string) {
	o.Status = &v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *Source) GetSince() string {
	if o == nil || o.Since == nil {
		var ret string
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetSinceOk() (*string, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *Source) HasSince() bool {
	if o != nil && o.Since != nil {
		return true
	}

	return false
}

// SetSince gets a reference to the given string and assigns it to the Since field.
func (o *Source) SetSince(v string) {
	o.Since = &v
}

// GetConnectorId returns the ConnectorId field value if set, zero value otherwise.
func (o *Source) GetConnectorId() string {
	if o == nil || o.ConnectorId == nil {
		var ret string
		return ret
	}
	return *o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectorIdOk() (*string, bool) {
	if o == nil || o.ConnectorId == nil {
		return nil, false
	}
	return o.ConnectorId, true
}

// HasConnectorId returns a boolean if a field has been set.
func (o *Source) HasConnectorId() bool {
	if o != nil && o.ConnectorId != nil {
		return true
	}

	return false
}

// SetConnectorId gets a reference to the given string and assigns it to the ConnectorId field.
func (o *Source) SetConnectorId(v string) {
	o.ConnectorId = &v
}

// GetConnectorName returns the ConnectorName field value if set, zero value otherwise.
func (o *Source) GetConnectorName() string {
	if o == nil || o.ConnectorName == nil {
		var ret string
		return ret
	}
	return *o.ConnectorName
}

// GetConnectorNameOk returns a tuple with the ConnectorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectorNameOk() (*string, bool) {
	if o == nil || o.ConnectorName == nil {
		return nil, false
	}
	return o.ConnectorName, true
}

// HasConnectorName returns a boolean if a field has been set.
func (o *Source) HasConnectorName() bool {
	if o != nil && o.ConnectorName != nil {
		return true
	}

	return false
}

// SetConnectorName gets a reference to the given string and assigns it to the ConnectorName field.
func (o *Source) SetConnectorName(v string) {
	o.ConnectorName = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *Source) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *Source) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *Source) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetConnectorImplementstionId returns the ConnectorImplementstionId field value if set, zero value otherwise.
func (o *Source) GetConnectorImplementstionId() string {
	if o == nil || o.ConnectorImplementstionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectorImplementstionId
}

// GetConnectorImplementstionIdOk returns a tuple with the ConnectorImplementstionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConnectorImplementstionIdOk() (*string, bool) {
	if o == nil || o.ConnectorImplementstionId == nil {
		return nil, false
	}
	return o.ConnectorImplementstionId, true
}

// HasConnectorImplementstionId returns a boolean if a field has been set.
func (o *Source) HasConnectorImplementstionId() bool {
	if o != nil && o.ConnectorImplementstionId != nil {
		return true
	}

	return false
}

// SetConnectorImplementstionId gets a reference to the given string and assigns it to the ConnectorImplementstionId field.
func (o *Source) SetConnectorImplementstionId(v string) {
	o.ConnectorImplementstionId = &v
}

func (o Source) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.AccountCorrelationConfig != nil {
		toSerialize["accountCorrelationConfig"] = o.AccountCorrelationConfig
	}
	if o.AccountCorrelationRule != nil {
		toSerialize["accountCorrelationRule"] = o.AccountCorrelationRule
	}
	if o.ManagerCorrelationMapping != nil {
		toSerialize["managerCorrelationMapping"] = o.ManagerCorrelationMapping
	}
	if o.ManagerCorrelationRule != nil {
		toSerialize["managerCorrelationRule"] = o.ManagerCorrelationRule
	}
	if o.BeforeProvisioningRule != nil {
		toSerialize["beforeProvisioningRule"] = o.BeforeProvisioningRule
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.PasswordPolicies != nil {
		toSerialize["passwordPolicies"] = o.PasswordPolicies
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.ConnectorClass != nil {
		toSerialize["connectorClass"] = o.ConnectorClass
	}
	if o.ConnectorAttributes != nil {
		toSerialize["connectorAttributes"] = o.ConnectorAttributes
	}
	if o.DeleteThreshold != nil {
		toSerialize["deleteThreshold"] = o.DeleteThreshold
	}
	if o.Authoritative != nil {
		toSerialize["authoritative"] = o.Authoritative
	}
	if o.ManagementWorkgroup != nil {
		toSerialize["managementWorkgroup"] = o.ManagementWorkgroup
	}
	if o.Healthy != nil {
		toSerialize["healthy"] = o.Healthy
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	if o.ConnectorId != nil {
		toSerialize["connectorId"] = o.ConnectorId
	}
	if o.ConnectorName != nil {
		toSerialize["connectorName"] = o.ConnectorName
	}
	if o.ConnectionType != nil {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if o.ConnectorImplementstionId != nil {
		toSerialize["connectorImplementstionId"] = o.ConnectorImplementstionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Source) UnmarshalJSON(bytes []byte) (err error) {
	varSource := _Source{}

	if err = json.Unmarshal(bytes, &varSource); err == nil {
		*o = Source(varSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "accountCorrelationConfig")
		delete(additionalProperties, "accountCorrelationRule")
		delete(additionalProperties, "managerCorrelationMapping")
		delete(additionalProperties, "managerCorrelationRule")
		delete(additionalProperties, "beforeProvisioningRule")
		delete(additionalProperties, "schemas")
		delete(additionalProperties, "passwordPolicies")
		delete(additionalProperties, "features")
		delete(additionalProperties, "type")
		delete(additionalProperties, "connector")
		delete(additionalProperties, "connectorClass")
		delete(additionalProperties, "connectorAttributes")
		delete(additionalProperties, "deleteThreshold")
		delete(additionalProperties, "authoritative")
		delete(additionalProperties, "managementWorkgroup")
		delete(additionalProperties, "healthy")
		delete(additionalProperties, "status")
		delete(additionalProperties, "since")
		delete(additionalProperties, "connectorId")
		delete(additionalProperties, "connectorName")
		delete(additionalProperties, "connectionType")
		delete(additionalProperties, "connectorImplementstionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

