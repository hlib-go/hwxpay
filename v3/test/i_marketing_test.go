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
	result, err := v3.MarketingFavorUsersOpenidCouponsPost(ConfigFws(), "oW-G-0nH7QgggeekmXrr8hweyFFs", &v3.MarketingFavorUsersOpenidCouponsParams{
		StockId:      "16038626",
		OutRequestNo: strconv.Itoa(rand.Int()),
		//Appid:             "wx239c521c61221a8a",
		//StockCreatorMchid: "1616529748",
		CouponValue:   0,
		CouponMinimum: 0,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success CouponId:", result.CouponId)
}

// 查询代金券
func TestMarketingFavorUsersOpenidCouponsGet(t *testing.T) {
	//couponId:29038964251
	result, err := v3.MarketingFavorUsersOpenidCouponsGet(ConfigFws(), "wx239c521c61221a8a", "oW-G-0nH7QgggeekmXrr8hweyFFs", "29038964251")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
