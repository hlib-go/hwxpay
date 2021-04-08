package v3

/*
提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，建议在提交退款申请后1分钟发起查询退款状态，一般来说零钱支付的退款5分钟内到账，银行卡支付的退款1-3个工作日到账。
*/

// 查询单笔退款API
func RefundQuery(conf *Config, out_refund_no string) (r RefundResult, err error) {
	err = GET(conf, "/v3/refund/domestic/refunds/"+out_refund_no, nil, &r)
	return
}
