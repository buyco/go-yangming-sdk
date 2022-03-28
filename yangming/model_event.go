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
	"fmt"
)

// Event struct for Event
type Event struct {
	EquipmentEvent *EquipmentEvent
	TransportEquipmentEvent *TransportEquipmentEvent
	TransportEvent *TransportEvent
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Event) UnmarshalJSON(data []byte) error {
	var err error

	// try to unmarshal JSON data into TransportEquipmentEvent
	err = json.Unmarshal(data, &dst.TransportEquipmentEvent);
	if err == nil {
		jsonTransportEquipmentEvent, _ := json.Marshal(dst.TransportEquipmentEvent)
		if string(jsonTransportEquipmentEvent) == "{}" { // empty struct
			dst.TransportEquipmentEvent = nil
		} else {
			if dst.TransportEquipmentEvent.EquipmentReference != "" && dst.TransportEquipmentEvent.TransportReference != "" {
				return nil // data stored in dst.TransportEquipmentEvent, return on the first match
			}
			dst.TransportEquipmentEvent = nil
		}
	} else {
		dst.TransportEquipmentEvent = nil
	}

	// try to unmarshal JSON data into EquipmentEvent
	err = json.Unmarshal(data, &dst.EquipmentEvent);
	if err == nil {
		jsonEquipmentEvent, _ := json.Marshal(dst.EquipmentEvent)
		if string(jsonEquipmentEvent) == "{}" { // empty struct
			dst.EquipmentEvent = nil
		} else {
			if dst.EquipmentEvent.EquipmentReference != "" {
				return nil // data stored in dst.EquipmentEvent, return on the first match
			}
			dst.EquipmentEvent = nil
		}
	} else {
		dst.EquipmentEvent = nil
	}

	// try to unmarshal JSON data into TransportEvent
	err = json.Unmarshal(data, &dst.TransportEvent);
	if err == nil {
		jsonTransportEvent, _ := json.Marshal(dst.TransportEvent)
		if string(jsonTransportEvent) == "{}" { // empty struct
			dst.TransportEvent = nil
		} else {
			if dst.TransportEvent.TransportReference != "" {
				return nil // data stored in dst.TransportEvent, return on the first match
			}
			dst.TransportEvent = nil
		}
	} else {
		dst.TransportEvent = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(Event)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Event) MarshalJSON() ([]byte, error) {
	if src.EquipmentEvent != nil {
		return json.Marshal(&src.EquipmentEvent)
	}

	if src.TransportEquipmentEvent != nil {
		return json.Marshal(&src.TransportEquipmentEvent)
	}

	if src.TransportEvent != nil {
		return json.Marshal(&src.TransportEvent)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


