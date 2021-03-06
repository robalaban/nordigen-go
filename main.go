/*
 * Nordigen Account Information Services API
 *
 * No description provided (generated by Openapi Generator https://github.com/nordigentools/openapi-generator)
 *
 * API version: 2.0 (v2)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	"github.com/robalaban/nordigen-go/nordigen"
)

func main() {
	log.Printf("Server started")

	AccountsApiService := nordigen.NewAccountsApiService()
	AccountsApiController := nordigen.NewAccountsApiController(AccountsApiService)

	AgreementsApiService := nordigen.NewAgreementsApiService()
	AgreementsApiController := nordigen.NewAgreementsApiController(AgreementsApiService)

	InstitutionsApiService := nordigen.NewInstitutionsApiService()
	InstitutionsApiController := nordigen.NewInstitutionsApiController(InstitutionsApiService)

	RequisitionsApiService := nordigen.NewRequisitionsApiService()
	RequisitionsApiController := nordigen.NewRequisitionsApiController(RequisitionsApiService)

	TokenApiService := nordigen.NewTokenApiService()
	TokenApiController := nordigen.NewTokenApiController(TokenApiService)

	router := nordigen.NewRouter(AccountsApiController, AgreementsApiController, InstitutionsApiController, RequisitionsApiController, TokenApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
