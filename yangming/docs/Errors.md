# Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | High level error message:   1.invalidQuery   2.noDataFound   3.systemException. | 
**Message** | **string** | Detailed error message:   1.invalidQuery: The request did not contain one of the three required query parameters.   2.noDataFound: No data found for given parameters.   3.systemException: Exception occurs when calling the API. | 

## Methods

### NewErrors

`func NewErrors(reason string, message string, ) *Errors`

NewErrors instantiates a new Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsWithDefaults

`func NewErrorsWithDefaults() *Errors`

NewErrorsWithDefaults instantiates a new Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Errors) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Errors) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Errors) SetReason(v string)`

SetReason sets Reason field to given value.


### GetMessage

`func (o *Errors) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Errors) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Errors) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


