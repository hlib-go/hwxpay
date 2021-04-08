package main

import (
	v3 "github.com/hlib-go/hwxpay/v3"
	"testing"
	"time"
)

// JSAPI 下单测试
func TestJsapiPrepay(t *testing.T) {
	c := Config()

	p, err := v3.JsapiPrepay(c, &v3.JsapiParams{
		Description: "mmm",
		OutTradeNo:  "23131345356",
		TimeExpire:  time.Now().Add(300 * time.Hour).Format(time.RFC3339),
		Attach:      "",
		NotifyUrl:   "https://msd.himkt.cn/wxpay/v3notify",
		GoodsTag:    "",
		Amount: &v3.JsapiParamsAmount{
			Total: 1,
		},
		Payer: &v3.JsapiParamsPayer{Openid: "oW-G-0hraaeGSfxrx_q9AfYRev60"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p.JSON())
}
