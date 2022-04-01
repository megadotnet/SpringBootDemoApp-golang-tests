# \UserAccountControllerApi

All URIs are relative to *https://localhost:7080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateAccountUsingGET**](UserAccountControllerApi.md#ActivateAccountUsingGET) | **Get** /api/activate | activateAccount
[**ChangePasswordUsingPOST**](UserAccountControllerApi.md#ChangePasswordUsingPOST) | **Post** /api/account/change_password | changePassword
[**FinishPasswordResetUsingPOST**](UserAccountControllerApi.md#FinishPasswordResetUsingPOST) | **Post** /api/account/reset_password/finish | finishPasswordReset
[**GetAccountUsingGET**](UserAccountControllerApi.md#GetAccountUsingGET) | **Get** /api/account | getAccount
[**IsAuthenticatedUsingGET**](UserAccountControllerApi.md#IsAuthenticatedUsingGET) | **Get** /api/authenticate | isAuthenticated
[**RegisterAccountUsingPOST**](UserAccountControllerApi.md#RegisterAccountUsingPOST) | **Post** /api/register | registerAccount
[**RequestPasswordResetUsingPOST**](UserAccountControllerApi.md#RequestPasswordResetUsingPOST) | **Post** /api/account/reset_password/init | requestPasswordReset
[**SaveAccountUsingPOST**](UserAccountControllerApi.md#SaveAccountUsingPOST) | **Post** /api/account | saveAccount


# **ActivateAccountUsingGET**
> string ActivateAccountUsingGET(ctx, key)
activateAccount

activate Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| key | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePasswordUsingPOST**
> ResponseEntity ChangePasswordUsingPOST(ctx, password)
changePassword

change Password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **password** | **string**| password | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FinishPasswordResetUsingPOST**
> string FinishPasswordResetUsingPOST(ctx, keyAndPassword)
finishPasswordReset

finish PasswordReset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyAndPassword** | [**KeyAndPasswordVm**](KeyAndPasswordVm.md)| keyAndPassword | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountUsingGET**
> UserDto GetAccountUsingGET(ctx, )
getAccount

get Account

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserDto**](UserDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsAuthenticatedUsingGET**
> string IsAuthenticatedUsingGET(ctx, )
isAuthenticated

is Authenticated

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterAccountUsingPOST**
> ResponseEntity RegisterAccountUsingPOST(ctx, managedUserVM)
registerAccount

register Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **managedUserVM** | [**ManagedUserVm**](ManagedUserVm.md)| managedUserVM | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPasswordResetUsingPOST**
> ResponseEntity RequestPasswordResetUsingPOST(ctx, mail)
requestPasswordReset

request PasswordReset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mail** | **string**| mail | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveAccountUsingPOST**
> ResponseEntity SaveAccountUsingPOST(ctx, userDTO)
saveAccount

save Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userDTO** | [**UserDto**](UserDto.md)| userDTO | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

