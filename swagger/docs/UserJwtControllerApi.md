# \UserJwtControllerApi

All URIs are relative to *https://localhost:7080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeUsingPOST**](UserJwtControllerApi.md#AuthorizeUsingPOST) | **Post** /api/authenticate | authorize


# **AuthorizeUsingPOST**
> ResponseEntity AuthorizeUsingPOST(ctx, loginVM)
authorize

authorize

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **loginVM** | [**LoginVm**](LoginVm.md)| loginVM | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

