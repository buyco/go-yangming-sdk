# \EventsApi

All URIs are relative to *https://api.yangming.com/open/dcsa/api/tnt*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V110GGetEvents**](EventsApi.md#V110GGetEvents) | **Get** /1.1.0/events | Find events by Booking Reference, Bill of Lading or Equipment Reference.



## V110GGetEvents

> Events V110GGetEvents(ctx).KeyId(keyId).BookingReference(bookingReference).BillOfLadingNumber(billOfLadingNumber).EquipmentReference(equipmentReference).Execute()

Find events by Booking Reference, Bill of Lading or Equipment Reference.



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
    keyId := "keyId_example" // string |  (optional) (default to "Assigned API Key")
    bookingReference := "E123456789" // string | The identifier for a shipment, which is issued by and unique within each of the carriers. (optional)
    billOfLadingNumber := "W123456789" // string | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier's receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. (optional)
    equipmentReference := "YMLU1234567" // string | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.V110GGetEvents(context.Background()).KeyId(keyId).BookingReference(bookingReference).BillOfLadingNumber(billOfLadingNumber).EquipmentReference(equipmentReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.V110GGetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V110GGetEvents`: Events
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.V110GGetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV110GGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyId** | **string** |  | [default to &quot;Assigned API Key&quot;]
 **bookingReference** | **string** | The identifier for a shipment, which is issued by and unique within each of the carriers. | 
 **billOfLadingNumber** | **string** | Bill of lading number is an identifier that links to a shipment. Bill of Lading is the legal document issued to the customer, which confirms the carrier&#39;s receipt of the cargo from the customer acknowledging goods being shipped and specifying the terms of delivery. | 
 **equipmentReference** | **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. | 

### Return type

[**Events**](Events.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

