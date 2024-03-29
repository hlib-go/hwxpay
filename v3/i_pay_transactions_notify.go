package v3

/*
微信支付通过支付通知接口将用户支付成功消息通知给商户
注意：
• 同样的通知可能会多次发送给商户系统。商户系统必须能够正确处理重复的通知。
	推荐的做法是，当商户系统收到通知进行处理时，先检查对应业务数据的状态，并判断该通知是否已经处理。
	如果未处理，则再进行处理；如果已处理，则直接返回结果成功。在对业务数据进行状态检查和处理之前，要采用数据锁进行并发控制，以避免函数重入造成的数据混乱。
• 如果在所有通知频率后没有收到微信侧回调，商户应调用查询订单接口确认订单状态。
特别提醒：商户系统对于开启结果通知的内容一定要做签名验证，并校验通知的信息是否与商户侧的信息一致，防止数据泄漏导致出现“假通知”，造成资金损失。

回调URL：该链接是通过统一下单接口中的请求参数“notify_url”来设置的，要求必须为https地址。请确保回调URL是外部可正常访问的，且不能携带后缀参数，否则可能导致商户无法接收到微信的回调通知信息。回调URL示例： “https://pay.weixin.qq.com/wxpay/pay.action”
*/

// 支付通知
func NotifyPay(conf *Config, data string) (p *PayNotifyParams, err error) {
	err = NotifyDecrypt(conf, data, &p)
	return
}

type PayNotifyParams struct {
	Appid          string                 `json:"appid"`
	Mchid          string                 `json:"mchid"`
	OutTradeNo     string                 `json:"out_trade_no"`
	TransactionId  string                 `json:"transaction_id"`
	TradeType      string                 `json:"trade_type"`
	TradeState     TradeState             `json:"trade_state"`
	TradeStateDesc string                 `json:"trade_state_desc"`
	BankType       string                 `json:"bank_type"`
	Attach         string                 `json:"attach"`
	SuccessTime    string                 `json:"success_time"`
	Payer          *PayNotifyParamsPayer  `json:"payer"`
	Amount         *PayNotifyParamsAmount `json:"amount"`
}

type PayNotifyParamsPayer struct {
	Openid string `json:"openid"` //用户在直连商户appid下的唯一标识。
}

type PayNotifyParamsAmount struct {
	Total         int64  `json:"total"`                   //订单总金额，单位为分。
	PayerTotal    int64  `json:"payer_total"`             //用户支付金额，单位为分。
	Currency      string `json:"currency,omitempt"`       //非必填。 CNY：人民币，境内商户号仅支持人民币。
	PayerCurrency string `json:"payer_currency,omitempt"` //用户支付币种
}
