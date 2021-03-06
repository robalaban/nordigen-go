/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nordigen

// JwtRefresh - Refresh access token.
type JwtRefresh struct {

	Refresh string `json:"refresh"`

	Access string `json:"access,omitempty"`
}

// AssertJwtRefreshRequired checks if the required fields are not zero-ed
func AssertJwtRefreshRequired(obj JwtRefresh) error {
	elements := map[string]interface{}{
		"refresh": obj.Refresh,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseJwtRefreshRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of JwtRefresh (e.g. [][]JwtRefresh), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseJwtRefreshRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aJwtRefresh, ok := obj.(JwtRefresh)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertJwtRefreshRequired(aJwtRefresh)
	})
}
