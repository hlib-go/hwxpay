package v3test

import (
	v3 "hwxpay/v3"
	"math/rand"
	"strconv"
	"testing"
)

/*
微信支付：营销接口测试
*/

// 发放代金券API测试
func TestMarketingFavorUsersOpenidCoupons(t *testing.T) {
	result, err := v3.MarketingFavorUsersOpenidCoupons(Config(), "oW-G-0hraaeGSfxrx_q9AfYRev60", &v3.MarketingFavorUsersOpenidCouponsParams{
		StockId:      "15245469",
		OutRequestNo: strconv.Itoa(rand.Int()),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success CouponId:", result.CouponId)
}
