# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandCountryIdRevenueDailyGet**](RevenueApi.md#ApiV2BrandCountryIdRevenueDailyGet) | **Get** /api/v2/{brand}/{countryId}/revenue/daily | 
[**ApiV2BrandRevenueDailyYearMonthDayGet**](RevenueApi.md#ApiV2BrandRevenueDailyYearMonthDayGet) | **Get** /api/v2/{brand}/revenue/daily/{year}/{month}/{day} | 
[**ApiV2BrandRevenueMonthsLastGet**](RevenueApi.md#ApiV2BrandRevenueMonthsLastGet) | **Get** /api/v2/{brand}/revenue/months/last | 
[**ApiV2BrandRevenueMonthsYearMonthGet**](RevenueApi.md#ApiV2BrandRevenueMonthsYearMonthGet) | **Get** /api/v2/{brand}/revenue/months/{year}/{month} | 
[**ApiV2BrandRevenueWeeklyTopGet**](RevenueApi.md#ApiV2BrandRevenueWeeklyTopGet) | **Get** /api/v2/{brand}/revenue/weekly/top | 
[**ApiV2RevenueDailyYearMonthDayGet**](RevenueApi.md#ApiV2RevenueDailyYearMonthDayGet) | **Get** /api/v2/revenue/daily/{year}/{month}/{day} | 
[**ApiV2RevenueMonthsLastGet**](RevenueApi.md#ApiV2RevenueMonthsLastGet) | **Get** /api/v2/revenue/months/last | 
[**ApiV2RevenueMonthsYearMonthGet**](RevenueApi.md#ApiV2RevenueMonthsYearMonthGet) | **Get** /api/v2/revenue/months/{year}/{month} | 
[**ApiV2RevenueWeeklyTopGet**](RevenueApi.md#ApiV2RevenueWeeklyTopGet) | **Get** /api/v2/revenue/weekly/top | 

# **ApiV2BrandCountryIdRevenueDailyGet**
> CountryDailyRevenueModelBrandData ApiV2BrandCountryIdRevenueDailyGet(ctx, brand, countryId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
 **optional** | ***RevenueApiApiV2BrandCountryIdRevenueDailyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RevenueApiApiV2BrandCountryIdRevenueDailyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Time**|  | 
 **to** | **optional.Time**|  | 

### Return type

[**CountryDailyRevenueModelBrandData**](CountryDailyRevenueModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueDailyYearMonthDayGet**
> BrandDailyRevenueListModel ApiV2BrandRevenueDailyYearMonthDayGet(ctx, brand, year, month, day)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **year** | **int32**|  | 
  **month** | **int32**|  | 
  **day** | **int32**|  | 

### Return type

[**BrandDailyRevenueListModel**](BrandDailyRevenueListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueMonthsLastGet**
> MonthTotalRevenueModelBrandData ApiV2BrandRevenueMonthsLastGet(ctx, brand)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 

### Return type

[**MonthTotalRevenueModelBrandData**](MonthTotalRevenueModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueMonthsYearMonthGet**
> TotalRevenueHistoryModelBrandData ApiV2BrandRevenueMonthsYearMonthGet(ctx, brand, year, month)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **year** | **int32**|  | 
  **month** | **int32**|  | 

### Return type

[**TotalRevenueHistoryModelBrandData**](TotalRevenueHistoryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandRevenueWeeklyTopGet**
> WeeklyTopModelBrandData ApiV2BrandRevenueWeeklyTopGet(ctx, brand, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
 **optional** | ***RevenueApiApiV2BrandRevenueWeeklyTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RevenueApiApiV2BrandRevenueWeeklyTopGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int32**|  | 

### Return type

[**WeeklyTopModelBrandData**](WeeklyTopModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2RevenueDailyYearMonthDayGet**
> BrandListDailyRevenueListModel ApiV2RevenueDailyYearMonthDayGet(ctx, year, month, day)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**|  | 
  **month** | **int32**|  | 
  **day** | **int32**|  | 

### Return type

[**BrandListDailyRevenueListModel**](BrandListDailyRevenueListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2RevenueMonthsLastGet**
> MonthTotalRevenueModelBrandListData ApiV2RevenueMonthsLastGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MonthTotalRevenueModelBrandListData**](MonthTotalRevenueModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2RevenueMonthsYearMonthGet**
> TotalRevenueHistoryModelBrandListData ApiV2RevenueMonthsYearMonthGet(ctx, year, month)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**|  | 
  **month** | **int32**|  | 

### Return type

[**TotalRevenueHistoryModelBrandListData**](TotalRevenueHistoryModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2RevenueWeeklyTopGet**
> WeeklyTopModelBrandListData ApiV2RevenueWeeklyTopGet(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RevenueApiApiV2RevenueWeeklyTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RevenueApiApiV2RevenueWeeklyTopGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**|  | 

### Return type

[**WeeklyTopModelBrandListData**](WeeklyTopModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

