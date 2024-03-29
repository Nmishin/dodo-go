/*
 * Dodo Global API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dodo

type RevenueModel struct {
	Year int32 `json:"year,omitempty"`
	Month int32 `json:"month,omitempty"`
	Revenue float64 `json:"revenue,omitempty"`
	StationaryRevenue float64 `json:"stationaryRevenue,omitempty"`
	DeliveryRevenue float64 `json:"deliveryRevenue,omitempty"`
	PickupRevenue float64 `json:"pickupRevenue,omitempty"`
}
