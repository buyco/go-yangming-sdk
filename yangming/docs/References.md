# References

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | **string** | The reference type codes defined by DCSA.   - FF (Freight Forwarder’s Reference)   - SI (Shipper’s Reference)   - PO (Purchase Order Reference)   - CR (Customer’s Reference)   - AAO (Consignee’s Reference)   - EQ (Equipment Reference)  Enum:  [ FF, SI, PO, CR, AAO, EQ ] | 
**ReferenceValue** | **string** | &lt;small&gt;maxLength: 100&lt;/small&gt;  The actual value of the reference. | 

## Methods

### NewReferences

`func NewReferences(referenceType string, referenceValue string, ) *References`

NewReferences instantiates a new References object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencesWithDefaults

`func NewReferencesWithDefaults() *References`

NewReferencesWithDefaults instantiates a new References object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *References) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *References) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *References) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetReferenceValue

`func (o *References) GetReferenceValue() string`

GetReferenceValue returns the ReferenceValue field if non-nil, zero value otherwise.

### GetReferenceValueOk

`func (o *References) GetReferenceValueOk() (*string, bool)`

GetReferenceValueOk returns a tuple with the ReferenceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceValue

`func (o *References) SetReferenceValue(v string)`

SetReferenceValue sets ReferenceValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


