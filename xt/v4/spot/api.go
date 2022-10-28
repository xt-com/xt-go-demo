package spot

import (
	"fmt"
	"sdk/response"
	v4 "sdk/xt/v4"
)

type XTPublicSpotHelper interface {
	// Public
	GetServerTime() *response.APIBody                              // Getting the server time
	GetCoinsInfo() *response.APIBody                               // Get the currency information
	GetMarketConfig(data map[string]interface{}) *response.APIBody // Get market configuration info
	GetAllMarketConfig() *response.APIBody                         // Get configuration information for all pairs
	GetDepth(data map[string]interface{}) *response.APIBody        //  Get the market depth
	GetKline(data map[string]interface{}) *response.APIBody        // Get k-line information
	GetTrades(data map[string]interface{}) *response.APIBody       // Get the most recent transaction
	GetTicker(data map[string]interface{}) *response.APIBody       // Get the latest market information
	GetFullTicker(data map[string]interface{}) *response.APIBody   // Get the latest market information for all currencies
	GetBestTicker(data map[string]interface{}) *response.APIBody   // Get the best market information
	Get24hTicker(data map[string]interface{}) *response.APIBody    // Get information about the 24-hour market
}

type XTPrivateSpotHelper interface {
	// Private
	GetOrder(data map[string]interface{}) *response.APIBody         // Query order information
	GetOrderList(data map[string]interface{}) *response.APIBody     // Query order information
	CancelOrder(orderId string) *response.APIBody                   // Cancel the order
	SendOrder(data map[string]interface{}) *response.APIBody        // order
	GetBatchOrder(data map[string]interface{}) *response.APIBody    // Get orders in bulk
	SendBatchOrder(data map[string]interface{}) *response.APIBody   //  Bulk order
	BatchCancelOrder(data map[string]interface{}) *response.APIBody // Batch cancel order
	GetOpenOrder(data map[string]interface{}) *response.APIBody     // Get the outstanding order
	CancelOpenOrder(data map[string]interface{}) *response.APIBody  // Cancel the outstanding order
	GetHistoryOrder(data map[string]interface{}) *response.APIBody  // Get order history
	GetUserTrade(data map[string]interface{}) *response.APIBody     // Get account transaction information
	GetBalance(data map[string]interface{}) *response.APIBody       // Get the balance
	GetListenKey() *response.APIBody                                // get ListenKey
}

type SignedHttpAPI struct {
	Accesskey string
	Secretkey string
}

func NewSignedHttpAPI(accesskey, secretkey string) *SignedHttpAPI {
	return &SignedHttpAPI{
		Accesskey: accesskey,
		Secretkey: secretkey,
	}
}

func (s *SignedHttpAPI) SetAuthOption(accesskey, secretkey string) {
	s.Accesskey = accesskey
	s.Secretkey = secretkey
}

/**
 *	@Param:
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedHttpAPI) GetListenKey() *response.APIBody {
	path := "/v4/ws-token"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)
	data := map[string]interface{}{}

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("POST", uri, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : orderId	    number	    true
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedHttpAPI) GetOrder(data map[string]interface{}) *response.APIBody {

	path := "/v4/order"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", uri, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : orderId	    number	    true
 *		::param : clientOrderId	string	    false
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedHttpAPI) GetOrderList(data map[string]interface{}) *response.APIBody {
	path := "/v4/order"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", uri, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : orderId	    number	    true
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": {
 *				"cancelId": "6216559590087220004"
 *			}
 *		}
**/
func (s SignedHttpAPI) CancelOrder(orderId string) *response.APIBody {
	path := "/v4/order"
	uri := fmt.Sprintf("%s/%s", path, orderId)
	url := BaseUrl + uri
	auth := NewAuth(s, uri, "DELETE")

	headers, err := auth.createPayload(map[string]interface{}{})
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesJson("DELETE", url, headers, map[string]interface{}{})

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : symbol	    string	    true
 *		::param : clientOrderId	string	    false	                    The longest number is 32 characters
 *		::param : side      	string	    true	                    BUY,SELL
 *		::param : type      	string	    true	                    order type:LIMIT,MARKET
 *		::param : timeInForce   string	    true	                    effective way:GTC, FOK, IOC, GTX
 *		::param : bizType       string	    true	                    SPOT, LEVER
 *		::param : price         number	    false	                    price. Required if it is the LIMIT price; blank if it is the MARKET price
 *		::param : quantity      number	    false	                    quantity. Required if it is the LIMIT price or the order is placed at the market price by quantity
 *		::param : quoteQty      number	    false	                    amount. Required if it is the LIMIT price or the order is the market price when placing an order by amount
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": {
 *				"orderId": "6216559590087220004"
 *			}
 *		}
**/
func (s SignedHttpAPI) SendOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "POST")

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesJson("POST", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : orderId	    string	    true		                order Ids eg: 6216559590087220004,6216559590087220004
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": {
 *				"cancelId": "6216559590087220004"
 *			}
 *		}
**/
func (s SignedHttpAPI) GetBatchOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/batch-order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	        Type	    mandatory    Default	    Description
 *		::param : clientBatchId	    string	    false		                Client batch number
 *		::param : items	            array	    true		                array
 *		::param : item.symbol	    string	    true
 *		::param : item.clientOrderIdstring	    false		                The longest number is 32 characters
 *		::param : item.side         string	    true		                BUY,SELL
 *		::param : item.type         string	    true		                order type:LIMIT,MARKET
 *		::param : item.timeInForce  string	    true		                effective way:GTC, FOK, IOC, GTX
 *		::param : item.bizType      string	    true		                SPOT, LEVER
 *		::param : item.price        number	    false		                price. Required if it is the LIMIT price; blank if it is the MARKET price
 *		::param : item.quantity     number	    false		                quantity. Required if it is the LIMIT price or the order is placed at the market price by quantity
 *		::param : item.quoteQty     number	    false		                amount. Required if it is the LIMIT price or the order is the market price when placing an order by amount
 *	@Return
 *		{
 *		"clientBatchId": "51232",
 *		"items": [
 *				{
 *				"symbol": "BTC_USDT",
 *				"clientOrderId": "16559590087220001",
 *				"side": "BUY",
 *				"type": "LIMIT",
 *				"timeInForce": "GTC",
 *				"bizType": "SPOT",
 *				"price": 40000,
 *				"quantity": 2,
 *				"quoteQty": 80000
 *				}
 *			]
 *		}
**/
func (s SignedHttpAPI) SendBatchOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/batch-order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "POST")

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesJson("POST", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : clientBatchId	string	    false		                client batch id
 *		::param : orderIds	    array	    true		                [6216559590087220004,6216559590087220005]
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": {}
 *		}
**/
func (s SignedHttpAPI) BatchCancelOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/batch-order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "DELETE")

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesJson("DELETE", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : symbol    	string	    false		                Trading pair, if not filled in, represents all
 *		::param : bizType	    string	    false		                SPOT, LEVER
 *		::param : side  	    string	    false		                BUY,SELL
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedHttpAPI) GetOpenOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/open-order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : symbol    	string	    false		                Trading pair, if not filled in, represents all
 *		::param : bizType	    string	    false		                SPOT, LEVER
 *		::param : side  	    string	    false		                BUY,SELL
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": {}
 *		}
**/
func (s SignedHttpAPI) CancelOpenOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/open-order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "DELETE")

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesJson("DELETE", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	        Type	    mandatory    Default	    Description
 *		::param : symbol    	    string	    false		                Trading pair, if not filled in, represents all
 *		::param : bizType	        string	    false		                SPOT, LEVER
 *		::param : side      	    string	    false		                BUY,SELL
 *		::param : type              string	    false		                LIMIT, MARKET
 *		::param : state             string	    false		                PARTIALLY_FILLED,FILLED,CANCELED,REJECTED,EXPIRED
 *		::param : fromId            number	    false		                start id
 *		::param : direction         string	    false		                query direction:PREV, NEXT
 *		::param : limit             number	    false		 20             Limit number, max 100
 *		::param : startTime         number	    false		                eg:1657682804112
 *		::param : endTime           number	    false
 *		::param : hiddenCanceled    number	    bool
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedHttpAPI) GetHistoryOrder(data map[string]interface{}) *response.APIBody {
	path := "/v4/history-order"
	url := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	        Type	    mandatory    Default	    Description
 *		::param : symbol    	    string	    false		                Trading pair, if not filled in, represents all
 *		::param : bizType	        string	    false		                SPOT, LEVER
 *		::param : orderSide      	string	    false		                BUY,SELL
 *		::param : orderType         string	    false		                LIMIT, MARKET
 *		::param : orderId           number	    false
 *		::param : fromId            number	    false		                start id
 *		::param : direction         string	    false		                query direction:PREV, NEXT
 *		::param : limit             number	    false		 20             Limit number, max 100
 *		::param : startTime         number	    false		                eg:1657682804112
 *		::param : endTime           number	    false
 *		::param : hiddenCanceled    number	    bool
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedHttpAPI) GetUserTrade(data map[string]interface{}) *response.APIBody {
	path := "/v4/trade"
	url := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", url, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : currency    	string	    true		                eg:usdt
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": {
 *				"currency": "usdt",
 *				"currencyId": 0,
 *				"frozenAmount": 0,
 *				"availableAmount": 0,
 *				"totalAmount": 0,
 *				"convertBtcAmount": 0  //Converted BTC amount
 *			}
 *		}
**/
func (s SignedHttpAPI) GetBalance(data map[string]interface{}) *response.APIBody {
	path := "/v4/balances"
	url := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", url, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", url, headers, data)

	return rep
}

type PublicHttpAPI struct {
}

/**
 * @Param: None
 * @Return
 * {
 *   "rc": 0,
 *   "mc": "SUCCESS",
 *   "ma": [],
 *   "result": {
 *       "serverTime": 1662435658062
 *    }
 *  }
 */
func (p PublicHttpAPI) GetServerTime() *response.APIBody {
	path := "/v4/public/time"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
* @Param: None
* @Return
* 	{
* 	"rc": 0,
* 	"mc": "string",
* 	"ma": [
* 		{}
* 	],
* 	"result": [
* 			{
* 			"id": 11,  //currency id
* 			"currency": "usdt", //currency name
* 			"fullName": "usdt",  //currency full name
* 			"logo": null,   //currency logo
* 			"cmcLink": null,  //cmc link
* 			"weight": 100,
* 			"maxPrecision": 6,
* 			"depositStatus": 1,  //Recharge status(0 close 1 open)
* 			"withdrawStatus": 1,  //Withdrawal status(0 close 1 open)
* 			"convertEnabled": 1,  //Small asset exchange switch[0=close;1=open]
* 			"transferEnabled": 1  //swipe switch[0=close;1=open]
* 			}
* 		]
* 	}
**/
func (p PublicHttpAPI) GetCoinsInfo() *response.APIBody {
	path := "/v4/public/currencies"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
* @Param: None
* @Return
* 	See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (p PublicHttpAPI) GetAllMarketConfig() *response.APIBody {
	path := "/future/market/v1/public/symbol/list"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *  @Param:
 *       @Desc     Parameter	    Type	    mandatory    Default	    Description
 *       ::param : symbol	    string	    false		          	    trading pair eg:btc_usdt
 *       ::param : symbols	    string	    false		          	    Collection of trading pairs. Priority is higher than symbol. eg: btc_usdt,eth_usdt
 *       ::param : version	    string	    false		                Version number, when the request version number is consistent with the response content version, the list will not be returned, reducing IO eg: 2e14d2cd5czcb2c2af2c1db65078d075
 *   @Return
 *       See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (p PublicHttpAPI) GetMarketConfig(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/symbol"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 * 	@Desc     Parameter	    Type	    mandatory    Default	    Description        Ranges
 * 	::param : symbol	    string	    true		          	    trading pair
 * 	::param : limit	        number	    false		 200                               1,1000
 * @Return
 * 	{
 * 	"rc": 0,
 * 	"mc": "string",
 * 	"ma": [
 * 		{}
 * 	],
 * 	"result": [
 * 			{
 * 			"i": 0,   //ID
 * 			"t": 0,   //transaction time
 * 			"p": "string", //transaction price
 * 			"q": "string",  //transaction quantity
 * 			"v": "string",  //transaction volume
 * 			"b": true   //buyerMaker
 * 			}
 * 		]
 * 	}
**/
func (p PublicHttpAPI) GetDepth(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/depth"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 *    @Desc     Parameter	    Type	    mandatory    Default	    Description
 *    ::param : symbol	    string	    true		          	    trading pair eg:btc_usdt
 *    ::param : interval	    string	    true		                K line type ,1m;3m;5m;15m;30m;1h;2h;4h;6h;8h;12h;1d;3d;1w;1M eg:1m
 *    ::param : startTime     number      false                       start timestamp
 *    ::param : endTime       number      false                       start timestamp
 *    ::param : limit         number      false        100
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "string",
 *		"ma": [
 *			{}
 *		],
 *		"result": [
 *				{
 *				"i": 0,   //ID
 *				"t": 0,   //transaction time
 *				"p": "string", //transaction price
 *				"q": "string",  //transaction quantity
 *				"v": "string",  //transaction volume
 *				"b": true   //buyerMaker
 *				}
 *			]
 *		}
**/
func (p PublicHttpAPI) GetKline(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/kline"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 *     @Desc     Parameter	    Type	    mandatory    Default	    Description
 *     ::param : symbol	    string	    true		          	    trading pair eg:btc_usdt
 *     ::param : limit         number      false        200	        1,1000
 * @Return
 *     {
 *     "rc": 0,
 *     "mc": "string",
 *     "ma": [
 *         {}
 *     ],
 *     "result": [
 *             {
 *             "i": 0,   //ID
 *             "t": 0,   //transaction time
 *             "p": "string", //transaction price
 *             "q": "string",  //transaction quantity
 *             "v": "string",  //transaction volume
 *             "b": true   //buyerMaker
 *             }
 *         ]
 *     }
**/
func (p PublicHttpAPI) GetTrades(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/trade/recent"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 *     @Desc     Parameter	    Type	    mandatory    Default	    Description                 Ranges
 *     ::param : symbol	    string	    false		          	    trading pair eg:btc_usdt
 *     ::param : symbols	    array	    false		          	    Collection of trading pairs. Priority is higher than symbol. eg: btc_usdt,eth_usdt
 * @Return
 *     {
 *     "rc": 0,
 *     "mc": "SUCCESS",
 *     "ma": [],
 *     "result": [
 *             {
 *             "s": "btc_usdt",   //symbol
 *             "p": "9000.0000",   //price
 *             "t": 1661856036925   //time
 *             }
 *         ]
 *     }
**/
func (p PublicHttpAPI) GetTicker(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/ticker/price"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description                 Ranges
 *		::param : symbol	    string	    false		          	    trading pair eg:btc_usdt
 *		::param : symbols	    array	    false		          	    Collection of trading pairs. Priority is higher than symbol. eg: btc_usdt,eth_usdt
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "SUCCESS",
 *		"ma": [],
 *		"result": [
 *				{
 *				"s": "btc_usdt",     //symbol
 *				"t": 1661856036925,  //time
 *				"cv": "0.0000",      //change value
 *				"cr": "0.00",        //change rate
 *				"o": "9000.0000",    //open price
 *				"l": "9000.0000",    //low
 *				"h": "9000.0000",    //high
 *				"c": "9000.0000",    //close price
 *				"q": "0.0136",       //quantity
 *				"v": "122.9940",     //volume
 *				"ap": null,          //asks price(sell one price)
 *				"aq": null,          //asks qty(sell one quantity)
 *				"bp": null,           //bids price(buy one price)
 *				"bq": null           //bids qty(buy one quantity)
 *				}
 *			]
 *		}
**/
func (p PublicHttpAPI) GetFullTicker(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/ticker/full"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description                 Ranges
 *		::param : symbol	    string	    false		          	    trading pair eg:btc_usdt
 *		::param : symbols	    array	    false		          	    Collection of trading pairs. Priority is higher than symbol. eg: btc_usdt,eth_usdt
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "SUCCESS",
 *		"ma": [],
 *		"result": [
 *				{
 *				"s": "btc_usdt",  //symbol
 *				"ap": null,  //asks price(sell one price)
 *				"aq": null,  //asks qty(sell one quantity)
 *				"bp": null,   //bids price(buy one price)
 *				"bq": null    //bids qty(buy one quantity)
 *				}
 *			]
 *		}
**/
func (p PublicHttpAPI) GetBestTicker(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/ticker/book"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description                 Ranges
 *		::param : symbol	    string	    false		          	    trading pair eg:btc_usdt
 *		::param : symbols	    array	    false		          	    Collection of trading pairs. Priority is higher than symbol. eg: btc_usdt,eth_usdt
 *	@Return
 *		{
 *		"rc": 0,
 *		"mc": "SUCCESS",
 *		"ma": [],
 *		"result": [
 *				{
 *				"s": "btc_usdt",   //symbol
 *				"cv": "0.0000",   //price change value
 *				"cr": "0.00",     //price change rate
 *				"o": "9000.0000",   //open price
 *				"l": "9000.0000",   //lowest price
 *				"h": "9000.0000",   //highest price
 *				"c": "9000.0000",   //close price
 *				"q": "0.0136",      //transaction quantity
 *				"v": "122.9940"    //transaction volume
 *				}
 *			]
 *		}
**/
func (p PublicHttpAPI) Get24hTicker(data map[string]interface{}) *response.APIBody {
	path := "/v4/public/ticker/24h"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}
