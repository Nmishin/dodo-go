/*
 * Dodo Global API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dodo

type DailyMetricsModel struct {
	Day string `json:"day,omitempty"`
	Metrics []MetricsModel `json:"metrics,omitempty"`
}