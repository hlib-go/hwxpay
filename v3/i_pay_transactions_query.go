package v3

// 查询订单API
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transactions/chapter3_5.shtml
func Query(conf *Config, transactionId, outTradeNo, mchid string) (r *QueryResult, err error) {
	// 微信支付订单号查询
	if transactionId != "" {
		err = GET(conf, "/v3/pay/transactions/id/"+transactionId+"?mchid="+mchid, nil, &r)
		return
	}
	//商户订单号查询
	err = GET(conf, "/v3/pay/transactions/out-trade-no/"+outTradeNo+"?mchid="+mchid, nil, &r)
	return
}

type QueryResult struct {
	Appid           string      `json:"appid"`
	Mchid           string      `json:"mchid"`
	OutTradeNo      string      `json:"out_trade_no"`
	TransactionId   string      `json:"transaction_id"`
	TradeType       string      `json:"trade_type"`
	TradeState      string      `json:"trade_state"`
	TradeStateDesc  string      `json:"trade_state_desc"`
	BankType        string      `json:"bank_type"`
	Attach          string      `json:"attach"`
	SuccessTime     string      `json:"success_time"`
	Payer           interface{} `json:"payer"`
	Amount          interface{} `json:"amount"`
	PromotionDetail interface{} `json:"promotion_detail"`
}
