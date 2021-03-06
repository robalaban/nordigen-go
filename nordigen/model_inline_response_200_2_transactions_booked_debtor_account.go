/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nordigen

type InlineResponse2002TransactionsBookedDebtorAccount struct {

	Bban string `json:"bban,omitempty"`
}

// AssertInlineResponse2002TransactionsBookedDebtorAccountRequired checks if the required fields are not zero-ed
func AssertInlineResponse2002TransactionsBookedDebtorAccountRequired(obj InlineResponse2002TransactionsBookedDebtorAccount) error {
	return nil
}

// AssertRecurseInlineResponse2002TransactionsBookedDebtorAccountRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InlineResponse2002TransactionsBookedDebtorAccount (e.g. [][]InlineResponse2002TransactionsBookedDebtorAccount), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineResponse2002TransactionsBookedDebtorAccountRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineResponse2002TransactionsBookedDebtorAccount, ok := obj.(InlineResponse2002TransactionsBookedDebtorAccount)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineResponse2002TransactionsBookedDebtorAccountRequired(aInlineResponse2002TransactionsBookedDebtorAccount)
	})
}
