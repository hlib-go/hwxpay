package main

import (
	v3 "github.com/hlib-go/hwxpay/v3"
	"math/rand"
	"strconv"
	"testing"
)

/*
微信支付：营销接口测试
*/

// 发放代金券API测试
func TestMarketingFavorUsersOpenidCoupons(t *testing.T) {
	// 朱虹波：oW-G-0nH7QgggeekmXrr8hweyFFs
	// zengs "oW-G-0hraaeGSfxrx_q9AfYRev60"
	result, err := v3.MarketingFavorUsersOpenidCoupons(Config(), "oW-G-0hraaeGSfxrx_q9AfYRev60", &v3.MarketingFavorUsersOpenidCouponsParams{
		StockId:      "15269237",
		OutRequestNo: strconv.Itoa(rand.Int()),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success CouponId:", result.CouponId)
}
