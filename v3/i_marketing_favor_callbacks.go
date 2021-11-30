package v3

// 设置消息通知地址  https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_1_12.shtml
func MarketingFavorCallbacks(cfg *Config, params *MarketingFavorCallbacksParams) (result *MarketingFavorCallbacksResult, err error) {
	err = POST(cfg, "/v3/marketing/favor/callbacks", params, &result)
	return
}

type MarketingFavorCallbacksParams struct {
	Mchid     string `json:"mchid"`            //该商户号必须为创建商户号进行配置
	NotifyUrl string `json:"notify_url"`       //不能携带参数
	Switch    bool   `json:"switch,omitempty"` //true：开启推送  false：停止推送
}

type MarketingFavorCallbacksResult struct {
	NotifyUrl  string `json:"notify_url"`
	UpdateTime string `json:"update_time"` // 修改时间，遵循rfc3339标准格式
}
