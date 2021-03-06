/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nordigen

type Status6e6Enum string

// List of Status6e6Enum
const (
	CRLONGCREATEDDESCRIPTIONREQUISITION_HAS_BEEN_SUCCESSFULLY_CREATED                  Status6e6Enum = "{\"short\":\"CR\",\"long\":\"CREATED\",\"description\":\"Requisition has been successfully created\"}"
	LNLONGLINKEDDESCRIPTIONACCOUNT_HAS_BEEN_SUCCESSFULLY_LINKED_TO_REQUISITION         Status6e6Enum = "{\"short\":\"LN\",\"long\":\"LINKED\",\"description\":\"Account has been successfully linked to requisition\"}"
	EXLONGEXPIREDDESCRIPTIONACCESS_TO_ACCOUNT_HAS_EXPIRED_AS_SET_IN_END_USER_AGREEMENT Status6e6Enum = "{\"short\":\"EX\",\"long\":\"EXPIRED\",\"description\":\"Access to account has expired as set in End User Agreement\"}"
	RJLONGREJECTEDDESCRIPTIONSSN_VERIFICATION_HAS_FAILED                               Status6e6Enum = "{\"short\":\"RJ\",\"long\":\"REJECTED\",\"description\":\"SSN verification has failed\"}"
)

// AssertStatus6e6EnumRequired checks if the required fields are not zero-ed
func AssertStatus6e6EnumRequired(obj Status6e6Enum) error {
	return nil
}

// AssertRecurseStatus6e6EnumRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Status6e6Enum (e.g. [][]Status6e6Enum), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseStatus6e6EnumRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aStatus6e6Enum, ok := obj.(Status6e6Enum)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertStatus6e6EnumRequired(aStatus6e6Enum)
	})
}
