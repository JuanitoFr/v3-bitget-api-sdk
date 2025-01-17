# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

### Code Demo
```golang
package openapi

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    openapiclient "github.com/bitget/openapi"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "os"
    "testing"
)

func Test_openapi_ApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test Api", func(t *testing.T) {
        configuration := openapiclient.NewConfiguration()
        configuration.AddDefaultHeader("ACCESS-KEY", "xxx")
        configuration.AddDefaultHeader("ACCESS-PASSPHRASE", "xxx")
        configuration.AddDefaultHeader("SECRET-KEY", "xxx")
        configuration.Host = "api.bitget.com"
        configuration.Scheme = "https"

        apiClient := openapiclient.NewAPIClient(configuration)
        param := *openapiclient.NewMixPlaceOrderRequest("USDT", "market", "open_long", float32(1.0), "BTCUSDT_UMCBL") // MixPlaceOrderRequest | param
        resp, r, err := apiClient.MixOrderApi.PlaceOrder(context.Background()).Param(param).Execute()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error when calling `MixOrderApi.PlaceOrder``: %v\n", err)
            fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        }
        bs, _ := json.Marshal(resp)
        var out bytes.Buffer
        json.Indent(&out, bs, "", "\t")
        fmt.Printf("student=%v\n", out.String())

        resp1, r, err := apiClient.MixOrderApi.MarginCoinCurrent(context.Background()).ProductType("umcbl").MarginCoin("USDT").Execute()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error when calling `MixOrderApi.GetMarginCoinCurrent``: %v\n", err)
            fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        }
        bs1, _ := json.Marshal(resp1)
        var out1 bytes.Buffer
        json.Indent(&out1, bs1, "", "\t")
        fmt.Printf("student=%v\n", out1.String())

        resp2, r, err := apiClient.MixOrderApi.MixOrderFills(context.Background()).Symbol("BTCUSDT_UMCBL").StartTime("1671402175000").EndTime("1673517445000").Execute()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error when calling `MixOrderApi.GetMarginCoinCurrent``: %v\n", err)
            fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        }
        bs2, _ := json.Marshal(resp2)
        var out2 bytes.Buffer
        json.Indent(&out2, bs2, "", "\t")
        fmt.Printf("student=%v\n", out2.String())
    })
}
````

## Documentation for API Endpoints

All URIs are relative to *https://api.bitget.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MarginCrossAccountApi* | [**MarginCrossAccountAssets**](docs/MarginCrossAccountApi.md#margincrossaccountassets) | **Get** /api/margin/v1/cross/account/assets | assets
*MarginCrossAccountApi* | [**MarginCrossAccountBorrow**](docs/MarginCrossAccountApi.md#margincrossaccountborrow) | **Post** /api/margin/v1/cross/account/borrow | borrow
*MarginCrossAccountApi* | [**MarginCrossAccountMaxBorrowableAmount**](docs/MarginCrossAccountApi.md#margincrossaccountmaxborrowableamount) | **Post** /api/margin/v1/cross/account/maxBorrowableAmount | maxBorrowableAmount
*MarginCrossAccountApi* | [**MarginCrossAccountMaxTransferOutAmount**](docs/MarginCrossAccountApi.md#margincrossaccountmaxtransferoutamount) | **Get** /api/margin/v1/cross/account/maxTransferOutAmount | maxTransferOutAmount
*MarginCrossAccountApi* | [**MarginCrossAccountRepay**](docs/MarginCrossAccountApi.md#margincrossaccountrepay) | **Post** /api/margin/v1/cross/account/repay | repay
*MarginCrossAccountApi* | [**MarginCrossAccountRiskRate**](docs/MarginCrossAccountApi.md#margincrossaccountriskrate) | **Get** /api/margin/v1/cross/account/riskRate | riskRate
*MarginCrossAccountApi* | [**Void**](docs/MarginCrossAccountApi.md#void) | **Get** /api/margin/v1/cross/account/void | void
*MarginCrossBorrowApi* | [**CrossLoanList**](docs/MarginCrossBorrowApi.md#crossloanlist) | **Get** /api/margin/v1/cross/loan/list | list
*MarginCrossFinflowApi* | [**CrossFinList**](docs/MarginCrossFinflowApi.md#crossfinlist) | **Get** /api/margin/v1/cross/fin/list | list
*MarginCrossInterestApi* | [**CrossInterestList**](docs/MarginCrossInterestApi.md#crossinterestlist) | **Get** /api/margin/v1/cross/interest/list | list
*MarginCrossLiquidationApi* | [**CrossLiquidationList**](docs/MarginCrossLiquidationApi.md#crossliquidationlist) | **Get** /api/margin/v1/cross/liquidation/list | list
*MarginCrossOrderApi* | [**MarginCrossBatchCancelOrder**](docs/MarginCrossOrderApi.md#margincrossbatchcancelorder) | **Post** /api/margin/v1/cross/order/batchCancelOrder | batchCancelOrder
*MarginCrossOrderApi* | [**MarginCrossBatchPlaceOrder**](docs/MarginCrossOrderApi.md#margincrossbatchplaceorder) | **Post** /api/margin/v1/cross/order/batchPlaceOrder | batchPlaceOrder
*MarginCrossOrderApi* | [**MarginCrossCancelOrder**](docs/MarginCrossOrderApi.md#margincrosscancelorder) | **Post** /api/margin/v1/cross/order/cancelOrder | cancelOrder
*MarginCrossOrderApi* | [**MarginCrossFills**](docs/MarginCrossOrderApi.md#margincrossfills) | **Get** /api/margin/v1/cross/order/fills | fills
*MarginCrossOrderApi* | [**MarginCrossHistoryOrders**](docs/MarginCrossOrderApi.md#margincrosshistoryorders) | **Get** /api/margin/v1/cross/order/history | history
*MarginCrossOrderApi* | [**MarginCrossOpenOrders**](docs/MarginCrossOrderApi.md#margincrossopenorders) | **Get** /api/margin/v1/cross/order/openOrders | openOrders
*MarginCrossOrderApi* | [**MarginCrossPlaceOrder**](docs/MarginCrossOrderApi.md#margincrossplaceorder) | **Post** /api/margin/v1/cross/order/placeOrder | placeOrder
*MarginCrossPublicApi* | [**MarginCrossPublicInterestRateAndLimit**](docs/MarginCrossPublicApi.md#margincrosspublicinterestrateandlimit) | **Get** /api/margin/v1/cross/public/interestRateAndLimit | interestRateAndLimit
*MarginCrossPublicApi* | [**MarginCrossPublicTierData**](docs/MarginCrossPublicApi.md#margincrosspublictierdata) | **Get** /api/margin/v1/cross/public/tierData | tierData
*MarginCrossRepayApi* | [**CrossRepayList**](docs/MarginCrossRepayApi.md#crossrepaylist) | **Get** /api/margin/v1/cross/repay/list | list
*MarginIsolatedAccountApi* | [**MarginIsolatedAccountAssets**](docs/MarginIsolatedAccountApi.md#marginisolatedaccountassets) | **Get** /api/margin/v1/isolated/account/assets | assets
*MarginIsolatedAccountApi* | [**MarginIsolatedAccountBorrow**](docs/MarginIsolatedAccountApi.md#marginisolatedaccountborrow) | **Post** /api/margin/v1/isolated/account/borrow | borrow
*MarginIsolatedAccountApi* | [**MarginIsolatedAccountMaxBorrowableAmount**](docs/MarginIsolatedAccountApi.md#marginisolatedaccountmaxborrowableamount) | **Post** /api/margin/v1/isolated/account/maxBorrowableAmount | maxBorrowableAmount
*MarginIsolatedAccountApi* | [**MarginIsolatedAccountMaxTransferOutAmount**](docs/MarginIsolatedAccountApi.md#marginisolatedaccountmaxtransferoutamount) | **Get** /api/margin/v1/isolated/account/maxTransferOutAmount | maxTransferOutAmount
*MarginIsolatedAccountApi* | [**MarginIsolatedAccountRepay**](docs/MarginIsolatedAccountApi.md#marginisolatedaccountrepay) | **Post** /api/margin/v1/isolated/account/repay | repay
*MarginIsolatedAccountApi* | [**MarginIsolatedAccountRiskRate**](docs/MarginIsolatedAccountApi.md#marginisolatedaccountriskrate) | **Post** /api/margin/v1/isolated/account/riskRate | riskRate
*MarginIsolatedBorrowApi* | [**IsolatedLoanList**](docs/MarginIsolatedBorrowApi.md#isolatedloanlist) | **Get** /api/margin/v1/isolated/loan/list | list
*MarginIsolatedFinflowApi* | [**IsolatedFinList**](docs/MarginIsolatedFinflowApi.md#isolatedfinlist) | **Get** /api/margin/v1/isolated/fin/list | list
*MarginIsolatedInterestApi* | [**IsolatedInterestList**](docs/MarginIsolatedInterestApi.md#isolatedinterestlist) | **Get** /api/margin/v1/isolated/interest/list | list
*MarginIsolatedLiquidationApi* | [**IsolatedLiquidationList**](docs/MarginIsolatedLiquidationApi.md#isolatedliquidationlist) | **Get** /api/margin/v1/isolated/liquidation/list | list
*MarginIsolatedOrderApi* | [**MarginIsolatedBatchCancelOrder**](docs/MarginIsolatedOrderApi.md#marginisolatedbatchcancelorder) | **Post** /api/margin/v1/isolated/order/batchCancelOrder | batchCancelOrder
*MarginIsolatedOrderApi* | [**MarginIsolatedBatchPlaceOrder**](docs/MarginIsolatedOrderApi.md#marginisolatedbatchplaceorder) | **Post** /api/margin/v1/isolated/order/batchPlaceOrder | batchPlaceOrder
*MarginIsolatedOrderApi* | [**MarginIsolatedCancelOrder**](docs/MarginIsolatedOrderApi.md#marginisolatedcancelorder) | **Post** /api/margin/v1/isolated/order/cancelOrder | cancelOrder
*MarginIsolatedOrderApi* | [**MarginIsolatedFills**](docs/MarginIsolatedOrderApi.md#marginisolatedfills) | **Get** /api/margin/v1/isolated/order/fills | fills
*MarginIsolatedOrderApi* | [**MarginIsolatedHistoryOrders**](docs/MarginIsolatedOrderApi.md#marginisolatedhistoryorders) | **Get** /api/margin/v1/isolated/order/history | history
*MarginIsolatedOrderApi* | [**MarginIsolatedOpenOrders**](docs/MarginIsolatedOrderApi.md#marginisolatedopenorders) | **Get** /api/margin/v1/isolated/order/openOrders | openOrders
*MarginIsolatedOrderApi* | [**MarginIsolatedPlaceOrder**](docs/MarginIsolatedOrderApi.md#marginisolatedplaceorder) | **Post** /api/margin/v1/isolated/order/placeOrder | placeOrder
*MarginIsolatedPublicApi* | [**MarginIsolatedPublicInterestRateAndLimit**](docs/MarginIsolatedPublicApi.md#marginisolatedpublicinterestrateandlimit) | **Get** /api/margin/v1/isolated/public/interestRateAndLimit | interestRateAndLimit
*MarginIsolatedPublicApi* | [**MarginIsolatedPublicTierData**](docs/MarginIsolatedPublicApi.md#marginisolatedpublictierdata) | **Get** /api/margin/v1/isolated/public/tierData | tierData
*MarginIsolatedRepayApi* | [**IsolateRepayList**](docs/MarginIsolatedRepayApi.md#isolaterepaylist) | **Get** /api/margin/v1/isolated/repay/list | list
*MarginPublicApi* | [**MarginPublicCurrencies**](docs/MarginPublicApi.md#marginpubliccurrencies) | **Get** /api/margin/v1/public/currencies | currencies
*P2pMerchantApi* | [**MerchantAdvList**](docs/P2pMerchantApi.md#merchantadvlist) | **Get** /api/p2p/v1/merchant/advList | advList
*P2pMerchantApi* | [**MerchantInfo**](docs/P2pMerchantApi.md#merchantinfo) | **Get** /api/p2p/v1/merchant/merchantInfo | merchantInfo
*P2pMerchantApi* | [**MerchantList**](docs/P2pMerchantApi.md#merchantlist) | **Get** /api/p2p/v1/merchant/merchantList | merchantList
*P2pMerchantApi* | [**MerchantOrderList**](docs/P2pMerchantApi.md#merchantorderlist) | **Get** /api/p2p/v1/merchant/orderList | orderList


## Documentation For Models

 - [ApiResponseResultOfListOfMarginCrossAssetsPopulationResult](docs/ApiResponseResultOfListOfMarginCrossAssetsPopulationResult.md)
 - [ApiResponseResultOfListOfMarginCrossLevelResult](docs/ApiResponseResultOfListOfMarginCrossLevelResult.md)
 - [ApiResponseResultOfListOfMarginCrossRateAndLimitResult](docs/ApiResponseResultOfListOfMarginCrossRateAndLimitResult.md)
 - [ApiResponseResultOfListOfMarginIsolatedAssetsPopulationResult](docs/ApiResponseResultOfListOfMarginIsolatedAssetsPopulationResult.md)
 - [ApiResponseResultOfListOfMarginIsolatedAssetsRiskResult](docs/ApiResponseResultOfListOfMarginIsolatedAssetsRiskResult.md)
 - [ApiResponseResultOfListOfMarginIsolatedLevelResult](docs/ApiResponseResultOfListOfMarginIsolatedLevelResult.md)
 - [ApiResponseResultOfListOfMarginIsolatedRateAndLimitResult](docs/ApiResponseResultOfListOfMarginIsolatedRateAndLimitResult.md)
 - [ApiResponseResultOfListOfMarginSystemResult](docs/ApiResponseResultOfListOfMarginSystemResult.md)
 - [ApiResponseResultOfMarginBatchCancelOrderResult](docs/ApiResponseResultOfMarginBatchCancelOrderResult.md)
 - [ApiResponseResultOfMarginBatchPlaceOrderResult](docs/ApiResponseResultOfMarginBatchPlaceOrderResult.md)
 - [ApiResponseResultOfMarginCrossAssetsResult](docs/ApiResponseResultOfMarginCrossAssetsResult.md)
 - [ApiResponseResultOfMarginCrossAssetsRiskResult](docs/ApiResponseResultOfMarginCrossAssetsRiskResult.md)
 - [ApiResponseResultOfMarginCrossBorrowLimitResult](docs/ApiResponseResultOfMarginCrossBorrowLimitResult.md)
 - [ApiResponseResultOfMarginCrossFinFlowResult](docs/ApiResponseResultOfMarginCrossFinFlowResult.md)
 - [ApiResponseResultOfMarginCrossMaxBorrowResult](docs/ApiResponseResultOfMarginCrossMaxBorrowResult.md)
 - [ApiResponseResultOfMarginCrossRepayResult](docs/ApiResponseResultOfMarginCrossRepayResult.md)
 - [ApiResponseResultOfMarginInterestInfoResult](docs/ApiResponseResultOfMarginInterestInfoResult.md)
 - [ApiResponseResultOfMarginIsolatedAssetsResult](docs/ApiResponseResultOfMarginIsolatedAssetsResult.md)
 - [ApiResponseResultOfMarginIsolatedBorrowLimitResult](docs/ApiResponseResultOfMarginIsolatedBorrowLimitResult.md)
 - [ApiResponseResultOfMarginIsolatedFinFlowResult](docs/ApiResponseResultOfMarginIsolatedFinFlowResult.md)
 - [ApiResponseResultOfMarginIsolatedInterestInfoResult](docs/ApiResponseResultOfMarginIsolatedInterestInfoResult.md)
 - [ApiResponseResultOfMarginIsolatedLiquidationInfoResult](docs/ApiResponseResultOfMarginIsolatedLiquidationInfoResult.md)
 - [ApiResponseResultOfMarginIsolatedLoanInfoResult](docs/ApiResponseResultOfMarginIsolatedLoanInfoResult.md)
 - [ApiResponseResultOfMarginIsolatedMaxBorrowResult](docs/ApiResponseResultOfMarginIsolatedMaxBorrowResult.md)
 - [ApiResponseResultOfMarginIsolatedRepayInfoResult](docs/ApiResponseResultOfMarginIsolatedRepayInfoResult.md)
 - [ApiResponseResultOfMarginIsolatedRepayResult](docs/ApiResponseResultOfMarginIsolatedRepayResult.md)
 - [ApiResponseResultOfMarginLiquidationInfoResult](docs/ApiResponseResultOfMarginLiquidationInfoResult.md)
 - [ApiResponseResultOfMarginLoanInfoResult](docs/ApiResponseResultOfMarginLoanInfoResult.md)
 - [ApiResponseResultOfMarginOpenOrderInfoResult](docs/ApiResponseResultOfMarginOpenOrderInfoResult.md)
 - [ApiResponseResultOfMarginPlaceOrderResult](docs/ApiResponseResultOfMarginPlaceOrderResult.md)
 - [ApiResponseResultOfMarginRepayInfoResult](docs/ApiResponseResultOfMarginRepayInfoResult.md)
 - [ApiResponseResultOfMarginTradeDetailInfoResult](docs/ApiResponseResultOfMarginTradeDetailInfoResult.md)
 - [ApiResponseResultOfMerchantAdvResult](docs/ApiResponseResultOfMerchantAdvResult.md)
 - [ApiResponseResultOfMerchantInfoResult](docs/ApiResponseResultOfMerchantInfoResult.md)
 - [ApiResponseResultOfMerchantOrderResult](docs/ApiResponseResultOfMerchantOrderResult.md)
 - [ApiResponseResultOfMerchantPersonInfo](docs/ApiResponseResultOfMerchantPersonInfo.md)
 - [ApiResponseResultOfVoid](docs/ApiResponseResultOfVoid.md)
 - [FiatPaymentDetailInfo](docs/FiatPaymentDetailInfo.md)
 - [FiatPaymentInfo](docs/FiatPaymentInfo.md)
 - [MarginBatchCancelOrderRequest](docs/MarginBatchCancelOrderRequest.md)
 - [MarginBatchCancelOrderResult](docs/MarginBatchCancelOrderResult.md)
 - [MarginBatchOrdersRequest](docs/MarginBatchOrdersRequest.md)
 - [MarginBatchPlaceOrderFailureResult](docs/MarginBatchPlaceOrderFailureResult.md)
 - [MarginBatchPlaceOrderResult](docs/MarginBatchPlaceOrderResult.md)
 - [MarginCancelOrderFailureResult](docs/MarginCancelOrderFailureResult.md)
 - [MarginCancelOrderRequest](docs/MarginCancelOrderRequest.md)
 - [MarginCancelOrderResult](docs/MarginCancelOrderResult.md)
 - [MarginCrossAssetsPopulationResult](docs/MarginCrossAssetsPopulationResult.md)
 - [MarginCrossAssetsResult](docs/MarginCrossAssetsResult.md)
 - [MarginCrossAssetsRiskResult](docs/MarginCrossAssetsRiskResult.md)
 - [MarginCrossBorrowLimitResult](docs/MarginCrossBorrowLimitResult.md)
 - [MarginCrossFinFlowInfo](docs/MarginCrossFinFlowInfo.md)
 - [MarginCrossFinFlowResult](docs/MarginCrossFinFlowResult.md)
 - [MarginCrossLevelResult](docs/MarginCrossLevelResult.md)
 - [MarginCrossLimitRequest](docs/MarginCrossLimitRequest.md)
 - [MarginCrossMaxBorrowRequest](docs/MarginCrossMaxBorrowRequest.md)
 - [MarginCrossMaxBorrowResult](docs/MarginCrossMaxBorrowResult.md)
 - [MarginCrossRateAndLimitResult](docs/MarginCrossRateAndLimitResult.md)
 - [MarginCrossRepayRequest](docs/MarginCrossRepayRequest.md)
 - [MarginCrossRepayResult](docs/MarginCrossRepayResult.md)
 - [MarginCrossVipResult](docs/MarginCrossVipResult.md)
 - [MarginInterestInfo](docs/MarginInterestInfo.md)
 - [MarginInterestInfoResult](docs/MarginInterestInfoResult.md)
 - [MarginIsolatedAssetsPopulationResult](docs/MarginIsolatedAssetsPopulationResult.md)
 - [MarginIsolatedAssetsResult](docs/MarginIsolatedAssetsResult.md)
 - [MarginIsolatedAssetsRiskRequest](docs/MarginIsolatedAssetsRiskRequest.md)
 - [MarginIsolatedAssetsRiskResult](docs/MarginIsolatedAssetsRiskResult.md)
 - [MarginIsolatedBorrowLimitResult](docs/MarginIsolatedBorrowLimitResult.md)
 - [MarginIsolatedFinFlowInfo](docs/MarginIsolatedFinFlowInfo.md)
 - [MarginIsolatedFinFlowResult](docs/MarginIsolatedFinFlowResult.md)
 - [MarginIsolatedInterestInfo](docs/MarginIsolatedInterestInfo.md)
 - [MarginIsolatedInterestInfoResult](docs/MarginIsolatedInterestInfoResult.md)
 - [MarginIsolatedLevelResult](docs/MarginIsolatedLevelResult.md)
 - [MarginIsolatedLimitRequest](docs/MarginIsolatedLimitRequest.md)
 - [MarginIsolatedLiquidationInfo](docs/MarginIsolatedLiquidationInfo.md)
 - [MarginIsolatedLiquidationInfoResult](docs/MarginIsolatedLiquidationInfoResult.md)
 - [MarginIsolatedLoanInfo](docs/MarginIsolatedLoanInfo.md)
 - [MarginIsolatedLoanInfoResult](docs/MarginIsolatedLoanInfoResult.md)
 - [MarginIsolatedMaxBorrowRequest](docs/MarginIsolatedMaxBorrowRequest.md)
 - [MarginIsolatedMaxBorrowResult](docs/MarginIsolatedMaxBorrowResult.md)
 - [MarginIsolatedRateAndLimitResult](docs/MarginIsolatedRateAndLimitResult.md)
 - [MarginIsolatedRepayInfo](docs/MarginIsolatedRepayInfo.md)
 - [MarginIsolatedRepayInfoResult](docs/MarginIsolatedRepayInfoResult.md)
 - [MarginIsolatedRepayRequest](docs/MarginIsolatedRepayRequest.md)
 - [MarginIsolatedRepayResult](docs/MarginIsolatedRepayResult.md)
 - [MarginIsolatedVipResult](docs/MarginIsolatedVipResult.md)
 - [MarginLiquidationInfo](docs/MarginLiquidationInfo.md)
 - [MarginLiquidationInfoResult](docs/MarginLiquidationInfoResult.md)
 - [MarginLoanInfo](docs/MarginLoanInfo.md)
 - [MarginLoanInfoResult](docs/MarginLoanInfoResult.md)
 - [MarginOpenOrderInfoResult](docs/MarginOpenOrderInfoResult.md)
 - [MarginOrderInfo](docs/MarginOrderInfo.md)
 - [MarginOrderRequest](docs/MarginOrderRequest.md)
 - [MarginPlaceOrderResult](docs/MarginPlaceOrderResult.md)
 - [MarginRepayInfo](docs/MarginRepayInfo.md)
 - [MarginRepayInfoResult](docs/MarginRepayInfoResult.md)
 - [MarginSystemResult](docs/MarginSystemResult.md)
 - [MarginTradeDetailInfo](docs/MarginTradeDetailInfo.md)
 - [MarginTradeDetailInfoResult](docs/MarginTradeDetailInfoResult.md)
 - [MerchantAdvInfo](docs/MerchantAdvInfo.md)
 - [MerchantAdvResult](docs/MerchantAdvResult.md)
 - [MerchantAdvUserLimitInfo](docs/MerchantAdvUserLimitInfo.md)
 - [MerchantInfo](docs/MerchantInfo.md)
 - [MerchantInfoResult](docs/MerchantInfoResult.md)
 - [MerchantOrderInfo](docs/MerchantOrderInfo.md)
 - [MerchantOrderPaymentInfo](docs/MerchantOrderPaymentInfo.md)
 - [MerchantOrderResult](docs/MerchantOrderResult.md)
 - [MerchantPersonInfo](docs/MerchantPersonInfo.md)
 - [OrderPaymentDetailInfo](docs/OrderPaymentDetailInfo.md)


## Documentation For Authorization



### ACCESS_KEY

- **Type**: API key
- **API key parameter name**: ACCESS-KEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ACCESS-KEY and passed in as the auth context for each request.


### ACCESS_PASSPHRASE

- **Type**: API key
- **API key parameter name**: ACCESS-PASSPHRASE
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ACCESS-PASSPHRASE and passed in as the auth context for each request.


### ACCESS_SIGN

- **Type**: API key
- **API key parameter name**: ACCESS-SIGN
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ACCESS-SIGN and passed in as the auth context for each request.


### ACCESS_TIMESTAMP

- **Type**: API key
- **API key parameter name**: ACCESS-TIMESTAMP
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ACCESS-TIMESTAMP and passed in as the auth context for each request.


### SECRET_KEY

- **Type**: API key
- **API key parameter name**: SECRET-KEY
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: SECRET-KEY and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



