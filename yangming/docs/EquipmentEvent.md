# EquipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source).&lt;br&gt;NB: This field should be considered Metadata | [optional] 
**EventCreatedDateTime** | **time.Time** | The timestamp of when the event was created.&lt;br&gt;NB: This field should be considered Metadata | 
**EventType** | **string** | The Event Type of the object - to be used as a discriminator. &lt;br&gt;NB: This field should be considered Metadata  Enum:  [ TRANSPORT ] | 
**EventClassifierCode** | **string** | Code for the event classifier can be    - ACT (Actual)   - PLN (Planned)   - EST (Estimated)  Enum:&lt;br&gt;[ ACT, PLN, EST ] | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 
**EquipmentEventTypeCode** | **string** | Unique identifier for equipmentEventTypeCode.   - LOAD (Loaded)   - DISC (Discharged)   - GTIN (Gated in)   - GTOT (Gated out)   - STUF (Stuffed)   - STRP (Stripped)   - PICK (Pick-up)   - DROP (Drop-off)   - INSP (Inspected)   - RSEA (Resealed)   - RMVD (Removed)  More details can be found on GitHub &lt;br&gt;&lt;br&gt;Enum:&lt;br&gt;[ LOAD, DISC, GTIN, GTOT, STUF, STRP, PICK, DROP, INSP, RSEA, RMVD ] | 
**EquipmentReference** | Pointer to **string** | &lt;small&gt;maxLength: 15&lt;/small&gt;&lt;br&gt;The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG. | [optional] 
**ISOEquipmentCode** | Pointer to **string** | &lt;small&gt;maxLength: 4&lt;/small&gt;&lt;br&gt;Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**EmptyIndicatorCode** | **string** | Code to denote whether the equipment is empty or laden.  Enum:  [ EMPTY, LADEN ] | 
**EventLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**TransportCall** | Pointer to [**TransportCall**](TransportCall.md) |  | [optional] 
**DocumentReferences** | Pointer to [**[]DocumentReferences**](DocumentReferences.md) |  | [optional] 
**References** | Pointer to [**[]References**](References.md) |  | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) | addresses the seal-related information associated with the shipment equipment. A seal is put on a shipment equipment once it is loaded. This seal is meant to stay on until the shipment equipment reaches its final destination. Only available when querying with Transport Document Reference. | [optional] 

## Methods

### NewEquipmentEvent

`func NewEquipmentEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time, equipmentEventTypeCode string, emptyIndicatorCode string, ) *EquipmentEvent`

NewEquipmentEvent instantiates a new EquipmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentEventWithDefaults

`func NewEquipmentEventWithDefaults() *EquipmentEvent`

NewEquipmentEventWithDefaults instantiates a new EquipmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *EquipmentEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *EquipmentEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *EquipmentEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *EquipmentEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *EquipmentEvent) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *EquipmentEvent) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *EquipmentEvent) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *EquipmentEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EquipmentEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EquipmentEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *EquipmentEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *EquipmentEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *EquipmentEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *EquipmentEvent) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *EquipmentEvent) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *EquipmentEvent) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetEquipmentEventTypeCode

`func (o *EquipmentEvent) GetEquipmentEventTypeCode() string`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *EquipmentEvent) GetEquipmentEventTypeCodeOk() (*string, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *EquipmentEvent) SetEquipmentEventTypeCode(v string)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.


### GetEquipmentReference

`func (o *EquipmentEvent) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *EquipmentEvent) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *EquipmentEvent) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *EquipmentEvent) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *EquipmentEvent) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *EquipmentEvent) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *EquipmentEvent) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *EquipmentEvent) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *EquipmentEvent) GetEmptyIndicatorCode() string`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *EquipmentEvent) GetEmptyIndicatorCodeOk() (*string, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *EquipmentEvent) SetEmptyIndicatorCode(v string)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *EquipmentEvent) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *EquipmentEvent) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *EquipmentEvent) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *EquipmentEvent) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetTransportCall

`func (o *EquipmentEvent) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *EquipmentEvent) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *EquipmentEvent) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.

### HasTransportCall

`func (o *EquipmentEvent) HasTransportCall() bool`

HasTransportCall returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *EquipmentEvent) GetDocumentReferences() []DocumentReferences`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *EquipmentEvent) GetDocumentReferencesOk() (*[]DocumentReferences, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *EquipmentEvent) SetDocumentReferences(v []DocumentReferences)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *EquipmentEvent) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetReferences

`func (o *EquipmentEvent) GetReferences() []References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EquipmentEvent) GetReferencesOk() (*[]References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EquipmentEvent) SetReferences(v []References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *EquipmentEvent) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeals

`func (o *EquipmentEvent) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *EquipmentEvent) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *EquipmentEvent) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *EquipmentEvent) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


