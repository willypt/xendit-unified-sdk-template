/*
Xendit Balance API

Utilize this API to get your Xendit's balance.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Balance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Balance{}

// Balance struct for Balance
type Balance struct {
	Balance *float32 `json:"balance,omitempty"`
}

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance() *Balance {
	this := Balance{}
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Balance) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Balance) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Balance) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *Balance) SetBalance(v float32) {
	o.Balance = &v
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	return toSerialize, nil
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


