package v3

import "fmt"

// MarketingFavorUsersOpenidCouponsGet 查询代金券详情API  https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_6.shtml
func MarketingFavorUsersOpenidCouponsGet(cfg *Config, appid, openid, couponId string) (result map[string]interface{}, err error) {
	///v3/marketing/favor/users/{openid}/coupons/{coupon_id}
	err = GET(cfg, fmt.Sprintf("/v3/marketing/favor/users/%s/coupons/%s?appid=%s", openid, couponId, appid), nil, &result)
	return
}
