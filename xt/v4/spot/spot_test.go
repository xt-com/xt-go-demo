package spot

import (
	"fmt"
	"testing"
)

const (
	accesskey = ""
	secretkey = ""
)

func TestSignAPI(t *testing.T) {

	// -------test--------
	req := SignedHttpAPI{}
	req.SetAuthOption(accesskey, secretkey)
	// ---------------

	// res := req.GetListenKey()
	// fmt.Println(res)

	// get_balance
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"currency": "usdt",
	// }
	// res := req.GetBalance(data)
	// fmt.Println(res)

	// get_trade
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetUserTrade(data)
	// fmt.Println(res)

	// get_history_order
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetHistoryOrder(data)
	// fmt.Println(res)

	// get_open_order
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.CancelOpenOrder(data)
	// fmt.Println(res)

	// get_open_order
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetOpenOrder(data)
	// fmt.Println(res)

	// batch_cancel_order
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"orderIds": []string{"140600606177600768"},
	// }
	// res := req.BatchCancelOrder(data)
	// fmt.Println(res)

	// batch_order
	// req := NewSignedHttpAPI(accesskey, secretkey)

	order := []map[string]string{
		map[string]string{"symbol": "btc_usdt", "price": "18000", "quantity": "0.001",
			"side": "BUY", "type": "LIMIT", "timeInForce": "GTC", "bizType": "SPOT"},
		map[string]string{"symbol": "btc_usdt", "price": "18001", "quantity": "0.001",
			"side": "BUY", "type": "LIMIT", "timeInForce": "GTC", "bizType": "SPOT"},
	}

	data := map[string]interface{}{
		"items": order,
	}
	res := req.SendBatchOrder(data)
	fmt.Println(res)

	// get_batch_order
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]interface{}{
	// 	"orderIds": "139865641251887680",
	// }
	// res := req.GetBatchOrder(data)
	// fmt.Println(res)

	// send - order
	// req := NewSignedHttpAPI(accesskey, secretkey)

	// data := map[string]interface{}{
	// 	"symbol":      "btc_usdt",
	// 	"side":        "BUY",
	// 	"type":        "LIMIT",
	// 	"timeInForce": "GTC",
	// 	"bizType":     "SPOT",
	// 	"price":       "18570",
	// 	"quantity":    "0.1",
	// }
	// res := req.SendOrder(data)
	// fmt.Println(res)

	// cancel - order
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// orderId := "140541112005403904"
	// res := req.CancelOrder(orderId)
	// fmt.Println(res)

	// get-order-list
	// req := NewSignedHttpAPI(accesskey, secretkey)
	// data := map[string]string{
	// 	"orderId": "140541112005403904",
	// }

	// res := req.GetOrderList(data)
	// fmt.Println(res)

	// get-order
	// req := NewSignedHttpAPI(accesskey, secretkey)

	// data := map[string]interface{}{
	// 	"orderId": "140541112005403904",
	// }

	// res := req.GetOrder(data)
	// fmt.Println(res)
}

func TestAPi(t *testing.T) {

	// get-24h-ticker
	// req := PublicHttpAPI{}
	// data := map[string]interface{}{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.Get24hTicker(data)
	// fmt.Println(res)

	// get-full-ticker
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetBestTicker(data)
	// fmt.Println(res)

	// get-full-ticker
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetFullTicker(data)
	// fmt.Println(res)

	// get-ticker
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetTicker(data)
	// fmt.Println(res)

	// get-trades
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetTrades(data)
	// fmt.Println(res)

	// get - server - time
	// req := PublicHttpAPI{}
	// res := req.GetServerTime()
	// fmt.Println(res)

	// get-coins-info
	// req := PublicHttpAPI{}
	// res := req.GetCoinsInfo()
	// fmt.Println(res)

	// get-market-config
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetMarketConfig(data)
	// fmt.Println(res)

	// get-depth
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol": "btc_usdt",
	// }
	// res := req.GetDepth(data)
	// fmt.Println(res)

	// get-kline
	// req := PublicHttpAPI{}
	// data := map[string]string{
	// 	"symbol":   "btc_usdt",
	// 	"interval": "1m",
	// }
	// res := req.GetKline(data)
	// fmt.Println(res)
}
