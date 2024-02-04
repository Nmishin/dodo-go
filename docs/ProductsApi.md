# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandCountryIdUnitIdProductsTopGet**](ProductsApi.md#ApiV2BrandCountryIdUnitIdProductsTopGet) | **Get** /api/v2/{brand}/{countryId}/{unitId}/products/top | 
[**ApiV2BrandProductsTopGet**](ProductsApi.md#ApiV2BrandProductsTopGet) | **Get** /api/v2/{brand}/products/top | 

# **ApiV2BrandCountryIdUnitIdProductsTopGet**
> ProductStatisticsModelBrandData ApiV2BrandCountryIdUnitIdProductsTopGet(ctx, brand, countryId, unitId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
  **countryId** | **int32**|  | 
  **unitId** | **int32**|  | 
 **optional** | ***ProductsApiApiV2BrandCountryIdUnitIdProductsTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiApiV2BrandCountryIdUnitIdProductsTopGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.String**|  | 
 **end** | **optional.String**|  | 

### Return type

[**ProductStatisticsModelBrandData**](ProductStatisticsModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2BrandProductsTopGet**
> ProductStatisticsModelBrandData ApiV2BrandProductsTopGet(ctx, brand, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 
 **optional** | ***ProductsApiApiV2BrandProductsTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiApiV2BrandProductsTopGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**|  | 
 **end** | **optional.String**|  | 

### Return type

[**ProductStatisticsModelBrandData**](ProductStatisticsModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

