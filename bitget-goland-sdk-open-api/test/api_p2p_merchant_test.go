/*
Bitget Open API

Testing P2pMerchantApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	openapiclient "github.com/bitget/openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_P2pMerchantApiService(t *testing.T) {

	apiClient := openapiclient.NewAPIClient(openapiclient.NewDefaultConfiguration())

	t.Run("Test P2pMerchantApiService MerchantAdvList", func(t *testing.T) {
		resp, httpRes, err := apiClient.P2pMerchantApi.MerchantAdvList(context.Background()).StartTime("1676260773000").Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		assert.NotNil(t, resp.GetData())
		data := resp.GetData()
		for i, item := range data.GetAdvList() {
			fmt.Printf("%d %v\n", i, item)
			assert.NotEmpty(t, item.GetCoin())
			assert.NotEmpty(t, item.GetAmount())
			assert.NotEmpty(t, item.GetAdvNo())
			assert.NotEmpty(t, item.GetType())
			assert.NotEmpty(t, item.GetFiatCode())
			assert.NotEmpty(t, item.GetPrice())
		}
	})

	t.Run("Test P2pMerchantApiService MerchantInfo", func(t *testing.T) {
		resp, httpRes, err := apiClient.P2pMerchantApi.MerchantInfo(context.Background()).Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		assert.NotNil(t, resp.GetData())
		assert.NotEmpty(t, resp.GetData().Email)
		assert.NotEmpty(t, resp.GetData().AveragePayment)
		assert.NotEmpty(t, resp.GetData().Mobile)
		assert.NotEmpty(t, resp.GetData().NickName)
	})

	t.Run("Test P2pMerchantApiService MerchantList", func(t *testing.T) {
		resp, httpRes, err := apiClient.P2pMerchantApi.MerchantList(context.Background()).Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		assert.NotNil(t, resp.GetData())
		data := resp.GetData()
		for i, item := range data.GetResultList() {
			fmt.Printf("%d %v\n", i, item)
			assert.NotEmpty(t, item.GetMerchantId())
			assert.NotEmpty(t, item.GetIsOnline())
			assert.NotEmpty(t, item.GetNickName())
		}
	})

	t.Run("Test P2pMerchantApiService MerchantOrderList", func(t *testing.T) {
		resp, httpRes, err := apiClient.P2pMerchantApi.MerchantOrderList(context.Background()).StartTime("1680598302000").Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		assert.NotNil(t, resp.GetData())
		data := resp.GetData()
		for i, item := range data.GetOrderList() {
			fmt.Printf("%d %v\n", i, item)
			assert.NotEmpty(t, item.GetCoin())
			assert.NotEmpty(t, item.GetAmount())
			assert.NotEmpty(t, item.GetAdvNo())
			assert.NotEmpty(t, item.GetType())
			assert.NotEmpty(t, item.GetOrderId())
			assert.NotEmpty(t, item.GetPrice())
		}
	})

}
