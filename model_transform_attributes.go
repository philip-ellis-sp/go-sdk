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

// TransformAttributes - Meta-data about the transform. Values in this list are specific to the type of transform to be executed.
type TransformAttributes struct {
	AccountAttribute *AccountAttribute
	Base64Decode *Base64Decode
	Base64Encode *Base64Encode
	Concatenation *Concatenation
	Conditional *Conditional
	DateCompare *DateCompare
	DateFormat *DateFormat
	DateMath *DateMath
	DecomposeDiacriticalMarks *DecomposeDiacriticalMarks
	E164phone *E164phone
	FirstValid *FirstValid
	ISO3166 *ISO3166
	IdentityAttribute *IdentityAttribute
	IndexOf *IndexOf
	LeftPad *LeftPad
	Lookup *Lookup
	Lower *Lower
	NameNormalizer *NameNormalizer
	RandomAlphaNumeric *RandomAlphaNumeric
	RandomNumeric *RandomNumeric
	Reference *Reference
	Replace *Replace
	ReplaceAll *ReplaceAll
	RightPad *RightPad
	Rule *Rule
	Split *Split
	Static *Static
	Substring *Substring
	Trim *Trim
	UUIDGenerator *UUIDGenerator
	Upper *Upper
}

// AccountAttributeAsTransformAttributes is a convenience function that returns AccountAttribute wrapped in TransformAttributes
func AccountAttributeAsTransformAttributes(v *AccountAttribute) TransformAttributes {
	return TransformAttributes{
		AccountAttribute: v,
	}
}

// Base64DecodeAsTransformAttributes is a convenience function that returns Base64Decode wrapped in TransformAttributes
func Base64DecodeAsTransformAttributes(v *Base64Decode) TransformAttributes {
	return TransformAttributes{
		Base64Decode: v,
	}
}

// Base64EncodeAsTransformAttributes is a convenience function that returns Base64Encode wrapped in TransformAttributes
func Base64EncodeAsTransformAttributes(v *Base64Encode) TransformAttributes {
	return TransformAttributes{
		Base64Encode: v,
	}
}

// ConcatenationAsTransformAttributes is a convenience function that returns Concatenation wrapped in TransformAttributes
func ConcatenationAsTransformAttributes(v *Concatenation) TransformAttributes {
	return TransformAttributes{
		Concatenation: v,
	}
}

// ConditionalAsTransformAttributes is a convenience function that returns Conditional wrapped in TransformAttributes
func ConditionalAsTransformAttributes(v *Conditional) TransformAttributes {
	return TransformAttributes{
		Conditional: v,
	}
}

// DateCompareAsTransformAttributes is a convenience function that returns DateCompare wrapped in TransformAttributes
func DateCompareAsTransformAttributes(v *DateCompare) TransformAttributes {
	return TransformAttributes{
		DateCompare: v,
	}
}

// DateFormatAsTransformAttributes is a convenience function that returns DateFormat wrapped in TransformAttributes
func DateFormatAsTransformAttributes(v *DateFormat) TransformAttributes {
	return TransformAttributes{
		DateFormat: v,
	}
}

// DateMathAsTransformAttributes is a convenience function that returns DateMath wrapped in TransformAttributes
func DateMathAsTransformAttributes(v *DateMath) TransformAttributes {
	return TransformAttributes{
		DateMath: v,
	}
}

// DecomposeDiacriticalMarksAsTransformAttributes is a convenience function that returns DecomposeDiacriticalMarks wrapped in TransformAttributes
func DecomposeDiacriticalMarksAsTransformAttributes(v *DecomposeDiacriticalMarks) TransformAttributes {
	return TransformAttributes{
		DecomposeDiacriticalMarks: v,
	}
}

// E164phoneAsTransformAttributes is a convenience function that returns E164phone wrapped in TransformAttributes
func E164phoneAsTransformAttributes(v *E164phone) TransformAttributes {
	return TransformAttributes{
		E164phone: v,
	}
}

// FirstValidAsTransformAttributes is a convenience function that returns FirstValid wrapped in TransformAttributes
func FirstValidAsTransformAttributes(v *FirstValid) TransformAttributes {
	return TransformAttributes{
		FirstValid: v,
	}
}

// ISO3166AsTransformAttributes is a convenience function that returns ISO3166 wrapped in TransformAttributes
func ISO3166AsTransformAttributes(v *ISO3166) TransformAttributes {
	return TransformAttributes{
		ISO3166: v,
	}
}

// IdentityAttributeAsTransformAttributes is a convenience function that returns IdentityAttribute wrapped in TransformAttributes
func IdentityAttributeAsTransformAttributes(v *IdentityAttribute) TransformAttributes {
	return TransformAttributes{
		IdentityAttribute: v,
	}
}

// IndexOfAsTransformAttributes is a convenience function that returns IndexOf wrapped in TransformAttributes
func IndexOfAsTransformAttributes(v *IndexOf) TransformAttributes {
	return TransformAttributes{
		IndexOf: v,
	}
}

// LeftPadAsTransformAttributes is a convenience function that returns LeftPad wrapped in TransformAttributes
func LeftPadAsTransformAttributes(v *LeftPad) TransformAttributes {
	return TransformAttributes{
		LeftPad: v,
	}
}

// LookupAsTransformAttributes is a convenience function that returns Lookup wrapped in TransformAttributes
func LookupAsTransformAttributes(v *Lookup) TransformAttributes {
	return TransformAttributes{
		Lookup: v,
	}
}

// LowerAsTransformAttributes is a convenience function that returns Lower wrapped in TransformAttributes
func LowerAsTransformAttributes(v *Lower) TransformAttributes {
	return TransformAttributes{
		Lower: v,
	}
}

// NameNormalizerAsTransformAttributes is a convenience function that returns NameNormalizer wrapped in TransformAttributes
func NameNormalizerAsTransformAttributes(v *NameNormalizer) TransformAttributes {
	return TransformAttributes{
		NameNormalizer: v,
	}
}

// RandomAlphaNumericAsTransformAttributes is a convenience function that returns RandomAlphaNumeric wrapped in TransformAttributes
func RandomAlphaNumericAsTransformAttributes(v *RandomAlphaNumeric) TransformAttributes {
	return TransformAttributes{
		RandomAlphaNumeric: v,
	}
}

// RandomNumericAsTransformAttributes is a convenience function that returns RandomNumeric wrapped in TransformAttributes
func RandomNumericAsTransformAttributes(v *RandomNumeric) TransformAttributes {
	return TransformAttributes{
		RandomNumeric: v,
	}
}

// ReferenceAsTransformAttributes is a convenience function that returns Reference wrapped in TransformAttributes
func ReferenceAsTransformAttributes(v *Reference) TransformAttributes {
	return TransformAttributes{
		Reference: v,
	}
}

// ReplaceAsTransformAttributes is a convenience function that returns Replace wrapped in TransformAttributes
func ReplaceAsTransformAttributes(v *Replace) TransformAttributes {
	return TransformAttributes{
		Replace: v,
	}
}

// ReplaceAllAsTransformAttributes is a convenience function that returns ReplaceAll wrapped in TransformAttributes
func ReplaceAllAsTransformAttributes(v *ReplaceAll) TransformAttributes {
	return TransformAttributes{
		ReplaceAll: v,
	}
}

// RightPadAsTransformAttributes is a convenience function that returns RightPad wrapped in TransformAttributes
func RightPadAsTransformAttributes(v *RightPad) TransformAttributes {
	return TransformAttributes{
		RightPad: v,
	}
}

// RuleAsTransformAttributes is a convenience function that returns Rule wrapped in TransformAttributes
func RuleAsTransformAttributes(v *Rule) TransformAttributes {
	return TransformAttributes{
		Rule: v,
	}
}

// SplitAsTransformAttributes is a convenience function that returns Split wrapped in TransformAttributes
func SplitAsTransformAttributes(v *Split) TransformAttributes {
	return TransformAttributes{
		Split: v,
	}
}

// StaticAsTransformAttributes is a convenience function that returns Static wrapped in TransformAttributes
func StaticAsTransformAttributes(v *Static) TransformAttributes {
	return TransformAttributes{
		Static: v,
	}
}

// SubstringAsTransformAttributes is a convenience function that returns Substring wrapped in TransformAttributes
func SubstringAsTransformAttributes(v *Substring) TransformAttributes {
	return TransformAttributes{
		Substring: v,
	}
}

// TrimAsTransformAttributes is a convenience function that returns Trim wrapped in TransformAttributes
func TrimAsTransformAttributes(v *Trim) TransformAttributes {
	return TransformAttributes{
		Trim: v,
	}
}

// UUIDGeneratorAsTransformAttributes is a convenience function that returns UUIDGenerator wrapped in TransformAttributes
func UUIDGeneratorAsTransformAttributes(v *UUIDGenerator) TransformAttributes {
	return TransformAttributes{
		UUIDGenerator: v,
	}
}

// UpperAsTransformAttributes is a convenience function that returns Upper wrapped in TransformAttributes
func UpperAsTransformAttributes(v *Upper) TransformAttributes {
	return TransformAttributes{
		Upper: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransformAttributes) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AccountAttribute
	err = newStrictDecoder(data).Decode(&dst.AccountAttribute)
	if err == nil {
		jsonAccountAttribute, _ := json.Marshal(dst.AccountAttribute)
		if string(jsonAccountAttribute) == "{}" { // empty struct
			dst.AccountAttribute = nil
		} else {
			match++
		}
	} else {
		dst.AccountAttribute = nil
	}

	// try to unmarshal data into Base64Decode
	err = newStrictDecoder(data).Decode(&dst.Base64Decode)
	if err == nil {
		jsonBase64Decode, _ := json.Marshal(dst.Base64Decode)
		if string(jsonBase64Decode) == "{}" { // empty struct
			dst.Base64Decode = nil
		} else {
			match++
		}
	} else {
		dst.Base64Decode = nil
	}

	// try to unmarshal data into Base64Encode
	err = newStrictDecoder(data).Decode(&dst.Base64Encode)
	if err == nil {
		jsonBase64Encode, _ := json.Marshal(dst.Base64Encode)
		if string(jsonBase64Encode) == "{}" { // empty struct
			dst.Base64Encode = nil
		} else {
			match++
		}
	} else {
		dst.Base64Encode = nil
	}

	// try to unmarshal data into Concatenation
	err = newStrictDecoder(data).Decode(&dst.Concatenation)
	if err == nil {
		jsonConcatenation, _ := json.Marshal(dst.Concatenation)
		if string(jsonConcatenation) == "{}" { // empty struct
			dst.Concatenation = nil
		} else {
			match++
		}
	} else {
		dst.Concatenation = nil
	}

	// try to unmarshal data into Conditional
	err = newStrictDecoder(data).Decode(&dst.Conditional)
	if err == nil {
		jsonConditional, _ := json.Marshal(dst.Conditional)
		if string(jsonConditional) == "{}" { // empty struct
			dst.Conditional = nil
		} else {
			match++
		}
	} else {
		dst.Conditional = nil
	}

	// try to unmarshal data into DateCompare
	err = newStrictDecoder(data).Decode(&dst.DateCompare)
	if err == nil {
		jsonDateCompare, _ := json.Marshal(dst.DateCompare)
		if string(jsonDateCompare) == "{}" { // empty struct
			dst.DateCompare = nil
		} else {
			match++
		}
	} else {
		dst.DateCompare = nil
	}

	// try to unmarshal data into DateFormat
	err = newStrictDecoder(data).Decode(&dst.DateFormat)
	if err == nil {
		jsonDateFormat, _ := json.Marshal(dst.DateFormat)
		if string(jsonDateFormat) == "{}" { // empty struct
			dst.DateFormat = nil
		} else {
			match++
		}
	} else {
		dst.DateFormat = nil
	}

	// try to unmarshal data into DateMath
	err = newStrictDecoder(data).Decode(&dst.DateMath)
	if err == nil {
		jsonDateMath, _ := json.Marshal(dst.DateMath)
		if string(jsonDateMath) == "{}" { // empty struct
			dst.DateMath = nil
		} else {
			match++
		}
	} else {
		dst.DateMath = nil
	}

	// try to unmarshal data into DecomposeDiacriticalMarks
	err = newStrictDecoder(data).Decode(&dst.DecomposeDiacriticalMarks)
	if err == nil {
		jsonDecomposeDiacriticalMarks, _ := json.Marshal(dst.DecomposeDiacriticalMarks)
		if string(jsonDecomposeDiacriticalMarks) == "{}" { // empty struct
			dst.DecomposeDiacriticalMarks = nil
		} else {
			match++
		}
	} else {
		dst.DecomposeDiacriticalMarks = nil
	}

	// try to unmarshal data into E164phone
	err = newStrictDecoder(data).Decode(&dst.E164phone)
	if err == nil {
		jsonE164phone, _ := json.Marshal(dst.E164phone)
		if string(jsonE164phone) == "{}" { // empty struct
			dst.E164phone = nil
		} else {
			match++
		}
	} else {
		dst.E164phone = nil
	}

	// try to unmarshal data into FirstValid
	err = newStrictDecoder(data).Decode(&dst.FirstValid)
	if err == nil {
		jsonFirstValid, _ := json.Marshal(dst.FirstValid)
		if string(jsonFirstValid) == "{}" { // empty struct
			dst.FirstValid = nil
		} else {
			match++
		}
	} else {
		dst.FirstValid = nil
	}

	// try to unmarshal data into ISO3166
	err = newStrictDecoder(data).Decode(&dst.ISO3166)
	if err == nil {
		jsonISO3166, _ := json.Marshal(dst.ISO3166)
		if string(jsonISO3166) == "{}" { // empty struct
			dst.ISO3166 = nil
		} else {
			match++
		}
	} else {
		dst.ISO3166 = nil
	}

	// try to unmarshal data into IdentityAttribute
	err = newStrictDecoder(data).Decode(&dst.IdentityAttribute)
	if err == nil {
		jsonIdentityAttribute, _ := json.Marshal(dst.IdentityAttribute)
		if string(jsonIdentityAttribute) == "{}" { // empty struct
			dst.IdentityAttribute = nil
		} else {
			match++
		}
	} else {
		dst.IdentityAttribute = nil
	}

	// try to unmarshal data into IndexOf
	err = newStrictDecoder(data).Decode(&dst.IndexOf)
	if err == nil {
		jsonIndexOf, _ := json.Marshal(dst.IndexOf)
		if string(jsonIndexOf) == "{}" { // empty struct
			dst.IndexOf = nil
		} else {
			match++
		}
	} else {
		dst.IndexOf = nil
	}

	// try to unmarshal data into LeftPad
	err = newStrictDecoder(data).Decode(&dst.LeftPad)
	if err == nil {
		jsonLeftPad, _ := json.Marshal(dst.LeftPad)
		if string(jsonLeftPad) == "{}" { // empty struct
			dst.LeftPad = nil
		} else {
			match++
		}
	} else {
		dst.LeftPad = nil
	}

	// try to unmarshal data into Lookup
	err = newStrictDecoder(data).Decode(&dst.Lookup)
	if err == nil {
		jsonLookup, _ := json.Marshal(dst.Lookup)
		if string(jsonLookup) == "{}" { // empty struct
			dst.Lookup = nil
		} else {
			match++
		}
	} else {
		dst.Lookup = nil
	}

	// try to unmarshal data into Lower
	err = newStrictDecoder(data).Decode(&dst.Lower)
	if err == nil {
		jsonLower, _ := json.Marshal(dst.Lower)
		if string(jsonLower) == "{}" { // empty struct
			dst.Lower = nil
		} else {
			match++
		}
	} else {
		dst.Lower = nil
	}

	// try to unmarshal data into NameNormalizer
	err = newStrictDecoder(data).Decode(&dst.NameNormalizer)
	if err == nil {
		jsonNameNormalizer, _ := json.Marshal(dst.NameNormalizer)
		if string(jsonNameNormalizer) == "{}" { // empty struct
			dst.NameNormalizer = nil
		} else {
			match++
		}
	} else {
		dst.NameNormalizer = nil
	}

	// try to unmarshal data into RandomAlphaNumeric
	err = newStrictDecoder(data).Decode(&dst.RandomAlphaNumeric)
	if err == nil {
		jsonRandomAlphaNumeric, _ := json.Marshal(dst.RandomAlphaNumeric)
		if string(jsonRandomAlphaNumeric) == "{}" { // empty struct
			dst.RandomAlphaNumeric = nil
		} else {
			match++
		}
	} else {
		dst.RandomAlphaNumeric = nil
	}

	// try to unmarshal data into RandomNumeric
	err = newStrictDecoder(data).Decode(&dst.RandomNumeric)
	if err == nil {
		jsonRandomNumeric, _ := json.Marshal(dst.RandomNumeric)
		if string(jsonRandomNumeric) == "{}" { // empty struct
			dst.RandomNumeric = nil
		} else {
			match++
		}
	} else {
		dst.RandomNumeric = nil
	}

	// try to unmarshal data into Reference
	err = newStrictDecoder(data).Decode(&dst.Reference)
	if err == nil {
		jsonReference, _ := json.Marshal(dst.Reference)
		if string(jsonReference) == "{}" { // empty struct
			dst.Reference = nil
		} else {
			match++
		}
	} else {
		dst.Reference = nil
	}

	// try to unmarshal data into Replace
	err = newStrictDecoder(data).Decode(&dst.Replace)
	if err == nil {
		jsonReplace, _ := json.Marshal(dst.Replace)
		if string(jsonReplace) == "{}" { // empty struct
			dst.Replace = nil
		} else {
			match++
		}
	} else {
		dst.Replace = nil
	}

	// try to unmarshal data into ReplaceAll
	err = newStrictDecoder(data).Decode(&dst.ReplaceAll)
	if err == nil {
		jsonReplaceAll, _ := json.Marshal(dst.ReplaceAll)
		if string(jsonReplaceAll) == "{}" { // empty struct
			dst.ReplaceAll = nil
		} else {
			match++
		}
	} else {
		dst.ReplaceAll = nil
	}

	// try to unmarshal data into RightPad
	err = newStrictDecoder(data).Decode(&dst.RightPad)
	if err == nil {
		jsonRightPad, _ := json.Marshal(dst.RightPad)
		if string(jsonRightPad) == "{}" { // empty struct
			dst.RightPad = nil
		} else {
			match++
		}
	} else {
		dst.RightPad = nil
	}

	// try to unmarshal data into Rule
	err = newStrictDecoder(data).Decode(&dst.Rule)
	if err == nil {
		jsonRule, _ := json.Marshal(dst.Rule)
		if string(jsonRule) == "{}" { // empty struct
			dst.Rule = nil
		} else {
			match++
		}
	} else {
		dst.Rule = nil
	}

	// try to unmarshal data into Split
	err = newStrictDecoder(data).Decode(&dst.Split)
	if err == nil {
		jsonSplit, _ := json.Marshal(dst.Split)
		if string(jsonSplit) == "{}" { // empty struct
			dst.Split = nil
		} else {
			match++
		}
	} else {
		dst.Split = nil
	}

	// try to unmarshal data into Static
	err = newStrictDecoder(data).Decode(&dst.Static)
	if err == nil {
		jsonStatic, _ := json.Marshal(dst.Static)
		if string(jsonStatic) == "{}" { // empty struct
			dst.Static = nil
		} else {
			match++
		}
	} else {
		dst.Static = nil
	}

	// try to unmarshal data into Substring
	err = newStrictDecoder(data).Decode(&dst.Substring)
	if err == nil {
		jsonSubstring, _ := json.Marshal(dst.Substring)
		if string(jsonSubstring) == "{}" { // empty struct
			dst.Substring = nil
		} else {
			match++
		}
	} else {
		dst.Substring = nil
	}

	// try to unmarshal data into Trim
	err = newStrictDecoder(data).Decode(&dst.Trim)
	if err == nil {
		jsonTrim, _ := json.Marshal(dst.Trim)
		if string(jsonTrim) == "{}" { // empty struct
			dst.Trim = nil
		} else {
			match++
		}
	} else {
		dst.Trim = nil
	}

	// try to unmarshal data into UUIDGenerator
	err = newStrictDecoder(data).Decode(&dst.UUIDGenerator)
	if err == nil {
		jsonUUIDGenerator, _ := json.Marshal(dst.UUIDGenerator)
		if string(jsonUUIDGenerator) == "{}" { // empty struct
			dst.UUIDGenerator = nil
		} else {
			match++
		}
	} else {
		dst.UUIDGenerator = nil
	}

	// try to unmarshal data into Upper
	err = newStrictDecoder(data).Decode(&dst.Upper)
	if err == nil {
		jsonUpper, _ := json.Marshal(dst.Upper)
		if string(jsonUpper) == "{}" { // empty struct
			dst.Upper = nil
		} else {
			match++
		}
	} else {
		dst.Upper = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AccountAttribute = nil
		dst.Base64Decode = nil
		dst.Base64Encode = nil
		dst.Concatenation = nil
		dst.Conditional = nil
		dst.DateCompare = nil
		dst.DateFormat = nil
		dst.DateMath = nil
		dst.DecomposeDiacriticalMarks = nil
		dst.E164phone = nil
		dst.FirstValid = nil
		dst.ISO3166 = nil
		dst.IdentityAttribute = nil
		dst.IndexOf = nil
		dst.LeftPad = nil
		dst.Lookup = nil
		dst.Lower = nil
		dst.NameNormalizer = nil
		dst.RandomAlphaNumeric = nil
		dst.RandomNumeric = nil
		dst.Reference = nil
		dst.Replace = nil
		dst.ReplaceAll = nil
		dst.RightPad = nil
		dst.Rule = nil
		dst.Split = nil
		dst.Static = nil
		dst.Substring = nil
		dst.Trim = nil
		dst.UUIDGenerator = nil
		dst.Upper = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TransformAttributes)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TransformAttributes)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransformAttributes) MarshalJSON() ([]byte, error) {
	if src.AccountAttribute != nil {
		return json.Marshal(&src.AccountAttribute)
	}

	if src.Base64Decode != nil {
		return json.Marshal(&src.Base64Decode)
	}

	if src.Base64Encode != nil {
		return json.Marshal(&src.Base64Encode)
	}

	if src.Concatenation != nil {
		return json.Marshal(&src.Concatenation)
	}

	if src.Conditional != nil {
		return json.Marshal(&src.Conditional)
	}

	if src.DateCompare != nil {
		return json.Marshal(&src.DateCompare)
	}

	if src.DateFormat != nil {
		return json.Marshal(&src.DateFormat)
	}

	if src.DateMath != nil {
		return json.Marshal(&src.DateMath)
	}

	if src.DecomposeDiacriticalMarks != nil {
		return json.Marshal(&src.DecomposeDiacriticalMarks)
	}

	if src.E164phone != nil {
		return json.Marshal(&src.E164phone)
	}

	if src.FirstValid != nil {
		return json.Marshal(&src.FirstValid)
	}

	if src.ISO3166 != nil {
		return json.Marshal(&src.ISO3166)
	}

	if src.IdentityAttribute != nil {
		return json.Marshal(&src.IdentityAttribute)
	}

	if src.IndexOf != nil {
		return json.Marshal(&src.IndexOf)
	}

	if src.LeftPad != nil {
		return json.Marshal(&src.LeftPad)
	}

	if src.Lookup != nil {
		return json.Marshal(&src.Lookup)
	}

	if src.Lower != nil {
		return json.Marshal(&src.Lower)
	}

	if src.NameNormalizer != nil {
		return json.Marshal(&src.NameNormalizer)
	}

	if src.RandomAlphaNumeric != nil {
		return json.Marshal(&src.RandomAlphaNumeric)
	}

	if src.RandomNumeric != nil {
		return json.Marshal(&src.RandomNumeric)
	}

	if src.Reference != nil {
		return json.Marshal(&src.Reference)
	}

	if src.Replace != nil {
		return json.Marshal(&src.Replace)
	}

	if src.ReplaceAll != nil {
		return json.Marshal(&src.ReplaceAll)
	}

	if src.RightPad != nil {
		return json.Marshal(&src.RightPad)
	}

	if src.Rule != nil {
		return json.Marshal(&src.Rule)
	}

	if src.Split != nil {
		return json.Marshal(&src.Split)
	}

	if src.Static != nil {
		return json.Marshal(&src.Static)
	}

	if src.Substring != nil {
		return json.Marshal(&src.Substring)
	}

	if src.Trim != nil {
		return json.Marshal(&src.Trim)
	}

	if src.UUIDGenerator != nil {
		return json.Marshal(&src.UUIDGenerator)
	}

	if src.Upper != nil {
		return json.Marshal(&src.Upper)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransformAttributes) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AccountAttribute != nil {
		return obj.AccountAttribute
	}

	if obj.Base64Decode != nil {
		return obj.Base64Decode
	}

	if obj.Base64Encode != nil {
		return obj.Base64Encode
	}

	if obj.Concatenation != nil {
		return obj.Concatenation
	}

	if obj.Conditional != nil {
		return obj.Conditional
	}

	if obj.DateCompare != nil {
		return obj.DateCompare
	}

	if obj.DateFormat != nil {
		return obj.DateFormat
	}

	if obj.DateMath != nil {
		return obj.DateMath
	}

	if obj.DecomposeDiacriticalMarks != nil {
		return obj.DecomposeDiacriticalMarks
	}

	if obj.E164phone != nil {
		return obj.E164phone
	}

	if obj.FirstValid != nil {
		return obj.FirstValid
	}

	if obj.ISO3166 != nil {
		return obj.ISO3166
	}

	if obj.IdentityAttribute != nil {
		return obj.IdentityAttribute
	}

	if obj.IndexOf != nil {
		return obj.IndexOf
	}

	if obj.LeftPad != nil {
		return obj.LeftPad
	}

	if obj.Lookup != nil {
		return obj.Lookup
	}

	if obj.Lower != nil {
		return obj.Lower
	}

	if obj.NameNormalizer != nil {
		return obj.NameNormalizer
	}

	if obj.RandomAlphaNumeric != nil {
		return obj.RandomAlphaNumeric
	}

	if obj.RandomNumeric != nil {
		return obj.RandomNumeric
	}

	if obj.Reference != nil {
		return obj.Reference
	}

	if obj.Replace != nil {
		return obj.Replace
	}

	if obj.ReplaceAll != nil {
		return obj.ReplaceAll
	}

	if obj.RightPad != nil {
		return obj.RightPad
	}

	if obj.Rule != nil {
		return obj.Rule
	}

	if obj.Split != nil {
		return obj.Split
	}

	if obj.Static != nil {
		return obj.Static
	}

	if obj.Substring != nil {
		return obj.Substring
	}

	if obj.Trim != nil {
		return obj.Trim
	}

	if obj.UUIDGenerator != nil {
		return obj.UUIDGenerator
	}

	if obj.Upper != nil {
		return obj.Upper
	}

	// all schemas are nil
	return nil
}

type NullableTransformAttributes struct {
	value *TransformAttributes
	isSet bool
}

func (v NullableTransformAttributes) Get() *TransformAttributes {
	return v.value
}

func (v *NullableTransformAttributes) Set(val *TransformAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformAttributes(val *TransformAttributes) *NullableTransformAttributes {
	return &NullableTransformAttributes{value: val, isSet: true}
}

func (v NullableTransformAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


