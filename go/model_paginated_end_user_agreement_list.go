/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PaginatedEndUserAgreementList struct {

	Count int32 `json:"count,omitempty"`

	Next *string `json:"next,omitempty"`

	Previous *string `json:"previous,omitempty"`

	Results []EndUserAgreement `json:"results,omitempty"`
}

// AssertPaginatedEndUserAgreementListRequired checks if the required fields are not zero-ed
func AssertPaginatedEndUserAgreementListRequired(obj PaginatedEndUserAgreementList) error {
	for _, el := range obj.Results {
		if err := AssertEndUserAgreementRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecursePaginatedEndUserAgreementListRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PaginatedEndUserAgreementList (e.g. [][]PaginatedEndUserAgreementList), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePaginatedEndUserAgreementListRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPaginatedEndUserAgreementList, ok := obj.(PaginatedEndUserAgreementList)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPaginatedEndUserAgreementListRequired(aPaginatedEndUserAgreementList)
	})
}