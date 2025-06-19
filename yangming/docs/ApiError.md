# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMethod** | **string** | The HTTP request method type. | 
**RequestUri** | **string** | The request URI. | 
**Errors** | [**[]SubErrors**](SubErrors.md) |  | 
**StatusCode** | **int32** | The HTTP status code. | 
**StatusCodeText** | **string** | The textual representation of the response status. | 
**ErrorDateTime** | **time.Time** | The date and time (in ISO 8601 format) the error occured. | 

## Methods

### NewApiError

`func NewApiError(httpMethod string, requestUri string, errors []SubErrors, statusCode int32, statusCodeText string, errorDateTime time.Time, ) *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpMethod

`func (o *ApiError) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ApiError) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ApiError) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.


### GetRequestUri

`func (o *ApiError) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *ApiError) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *ApiError) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.


### GetErrors

`func (o *ApiError) GetErrors() []SubErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiError) GetErrorsOk() (*[]SubErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiError) SetErrors(v []SubErrors)`

SetErrors sets Errors field to given value.


### GetStatusCode

`func (o *ApiError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApiError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApiError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusCodeText

`func (o *ApiError) GetStatusCodeText() string`

GetStatusCodeText returns the StatusCodeText field if non-nil, zero value otherwise.

### GetStatusCodeTextOk

`func (o *ApiError) GetStatusCodeTextOk() (*string, bool)`

GetStatusCodeTextOk returns a tuple with the StatusCodeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeText

`func (o *ApiError) SetStatusCodeText(v string)`

SetStatusCodeText sets StatusCodeText field to given value.


### GetErrorDateTime

`func (o *ApiError) GetErrorDateTime() time.Time`

GetErrorDateTime returns the ErrorDateTime field if non-nil, zero value otherwise.

### GetErrorDateTimeOk

`func (o *ApiError) GetErrorDateTimeOk() (*time.Time, bool)`

GetErrorDateTimeOk returns a tuple with the ErrorDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDateTime

`func (o *ApiError) SetErrorDateTime(v time.Time)`

SetErrorDateTime sets ErrorDateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


