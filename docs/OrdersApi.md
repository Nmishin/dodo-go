# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandOrdersCountBySourceYearMonthDayGet**](OrdersApi.md#ApiV2BrandOrdersCountBySourceYearMonthDayGet) | **Get** /api/v2/{brand}/orders/count-by-source/{year}/{month}/{day} | 
[**ApiV2OrdersCountBySourceYearMonthDayGet**](OrdersApi.md#ApiV2OrdersCountBySourceYearMonthDayGet) | **Get** /api/v2/orders/count-by-source/{year}/{month}/{day} | 

# **ApiV2BrandOrdersCountBySourceYearMonthDayGet**
> OrderModelBrandData ApiV2BrandOrdersCountBySourceYearMonthDayGet(ctx, brand, year, month, day)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **year** | **int32**|  | 
  **month** | **int32**|  | 
  **day** | **int32**|  | 

### Return type

[**OrderModelBrandData**](OrderModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2OrdersCountBySourceYearMonthDayGet**
> OrderModelBrandListData ApiV2OrdersCountBySourceYearMonthDayGet(ctx, year, month, day)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **year** | **int32**|  | 
  **month** | **int32**|  | 
  **day** | **int32**|  | 

### Return type

[**OrderModelBrandListData**](OrderModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

