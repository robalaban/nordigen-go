/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// InstitutionsApiService is a service that implements the logic for the InstitutionsApiServicer
// This service should implement the business logic for every endpoint for the InstitutionsApi API.
// Include any external packages or services that will be required by this service.
type InstitutionsApiService struct {
}

// NewInstitutionsApiService creates a default api service
func NewInstitutionsApiService() InstitutionsApiServicer {
	return &InstitutionsApiService{}
}

// RetrieveAllSupportedInstitutionsInAGivenCountry - 
func (s *InstitutionsApiService) RetrieveAllSupportedInstitutionsInAGivenCountry(ctx context.Context, country string) (ImplResponse, error) {
	// TODO - update RetrieveAllSupportedInstitutionsInAGivenCountry with the required logic for this service method.
	// Add api_institutions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Aspsp{}) or use other options such as http.Ok ...
	//return Response(200, []Aspsp{}), nil

	//TODO: Uncomment the next line to return response Response(400, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(400, map[string]interface{}{}), nil

	//TODO: Uncomment the next line to return response Response(403, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(403, map[string]interface{}{}), nil

	//TODO: Uncomment the next line to return response Response(404, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(404, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RetrieveAllSupportedInstitutionsInAGivenCountry method not implemented")
}

// RetrieveInstitution - 
func (s *InstitutionsApiService) RetrieveInstitution(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update RetrieveInstitution with the required logic for this service method.
	// Add api_institutions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Aspsp{}) or use other options such as http.Ok ...
	//return Response(200, Aspsp{}), nil

	//TODO: Uncomment the next line to return response Response(404, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(404, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RetrieveInstitution method not implemented")
}
