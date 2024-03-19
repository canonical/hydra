# \V0alpha2Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OAuth2DeviceFlow**](V0alpha2Api.md#OAuth2DeviceFlow) | **Post** /oauth2/device/auth | The OAuth 2.0 Device Authorize Endpoint
[**PerformOAuth2DeviceVerificationFlow**](V0alpha2Api.md#PerformOAuth2DeviceVerificationFlow) | **Get** /oauth2/device/verify | OAuth 2.0 Device Verification Endpoint



## OAuth2DeviceFlow

> DeviceAuthorization OAuth2DeviceFlow(ctx).Execute()

The OAuth 2.0 Device Authorize Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V0alpha2Api.OAuth2DeviceFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.OAuth2DeviceFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OAuth2DeviceFlow`: DeviceAuthorization
    fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.OAuth2DeviceFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOAuth2DeviceFlowRequest struct via the builder pattern


### Return type

[**DeviceAuthorization**](DeviceAuthorization.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformOAuth2DeviceVerificationFlow

> ErrorOAuth2 PerformOAuth2DeviceVerificationFlow(ctx).Execute()

OAuth 2.0 Device Verification Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V0alpha2Api.PerformOAuth2DeviceVerificationFlow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V0alpha2Api.PerformOAuth2DeviceVerificationFlow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformOAuth2DeviceVerificationFlow`: ErrorOAuth2
    fmt.Fprintf(os.Stdout, "Response from `V0alpha2Api.PerformOAuth2DeviceVerificationFlow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPerformOAuth2DeviceVerificationFlowRequest struct via the builder pattern


### Return type

[**ErrorOAuth2**](ErrorOAuth2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

