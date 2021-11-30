package v3

/*
交易状态，枚举值：
SUCCESS：支付成功
REFUND：转入退款
NOTPAY：未支付
CLOSED：已关闭
REVOKED：已撤销（付款码支付）
USERPAYING：用户支付中（付款码支付）
PAYERROR：支付失败(其他原因，如银行返回失败)
*/
type TradeState string

const (
	TRADE_STATE_SUCCESS    TradeState = "SUCCESS"    //支付成功
	TRADE_STATE_REFUND     TradeState = "REFUND"     //转入退款
	TRADE_STATE_NOTPAY     TradeState = "NOTPAY"     //未支付
	TRADE_STATE_CLOSED     TradeState = "CLOSED"     //已关闭
	TRADE_STATE_REVOKED    TradeState = "REVOKED"    //已撤销（付款码支付）
	TRADE_STATE_USERPAYING TradeState = "USERPAYING" //用户支付中（付款码支付）
	TRADE_STATE_PAYERROR   TradeState = "PAYERROR"   //支付失败(其他原因，如银行返回失败)
)

// 代金券状态
type CouponStatus string

const (
	COUPON_SENDED  CouponStatus = "SENDED"  //可用
	COUPON_USED    CouponStatus = "USED"    //已实扣
	COUPON_EXPIRED CouponStatus = "EXPIRED" //已过期
)

//券类型
type CouponType string

const (
	COUPON_NORMAL CouponType = "NORMAL" // 满减券
	COUPON_CUT_TO CouponType = "CUT_TO" // 满减券
)
