package v3

// MarketingFavorUsersOpenidCoupons 发放代金券 文档 https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/marketing/convention/chapter3_2.shtml
func MarketingFavorUsersOpenidCoupons(cfg *Config, openid string, i *MarketingFavorUsersOpenidCouponsReq) (o *MarketingFavorUsersOpenidCouponsRes, err error) {
	err = Call(cfg, "/v3/marketing/favor/users/"+openid+"/coupons", "POST", i, &o)
	return o, err
}

type MarketingFavorUsersOpenidCouponsReq struct {
	StockId           string `json:"stock_id"`            // 微信为每个批次分配的唯一id。   校验规则：必须为代金券（全场券或单品券）批次号，不支持立减与折扣。 示例值：9856000
	OutRequestNo      string `json:"out_request_no"`      //商户此次发放凭据号（格式：商户id+日期+流水号），可包含英文字母，数字，|，_，*，-等内容，不允许出现其他不合法符号，商户侧需保持唯一性。   示例值： 89560002019101000121
	Appid             string `json:"appid"`               //微信为发券方商户分配的公众账号ID，接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
	StockCreatorMchid string `json:"stock_creator_mchid"` //批次创建方商户号。需要与appid关联
	CouponValue       uint64 `json:"coupon_value"`        // 非必填。 指定面额发券场景，券面额，其他场景不需要填，单位：分。
	CouponMinimum     uint64 `json:"coupon_minimum"`      // 非必填。 指定面额发券批次门槛，其他场景不需要，单位：分。
}

type MarketingFavorUsersOpenidCouponsRes struct {
	CouponId string `json:"coupon_id"`
}
