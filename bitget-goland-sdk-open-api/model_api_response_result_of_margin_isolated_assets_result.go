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

// ApiResponseResultOfMarginIsolatedAssetsResult struct for ApiResponseResultOfMarginIsolatedAssetsResult
type ApiResponseResultOfMarginIsolatedAssetsResult struct {
	// code
	Code *string                     `json:"code,omitempty"`
	Data *MarginIsolatedAssetsResult `json:"data,omitempty"`
	// msg
	Msg *string `json:"msg,omitempty"`
	// requestTime
	RequestTime          *int64 `json:"requestTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiResponseResultOfMarginIsolatedAssetsResult ApiResponseResultOfMarginIsolatedAssetsResult

// NewApiResponseResultOfMarginIsolatedAssetsResult instantiates a new ApiResponseResultOfMarginIsolatedAssetsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponseResultOfMarginIsolatedAssetsResult() *ApiResponseResultOfMarginIsolatedAssetsResult {
	this := ApiResponseResultOfMarginIsolatedAssetsResult{}
	return &this
}

// NewApiResponseResultOfMarginIsolatedAssetsResultWithDefaults instantiates a new ApiResponseResultOfMarginIsolatedAssetsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponseResultOfMarginIsolatedAssetsResultWithDefaults() *ApiResponseResultOfMarginIsolatedAssetsResult {
	this := ApiResponseResultOfMarginIsolatedAssetsResult{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetData() MarginIsolatedAssetsResult {
	if o == nil || isNil(o.Data) {
		var ret MarginIsolatedAssetsResult
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetDataOk() (*MarginIsolatedAssetsResult, bool) {
	if o == nil || isNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MarginIsolatedAssetsResult and assigns it to the Data field.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) SetData(v MarginIsolatedAssetsResult) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetMsg() string {
	if o == nil || isNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetMsgOk() (*string, bool) {
	if o == nil || isNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) HasMsg() bool {
	if o != nil && !isNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) SetMsg(v string) {
	o.Msg = &v
}

// GetRequestTime returns the RequestTime field value if set, zero value otherwise.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetRequestTime() int64 {
	if o == nil || isNil(o.RequestTime) {
		var ret int64
		return ret
	}
	return *o.RequestTime
}

// GetRequestTimeOk returns a tuple with the RequestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) GetRequestTimeOk() (*int64, bool) {
	if o == nil || isNil(o.RequestTime) {
		return nil, false
	}
	return o.RequestTime, true
}

// HasRequestTime returns a boolean if a field has been set.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) HasRequestTime() bool {
	if o != nil && !isNil(o.RequestTime) {
		return true
	}

	return false
}

// SetRequestTime gets a reference to the given int64 and assigns it to the RequestTime field.
func (o *ApiResponseResultOfMarginIsolatedAssetsResult) SetRequestTime(v int64) {
	o.RequestTime = &v
}

func (o ApiResponseResultOfMarginIsolatedAssetsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !isNil(o.RequestTime) {
		toSerialize["requestTime"] = o.RequestTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApiResponseResultOfMarginIsolatedAssetsResult) UnmarshalJSON(bytes []byte) (err error) {
	varApiResponseResultOfMarginIsolatedAssetsResult := _ApiResponseResultOfMarginIsolatedAssetsResult{}

	if err = json.Unmarshal(bytes, &varApiResponseResultOfMarginIsolatedAssetsResult); err == nil {
		*o = ApiResponseResultOfMarginIsolatedAssetsResult(varApiResponseResultOfMarginIsolatedAssetsResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "data")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "requestTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiResponseResultOfMarginIsolatedAssetsResult struct {
	value *ApiResponseResultOfMarginIsolatedAssetsResult
	isSet bool
}

func (v NullableApiResponseResultOfMarginIsolatedAssetsResult) Get() *ApiResponseResultOfMarginIsolatedAssetsResult {
	return v.value
}

func (v *NullableApiResponseResultOfMarginIsolatedAssetsResult) Set(val *ApiResponseResultOfMarginIsolatedAssetsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponseResultOfMarginIsolatedAssetsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponseResultOfMarginIsolatedAssetsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponseResultOfMarginIsolatedAssetsResult(val *ApiResponseResultOfMarginIsolatedAssetsResult) *NullableApiResponseResultOfMarginIsolatedAssetsResult {
	return &NullableApiResponseResultOfMarginIsolatedAssetsResult{value: val, isSet: true}
}

func (v NullableApiResponseResultOfMarginIsolatedAssetsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponseResultOfMarginIsolatedAssetsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}