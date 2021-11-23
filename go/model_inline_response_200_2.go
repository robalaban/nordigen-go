/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse2002 struct {

	Transactions InlineResponse2002Transactions `json:"transactions,omitempty"`
}

// AssertInlineResponse2002Required checks if the required fields are not zero-ed
func AssertInlineResponse2002Required(obj InlineResponse2002) error {
	if err := AssertInlineResponse2002TransactionsRequired(obj.Transactions); err != nil {
		return err
	}
	return nil
}

// AssertRecurseInlineResponse2002Required recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of InlineResponse2002 (e.g. [][]InlineResponse2002), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseInlineResponse2002Required(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aInlineResponse2002, ok := obj.(InlineResponse2002)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertInlineResponse2002Required(aInlineResponse2002)
	})
}
