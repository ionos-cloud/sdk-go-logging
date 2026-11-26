# MetadataForCentralLoggingAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**GrafanaEndpoint** | **string** | The endpoint of the Grafana instance.  | [readonly] |
|**Products** | Pointer to **[]string** | Products with central logging enabled. This is a comma-separated list of product names.  | [optional] [readonly] |

## Methods

### NewMetadataForCentralLoggingAllOf

`func NewMetadataForCentralLoggingAllOf(grafanaEndpoint string, ) *MetadataForCentralLoggingAllOf`

NewMetadataForCentralLoggingAllOf instantiates a new MetadataForCentralLoggingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataForCentralLoggingAllOfWithDefaults

`func NewMetadataForCentralLoggingAllOfWithDefaults() *MetadataForCentralLoggingAllOf`

NewMetadataForCentralLoggingAllOfWithDefaults instantiates a new MetadataForCentralLoggingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrafanaEndpoint

`func (o *MetadataForCentralLoggingAllOf) GetGrafanaEndpoint() string`

GetGrafanaEndpoint returns the GrafanaEndpoint field if non-nil, zero value otherwise.

### GetGrafanaEndpointOk

`func (o *MetadataForCentralLoggingAllOf) GetGrafanaEndpointOk() (*string, bool)`

GetGrafanaEndpointOk returns a tuple with the GrafanaEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaEndpoint

`func (o *MetadataForCentralLoggingAllOf) SetGrafanaEndpoint(v string)`

SetGrafanaEndpoint sets GrafanaEndpoint field to given value.


### GetProducts

`func (o *MetadataForCentralLoggingAllOf) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *MetadataForCentralLoggingAllOf) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *MetadataForCentralLoggingAllOf) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *MetadataForCentralLoggingAllOf) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


