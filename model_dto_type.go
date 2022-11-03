/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"fmt"
)

// DtoType An enumeration of the types of DTOs supported within the IdentityNow infrastructure.
type DtoType string

// List of DtoType
const (
	ACCOUNT_CORRELATION_CONFIGDtoType = "ACCOUNT_CORRELATION_CONFIG"
	ACCESS_PROFILEDtoType = "ACCESS_PROFILE"
	ACCESS_REQUEST_APPROVALDtoType = "ACCESS_REQUEST_APPROVAL"
	ACCOUNTDtoType = "ACCOUNT"
	APPLICATIONDtoType = "APPLICATION"
	CAMPAIGNDtoType = "CAMPAIGN"
	CAMPAIGN_FILTERDtoType = "CAMPAIGN_FILTER"
	CERTIFICATIONDtoType = "CERTIFICATION"
	CLUSTERDtoType = "CLUSTER"
	CONNECTOR_SCHEMADtoType = "CONNECTOR_SCHEMA"
	ENTITLEMENTDtoType = "ENTITLEMENT"
	GOVERNANCE_GROUPDtoType = "GOVERNANCE_GROUP"
	IDENTITYDtoType = "IDENTITY"
	IDENTITY_PROFILEDtoType = "IDENTITY_PROFILE"
	IDENTITY_REQUESTDtoType = "IDENTITY_REQUEST"
	LIFECYCLE_STATEDtoType = "LIFECYCLE_STATE"
	PASSWORD_POLICYDtoType = "PASSWORD_POLICY"
	ROLEDtoType = "ROLE"
	RULEDtoType = "RULE"
	SOD_POLICYDtoType = "SOD_POLICY"
	SOURCEDtoType = "SOURCE"
	TAG_CATEGORYDtoType = "TAG_CATEGORY"
	TASK_RESULTDtoType = "TASK_RESULT"
	REPORT_RESULTDtoType = "REPORT_RESULT"
	SOD_VIOLATIONDtoType = "SOD_VIOLATION"
	ACCOUNT_ACTIVITYDtoType = "ACCOUNT_ACTIVITY"
)

// All allowed values of DtoType enum
var AllowedDtoTypeEnumValues = []DtoType{
	"ACCOUNT_CORRELATION_CONFIG",
	"ACCESS_PROFILE",
	"ACCESS_REQUEST_APPROVAL",
	"ACCOUNT",
	"APPLICATION",
	"CAMPAIGN",
	"CAMPAIGN_FILTER",
	"CERTIFICATION",
	"CLUSTER",
	"CONNECTOR_SCHEMA",
	"ENTITLEMENT",
	"GOVERNANCE_GROUP",
	"IDENTITY",
	"IDENTITY_PROFILE",
	"IDENTITY_REQUEST",
	"LIFECYCLE_STATE",
	"PASSWORD_POLICY",
	"ROLE",
	"RULE",
	"SOD_POLICY",
	"SOURCE",
	"TAG_CATEGORY",
	"TASK_RESULT",
	"REPORT_RESULT",
	"SOD_VIOLATION",
	"ACCOUNT_ACTIVITY",
}

func (v *DtoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DtoType(value)
	for _, existing := range AllowedDtoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DtoType", value)
}

// NewDtoTypeFromValue returns a pointer to a valid DtoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDtoTypeFromValue(v string) (*DtoType, error) {
	ev := DtoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DtoType: valid values are %v", v, AllowedDtoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DtoType) IsValid() bool {
	for _, existing := range AllowedDtoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DtoType value
func (v DtoType) Ptr() *DtoType {
	return &v
}

type NullableDtoType struct {
	value *DtoType
	isSet bool
}

func (v NullableDtoType) Get() *DtoType {
	return v.value
}

func (v *NullableDtoType) Set(val *DtoType) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoType) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoType(val *DtoType) *NullableDtoType {
	return &NullableDtoType{value: val, isSet: true}
}

func (v NullableDtoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
