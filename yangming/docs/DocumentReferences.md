# DocumentReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentReferenceType** | Pointer to **string** | Describes where the documentReferenceValue is pointing to  Enum:  [ BKG (Booking), TRD (Transport Document) ] | [optional] 
**DocumentReferenceValue** | Pointer to **string** | The value of the identifier the documentReferenceType is describing | [optional] 

## Methods

### NewDocumentReferences

`func NewDocumentReferences() *DocumentReferences`

NewDocumentReferences instantiates a new DocumentReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentReferencesWithDefaults

`func NewDocumentReferencesWithDefaults() *DocumentReferences`

NewDocumentReferencesWithDefaults instantiates a new DocumentReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentReferenceType

`func (o *DocumentReferences) GetDocumentReferenceType() string`

GetDocumentReferenceType returns the DocumentReferenceType field if non-nil, zero value otherwise.

### GetDocumentReferenceTypeOk

`func (o *DocumentReferences) GetDocumentReferenceTypeOk() (*string, bool)`

GetDocumentReferenceTypeOk returns a tuple with the DocumentReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferenceType

`func (o *DocumentReferences) SetDocumentReferenceType(v string)`

SetDocumentReferenceType sets DocumentReferenceType field to given value.

### HasDocumentReferenceType

`func (o *DocumentReferences) HasDocumentReferenceType() bool`

HasDocumentReferenceType returns a boolean if a field has been set.

### GetDocumentReferenceValue

`func (o *DocumentReferences) GetDocumentReferenceValue() string`

GetDocumentReferenceValue returns the DocumentReferenceValue field if non-nil, zero value otherwise.

### GetDocumentReferenceValueOk

`func (o *DocumentReferences) GetDocumentReferenceValueOk() (*string, bool)`

GetDocumentReferenceValueOk returns a tuple with the DocumentReferenceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferenceValue

`func (o *DocumentReferences) SetDocumentReferenceValue(v string)`

SetDocumentReferenceValue sets DocumentReferenceValue field to given value.

### HasDocumentReferenceValue

`func (o *DocumentReferences) HasDocumentReferenceValue() bool`

HasDocumentReferenceValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


