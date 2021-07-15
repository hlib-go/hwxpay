package main

import (
	"encoding/json"
	v3 "github.com/hlib-go/hwxpay/v3"
	"testing"
)

func TestRefund(t *testing.T) {
	c := Config()

	r, err := v3.Refund(c, &v3.RefundParams{
		TransactionId: "4200001146202107159403183386",
		OutTradeNo:    "14156014837805301780",
		OutRefundNo:   "14156014837805301780_01",
		Reason:        "Reason",
		NotifyUrl:     "",
		FundsAccount:  "",
		Amount: &v3.RefundParamsAmount{
			Refund:   1,
			Total:    1,
			Currency: "CNY",
		},
		GoodsDetail: nil,
	})
	if err != nil {
		t.Error(err)
		return
	}
	rs, _ := json.Marshal(r)
	t.Log(string(rs))
}
