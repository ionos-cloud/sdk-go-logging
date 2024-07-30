# \CentralApi

All URIs are relative to *https://logging.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CentralLoggingGet**](CentralApi.md#CentralLoggingGet) | **Get** /central | Gets the central logging properties.|
|[**CentralLoggingToggle**](CentralApi.md#CentralLoggingToggle) | **Put** /central | Toggles the central logging.|



## CentralLoggingGet

```go
var result CentralLoggingResponse = CentralLoggingGet(ctx)
                      .Execute()
```

Gets the central logging properties.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-logging"
)

func main() {

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralLoggingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralLoggingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralLoggingGet`: CentralLoggingResponse
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralLoggingGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiCentralLoggingGetRequest struct via the builder pattern


### Return type

[**CentralLoggingResponse**](../models/CentralLoggingResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"CentralApiService.CentralLoggingGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "CentralApiService.CentralLoggingGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "CentralApiService.CentralLoggingGet": {
    "port": "8443",
},
})
```


## CentralLoggingToggle

```go
var result CentralLoggingResponse = CentralLoggingToggle(ctx)
                      .CentralLoggingToggle(centralLoggingToggle)
                      .Execute()
```

Toggles the central logging.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-logging"
)

func main() {
    centralLoggingToggle := *openapiclient.NewCentralLoggingToggle() // CentralLoggingToggle | Toggle central logging.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralLoggingToggle(context.Background()).CentralLoggingToggle(centralLoggingToggle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralLoggingToggle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralLoggingToggle`: CentralLoggingResponse
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralLoggingToggle`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCentralLoggingToggleRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **centralLoggingToggle** | [**CentralLoggingToggle**](../models/CentralLoggingToggle.md) | Toggle central logging. | |

### Return type

[**CentralLoggingResponse**](../models/CentralLoggingResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"CentralApiService.CentralLoggingToggle"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "CentralApiService.CentralLoggingToggle": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "CentralApiService.CentralLoggingToggle": {
    "port": "8443",
},
})
```

