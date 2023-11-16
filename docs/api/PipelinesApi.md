# \PipelinesApi

All URIs are relative to *https://logging.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**PipelinesDelete**](PipelinesApi.md#PipelinesDelete) | **Delete** /pipelines/{pipelineId} | Delete a pipeline|
|[**PipelinesFindById**](PipelinesApi.md#PipelinesFindById) | **Get** /pipelines/{pipelineId} | Fetch a pipeline|
|[**PipelinesGet**](PipelinesApi.md#PipelinesGet) | **Get** /pipelines | List pipelines|
|[**PipelinesKeyPost**](PipelinesApi.md#PipelinesKeyPost) | **Post** /pipelines/{pipelineId}/key | Renews the key of a Pipeline|
|[**PipelinesPatch**](PipelinesApi.md#PipelinesPatch) | **Patch** /pipelines/{pipelineId} | Patch a pipeline|
|[**PipelinesPost**](PipelinesApi.md#PipelinesPost) | **Post** /pipelines | Create a pipeline|



## PipelinesDelete

```go
var result Pipeline = PipelinesDelete(ctx, pipelineId)
                      .Execute()
```

Delete a pipeline



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
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.PipelinesApi.PipelinesDelete(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesDelete`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PipelinesApiService.PipelinesDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "PipelinesApiService.PipelinesDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "PipelinesApiService.PipelinesDelete": {
    "port": "8443",
},
})
```


## PipelinesFindById

```go
var result Pipeline = PipelinesFindById(ctx, pipelineId)
                      .Execute()
```

Fetch a pipeline



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
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesFindById(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesFindById`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PipelinesApiService.PipelinesFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "PipelinesApiService.PipelinesFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "PipelinesApiService.PipelinesFindById": {
    "port": "8443",
},
})
```


## PipelinesGet

```go
var result PipelineListResponse = PipelinesGet(ctx)
                      .Limit(limit)
                      .Offset(offset)
                      .OrderBy(orderBy)
                      .Execute()
```

List pipelines



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
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination). Default to 100 (optional) (default to 0)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination). Default to 0 (optional) (default to 0)
    orderBy := "orderBy_example" // string | Sorts the results alphanumerically in ascending order based on the specified property (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesGet(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesGet`: PipelineListResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | the maximum number of elements to return (use together with offset for pagination). Default to 100 | [default to 0]|
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with limit for pagination). Default to 0 | [default to 0]|
| **orderBy** | **string** | Sorts the results alphanumerically in ascending order based on the specified property | |

### Return type

[**PipelineListResponse**](../models/PipelineListResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PipelinesApiService.PipelinesGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "PipelinesApiService.PipelinesGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "PipelinesApiService.PipelinesGet": {
    "port": "8443",
},
})
```


## PipelinesKeyPost

```go
var result PipelinesKeyPost200Response = PipelinesKeyPost(ctx, pipelineId)
                      .Execute()
```

Renews the key of a Pipeline



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
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesKeyPost(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesKeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesKeyPost`: PipelinesKeyPost200Response
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesKeyPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesKeyPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**PipelinesKeyPost200Response**](../models/PipelinesKeyPost200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PipelinesApiService.PipelinesKeyPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "PipelinesApiService.PipelinesKeyPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "PipelinesApiService.PipelinesKeyPost": {
    "port": "8443",
},
})
```


## PipelinesPatch

```go
var result Pipeline = PipelinesPatch(ctx, pipelineId)
                      .Pipeline(pipeline)
                      .Execute()
```

Patch a pipeline



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
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline
    pipeline := *openapiclient.NewPipelinePatch(*openapiclient.NewPipelinePatchProperties()) // PipelinePatch | The modified pipeline.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPatch(context.Background(), pipelineId).Pipeline(pipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPatch`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipeline** | [**PipelinePatch**](../models/PipelinePatch.md) | The modified pipeline. | |

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PipelinesApiService.PipelinesPatch"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "PipelinesApiService.PipelinesPatch": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "PipelinesApiService.PipelinesPatch": {
    "port": "8443",
},
})
```


## PipelinesPost

```go
var result Pipeline = PipelinesPost(ctx)
                      .Pipeline(pipeline)
                      .Execute()
```

Create a pipeline



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
    pipeline := *openapiclient.NewPipelineCreate(*openapiclient.NewPipelineCreateProperties("Name_example", []openapiclient.PipelineCreatePropertiesLogs{*openapiclient.NewPipelineCreatePropertiesLogs()})) // PipelineCreate | The pipeline to be created.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPost(context.Background()).Pipeline(pipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPost`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipeline** | [**PipelineCreate**](../models/PipelineCreate.md) | The pipeline to be created. | |

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PipelinesApiService.PipelinesPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "PipelinesApiService.PipelinesPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "PipelinesApiService.PipelinesPost": {
    "port": "8443",
},
})
```

