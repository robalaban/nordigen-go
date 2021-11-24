/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nordigen

import (
	"time"
)

// AccountV2 - AccountV2Serializer.
type AccountV2 struct {

	// The ID of this Account, used to refer to this account in other API calls.
	Id string `json:"id,omitempty"`

	// The date & time at which the account object was created.
	Created time.Time `json:"created,omitempty"`

	// The date & time at which the account object was last accessed.
	LastAccessed *time.Time `json:"last_accessed,omitempty"`

	// The Account IBAN
	Iban string `json:"iban,omitempty"`

	// The ASPSP associated with this account.
	InstitutionId string `json:"institution_id,omitempty"`

	// The processing status of this account.
	Status *AccountV2StatusEnum `json:"status,omitempty"`
}

// AssertAccountV2Required checks if the required fields are not zero-ed
func AssertAccountV2Required(obj AccountV2) error {
	if obj.Status != nil {
		if err := AssertAccountV2StatusEnumRequired(*obj.Status); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseAccountV2Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AccountV2 (e.g. [][]AccountV2), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAccountV2Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAccountV2, ok := obj.(AccountV2)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAccountV2Required(aAccountV2)
	})
}