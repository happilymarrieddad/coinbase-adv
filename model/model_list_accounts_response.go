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

// ListAccountsResponse struct for ListAccountsResponse
type ListAccountsResponse struct {
	HasNext *bool `json:"has_next,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Accounts []ListAccountsResponseAccountsInner `json:"accounts,omitempty"`
}

// NewListAccountsResponse instantiates a new ListAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccountsResponse() *ListAccountsResponse {
	this := ListAccountsResponse{}
	return &this
}

// NewListAccountsResponseWithDefaults instantiates a new ListAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccountsResponseWithDefaults() *ListAccountsResponse {
	this := ListAccountsResponse{}
	return &this
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *ListAccountsResponse) GetHasNext() bool {
	if o == nil || isNil(o.HasNext) {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetHasNextOk() (*bool, bool) {
	if o == nil || isNil(o.HasNext) {
    return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *ListAccountsResponse) HasHasNext() bool {
	if o != nil && !isNil(o.HasNext) {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *ListAccountsResponse) SetHasNext(v bool) {
	o.HasNext = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListAccountsResponse) GetCursor() string {
	if o == nil || isNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetCursorOk() (*string, bool) {
	if o == nil || isNil(o.Cursor) {
    return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListAccountsResponse) HasCursor() bool {
	if o != nil && !isNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListAccountsResponse) SetCursor(v string) {
	o.Cursor = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ListAccountsResponse) GetSize() int32 {
	if o == nil || isNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetSizeOk() (*int32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ListAccountsResponse) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ListAccountsResponse) SetSize(v int32) {
	o.Size = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ListAccountsResponse) GetAccounts() []ListAccountsResponseAccountsInner {
	if o == nil || isNil(o.Accounts) {
		var ret []ListAccountsResponseAccountsInner
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccountsResponse) GetAccountsOk() ([]ListAccountsResponseAccountsInner, bool) {
	if o == nil || isNil(o.Accounts) {
    return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ListAccountsResponse) HasAccounts() bool {
	if o != nil && !isNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []ListAccountsResponseAccountsInner and assigns it to the Accounts field.
func (o *ListAccountsResponse) SetAccounts(v []ListAccountsResponseAccountsInner) {
	o.Accounts = v
}

func (o ListAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasNext) {
		toSerialize["has_next"] = o.HasNext
	}
	if !isNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableListAccountsResponse struct {
	value *ListAccountsResponse
	isSet bool
}

func (v NullableListAccountsResponse) Get() *ListAccountsResponse {
	return v.value
}

func (v *NullableListAccountsResponse) Set(val *ListAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccountsResponse(val *ListAccountsResponse) *NullableListAccountsResponse {
	return &NullableListAccountsResponse{value: val, isSet: true}
}

func (v NullableListAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


