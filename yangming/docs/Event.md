# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).&lt;br&gt;NB: This field should be considered Metadata | [optional] 
**EventCreatedDateTime** | **string** | The timestamp of when the event was created.&lt;br&gt;NB: This field should be considered Metadata | 
**EventType** | **string** | The Event Type of the object - to be used as a discriminator. &lt;br&gt;NB: This field should be considered Metadata  Enum:  [ TRANSPORT ] | 
**EventClassifierCode** | **string** | Code for the event classifier can be    - ACT (Actual)   - PLN (Planned)   - EST (Estimated)  Enum:&lt;br&gt;[ ACT, PLN, EST ] | 
**EventDateTime** | **string** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 
**TransportEventTypeCode** | **string** | Identifier for type of Transport event   - ARRI (Arrived)   - DEPA (Departed)  More details can be found on GitHub&lt;br&gt;&lt;br&gt;Enum:&lt;br&gt;[ ARRI, DEPA ] | 
**DelayReasonCode** | Pointer to **string** | &lt;small&gt;maxLength: 3&lt;/small&gt;&lt;br&gt;Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/ | [optional] 
**ChangeRemark** | Pointer to **string** | &lt;small&gt;maxLength: 250&lt;/small&gt;&lt;br&gt;Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 
**DocumentReferences** | Pointer to [**[]DocumentReferences**](DocumentReferences.md) |  | [optional] 
**References** | Pointer to [**[]References**](References.md) |  | [optional] 
**EquipmentEventTypeCode** | **string** | Unique identifier for equipmentEventTypeCode.   - LOAD (Loaded)   - DISC (Discharged)   - GTIN (Gated in)   - GTOT (Gated out)   - STUF (Stuffed)   - STRP (Stripped)   - PICK (Pick-up)   - DROP (Drop-off)   - INSP (Inspected)   - RSEA (Resealed)   - RMVD (Removed)  More details can be found on GitHub &lt;br&gt;&lt;br&gt;Enum:&lt;br&gt;[ LOAD, DISC, GTIN, GTOT, STUF, STRP, PICK, DROP, INSP, RSEA, RMVD ] | 
**EquipmentReference** | Pointer to **string** | &lt;small&gt;maxLength: 15&lt;/small&gt;&lt;br&gt;The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG. | [optional] 
**ISOEquipmentCode** | Pointer to **string** | &lt;small&gt;maxLength: 4&lt;/small&gt;&lt;br&gt;Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**EmptyIndicatorCode** | **string** | Code to denote whether the equipment is empty or laden.  Enum:  [ EMPTY, LADEN ] | 
**EventLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) | addresses the seal-related information associated with the shipment equipment. A seal is put on a shipment equipment once it is loaded. This seal is meant to stay on until the shipment equipment reaches its final destination. Only available when querying with Transport Document Reference. | [optional] 

## Methods

### NewEvent

`func NewEvent(eventCreatedDateTime string, eventType string, eventClassifierCode string, eventDateTime string, transportEventTypeCode string, transportCall TransportCall, equipmentEventTypeCode string, emptyIndicatorCode string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *Event) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *Event) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *Event) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *Event) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *Event) GetEventCreatedDateTime() string`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *Event) GetEventCreatedDateTimeOk() (*string, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *Event) SetEventCreatedDateTime(v string)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *Event) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *Event) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *Event) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *Event) GetEventDateTime() string`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *Event) GetEventDateTimeOk() (*string, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *Event) SetEventDateTime(v string)`

SetEventDateTime sets EventDateTime field to given value.


### GetTransportEventTypeCode

`func (o *Event) GetTransportEventTypeCode() string`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *Event) GetTransportEventTypeCodeOk() (*string, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *Event) SetTransportEventTypeCode(v string)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.


### GetDelayReasonCode

`func (o *Event) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *Event) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *Event) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *Event) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetChangeRemark

`func (o *Event) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *Event) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *Event) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *Event) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetTransportCall

`func (o *Event) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *Event) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *Event) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetDocumentReferences

`func (o *Event) GetDocumentReferences() []DocumentReferences`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *Event) GetDocumentReferencesOk() (*[]DocumentReferences, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *Event) SetDocumentReferences(v []DocumentReferences)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *Event) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetReferences

`func (o *Event) GetReferences() []References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Event) GetReferencesOk() (*[]References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Event) SetReferences(v []References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Event) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetEquipmentEventTypeCode

`func (o *Event) GetEquipmentEventTypeCode() string`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *Event) GetEquipmentEventTypeCodeOk() (*string, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *Event) SetEquipmentEventTypeCode(v string)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.


### GetEquipmentReference

`func (o *Event) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *Event) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *Event) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *Event) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *Event) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *Event) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *Event) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *Event) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *Event) GetEmptyIndicatorCode() string`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *Event) GetEmptyIndicatorCodeOk() (*string, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *Event) SetEmptyIndicatorCode(v string)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *Event) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *Event) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *Event) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *Event) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetSeals

`func (o *Event) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *Event) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *Event) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *Event) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


