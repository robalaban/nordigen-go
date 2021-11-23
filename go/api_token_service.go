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

// TokenApiService is a service that implements the logic for the TokenApiServicer
// This service should implement the business logic for every endpoint for the TokenApi API.
// Include any external packages or services that will be required by this service.
type TokenApiService struct {
}

// NewTokenApiService creates a default api service
func NewTokenApiService() TokenApiServicer {
	return &TokenApiService{}
}

// JWTObtain - 
func (s *TokenApiService) JWTObtain(ctx context.Context, jwtObtainPair JwtObtainPair) (ImplResponse, error) {
	// TODO - update JWTObtain with the required logic for this service method.
	// Add api_token_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, SpectacularJwtObtain{}) or use other options such as http.Ok ...
	//return Response(200, SpectacularJwtObtain{}), nil

	//TODO: Uncomment the next line to return response Response(401, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(401, map[string]interface{}{}), nil

	//TODO: Uncomment the next line to return response Response(403, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(403, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("JWTObtain method not implemented")
}

// JWTRefresh - 
func (s *TokenApiService) JWTRefresh(ctx context.Context, jwtRefresh JwtRefresh) (ImplResponse, error) {
	// TODO - update JWTRefresh with the required logic for this service method.
	// Add api_token_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, SpectacularJwtRefresh{}) or use other options such as http.Ok ...
	//return Response(200, SpectacularJwtRefresh{}), nil

	//TODO: Uncomment the next line to return response Response(401, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(401, map[string]interface{}{}), nil

	//TODO: Uncomment the next line to return response Response(403, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(403, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("JWTRefresh method not implemented")
}
