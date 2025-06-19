# \EventsAPI

All URIs are relative to *https://apidoc.yangming.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventsAPI.md#GetEvents) | **Get** /v2/events | Find events.



## GetEvents

> Events GetEvents(ctx).KeyId(keyId).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).TransportEventTypeCode(transportEventTypeCode).VesselIMONumber(vesselIMONumber).TransportCallID(transportCallID).ExportVoyageNumber(exportVoyageNumber).EquipmentEventTypeCode(equipmentEventTypeCode).CarrierServiceCode(carrierServiceCode).EventType(eventType).UNLocationCode(uNLocationCode).Execute()

Find events.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/buyco/go-yangming-sdk/yangming"
)

func main() {
	keyId := "keyId_example" // string |  (optional) (default to "Assigned API Key")
	carrierBookingReference := "E123456789" // string | A set of unique characters provided by carrier to identify a booking.<br>Specifying this filter will only return events related to this particular carrierBookingReference.<br><small>*maxLength: 35*</small> (optional)
	transportDocumentReference := "W123456789" // string | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment.<br>Specifying this filter will only return events related to this particular transportDocumentReference <br><small>*maxLength: 20*</small> (optional)
	equipmentReference := "YMLU1234567" // string | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible.<br>Specifying this filter will only return events related to this particular equipmentReference<br><small>*maxLength: 15*</small> (optional)
	transportEventTypeCode := []string{"Inner_example"} // []string | Identifier for type of Transport event to filter by   - ARRI (Arrived)   - DEPA (Departed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example *transportEventTypeCode=ARRI,DEPA* matches **both** Arrived (ARRI) and Departed (DEPA) transport events.<br>Default is all transportEventTypeCodes.<br>This filter is only relevant when filtering on TransportEvents<br>*Available values* : ARRI, DEPA<br>*Default value* : ARRI,DEPA<br>*Example* : List [\"ARRI\",\"DEPA\"] (optional)
	vesselIMONumber := "9648714" // string | The identifier of vessel for which schedule details are published. Depending on schedule type, this may not be available yet.<br>Specifying this filter will only return events related to this particular vesselIMONumber. <br><small>*maxLength: 7*</small> (optional)
	transportCallID := "123e4567-e89b-12d3-a456-426614174000" // string | ID uniquely identifying a transport call, to filter events by.<br>Specifying this filter will only return events related to this particular transportCallID <br><small>*maxLength: 100*</small> (optional)
	exportVoyageNumber := "2103S" // string | Filter on the vessel operator-specific identifier of the export Voyage.<br>Specifying this filter will only return events related to this particular exportVoyageNumber. <br><small>*maxLength: 50*</small> (optional)
	equipmentEventTypeCode := []string{"Inner_example"} // []string | Unique identifier for equipmentEventTypeCode.   - LOAD (Loaded)   - DISC (Discharged)   - GTIN (Gated in)   - GTOT (Gated out)   - STUF (Stuffed)   - STRP (Stripped)   - PICK (Pick-up)   - DROP (Drop-off)   - INSP (Inspected)   - RSEA (Resealed)   - RMVD (Removed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example *equipmentEventTypeCode=GTIN,GTOT* matches **both** Gated in (GTIN) and Gated out (GTOT) equipment events.<br>Default is all equipmentEventTypeCodes.<br>This filter is only relevant when filtering on EquipmentEvents<br>*Available values* : LOAD, DISC, GTIN, GTOT, STUF, STRP, PICK, DROP, INSP, RSEA, RMVD<br>*Default value* : LOAD,DISC,GTIN,GTOT,STUF,STRP,PICK,DROP,INSP,RSEA,RMVD<br>*Example* : List [\"GTOT\",\"GTIN\"] (optional)
	carrierServiceCode := "FE1" // string | Filter on the carrier specific identifier of the service.<br>Specifying this filter will only return events related to this particular carrierServiceCode.<br><small>*maxLength: 5*</small> (optional)
	eventType := []string{"Inner_example"} // []string | The type of event(s) to filter by. Possible values are   - TRANSPORT (Transport events)   - EQUIPMENT (Equipment events)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example eventType=TRANSPORT,EQUIPMENT matches both Transport and Equipment events.<br>Default value is all event types.<br>*Available values* : TRANSPORT, EQUIPMENT<br>*Example* : List [\"TRANSPORT\",\"EQUIPMENT\"] (optional)
	uNLocationCode := "FRPAR" // string | The UN Location code specifying where the place is located.<br>Specifying this filter will only return events related to this particular UN Location code.<br><small>*maxLength: 5*</small> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEvents(context.Background()).KeyId(keyId).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).TransportEventTypeCode(transportEventTypeCode).VesselIMONumber(vesselIMONumber).TransportCallID(transportCallID).ExportVoyageNumber(exportVoyageNumber).EquipmentEventTypeCode(equipmentEventTypeCode).CarrierServiceCode(carrierServiceCode).EventType(eventType).UNLocationCode(uNLocationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: Events
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyId** | **string** |  | [default to &quot;Assigned API Key&quot;]
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking.&lt;br&gt;Specifying this filter will only return events related to this particular carrierBookingReference.&lt;br&gt;&lt;small&gt;*maxLength: 35*&lt;/small&gt; | 
 **transportDocumentReference** | **string** | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment.&lt;br&gt;Specifying this filter will only return events related to this particular transportDocumentReference &lt;br&gt;&lt;small&gt;*maxLength: 20*&lt;/small&gt; | 
 **equipmentReference** | **string** | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible.&lt;br&gt;Specifying this filter will only return events related to this particular equipmentReference&lt;br&gt;&lt;small&gt;*maxLength: 15*&lt;/small&gt; | 
 **transportEventTypeCode** | **[]string** | Identifier for type of Transport event to filter by   - ARRI (Arrived)   - DEPA (Departed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example *transportEventTypeCode&#x3D;ARRI,DEPA* matches **both** Arrived (ARRI) and Departed (DEPA) transport events.&lt;br&gt;Default is all transportEventTypeCodes.&lt;br&gt;This filter is only relevant when filtering on TransportEvents&lt;br&gt;*Available values* : ARRI, DEPA&lt;br&gt;*Default value* : ARRI,DEPA&lt;br&gt;*Example* : List [\&quot;ARRI\&quot;,\&quot;DEPA\&quot;] | 
 **vesselIMONumber** | **string** | The identifier of vessel for which schedule details are published. Depending on schedule type, this may not be available yet.&lt;br&gt;Specifying this filter will only return events related to this particular vesselIMONumber. &lt;br&gt;&lt;small&gt;*maxLength: 7*&lt;/small&gt; | 
 **transportCallID** | **string** | ID uniquely identifying a transport call, to filter events by.&lt;br&gt;Specifying this filter will only return events related to this particular transportCallID &lt;br&gt;&lt;small&gt;*maxLength: 100*&lt;/small&gt; | 
 **exportVoyageNumber** | **string** | Filter on the vessel operator-specific identifier of the export Voyage.&lt;br&gt;Specifying this filter will only return events related to this particular exportVoyageNumber. &lt;br&gt;&lt;small&gt;*maxLength: 50*&lt;/small&gt; | 
 **equipmentEventTypeCode** | **[]string** | Unique identifier for equipmentEventTypeCode.   - LOAD (Loaded)   - DISC (Discharged)   - GTIN (Gated in)   - GTOT (Gated out)   - STUF (Stuffed)   - STRP (Stripped)   - PICK (Pick-up)   - DROP (Drop-off)   - INSP (Inspected)   - RSEA (Resealed)   - RMVD (Removed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example *equipmentEventTypeCode&#x3D;GTIN,GTOT* matches **both** Gated in (GTIN) and Gated out (GTOT) equipment events.&lt;br&gt;Default is all equipmentEventTypeCodes.&lt;br&gt;This filter is only relevant when filtering on EquipmentEvents&lt;br&gt;*Available values* : LOAD, DISC, GTIN, GTOT, STUF, STRP, PICK, DROP, INSP, RSEA, RMVD&lt;br&gt;*Default value* : LOAD,DISC,GTIN,GTOT,STUF,STRP,PICK,DROP,INSP,RSEA,RMVD&lt;br&gt;*Example* : List [\&quot;GTOT\&quot;,\&quot;GTIN\&quot;] | 
 **carrierServiceCode** | **string** | Filter on the carrier specific identifier of the service.&lt;br&gt;Specifying this filter will only return events related to this particular carrierServiceCode.&lt;br&gt;&lt;small&gt;*maxLength: 5*&lt;/small&gt; | 
 **eventType** | **[]string** | The type of event(s) to filter by. Possible values are   - TRANSPORT (Transport events)   - EQUIPMENT (Equipment events)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example eventType&#x3D;TRANSPORT,EQUIPMENT matches both Transport and Equipment events.&lt;br&gt;Default value is all event types.&lt;br&gt;*Available values* : TRANSPORT, EQUIPMENT&lt;br&gt;*Example* : List [\&quot;TRANSPORT\&quot;,\&quot;EQUIPMENT\&quot;] | 
 **uNLocationCode** | **string** | The UN Location code specifying where the place is located.&lt;br&gt;Specifying this filter will only return events related to this particular UN Location code.&lt;br&gt;&lt;small&gt;*maxLength: 5*&lt;/small&gt; | 

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

