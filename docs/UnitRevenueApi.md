# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandRevenueUnitCountryIdUnitIdDailyYearMonthDayGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdDailyYearMonthDayGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/daily/{year}/{month}/{day} | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdHistoryGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdHistoryGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/history | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdHistoryStartYearEndYearGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdHistoryStartYearEndYearGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/history/{startYear}/{endYear} | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdMonthlyYearMonthGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdMonthlyYearMonthGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/monthly/{year}/{month} | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdMonthsLastGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdMonthsLastGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/months/last | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdTodayGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdTodayGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/today | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdWeekHistoryStartYearEndYearGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdWeekHistoryStartYearEndYearGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/week/history/{startYear}/{endYear} | 
[**ApiV2BrandRevenueUnitCountryIdUnitIdYesterdayGet**](UnitRevenueApi.md#ApiV2BrandRevenueUnitCountryIdUnitIdYesterdayGet) | **Get** /api/v2/{brand}/revenue/unit/{countryId}/{unitId}/yesterday | 

# **ApiV2BrandRevenueUnitCountryIdUnitIdDailyYearMonthDayGet**
> DailyRevenueListModel ApiV2BrandRevenueUnitCountryIdUnitIdDailyYearMonthDayGet(ctx, brand, countryId, unitId, year, month, day)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 
  **year** | **int32**|  | 
  **month** | **int32**|  | 
  **day** | **int32**|  | 

### Return type

[**DailyRevenueListModel**](DailyRevenueListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdHistoryGet**
> RevenueHistoryModelBrandData ApiV2BrandRevenueUnitCountryIdUnitIdHistoryGet(ctx, brand, countryId, unitId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 

### Return type

[**RevenueHistoryModelBrandData**](RevenueHistoryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdHistoryStartYearEndYearGet**
> RevenueHistoryModelBrandData ApiV2BrandRevenueUnitCountryIdUnitIdHistoryStartYearEndYearGet(ctx, brand, countryId, unitId, startYear, endYear)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 
  **startYear** | **int32**|  | 
  **endYear** | **int32**|  | 

### Return type

[**RevenueHistoryModelBrandData**](RevenueHistoryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdMonthlyYearMonthGet**
> PeriodicRevenueListModel ApiV2BrandRevenueUnitCountryIdUnitIdMonthlyYearMonthGet(ctx, brand, countryId, unitId, year, month)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 
  **year** | **int32**|  | 
  **month** | **int32**|  | 

### Return type

[**PeriodicRevenueListModel**](PeriodicRevenueListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdMonthsLastGet**
> MonthRevenueHistoryModelBrandData ApiV2BrandRevenueUnitCountryIdUnitIdMonthsLastGet(ctx, brand, countryId, unitId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 

### Return type

[**MonthRevenueHistoryModelBrandData**](MonthRevenueHistoryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdTodayGet**
> StatisticsModelBrandData ApiV2BrandRevenueUnitCountryIdUnitIdTodayGet(ctx, brand, countryId, unitId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 

### Return type

[**StatisticsModelBrandData**](StatisticsModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdWeekHistoryStartYearEndYearGet**
> WeekRevenueHistoryModelBrandData ApiV2BrandRevenueUnitCountryIdUnitIdWeekHistoryStartYearEndYearGet(ctx, brand, countryId, unitId, startYear, endYear)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 
  **startYear** | **int32**|  | 
  **endYear** | **int32**|  | 

### Return type

[**WeekRevenueHistoryModelBrandData**](WeekRevenueHistoryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueUnitCountryIdUnitIdYesterdayGet**
> StatisticsModelBrandData ApiV2BrandRevenueUnitCountryIdUnitIdYesterdayGet(ctx, brand, countryId, unitId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 

### Return type

[**StatisticsModelBrandData**](StatisticsModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

