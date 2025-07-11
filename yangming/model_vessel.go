/*
Yang Ming Track and Trace API

API specifications for the Track and Trace interface standard

API version: DCSA Standard
Contact: itcs@yangming.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yangming

import (
	"encoding/json"
)

// checks if the Vessel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vessel{}

// Vessel describes a floating, sea going structure (mother vessels and feeder vessels) with either an internal or external mode of propulsion designed for the transport of cargo and/or passengers. Ocean vessels are uniquely identified by an IMO number consisting of 7 digits, or alternatively by their AIS signal with an MMSI number.
type Vessel struct {
	// <small>maxLength: 7</small><br>The unique reference for a registered Vessel. The reference is the International Maritime Organisation (IMO) number, also sometimes known as the Lloyd's register code, which does not change during the lifetime of the vessel
	VesselIMONumber *string `json:"vesselIMONumber,omitempty"`
	// <small>maxLength: 35</small><br>The name of the Vessel given by the Vessel Operator and registered with IMO.
	VesselName *string `json:"vesselName,omitempty"`
	// <small>maxLength: 2</small><br>The flag of the nation whose laws the vessel is registered under. This is the ISO 3166 two-letter country code
	VesselFlag *string `json:"vesselFlag,omitempty"`
	// <small>maxLength: 10</small><br>A unique alphanumeric identity that belongs to the vessel and is assigned by the International Telecommunication Union (ITU). It consists of a threeletter alphanumeric prefix that indicates nationality, followed by one to four characters to identify the individual vessel. For instance, vessels registered under Denmark are assigned the prefix ranges 5PA-5QZ, OUAOZZ, and XPA-XPZ. The Call Sign changes whenever a vessel changes its flag.
	VesselCallSignNumber *string `json:"vesselCallSignNumber,omitempty"`
	// <small>maxLength: 10</small><br><small>nullable: false</small>  The carrier who is in charge of the vessel operation based on either the SMDG or SCAC code lists
	VesselOperatorCarrierCode *string `json:"vesselOperatorCarrierCode,omitempty"`
	// <small>nullable: false</small><br>Identifies the code list provider used for the operator and partner carriercodes.<br><br>Enum:<br>[ SMDG, NMFTA ]
	VesselOperatorCarrierCodeListProvider *string `json:"vesselOperatorCarrierCodeListProvider,omitempty"`
}

// NewVessel instantiates a new Vessel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVessel() *Vessel {
	this := Vessel{}
	return &this
}

// NewVesselWithDefaults instantiates a new Vessel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVesselWithDefaults() *Vessel {
	this := Vessel{}
	return &this
}

// GetVesselIMONumber returns the VesselIMONumber field value if set, zero value otherwise.
func (o *Vessel) GetVesselIMONumber() string {
	if o == nil || IsNil(o.VesselIMONumber) {
		var ret string
		return ret
	}
	return *o.VesselIMONumber
}

// GetVesselIMONumberOk returns a tuple with the VesselIMONumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetVesselIMONumberOk() (*string, bool) {
	if o == nil || IsNil(o.VesselIMONumber) {
		return nil, false
	}
	return o.VesselIMONumber, true
}

// HasVesselIMONumber returns a boolean if a field has been set.
func (o *Vessel) HasVesselIMONumber() bool {
	if o != nil && !IsNil(o.VesselIMONumber) {
		return true
	}

	return false
}

// SetVesselIMONumber gets a reference to the given string and assigns it to the VesselIMONumber field.
func (o *Vessel) SetVesselIMONumber(v string) {
	o.VesselIMONumber = &v
}

// GetVesselName returns the VesselName field value if set, zero value otherwise.
func (o *Vessel) GetVesselName() string {
	if o == nil || IsNil(o.VesselName) {
		var ret string
		return ret
	}
	return *o.VesselName
}

// GetVesselNameOk returns a tuple with the VesselName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetVesselNameOk() (*string, bool) {
	if o == nil || IsNil(o.VesselName) {
		return nil, false
	}
	return o.VesselName, true
}

// HasVesselName returns a boolean if a field has been set.
func (o *Vessel) HasVesselName() bool {
	if o != nil && !IsNil(o.VesselName) {
		return true
	}

	return false
}

// SetVesselName gets a reference to the given string and assigns it to the VesselName field.
func (o *Vessel) SetVesselName(v string) {
	o.VesselName = &v
}

// GetVesselFlag returns the VesselFlag field value if set, zero value otherwise.
func (o *Vessel) GetVesselFlag() string {
	if o == nil || IsNil(o.VesselFlag) {
		var ret string
		return ret
	}
	return *o.VesselFlag
}

// GetVesselFlagOk returns a tuple with the VesselFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetVesselFlagOk() (*string, bool) {
	if o == nil || IsNil(o.VesselFlag) {
		return nil, false
	}
	return o.VesselFlag, true
}

// HasVesselFlag returns a boolean if a field has been set.
func (o *Vessel) HasVesselFlag() bool {
	if o != nil && !IsNil(o.VesselFlag) {
		return true
	}

	return false
}

// SetVesselFlag gets a reference to the given string and assigns it to the VesselFlag field.
func (o *Vessel) SetVesselFlag(v string) {
	o.VesselFlag = &v
}

// GetVesselCallSignNumber returns the VesselCallSignNumber field value if set, zero value otherwise.
func (o *Vessel) GetVesselCallSignNumber() string {
	if o == nil || IsNil(o.VesselCallSignNumber) {
		var ret string
		return ret
	}
	return *o.VesselCallSignNumber
}

// GetVesselCallSignNumberOk returns a tuple with the VesselCallSignNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetVesselCallSignNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VesselCallSignNumber) {
		return nil, false
	}
	return o.VesselCallSignNumber, true
}

// HasVesselCallSignNumber returns a boolean if a field has been set.
func (o *Vessel) HasVesselCallSignNumber() bool {
	if o != nil && !IsNil(o.VesselCallSignNumber) {
		return true
	}

	return false
}

// SetVesselCallSignNumber gets a reference to the given string and assigns it to the VesselCallSignNumber field.
func (o *Vessel) SetVesselCallSignNumber(v string) {
	o.VesselCallSignNumber = &v
}

// GetVesselOperatorCarrierCode returns the VesselOperatorCarrierCode field value if set, zero value otherwise.
func (o *Vessel) GetVesselOperatorCarrierCode() string {
	if o == nil || IsNil(o.VesselOperatorCarrierCode) {
		var ret string
		return ret
	}
	return *o.VesselOperatorCarrierCode
}

// GetVesselOperatorCarrierCodeOk returns a tuple with the VesselOperatorCarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetVesselOperatorCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VesselOperatorCarrierCode) {
		return nil, false
	}
	return o.VesselOperatorCarrierCode, true
}

// HasVesselOperatorCarrierCode returns a boolean if a field has been set.
func (o *Vessel) HasVesselOperatorCarrierCode() bool {
	if o != nil && !IsNil(o.VesselOperatorCarrierCode) {
		return true
	}

	return false
}

// SetVesselOperatorCarrierCode gets a reference to the given string and assigns it to the VesselOperatorCarrierCode field.
func (o *Vessel) SetVesselOperatorCarrierCode(v string) {
	o.VesselOperatorCarrierCode = &v
}

// GetVesselOperatorCarrierCodeListProvider returns the VesselOperatorCarrierCodeListProvider field value if set, zero value otherwise.
func (o *Vessel) GetVesselOperatorCarrierCodeListProvider() string {
	if o == nil || IsNil(o.VesselOperatorCarrierCodeListProvider) {
		var ret string
		return ret
	}
	return *o.VesselOperatorCarrierCodeListProvider
}

// GetVesselOperatorCarrierCodeListProviderOk returns a tuple with the VesselOperatorCarrierCodeListProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetVesselOperatorCarrierCodeListProviderOk() (*string, bool) {
	if o == nil || IsNil(o.VesselOperatorCarrierCodeListProvider) {
		return nil, false
	}
	return o.VesselOperatorCarrierCodeListProvider, true
}

// HasVesselOperatorCarrierCodeListProvider returns a boolean if a field has been set.
func (o *Vessel) HasVesselOperatorCarrierCodeListProvider() bool {
	if o != nil && !IsNil(o.VesselOperatorCarrierCodeListProvider) {
		return true
	}

	return false
}

// SetVesselOperatorCarrierCodeListProvider gets a reference to the given string and assigns it to the VesselOperatorCarrierCodeListProvider field.
func (o *Vessel) SetVesselOperatorCarrierCodeListProvider(v string) {
	o.VesselOperatorCarrierCodeListProvider = &v
}

func (o Vessel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vessel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VesselIMONumber) {
		toSerialize["vesselIMONumber"] = o.VesselIMONumber
	}
	if !IsNil(o.VesselName) {
		toSerialize["vesselName"] = o.VesselName
	}
	if !IsNil(o.VesselFlag) {
		toSerialize["vesselFlag"] = o.VesselFlag
	}
	if !IsNil(o.VesselCallSignNumber) {
		toSerialize["vesselCallSignNumber"] = o.VesselCallSignNumber
	}
	if !IsNil(o.VesselOperatorCarrierCode) {
		toSerialize["vesselOperatorCarrierCode"] = o.VesselOperatorCarrierCode
	}
	if !IsNil(o.VesselOperatorCarrierCodeListProvider) {
		toSerialize["vesselOperatorCarrierCodeListProvider"] = o.VesselOperatorCarrierCodeListProvider
	}
	return toSerialize, nil
}

type NullableVessel struct {
	value *Vessel
	isSet bool
}

func (v NullableVessel) Get() *Vessel {
	return v.value
}

func (v *NullableVessel) Set(val *Vessel) {
	v.value = val
	v.isSet = true
}

func (v NullableVessel) IsSet() bool {
	return v.isSet
}

func (v *NullableVessel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVessel(val *Vessel) *NullableVessel {
	return &NullableVessel{value: val, isSet: true}
}

func (v NullableVessel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVessel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
