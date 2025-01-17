/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MarginTradeDetailInfo struct for MarginTradeDetailInfo
type MarginTradeDetailInfo struct {
	Ctime                *string `json:"ctime,omitempty"`
	FeeCcy               *string `json:"feeCcy,omitempty"`
	Fees                 *string `json:"fees,omitempty"`
	FillId               *string `json:"fillId,omitempty"`
	FillPrice            *string `json:"fillPrice,omitempty"`
	FillQuantity         *string `json:"fillQuantity,omitempty"`
	FillTotalAmount      *string `json:"fillTotalAmount,omitempty"`
	OrderId              *string `json:"orderId,omitempty"`
	OrderType            *string `json:"orderType,omitempty"`
	Side                 *string `json:"side,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginTradeDetailInfo MarginTradeDetailInfo

// NewMarginTradeDetailInfo instantiates a new MarginTradeDetailInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginTradeDetailInfo() *MarginTradeDetailInfo {
	this := MarginTradeDetailInfo{}
	return &this
}

// NewMarginTradeDetailInfoWithDefaults instantiates a new MarginTradeDetailInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginTradeDetailInfoWithDefaults() *MarginTradeDetailInfo {
	this := MarginTradeDetailInfo{}
	return &this
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MarginTradeDetailInfo) SetCtime(v string) {
	o.Ctime = &v
}

// GetFeeCcy returns the FeeCcy field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetFeeCcy() string {
	if o == nil || isNil(o.FeeCcy) {
		var ret string
		return ret
	}
	return *o.FeeCcy
}

// GetFeeCcyOk returns a tuple with the FeeCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetFeeCcyOk() (*string, bool) {
	if o == nil || isNil(o.FeeCcy) {
		return nil, false
	}
	return o.FeeCcy, true
}

// HasFeeCcy returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasFeeCcy() bool {
	if o != nil && !isNil(o.FeeCcy) {
		return true
	}

	return false
}

// SetFeeCcy gets a reference to the given string and assigns it to the FeeCcy field.
func (o *MarginTradeDetailInfo) SetFeeCcy(v string) {
	o.FeeCcy = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetFees() string {
	if o == nil || isNil(o.Fees) {
		var ret string
		return ret
	}
	return *o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetFeesOk() (*string, bool) {
	if o == nil || isNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasFees() bool {
	if o != nil && !isNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given string and assigns it to the Fees field.
func (o *MarginTradeDetailInfo) SetFees(v string) {
	o.Fees = &v
}

// GetFillId returns the FillId field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetFillId() string {
	if o == nil || isNil(o.FillId) {
		var ret string
		return ret
	}
	return *o.FillId
}

// GetFillIdOk returns a tuple with the FillId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetFillIdOk() (*string, bool) {
	if o == nil || isNil(o.FillId) {
		return nil, false
	}
	return o.FillId, true
}

// HasFillId returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasFillId() bool {
	if o != nil && !isNil(o.FillId) {
		return true
	}

	return false
}

// SetFillId gets a reference to the given string and assigns it to the FillId field.
func (o *MarginTradeDetailInfo) SetFillId(v string) {
	o.FillId = &v
}

// GetFillPrice returns the FillPrice field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetFillPrice() string {
	if o == nil || isNil(o.FillPrice) {
		var ret string
		return ret
	}
	return *o.FillPrice
}

// GetFillPriceOk returns a tuple with the FillPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetFillPriceOk() (*string, bool) {
	if o == nil || isNil(o.FillPrice) {
		return nil, false
	}
	return o.FillPrice, true
}

// HasFillPrice returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasFillPrice() bool {
	if o != nil && !isNil(o.FillPrice) {
		return true
	}

	return false
}

// SetFillPrice gets a reference to the given string and assigns it to the FillPrice field.
func (o *MarginTradeDetailInfo) SetFillPrice(v string) {
	o.FillPrice = &v
}

// GetFillQuantity returns the FillQuantity field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetFillQuantity() string {
	if o == nil || isNil(o.FillQuantity) {
		var ret string
		return ret
	}
	return *o.FillQuantity
}

// GetFillQuantityOk returns a tuple with the FillQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetFillQuantityOk() (*string, bool) {
	if o == nil || isNil(o.FillQuantity) {
		return nil, false
	}
	return o.FillQuantity, true
}

// HasFillQuantity returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasFillQuantity() bool {
	if o != nil && !isNil(o.FillQuantity) {
		return true
	}

	return false
}

// SetFillQuantity gets a reference to the given string and assigns it to the FillQuantity field.
func (o *MarginTradeDetailInfo) SetFillQuantity(v string) {
	o.FillQuantity = &v
}

// GetFillTotalAmount returns the FillTotalAmount field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetFillTotalAmount() string {
	if o == nil || isNil(o.FillTotalAmount) {
		var ret string
		return ret
	}
	return *o.FillTotalAmount
}

// GetFillTotalAmountOk returns a tuple with the FillTotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetFillTotalAmountOk() (*string, bool) {
	if o == nil || isNil(o.FillTotalAmount) {
		return nil, false
	}
	return o.FillTotalAmount, true
}

// HasFillTotalAmount returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasFillTotalAmount() bool {
	if o != nil && !isNil(o.FillTotalAmount) {
		return true
	}

	return false
}

// SetFillTotalAmount gets a reference to the given string and assigns it to the FillTotalAmount field.
func (o *MarginTradeDetailInfo) SetFillTotalAmount(v string) {
	o.FillTotalAmount = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *MarginTradeDetailInfo) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetOrderType() string {
	if o == nil || isNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetOrderTypeOk() (*string, bool) {
	if o == nil || isNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasOrderType() bool {
	if o != nil && !isNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *MarginTradeDetailInfo) SetOrderType(v string) {
	o.OrderType = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *MarginTradeDetailInfo) GetSide() string {
	if o == nil || isNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginTradeDetailInfo) GetSideOk() (*string, bool) {
	if o == nil || isNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *MarginTradeDetailInfo) HasSide() bool {
	if o != nil && !isNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *MarginTradeDetailInfo) SetSide(v string) {
	o.Side = &v
}

func (o MarginTradeDetailInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.FeeCcy) {
		toSerialize["feeCcy"] = o.FeeCcy
	}
	if !isNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !isNil(o.FillId) {
		toSerialize["fillId"] = o.FillId
	}
	if !isNil(o.FillPrice) {
		toSerialize["fillPrice"] = o.FillPrice
	}
	if !isNil(o.FillQuantity) {
		toSerialize["fillQuantity"] = o.FillQuantity
	}
	if !isNil(o.FillTotalAmount) {
		toSerialize["fillTotalAmount"] = o.FillTotalAmount
	}
	if !isNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !isNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !isNil(o.Side) {
		toSerialize["side"] = o.Side
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginTradeDetailInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMarginTradeDetailInfo := _MarginTradeDetailInfo{}

	if err = json.Unmarshal(bytes, &varMarginTradeDetailInfo); err == nil {
		*o = MarginTradeDetailInfo(varMarginTradeDetailInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "feeCcy")
		delete(additionalProperties, "fees")
		delete(additionalProperties, "fillId")
		delete(additionalProperties, "fillPrice")
		delete(additionalProperties, "fillQuantity")
		delete(additionalProperties, "fillTotalAmount")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "side")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginTradeDetailInfo struct {
	value *MarginTradeDetailInfo
	isSet bool
}

func (v NullableMarginTradeDetailInfo) Get() *MarginTradeDetailInfo {
	return v.value
}

func (v *NullableMarginTradeDetailInfo) Set(val *MarginTradeDetailInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginTradeDetailInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginTradeDetailInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginTradeDetailInfo(val *MarginTradeDetailInfo) *NullableMarginTradeDetailInfo {
	return &NullableMarginTradeDetailInfo{value: val, isSet: true}
}

func (v NullableMarginTradeDetailInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginTradeDetailInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
