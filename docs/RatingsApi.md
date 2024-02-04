# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandRatingsGet**](RatingsApi.md#ApiV2BrandRatingsGet) | **Get** /api/v2/{brand}/ratings | 
[**ApiV2RatingsGet**](RatingsApi.md#ApiV2RatingsGet) | **Get** /api/v2/ratings | 

# **ApiV2BrandRatingsGet**
> RatingsModelBrandData ApiV2BrandRatingsGet(ctx, brand)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 

### Return type

[**RatingsModelBrandData**](RatingsModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2RatingsGet**
> RatingsModelBrandListData ApiV2RatingsGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RatingsModelBrandListData**](RatingsModelBrandListData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

