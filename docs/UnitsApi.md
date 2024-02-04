# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandUnitsAllCountryIdGet**](UnitsApi.md#ApiV2BrandUnitsAllCountryIdGet) | **Get** /api/v2/{brand}/units/all/{countryId} | 
[**ApiV2BrandUnitsCountGet**](UnitsApi.md#ApiV2BrandUnitsCountGet) | **Get** /api/v2/{brand}/units/count | 
[**ApiV2BrandUnitsCountMonthsYearMonthGet**](UnitsApi.md#ApiV2BrandUnitsCountMonthsYearMonthGet) | **Get** /api/v2/{brand}/units/count/months/{year}/{month} | 
[**ApiV2BrandUnitsCountryIdUnitIdGet**](UnitsApi.md#ApiV2BrandUnitsCountryIdUnitIdGet) | **Get** /api/v2/{brand}/units/{countryId}/{unitId} | 
[**ApiV2UnitsAllCountryIdGet**](UnitsApi.md#ApiV2UnitsAllCountryIdGet) | **Get** /api/v2/units/all/{countryId} | 
[**ApiV2UnitsCountGet**](UnitsApi.md#ApiV2UnitsCountGet) | **Get** /api/v2/units/count | 
[**ApiV2UnitsCountMonthsYearMonthGet**](UnitsApi.md#ApiV2UnitsCountMonthsYearMonthGet) | **Get** /api/v2/units/count/months/{year}/{month} | 

# **ApiV2BrandUnitsAllCountryIdGet**
> UnitListModelBrandData ApiV2BrandUnitsAllCountryIdGet(ctx, brand, countryId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 

### Return type

[**UnitListModelBrandData**](UnitListModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandUnitsCountGet**
> TotalUnitCountListModel ApiV2BrandUnitsCountGet(ctx, brand)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 

### Return type

[**TotalUnitCountListModel**](TotalUnitCountListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandUnitsCountMonthsYearMonthGet**
> MonthUnitCountHistoryModelBrandData ApiV2BrandUnitsCountMonthsYearMonthGet(ctx, brand, year, month)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **year** | **int32**|  | 
  **month** | **int32**|  | 

### Return type

[**MonthUnitCountHistoryModelBrandData**](MonthUnitCountHistoryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandUnitsCountryIdUnitIdGet**
> UnitSingleModelBrandData ApiV2BrandUnitsCountryIdUnitIdGet(ctx, brand, countryId, unitId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 

### Return type

[**UnitSingleModelBrandData**](UnitSingleModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2UnitsAllCountryIdGet**
> UnitListModelBrandListData ApiV2UnitsAllCountryIdGet(ctx, countryId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryId** | **int32**|  | 

### Return type

[**UnitListModelBrandListData**](UnitListModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2UnitsCountGet**
> BrandListTotalUnitCountListModel ApiV2UnitsCountGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BrandListTotalUnitCountListModel**](BrandListTotalUnitCountListModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2UnitsCountMonthsYearMonthGet**
> MonthUnitCountHistoryModelBrandListData ApiV2UnitsCountMonthsYearMonthGet(ctx, year, month)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**|  | 
  **month** | **int32**|  | 

### Return type

[**MonthUnitCountHistoryModelBrandListData**](MonthUnitCountHistoryModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

