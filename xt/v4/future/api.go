package future

import (
	"gopy/sdk/response"
	v4 "gopy/sdk/xt/v4"
)

type XTPublicFutureHelper interface {
	// Public
	GetServerTime() *response.APIBody                                // 获取服务器时间
	GetCoinsInfo() *response.APIBody                                 // 获取交易对币种
	GetAllMarketConfig() *response.APIBody                           // 获取所有交易对的配置信息
	GetMarketConfig(data map[string]interface{}) *response.APIBody   // 获取所有交易对的配置信息
	GetLeverageDetail(data map[string]interface{}) *response.APIBody // 查询单个交易对杠杆分层
	GetLeverageDetailList() *response.APIBody                        // 查询所有交易对杠杆分层
	GetMarketTicker(data map[string]interface{}) *response.APIBody   // 获取指定交易对的行情信息
	GetMarketTickers() *response.APIBody                             // 获取全交易对的行情信息
	GetMarketDeal(data map[string]interface{}) *response.APIBody     // 获取交易对的最新成交信息
	GetDepth(data map[string]interface{}) *response.APIBody          // 获取交易对的深度信息
	GetIndexPrice(data map[string]interface{}) *response.APIBody     // 获取单个交易对的指数价格
	GetAllIndexPrice() *response.APIBody                             // 获取所有交易对的指数价格
	GetMarketPrice(data map[string]interface{}) *response.APIBody    // 获取单个交易对的标记价格
	GetAllMarketPrice() *response.APIBody                            // 获取所有交易对的标记价格
	GetKline(data map[string]interface{}) *response.APIBody          // 获取交易对的k线信息
	GetAggTicker(data map[string]interface{}) *response.APIBody      // 获取指定交易对的聚合行情信息
	GetAllAggTicker() *response.APIBody                              // 获取全交易对的聚合行情信息
	GetFundRate(data map[string]interface{}) *response.APIBody       // 获取资金费率
	GetFundRateRecord(data map[string]interface{}) *response.APIBody // 获取资金费率记录
	GetRiskBalance(data map[string]interface{}) *response.APIBody    // 获取交易对风险基金余额
	GetOpenInterest(data map[string]interface{}) *response.APIBody   // 获取交易对开仓量
}

type XTPrivateFutureHelper interface {
	// Private
	SendOrder(data map[string]interface{}) *response.APIBody        // 下单
	GetHistoryList(data map[string]interface{}) *response.APIBody   // 查询历史订单
	GetTradeList(data map[string]interface{}) *response.APIBody     // 查询成交明细
	SendBatchOrder(data map[string]interface{}) *response.APIBody   // 批量下单
	GetOrderDetail(data map[string]interface{}) *response.APIBody   // 根据id查询订单
	GetOrderList(data map[string]interface{}) *response.APIBody     // 查询订单
	CancelOrder(data map[string]interface{}) *response.APIBody      // 撤销订单
	CancelBatchOrder(data map[string]interface{}) *response.APIBody // 撤销所有订单

	GetAccountInfo() *response.APIBody                                      // 获取账户相关信息
	GetListenKey() *response.APIBody                                        // 获取listenKey
	AccountOpen() *response.APIBody                                         // 开通合约
	GetBalance(data map[string]interface{}) *response.APIBody               // 获取用户单币种资金
	GetBalanceList() *response.APIBody                                      // 获取用户资金
	GetBalanceBills(data map[string]interface{}) *response.APIBody          // 获取用户账务流水
	GetFundingRateList(data map[string]interface{}) *response.APIBody       // 获取资金费用
	GetPositionList(data map[string]interface{}) *response.APIBody          // 获取持仓信息
	GetAdjustLeverage(data map[string]interface{}) *response.APIBody        // 调整杠杆倍数
	UpdatePositionMargin(data map[string]interface{}) *response.APIBody     // 修改保证金
	UpdatePositionAutoMargin(data map[string]interface{}) *response.APIBody // 修改自动追加保证金
	AllPositionClose() *response.APIBody                                    // 一键平仓
	PositionADL() *response.APIBody                                         // 获取ADL信息
	CollectionAdd(data map[string]interface{}) *response.APIBody            // 收藏交易对
	CollectionCancel(data map[string]interface{}) *response.APIBody         // 取消收藏交易对
	CollectionList() *response.APIBody                                      // 收藏交易对列表

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
 *			startTime	integer	false		N/A		起始时间
 *			endTime		integer	false		N/A		结束时间
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
 *			"accountId": 0, //帐户id
 *			"allowOpenPosition": false, //是否可开仓
 *			"allowTrade": false, //是否可以交易
 *			"allowTransfer": false, //是否可以划转
 *			"openTime": "", //开通时间
 *			"state": 0, //用户状态
 *			"userId": 0 //用户id
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
 *			"longQuantile": 0, //多头adl
 *			"shortQuantile": 0, //空头adl
 *			"symbol": "" //交易对
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

/**
 *	@Param:
 *     See https://futuresee.github.io/github.io/#entrust_cncreatePlan
 *	@Return
 *		{
 *			"error": {
 *			"code": "",
 *			"msg": ""
 *			},
 *			"msgInfo": "",
 *			"result": {},
 *			"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) CreateEntrustPlan(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/entrust/create-plan"
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
 *		entrustId		integer	true		N/A		Trigger order ID
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
func (s SignedFutureHttpAPI) CancelEntrustPlan(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/entrust/cancel-plan"
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
 *		Parameter		Type	mandatory	Default		Description			Ranges
 *		symbol			string	false		N/A			Trading pair (cancel all trading pair orders if don't pass parameters)
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
func (s SignedFutureHttpAPI) CancelAllPlan(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/entrust/cancel-all-plan"
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
 *     See https://futuresee.github.io/github.io/#entrust_cncancelPlan
 *	@Return
 *	   See https://futuresee.github.io/github.io/#entrust_cncancelPlan
**/
func (s SignedFutureHttpAPI) GetEntrustPlanList(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/entrust/plan-list"
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
 *		Parameter	Type	mandatory	Default	Description			Ranges
 *		entrustId	integer	true		N/A		Order ID
 *	@Return
 *	   See https://futuresee.github.io/github.io/#entrust_cncancelPlan
**/
func (s SignedFutureHttpAPI) GetEntrustPlanDetail(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/entrust/plan-detail"
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
 *		Parameter		Type	mandatory	Default	Description												Ranges
 *		symbol			string	true		N/A		Trading pairs (queries all trading pairs if not passed)
 *		direction		string	false		NEXT	Direction（PREV:Previous page；NEXT:Next page）			PREV;NEXT
 *		id				integer	false		N/A		id
 *		limit			integer	false		10		Limit
 *		startTime		integer	false		N/A		Start time
 *		endTime			integer	false		N/A		End time
 *	@Return
 *	   See https://futuresee.github.io/github.io/#entrust_cngetPlanHistory
**/
func (s SignedFutureHttpAPI) GetEntrustHistoryPlanList(data map[string]interface{}) *response.APIBody {

	path := "/future/trade/v1/entrust/plan-list-history"
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
 *     See https://futuresee.github.io/github.io/#entrust_cncreatePlan
 *	@Return
 *		{
 *			"error": {
 *			"code": "",
 *			"msg": ""
 *			},
 *			"msgInfo": "",
 *			"result": {},
 *			"returnCode": 0
 *		}
**/
func (s SignedFutureHttpAPI) CreateEntrustProfitPlan(data map[string]interface{}) *response.APIBody {
	// TODO
	path := "/future/trade/v1/entrust/create-plan"
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
 *		Parameter	Type	mandatory	Default	Description		Ranges
 *		profitId	integer	true		N/A		Stop limit ID
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
func (s SignedFutureHttpAPI) CancelEntrustProfitStop(data map[string]interface{}) *response.APIBody {
	// TODO
	path := "/future/trade/v1/entrust/cancel-profit-stop"
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
 *		symbol			string	false		N/A		Trading pair (cancel all trading pair orders if don't pass parameters)
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
func (s SignedFutureHttpAPI) CancelAllEntrustProfitStop(data map[string]interface{}) *response.APIBody {
	// TODO
	path := "/future/trade/v1/entrust/cancel-all-profit-stop"
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
 *		symbol			string	false		N/A		Trading pair (cancel all trading pair orders if don't pass parameters)
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
func (s SignedFutureHttpAPI) GetEntrustProfitList(data map[string]interface{}) *response.APIBody {
	// TODO
	path := "/future/trade/v1/entrust/profit-list"
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
 *		Parameter	Type	mandatory	Default	Description			Ranges
 *		profitId	integer	true		N/A		Stop limit ID
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
func (s SignedFutureHttpAPI) GetEntrustProfitDetail(data map[string]interface{}) *response.APIBody {
	// TODO
	path := "/future/trade/v1/entrust/profit-detail"
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
 *		Parameter			Type	mandatory	Default		Description			Ranges
 *		profitId			integer	true		N/A			Stop limit ID
 *		triggerProfitPrice	number	false		N/A			TP trigger price
 *		triggerStopPrice	number	false		N/A			SL trigger price
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
func (s SignedFutureHttpAPI) UpdateEntrustProfitStop(data map[string]interface{}) *response.APIBody {
	// TODO
	path := "/future/trade/v1/entrust/update-profit-stop"
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
 *				"bracket": 0, //档位
 *				"maintMarginRate": 0, //维持保证金率
 *				"maxLeverage": 0, //最大杠杆倍数
 *				"maxNominalValue": 0, //该层最大名义价值
 *				"maxStartMarginRate": 0, //最大起始保证金率
 *				"minLeverage": 0, //最小杠杆倍数
 *				"startMarginRate": 0, //起始保证金率
 *				"symbol": "" //交易对
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
 *			"a": "", //24小时成交量
 *			"c": "", //最新价
 *			"h": "", //24小时最高价
 *			"l": "", //24小时最低价
 *			"o": "", //24小时前第一笔成交价
 *			"r": "", //24小时涨跌幅
 *			"s": "", //交易对
 *			"t": 0, //时间
 *			"v": "" //24小时成交额
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
 *			"p": 0, //价格
 *			"s": "", //交易对
 *			"t": 0 //时间
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
 *			"p": 0, //价格
 *			"s": "", //交易对
 *			"t": 0 //时间
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
 *			"p": 0, //价格
 *			"s": "", //交易对
 *			"t": 0 //时间
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
 *			"p": 0, //价格
 *			"s": "", //交易对
 *			"t": 0 //时间
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
