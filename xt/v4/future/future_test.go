package future

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	accesskey = "xxxxxxxxxxxxxxxxxxxxxxxx"
	secretkey = "yyyyyyyyyyyyyyyyyyyyyyyy"
)

func TestEntrustPlan(t *testing.T) {

	// signedFutureHttpAPI := NewSignedFutureHttpAPI(accesskey, secretkey)

	RepBody = &Body{}
	publicFutureHttpAPI := PublicFutureHttpAPI{}
	rep := publicFutureHttpAPI.GetServerTime()
	json.Unmarshal([]byte(rep.Data), RepBody)
	fmt.Println(RepBody)

	// UpdateEntrustProfitStop
	// data := map[string]interface{}{
	// 	"profitId": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.UpdateEntrustProfitStop(data)
	// fmt.Println(rep)

	// GetEntrustProfitDetail
	// data := map[string]interface{}{
	// 	"profitId": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.GetEntrustProfitDetail(data)
	// fmt.Println(rep)

	// GetEntrustProfitList
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// 	"state":  "NOT_TRIGGERED",
	// }
	// rep = signedFutureHttpAPI.GetEntrustProfitList(data)
	// fmt.Println(rep)

	// CancelAllEntrustProfitStop
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CancelAllEntrustProfitStop(data)
	// fmt.Println(rep)

	// CancelEntrustProfitStop
	// data := map[string]interface{}{
	// 	"profitId": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CancelEntrustProfitStop(data)
	// fmt.Println(rep)

	// CreateEntrustProfitPlan
	// data := map[string]interface{}{
	// 	"symbol":             "btc_usdt",
	// 	"origQty":            "100",
	// 	"triggerProfitPrice": "19300",
	// 	"triggerStopPrice":   "19100",
	// 	"expireTime":         "30",
	// 	"positionSide":       "LONG",
	// }
	// rep = signedFutureHttpAPI.CreateEntrustProfitPlan(data)
	// fmt.Println(rep)

	// GetEntrustHistoryPlanList
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.GetEntrustHistoryPlanList(data)
	// fmt.Println(rep)

	// GetEntrustPlanDetail
	// data := map[string]interface{}{
	// 	"entrustId": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.GetEntrustPlanDetail(data)
	// fmt.Println(rep)

	// GetEntrustPlanList
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// 	"state":  "NOT_TRIGGERED",
	// }
	// rep = signedFutureHttpAPI.GetEntrustPlanList(data)
	// fmt.Println(rep)

	// CancelAllPlan
	// data := map[string]interface{}{
	// 	"symbol":           "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CancelAllPlan(data)
	// fmt.Println(rep)

	// CancelEntrustPlan
	// data := map[string]interface{}{
	// 	"entrustId":           "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CancelEntrustPlan(data)
	// fmt.Println(rep)

	// CollectionList
	// data := map[string]interface{}{
	// 	"symbol":           "btc_usdt",
	// 	"orderSide":        "BUY",
	// 	"price":            "19013",
	// 	"stopPrice":        "19012",
	// 	"entrustType":      "TAKE_PROFIT",
	// 	"origQty":          "100",
	// 	"timeInForce":      "IOC",
	// 	"triggerPriceType": "INDEX_PRICE",
	// 	"positionSide":     "LONG",
	// }
	// rep = signedFutureHttpAPI.CreateEntrustPlan(data)
	// fmt.Println(rep)
}

/**
 * ORDER: 142681429300484160 , 142681428734282624, 142681428642007937
**/
func TestPrivateFutureAPI(t *testing.T) {

	signedFutureHttpAPI := NewSignedFutureHttpAPI(accesskey, secretkey)

	RepBody = &Body{}
	publicFutureHttpAPI := PublicFutureHttpAPI{}
	rep := publicFutureHttpAPI.GetServerTime()
	json.Unmarshal([]byte(rep.Data), RepBody)
	fmt.Println(RepBody)

	// CollectionList
	// rep = signedFutureHttpAPI.CollectionList()
	// fmt.Println(rep)

	// CollectionAdd
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CollectionCancel(data)
	// fmt.Println(rep)

	// CollectionAdd
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CollectionAdd(data)
	// fmt.Println(rep)

	// // AllPositionClose
	// rep = signedFutureHttpAPI.PositionADL()
	// fmt.Println(rep)

	// AllPositionClose
	// rep = signedFutureHttpAPI.AllPositionClose()
	// fmt.Println(rep)

	// UpdatePositionAutoMargin
	// data := map[string]interface{}{
	// 	"symbol":       "btc_usdt",
	// 	"autoMargin":   false,
	// 	"positionSide": "LONG",
	// }
	// rep = signedFutureHttpAPI.UpdatePositionAutoMargin(data)
	// fmt.Println(rep)

	// UpdatePositionMargin
	// data := map[string]interface{}{
	// 	"symbol":       "btc_usdt",
	// 	"type":         "ADD",
	// 	"margin":       10,
	// 	"positionSide": "LONG",
	// }
	// rep = signedFutureHttpAPI.UpdatePositionMargin(data)
	// fmt.Println(rep)

	// GetAdjustLeverage
	// data := map[string]interface{}{
	// 	"symbol":       "btc_usdt",
	// 	"positionSide": "SHORT",
	// 	"leverage":     1,
	// }
	// rep = signedFutureHttpAPI.GetAdjustLeverage(data)
	// fmt.Println(rep)

	// // GetPositionList
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.GetPositionList(data)
	// fmt.Println(rep)

	// GetBalanceList
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.GetBalanceBills(data)
	// fmt.Println(rep)

	// GetBalance
	// data := map[string]interface{}{
	// 	"coin": "usdt",
	// }
	// rep = signedFutureHttpAPI.GetBalance(data)
	// fmt.Println(rep)

	// AccountOpen
	// rep = signedFutureHttpAPI.AccountOpen()
	// fmt.Println(rep)

	// GetListenKey
	// rep = signedFutureHttpAPI.GetListenKey()
	// fmt.Println(rep)

	// GetAccountInfo
	// rep = signedFutureHttpAPI.GetAccountInfo()
	// fmt.Println(rep)

	// CancelBatchOrder
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.CancelBatchOrder(data)
	// fmt.Println(rep)

	// CancelOrder
	// data := map[string]interface{}{
	// 	"orderId": "142681429300484160",
	// }
	// rep = signedFutureHttpAPI.CancelOrder(data)
	// fmt.Println(rep)

	// GetOrderList
	// data := map[string]interface{}{
	// 	"state": "NEW",
	// }
	// rep = signedFutureHttpAPI.GetOrderList(data)
	// fmt.Println(rep)

	// GetOrderDetail
	// data := map[string]interface{}{
	// 	"orderId": "83658649038409920",
	// }
	// rep = signedFutureHttpAPI.GetOrderDetail(data)
	// fmt.Println(rep)

	// GetTradeList
	// order := []interface{}{
	// 	map[string]interface{}{
	// 		"symbol":       "btc_usdt",
	// 		"price":        "19300",
	// 		"orderSide":    "BUY",
	// 		"orderType":    "LIMIT",
	// 		"timeInForce":  "GTC",
	// 		"origQty":      "100",
	// 		"positionSide": "LONG",
	// 	},
	// }
	// data := map[string]interface{}{
	// 	"list": order,
	// }
	// rep = signedFutureHttpAPI.SendBatchOrder(data)
	// fmt.Println(rep)

	// GetHistoryList
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep = signedFutureHttpAPI.GetHistoryList(data)
	// fmt.Println(rep)

	// SendOrder
	for i := 0; i < 1; i++ {
		RepBody = &Body{}
		publicFutureHttpAPI := PublicFutureHttpAPI{}
		rep := publicFutureHttpAPI.GetServerTime()
		json.Unmarshal([]byte(rep.Data), RepBody)
		fmt.Println(RepBody)

		time.Sleep(time.Microsecond * 200)
		data := map[string]interface{}{
			"symbol":       "btc_usdt",
			"price":        "19300",
			"orderSide":    "SELL",
			"orderType":    "LIMIT",
			"timeInForce":  "GTC",
			"origQty":      "1",
			"positionSide": "SHORT",
		}

		rep = signedFutureHttpAPI.SendOrder(data)
		fmt.Println(rep)
	}

	// GetAccountInfo
	// signedFutureHttpAPI := NewSignedFutureHttpAPI(accesskey, secretkey)
	// rep := signedFutureHttpAPI.GetAccountInfo()
	// fmt.Println(rep)

}

func TestPublicFutureAPI(t *testing.T) {

	// GetOpenInterest
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetOpenInterest(data)
	// fmt.Println(rep)

	// GetRiskBalance
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetRiskBalance(data)
	// fmt.Println(rep)

	// GetFundRate
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetFundRateRecord(data)
	// fmt.Println(rep)

	// GetAggTicker
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetAllAggTicker()
	// fmt.Println(rep)

	// GetAggTicker
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetAggTicker(data)
	// fmt.Println(rep)

	// GetKline
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol":   "btc_usdt",
	// 	"interval": "1m",
	// }
	// rep := publicFutureHttpAPI.GetKline(data)
	// fmt.Println(rep)

	// GetIndexPrice
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetAllMarketPrice()
	// fmt.Println(rep)

	// GetMarketPrice
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetMarketPrice(data)
	// fmt.Println(rep)

	// GetIndexPrice
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetAllIndexPrice()
	// fmt.Println(rep)

	// GetIndexPrice
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetIndexPrice(data)
	// fmt.Println(rep)

	// GetDepth
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// 	"level":  10,
	// }
	// rep := publicFutureHttpAPI.GetDepth(data)
	// fmt.Println(rep)

	// // GetMarketDeal
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// 	"num":    10,
	// }
	// rep := publicFutureHttpAPI.GetMarketDeal(data)
	// fmt.Println(rep)

	// GetMarketTickers
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetMarketTickers()
	// fmt.Println(rep)

	// GetMarketTicker
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// rep := publicFutureHttpAPI.GetMarketTicker(data)
	// fmt.Println(rep)

	// GetLeverageDetailList
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetLeverageDetailList()
	// fmt.Println(rep)

	// GetLeverageDetail
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetLeverageDetail(data)
	// fmt.Println(rep)

	// GetMarketConfig
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetAllMarketConfig()
	// fmt.Println(rep)

	// GetMarketConfig
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetMarketConfig(data)
	// fmt.Println(rep)

	// GetCoinsInfo
	// publicFutureHttpAPI := PublicFutureHttpAPI{}
	// rep := publicFutureHttpAPI.GetCoinsInfo()
	// fmt.Println(rep)
}
