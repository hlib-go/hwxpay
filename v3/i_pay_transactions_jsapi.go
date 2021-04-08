package v3

import (
	"encoding/json"
	"strconv"
	"time"
)

// 商户直连： JSAPI/小程序下单API
func JsapiPrepay(cfg *Config, params *JsapiParams) (p *PrepayParams, err error) {
	if params.Appid == "" {
		params.Appid = cfg.Appid
	}
	if params.Mchid == "" {
		params.Mchid = cfg.Mchid
	}

	r, err := TransactionsJsapi(cfg, params)
	if err != nil {
		return
	}
	p, err = Prepay(cfg, r.PrepayId)
	if err != nil {
		return
	}
	return
}

func TransactionsJsapi(conf *Config, params *JsapiParams) (r *JsapiResult, err error) {
	err = POST(conf, "/v3/pay/transactions/jsapi", params, &r)
	return
}

type JsapiParams struct {
	Appid       string             `json:"appid"`                //直连商户申请的公众号或者小程序appid。
	Mchid       string             `json:"mchid"`                //直连商户的商户号，由微信支付生成并下发。
	Description string             `json:"description"`          //商品描述
	OutTradeNo  string             `json:"out_trade_no"`         //商户系统内部订单号
	TimeExpire  string             `json:"time_expire,omitempt"` //非必填。订单失效时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，
	Attach      string             `json:"attach,omitempt"`      //非必填。附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用
	NotifyUrl   string             `json:"notify_url"`           //通知URL必须为直接可访问的URL，不允许携带查询串，要求必须为https地址。
	GoodsTag    string             `json:"goods_tag,omitempt"`   //非必填。订单优惠标记
	Amount      *JsapiParamsAmount `json:"amount"`               //订单金额信息
	Payer       *JsapiParamsPayer  `json:"payer"`                //支付者信息
	//Detail      *JsapiParamsDetail `json:"detail,omitempt"`      //非必填。优惠功能,正常支付订单不比此参数
	//SceneInfo   interface{}        `json:"scene_info,omitempt"`  //非必填。支付场景描述
	//SettleInfo  interface{}        `json:"settle_info,omitempt"` //非必填。结算信息
}

type JsapiParamsAmount struct {
	Total    int64  `json:"total"`             //订单总金额，单位为分。
	Currency string `json:"currency,omitempt"` //非必填。 CNY：人民币，境内商户号仅支持人民币。
}

type JsapiParamsPayer struct {
	Openid string `json:"openid"` //用户在直连商户appid下的唯一标识。
}

type JsapiParamsDetail struct {
	CostPrice   int64                     `json:"cost_price,omitempt"`   //订单原价
	InvoiceId   string                    `json:"invoice_id,omitempt"`   //非必填。商品小票ID
	GoodsDetail []*JsapiParamsDetailGoods `json:"goods_detail,omitempt"` //非必填。单品列表信息，条目个数限制：【1，6000】
}

type JsapiParamsDetailGoods struct {
	MerchantGoodsId  string `json:"merchant_goods_id"`           //商户侧商品编码
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempt"` //非必填。微信侧商品编码
	GoodsName        string `json:"goods_name,omitempt"`         //非必填。商品名称
	Quantity         int64  `json:"quantity"`                    //商品数量
	UnitPrice        int64  `json:"unit_price"`                  //商品单价
}

type JsapiResult struct {
	PrepayId string `json:"prepay_id"` //预支付交易会话标识。用于后续接口调用中使用，该值有效期为2小时
}

// JSAPI 与 小程序 调起支付参数
func Prepay(conf *Config, prepayId string) (p *PrepayParams, err error) {
	p = &PrepayParams{
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  RandomString(32),
		Package:   "prepay_id=" + prepayId,
		SignType:  "RSA",
		PaySign:   "",
	}
	s := conf.Appid + "\n" + p.TimeStamp + "\n" + p.NonceStr + "\n" + p.Package + "\n"
	priKey, err := conf.GetMchPriKey()
	if err != nil {
		return
	}
	sign, err := RsaSignWithSha256(s, priKey)
	if err != nil {
		return
	}
	p.PaySign = sign
	return
}

type PrepayParams struct {
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (p *PrepayParams) JSON() string {
	v, _ := json.Marshal(p)
	return string(v)
}
