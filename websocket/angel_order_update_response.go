package websocket

type AngelOrder struct {
	Variety             string  `json:"variety"`
	OrderType           string  `json:"ordertype"`
	Ordertag            string  `json:"ordertag"`
	Producttype         string  `json:"producttype"`
	Price               float64 `json:"price"`
	Triggerprice        float64 `json:"triggerprice"`
	Quantity            string  `json:"quantity"`
	Disclosedquantity   string  `json:"disclosedquantity"`
	Duration            string  `json:"duration"`
	Squareoff           float64 `json:"squareoff"`
	Stoploss            float64 `json:"stoploss"`
	Trailingstoploss    float64 `json:"trailingstoploss"`
	TradingSymbol       string  `json:"tradingsymbol"`
	TransactionType     string  `json:"transactiontype"`
	Exchange            string  `json:"exchange"`
	SymbolToken         string  `json:"symboltoken"`
	Instrumenttype      string  `json:"instrumenttype"`
	Strikeprice         float64 `json:"strikeprice"`
	OptionType          string  `json:"optiontype"`
	Expirydate          string  `json:"expirydate"`
	Lotsize             string  `json:"lotsize"`
	Cancelsize          string  `json:"cancelsize"`
	AveragePrice        float64 `json:"averageprice"`
	FilledShares        string  `json:"filledshares"`
	Unfilledshares      string  `json:"unfilledshares"`
	OrderID             string  `json:"orderid"`
	Text                string  `json:"text"`
	Status              string  `json:"status"`
	OrderStatus         string  `json:"orderstatus"`
	Updatetime          string  `json:"updatetime"`
	ExchangeTime        string  `json:"exchtime"`
	Exchorderupdatetime string  `json:"exchorderupdatetime"`
	Fillid              string  `json:"fillid"`
	Filltime            string  `json:"filltime"`
	Parentorderid       string  `json:"parentorderid"`
}

type AngelOrderUpdateResponse struct {
	//{"user-id": "R52269012","status-code": "200","order-status": "AB01","error-message": "","orderData": {"variety": "NORMAL","ordertype": "MARKET","ordertag": "","producttype": "CARRYFORWARD","price": 0.0,"triggerprice": 0.0,"quantity": "1120","disclosedquantity": "0","duration": "DAY","squareoff": 0.0,"stoploss": 0.0,"trailingstoploss": 0.0,"tradingsymbol": "FINNIFTY19DEC2321750CE","transactiontype": "SELL","exchange": "NFO","symboltoken": "35040","instrumenttype": "OPTIDX","strikeprice": 21750.0,"optiontype": "CE","expirydate": "19DEC2023","lotsize": "40","cancelsize": "0","averageprice": 0.0,"filledshares": "0","unfilledshares": "1120","orderid": "231219000044295","text": "","status": "open","orderstatus": "open","updatetime": "19-Dec-2023 09:18:10","exchtime": "19-Dec-2023 09:18:10","exchorderupdatetime": "19-Dec-2023 09:18:10","fillid": "","filltime": "","parentorderid": ""}}
	UserID       string     `json:"user-id"`
	StatusCode   string     `json:"status-code"`
	OrderStatus  string     `json:"order-status"`
	ErrorMessage string     `json:"error-message"`
	OrderData    AngelOrder `json:"orderData"`
}

func (a AngelOrderUpdateResponse) IsValid() bool {
	return a.StatusCode == "200" && a.OrderStatus != string(OrderStatus.Zero)
}

type OrderStatusEnum string

var OrderStatus = struct {
	Zero                              OrderStatusEnum
	Open                              OrderStatusEnum
	Cancelled                         OrderStatusEnum
	Rejected                          OrderStatusEnum
	Modified                          OrderStatusEnum
	Complete                          OrderStatusEnum
	AfterMarketOrderReqReceived       OrderStatusEnum
	CancelledAfterMarketOrder         OrderStatusEnum
	ModifyAfterMarketOrderReqReceived OrderStatusEnum
	OpenPending                       OrderStatusEnum
	TriggerPending                    OrderStatusEnum
	ModifyPending                     OrderStatusEnum
}{
	Zero:                              "AB00",
	Open:                              "AB01",
	Cancelled:                         "AB02",
	Rejected:                          "AB03",
	Modified:                          "AB04",
	Complete:                          "AB05",
	AfterMarketOrderReqReceived:       "AB06",
	CancelledAfterMarketOrder:         "AB07",
	ModifyAfterMarketOrderReqReceived: "AB08",
	OpenPending:                       "AB09",
	TriggerPending:                    "AB10",
	ModifyPending:                     "AB11",
}
