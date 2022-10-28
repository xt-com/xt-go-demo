package future

import (
	"sdk/response"
	v4 "sdk/xt/v4"
)

type XTPublicFutureHelper interface {
	// Public

	GetServerTime() *response.APIBody                                // Getting the server time
	GetCoinsInfo() *response.APIBody                                 // Get the currency of the transaction pair
	GetAllMarketConfig() *response.APIBody                           // Get configuration information for all transaction pairs
	GetMarketConfig(data map[string]interface{}) *response.APIBody   // Get configuration information for all transaction pairs
	GetLeverageDetail(data map[string]interface{}) *response.APIBody // Query a single transaction pair for leverage stratification
	GetLeverageDetailList() *response.APIBody                        // Query all trading pairs for leverage stratification
	GetMarketTicker(data map[string]interface{}) *response.APIBody   // Get market information for a specified trading pair
	GetMarketTickers() *response.APIBody                             // Get market information for all trading pairs
	GetMarketDeal(data map[string]interface{}) *response.APIBody     // Get the latest transaction information of the trade pair
	GetDepth(data map[string]interface{}) *response.APIBody          // Get the depth information of the transaction pairs
	GetIndexPrice(data map[string]interface{}) *response.APIBody     // Get the index price for a single trading pair
	GetAllIndexPrice() *response.APIBody                             // Get the index price for all trading pairs
	GetMarketPrice(data map[string]interface{}) *response.APIBody    // Get the tagged price for a single transaction pair
	GetAllMarketPrice() *response.APIBody                            // Get the tag price for all transaction pairs
	GetKline(data map[string]interface{}) *response.APIBody          // Get the k-line information of the trading pair
	GetAggTicker(data map[string]interface{}) *response.APIBody      // Get aggregate market information for a specified trade pair
	GetAllAggTicker() *response.APIBody                              // Get aggregate market information for all trading pairs
	GetFundRate(data map[string]interface{}) *response.APIBody       // Access to funds rate
	GetFundRateRecord(data map[string]interface{}) *response.APIBody // Get funding rate records
	GetRiskBalance(data map[string]interface{}) *response.APIBody    // Obtain trading pairs of venture fund balances
	GetOpenInterest(data map[string]interface{}) *response.APIBody   // Get the open position of the trade pair
}

type XTPrivateFutureHelper interface {
	// Private
	SendOrder(data map[string]interface{}) *response.APIBody        // Order
	GetHistoryList(data map[string]interface{}) *response.APIBody   // Querying order history
	GetTradeList(data map[string]interface{}) *response.APIBody     // Query transaction details
	SendBatchOrder(data map[string]interface{}) *response.APIBody   // Bulk order
	GetOrderDetail(data map[string]interface{}) *response.APIBody   // Query the order by id
	GetOrderList(data map[string]interface{}) *response.APIBody     // Order Details
	CancelOrder(data map[string]interface{}) *response.APIBody      // Cancel the order
	CancelBatchOrder(data map[string]interface{}) *response.APIBody // Cancel all orders

	GetAccountInfo() *response.APIBody                                      // Get information about the account
	GetListenKey() *response.APIBody                                        // Get listenKey
	AccountOpen() *response.APIBody                                         // Open the contract
	GetBalance(data map[string]interface{}) *response.APIBody               // Get the user's single-currency funds
	GetBalanceList() *response.APIBody                                      // Get user funds
	GetBalanceBills(data map[string]interface{}) *response.APIBody          // Get user account flow
	GetFundingRateList(data map[string]interface{}) *response.APIBody       // Fees for obtaining funds
	GetPositionList(data map[string]interface{}) *response.APIBody          // Get position information
	GetAdjustLeverage(data map[string]interface{}) *response.APIBody        // Adjusting the leverage ratio
	UpdatePositionMargin(data map[string]interface{}) *response.APIBody     // Modified margin
	UpdatePositionAutoMargin(data map[string]interface{}) *response.APIBody // Modified automatic margin calls
	AllPositionClose() *response.APIBody                                    // A key positions
	PositionADL() *response.APIBody                                         // Obtaining ADL information
	CollectionAdd(data map[string]interface{}) *response.APIBody            // Collection trade pair
	CollectionCancel(data map[string]interface{}) *response.APIBody         // Cancel collectible trading pairs
	CollectionList() *response.APIBody                                      // List of collectible trading pairs

}

type SignedFutureHttpAPI struct {
	Accesskey string
	Secretkey string
}

func NewSignedFutureHttpAPI(accesskey, secretkey string) *SignedFutureHttpAPI {
	return &SignedFutureHttpAPI{
		Accesskey: accesskey,
		Secretkey: secretkey,
	}
}

func (s *SignedFutureHttpAPI) SetAuthOption(accesskey, secretkey string) {
	s.Accesskey = accesskey
	s.Secretkey = secretkey
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : orderId	    number	    true
 *	@Return
 *		See: https://xt-com.github.io/xt4-api/#market_cn2symbol
**/
func (s SignedFutureHttpAPI) SendOrder(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/create"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("POST", uri, headers, data)
	// rep := requestPerpare.RequesJson("POST", uri, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc
 *			Parameter		Type	mandatory	Default	Description										Ranges
 *			symbol			string	true		N/A		Trading pairs (queries all trading pairs if not passed)
 *			direction		string	false		NEXT	Direction（PREV:Previous page；NEXT:Next page）	PREV;NEXT
 *			id				integer	false		N/A		id
 *			limit			integer	false		10		Limit
 *			startTime		integer	false		N/A		Start time
 *			endTime			integer	false		N/A		End time
 *	@Return
 *		See: https://futuresee.github.io/github.io/#ordergetHistory
**/
func (s SignedFutureHttpAPI) GetHistoryList(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/list-history"
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
 *		@Desc
 *			Parameter	Type	mandatory	Default	Description				Ranges
 *			orderId		integer	false		N/A		Order ID
 *			symbol		string	false		N/A		Trading pair
 *			page		integer	false		1		Page
 *			size		integer	false		10		Quantity of a single page
 *			startTime	integer	false		N/A		starting time
 *			endTime		integer	false		N/A		ClosingTime
 *	@Return
 *		See: https://futuresee.github.io/github.io/#ordergetHistory
**/
func (s SignedFutureHttpAPI) GetTradeList(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/trade-list"
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
 *		@Desc
 *			Parameter	Type	mandatory	Default	Description				Ranges
 *			list		string	true		N/A		List collection of order data
 *	@Return
 *		See: https://futuresee.github.io/github.io/#ordergetHistory
**/
func (s SignedFutureHttpAPI) SendBatchOrder(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/create-batch"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

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
 *		@Desc
 *			Parameter	Type	mandatory	Default	Description				Ranges
 *			orderId		integer	false		N/A		Order ID
 *	@Return
 *		See: https://futuresee.github.io/github.io/#order_cngetById
**/
func (s SignedFutureHttpAPI) GetOrderDetail(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/detail"
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
 *		@Desc
 *			Parameter		Type	mandatory	Default	Description			Ranges
 *			clientOrderId	String	false		N/A		Client order ID
 *			page			integer	false		1		Page
 *			size			integer	false		10		Quantity of a single page
 *			startTime		integer	false		N/A		Start time
 *			endTime			integer	false		N/A		End time
 *			state			string	true		NEW		Order state: NEW：New order (unfilled);PARTIALLY_FILLED:Partial deal;PARTIALLY_CANCELED:Partial revocation;FILLED:Filled;CANCELED:Cancled;REJECTED:Order failed;EXPIRED：Expired;UNFINISHED:Unfinished;HISTORY:(History)
 *			symbol			string	false		N/A		Trading pair
 *	@Return
 *		See: https://futuresee.github.io/github.io/#order_cngetOrders
**/
func (s SignedFutureHttpAPI) GetOrderList(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/list"
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
 *		@Desc
 *			Parameter		Type	mandatory	Default	Description			Ranges
 *			orderId			Integer	true		N/A		Order ID
 *	@Return
 *		See: https://futuresee.github.io/github.io/#order_cncancel
**/
func (s SignedFutureHttpAPI) CancelOrder(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/cancel"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

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
 *		@Desc
 *			Parameter		Type	mandatory	Default	Description			Ranges
 *			symbol			String	true		N/A		Trading pair (cancel all trading pair orders if don't pass parameters)
 *	@Return
 *		See: https://futuresee.github.io/github.io/#order_cncancel
**/
func (s SignedFutureHttpAPI) CancelBatchOrder(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/order/cancel-all"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("POST", uri, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"accountId": 0, //Id
 *			"allowOpenPosition": false, //AvailabilityOfOpenPositions
 *			"allowTrade": false, //Tradeability
 *			"allowTransfer": false, //Transferability
 *			"openTime": "", //OpeningTime
 *			"state": 0, //UserStatus
 *			"userId": 0 //UserId
 *		},
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) GetAccountInfo() *response.APIBody {

	path := "/future/user/v1/account/info"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)
	data := map[string]interface{}{}

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", uri, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {},
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) GetListenKey() *response.APIBody {

	path := "/future/user/v1/user/listen-key"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)
	data := map[string]interface{}{}

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", uri, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": true,
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) AccountOpen() *response.APIBody {

	path := "/future/user/v1/account/open"
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
 *		Parameter		Type	mandatory	Default	Description	Ranges
 *		coin			string	true		N/A		Currency
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetBalance
**/
func (s SignedFutureHttpAPI) GetBalance(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/balance/detail"
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
 *	@Param: None
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetBalances
**/
func (s SignedFutureHttpAPI) GetBalanceList() *response.APIBody {

	path := "/future/user/v1/balance/list"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)
	data := map[string]interface{}{}

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
 *		Parameter		Type	mandatory	Default	Description				Ranges
 *		symbol			string	true		N/A		Trading pairs (queries all trading pairs if not passed)
 *		direction		string	false		NEXT	Direction（PREV:Previous page；NEXT:Next page）	PREV;NEXT
 *		id				integer	false		N/A		id
 *		limit			integer	false		10		Limit
 *		startTime		integer	false		N/A		Start time
 *		endTime			integer	false		N/A		End time
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetBalanceBill
**/
func (s SignedFutureHttpAPI) GetBalanceBills(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/balance/bills"
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
 *		Parameter		Type	mandatory	Default	Description				Ranges
 *		symbol			string	true		N/A		Trading pairs (queries all trading pairs if not passed)
 *		direction		string	false		NEXT	Direction（PREV:Previous page；NEXT:Next page）	PREV;NEXT
 *		id				integer	false		N/A		id
 *		limit			integer	false		10		Limit
 *		startTime		integer	false		N/A		Start time
 *		endTime			integer	false		N/A		End time
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetFunding
**/
func (s SignedFutureHttpAPI) GetFundingRateList(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/balance/funding-rate-list"
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
 *		Parameter		Type	mandatory	Default	Description				Ranges
 *		symbol			string	true		N/A		Trading pairs (queries all trading pairs if not passed)
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetPosition
**/
func (s SignedFutureHttpAPI) GetPositionList(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/position/list"
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
 *		Parameter		Type	mandatory	Default	Description			Ranges
 *		symbol			string	true		N/A		Trading pair
 *		positionSide	string	true		N/A		Position side	LONG;SHORT
 *		leverage		integer	true		N/A		Leverage
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetPosition
**/
func (s SignedFutureHttpAPI) GetAdjustLeverage(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/position/adjust-leverage"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

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
 *		Parameter		Type	mandatory	Default	Description			Ranges
 *		symbol			string	true		N/A		Trading pair
 *		margin			number	false		N/A		Quantity
 *		positionSide	string	false		N/A		Position side:LONG;SHORT
 *		type			string	false		N/A		Adjust direction (add isolated margin, reduce isolated margin)	ADD;SUB
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetPosition
**/
func (s SignedFutureHttpAPI) UpdatePositionMargin(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/position/margin"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

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
 *		Parameter		Type	mandatory	Default	Description			Ranges
 *		symbol			string	true		N/A		Trading pair
 *		autoMargin		bool	false		N/A		Whether to automatically call margin
 *		positionSide	string	false		N/A		Position side:LONG;SHORT
 *	@Return
 * 		See https://futuresee.github.io/github.io/#user_cngetPosition
**/
func (s SignedFutureHttpAPI) UpdatePositionAutoMargin(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/position/auto-margin"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("POST", uri, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": true,
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) AllPositionClose() *response.APIBody {

	path := "/future/user/v1/position/close-all"
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
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": [
 *			{
 *			"longQuantile": 0, //MultipleAdl
 *			"shortQuantile": 0, //ShortAdl
 *			"symbol": "" //TradingPair
 *			}
 *		],
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) PositionADL() *response.APIBody {

	path := "/future/user/v1/position/adl"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)
	data := map[string]interface{}{}

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
 *		Parameter		Type	mandatory	Default	Description			Ranges
 *		symbol			string	true		N/A		Trading pair
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": true,
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) CollectionAdd(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/user/collection/add"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

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
 *		Parameter		Type	mandatory	Default	Description			Ranges
 *		symbol			string	true		N/A		Trading pair
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": true,
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) CollectionCancel(data map[string]interface{}) *response.APIBody {

	path := "/future/user/v1/user/collection/add"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "POST")
	auth.SetUrlencode(true)

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("POST", uri, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": [],
 *		"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) CollectionList() *response.APIBody {

	path := "/future/user/v1/user/collection/list"
	uri := BaseUrl + path
	auth := NewAuth(s, path, "GET")
	auth.SetUrlencode(true)
	data := map[string]interface{}{}

	headers, err := auth.createPayload(data)
	if err != nil {
		return response.APIResponse(err.Error(), "Failed", uri, false)
	}

	requestPerpare := v4.NewRequestPerpare()
	rep := requestPerpare.RequesParam("GET", uri, headers, data)

	return rep
}

type PublicFutureHttpAPI struct {
}

func (p PublicFutureHttpAPI) GetServerTime() *response.APIBody {
	path := "/future/market/v1/public/time"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param: None
 * @Return
 *	{
 *	"error": {
 *		"code": "",
 *		"msg": ""
 *	},
 *	"msgInfo": "",
 *	"result": [],
 *	"returnCode": 0
 *	}
 */
func (p PublicFutureHttpAPI) GetCoinsInfo() *response.APIBody {
	path := "/future/market/v1/public/symbol/coins"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		See: https://futuresee.github.io/github.io/#quotes_cngetLeverageBracket
**/
func (p PublicFutureHttpAPI) GetAllMarketConfig() *response.APIBody {
	path := "/future/market/v1/public/symbol/list"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : symbol	    string	    true
 *	@Return
 *		See: https://futuresee.github.io/github.io/#quotes_cngetLeverageBracket
**/
func (p PublicFutureHttpAPI) GetMarketConfig(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/symbol/detail"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : symbol	    string	    true
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"leverageBrackets": [
 *			{
 *				"bracket": 0, //Level
 *				"maintMarginRate": 0, //MaintenanceOfMarginRates
 *				"maxLeverage": 0, //MaximumLeverageMultiple
 *				"maxNominalValue": 0, //MaximumNominalValueOfTheLayer
 *				"maxStartMarginRate": 0, //MaximumStartingMarginRate
 *				"minLeverage": 0, //MinimumLeverageMultiple
 *				"startMarginRate": 0, //StartingMarginRate
 *				"symbol": "" //TradingPair
 *			}
 *			],
 *			"symbol": ""
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetLeverageDetail(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/leverage/bracket/detail"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		See https://futuresee.github.io/github.io/#quotes_cngetLeverageBrackets
**/
func (p PublicFutureHttpAPI) GetLeverageDetailList() *response.APIBody {
	path := "/future/market/v1/public/leverage/bracket/list"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		@Desc     Parameter	    Type	    mandatory    Default	    Description
 *		::param : symbol	    string	    true
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"a": "", //24HourVolume
 *			"c": "", //LatestPrice
 *			"h": "", //24HourMaximumPrice
 *			"l": "", //LowestPriceIn_24Hours
 *			"o": "", //FirstSalePrice_24HoursAgo
 *			"r": "", //24HourRiseAndFall
 *			"s": "", //TradingPair
 *			"t": 0, //Times
 *			"v": "" //24HourTurnover
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetMarketTicker(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/ticker"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		See https://futuresee.github.io/github.io/#quotes_cngetLeverageBrackets
**/
func (p PublicFutureHttpAPI) GetMarketTickers() *response.APIBody {
	path := "/future/market/v1/public/q/tickers"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		Parameter		Type		mandatory	Default		Description		Ranges
 *		symbol			string		true		N/A			Trading pair
 *		num				integer		true		N/A			Quantity
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": [
 *			{
 *			"a": 0, //Volume
 *			"m": "", //Order side
 *			"p": 0, //Price
 *			"s": "", //Trading pair
 *			"t": 0 //Time
 *			}
 *		],
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetMarketDeal(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/deal"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		Parameter		Type		mandatory	Default		Description		Ranges
 *		symbol			string		true		N/A			Trading pair
 *		level			integer		true		N/A			Level(min:1,max:50)
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"a": [], //Buy
 *			"b": [], //Sell
 *			"s": "", //Trading pair
 *			"t": 0, //Time
 *			"u": 0 //updateId
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetDepth(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/depth"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		Parameter		Type		mandatory	Default		Description		Ranges
 *		symbol			string		false		N/A			Trading pair
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"p": 0, //Price
 *			"s": "", //TradingPair
 *			"t": 0 //Times
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetIndexPrice(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/symbol-index-price"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param: None
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"p": 0, //Price
 *			"s": "", //TradingPair
 *			"t": 0 //Times
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetAllIndexPrice() *response.APIBody {
	path := "/future/market/v1/public/q/index-price"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		Parameter		Type		mandatory	Default		Description		Ranges
 *		symbol			string		false		N/A			Trading pair
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"p": 0, //Price
 *			"s": "", //TradingPair
 *			"t": 0 //Times
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetMarketPrice(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/symbol-mark-price"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 *	@Param:
 *		Parameter		Type		mandatory	Default		Description		Ranges
 *		symbol			string		false		N/A			Trading pair
 *	@Return
 *		{
 *		"error": {
 *			"code": "",
 *			"msg": ""
 *		},
 *		"msgInfo": "",
 *		"result": {
 *			"p": 0, //Price
 *			"s": "", //TradingPair
 *			"t": 0 //Times
 *		},
 *		"returnCode": 0
 *		}
**/
func (p PublicFutureHttpAPI) GetAllMarketPrice() *response.APIBody {
	path := "/future/market/v1/public/q/mark-price"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 * 	 Parameter		Type		mandatory	Default	Description		Ranges
 * 		symbol		string		true		N/A		Trading pair
 * 		interval	string		true		N/A		Time-interval	1m;5m;15m;30m;1h;4h;1d;1w
 * 		startTime	integer		false		N/A		Start time
 * 		endTime		integer		false		N/A		End time
 * 		limit		integer		false		N/A		Limit
 * @Return
 * 	{
 * 	"error": {
 * 		"code": "",
 * 		"msg": ""
 * 	},
 * 	"msgInfo": "",
 * 	"result": [
 * 		{
 * 		"a": 0, //Volume
 * 		"c": 0, //Close price
 * 		"h": 0, //Highest price
 * 		"l": 0, //Lowest price
 * 		"o": 0, //Open price
 * 		"s": "", //Trading pair
 * 		"t": 0, //Time
 * 		"v": 0 //Turnover
 * 		}
 * 	],
 * 	"returnCode": 0
 * 	}
 **/
func (p PublicFutureHttpAPI) GetKline(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/kline"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 * 	 Parameter		Type		mandatory	Default	Description		Ranges
 * 		symbol		string		true		N/A		Trading pair
 * @Return
 * 	See https://futuresee.github.io/github.io/#quotes_cngetAggTickers
 **/
func (p PublicFutureHttpAPI) GetAggTicker(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/agg-ticker"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param: None
 * @Return
 * 	See https://futuresee.github.io/github.io/#quotes_cngetAggTickers
 **/
func (p PublicFutureHttpAPI) GetAllAggTicker() *response.APIBody {
	path := "/future/market/v1/public/q/agg-tickers"
	requestPerpare := v4.NewRequestPerpare()
	headers, data := map[string]string{}, map[string]interface{}{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 * 	 Parameter		Type		mandatory	Default	Description		Ranges
 * 		symbol		string		true		N/A		Trading pair
 * @Return
 * 	See https://futuresee.github.io/github.io/#quotes_cngetFundingRate
 **/
func (p PublicFutureHttpAPI) GetFundRate(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/funding-rate"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 *	Parameter			Type		mandatory	Default	Description	Ranges
 *		symbol			string		true		N/A		Trading pair
 *		direction		string		false		NEXT	Direction（PREV:Previous page；NEXT:Next page）	PREV;NEXT
 *		id				integer		false		N/A		id
 *		limit			integer		false		10		Limit
 * @Return
 * 	See https://futuresee.github.io/github.io/#quotes_cngetFundingRateRecord
 **/
func (p PublicFutureHttpAPI) GetFundRateRecord(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/q/funding-rate-record"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 *	Parameter			Type		mandatory	Default	Description	Ranges
 *		symbol			string		true		N/A		Trading pair
 *		direction		string		false		NEXT	Direction（PREV:Previous page；NEXT:Next page）	PREV;NEXT
 *		id				integer		false		N/A		id
 *		limit			integer		false		10		Limit
 * @Return
 * 	See https://futuresee.github.io/github.io/#quotes_cngetRiskBalance
 **/
func (p PublicFutureHttpAPI) GetRiskBalance(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/contract/risk-balance"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}

/**
 * @Param:
 *	Parameter			Type		mandatory	Default	Description	Ranges
 *		symbol			string		true		N/A		Trading pair
 * @Return
 * 	See https://futuresee.github.io/github.io/#quotes_cngetOpenInterest
 **/
func (p PublicFutureHttpAPI) GetOpenInterest(data map[string]interface{}) *response.APIBody {
	path := "/future/market/v1/public/contract/open-interest"
	requestPerpare := v4.NewRequestPerpare()
	headers := map[string]string{}
	rep := requestPerpare.RequesParam("GET", BaseUrl+path, headers, data)

	return rep
}
