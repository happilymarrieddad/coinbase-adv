/*
Coinbase Advanced Trading API

OpenAPI 3.x specification for Coinbase Adavanced Trading

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// AccountAvailableBalance struct for AccountAvailableBalance
type AccountAvailableBalance struct {
	Value *float64 `json:"value,omitempty,string"`
	Currency *string `json:"currency,omitempty"`
}

// NewAccountAvailableBalance instantiates a new AccountAvailableBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAvailableBalance() *AccountAvailableBalance {
	this := AccountAvailableBalance{}
	return &this
}

// NewAccountAvailableBalanceWithDefaults instantiates a new AccountAvailableBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAvailableBalanceWithDefaults() *AccountAvailableBalance {
	this := AccountAvailableBalance{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AccountAvailableBalance) GetValue() float64 {
	if o == nil || isNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAvailableBalance) GetValueOk() (*float64, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AccountAvailableBalance) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *AccountAvailableBalance) SetValue(v float64) {
	o.Value = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountAvailableBalance) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAvailableBalance) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
    return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountAvailableBalance) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountAvailableBalance) SetCurrency(v string) {
	o.Currency = &v
}

func (o AccountAvailableBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAvailableBalance struct {
	value *AccountAvailableBalance
	isSet bool
}

func (v NullableAccountAvailableBalance) Get() *AccountAvailableBalance {
	return v.value
}

func (v *NullableAccountAvailableBalance) Set(val *AccountAvailableBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAvailableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAvailableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAvailableBalance(val *AccountAvailableBalance) *NullableAccountAvailableBalance {
	return &NullableAccountAvailableBalance{value: val, isSet: true}
}

func (v NullableAccountAvailableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAvailableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

