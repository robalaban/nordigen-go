/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RequisitionLinkParams - RequisitionLinkParamsSerializer.
type RequisitionLinkParams struct {

	// ASPSP ID. Should match one used while creating EUA, if it's linked to this requisition
	AspspId string `json:"aspsp_id"`
}

// AssertRequisitionLinkParamsRequired checks if the required fields are not zero-ed
func AssertRequisitionLinkParamsRequired(obj RequisitionLinkParams) error {
	elements := map[string]interface{}{
		"aspsp_id": obj.AspspId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseRequisitionLinkParamsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of RequisitionLinkParams (e.g. [][]RequisitionLinkParams), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseRequisitionLinkParamsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aRequisitionLinkParams, ok := obj.(RequisitionLinkParams)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertRequisitionLinkParamsRequired(aRequisitionLinkParams)
	})
}