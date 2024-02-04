# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2BrandCountriesGet**](CountriesApi.md#ApiV2BrandCountriesGet) | **Get** /api/v2/{brand}/countries | 
[**ApiV2CountriesGet**](CountriesApi.md#ApiV2CountriesGet) | **Get** /api/v2/countries | 

# **ApiV2BrandCountriesGet**
> []CountryModel ApiV2BrandCountriesGet(ctx, brand)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **brand** | **string**|  | 

### Return type

[**[]CountryModel**](CountryModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV2CountriesGet**
> []CountryModelBrandData ApiV2CountriesGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CountryModelBrandData**](CountryModelBrandData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

