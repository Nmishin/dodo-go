/*
 * Dodo Global API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dodo

type UnitCountModel struct {
	CountryId int32 `json:"countryId,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	CountryName string `json:"countryName,omitempty"`
	PizzeriaCount int32 `json:"pizzeriaCount,omitempty"`
}
