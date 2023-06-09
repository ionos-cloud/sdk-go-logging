# \TokensApi

All URIs are relative to *https://logging.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**TokensPost**](TokensApi.md#TokensPost) | **Post** /tokens | |



## TokensPost

```go
var result InlineResponse200 = TokensPost(ctx)
                      .Execute()
```



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-laas"
)

func main() {

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.TokensPost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.TokensPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TokensPost`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.TokensPost`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiTokensPostRequest struct via the builder pattern


### Return type

[**InlineResponse200**](../models/InlineResponse200.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


