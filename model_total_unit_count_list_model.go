/*
 * Dodo Global API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dodo

type TotalUnitCountListModel struct {
	Countries []UnitCountModel `json:"countries,omitempty"`
	Errors []ErrorModel `json:"errors,omitempty"`
	Total int32 `json:"total,omitempty"`
}