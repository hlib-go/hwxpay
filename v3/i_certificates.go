package v3

import (
	"io/ioutil"
	"log"
	"time"
)

/*
文档：https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu

获取商户当前可用的平台证书列表。微信支付提供该接口，帮助商户后台系统实现平台证书的平滑更换。

注意事项
如果自行实现验证平台签名逻辑的话，需要注意以下事项:
程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
定期调用该接口，间隔时间小于12 小时
加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）

说明：
在启用新的平台证书前，微信支付会提前24小时把新证书加入到平台证书列表中
接口的频率限制: 单个商户号1000 次/s
首次下载证书，可以使用微信支付提供的证书下载工具
*/

// Certificates 微信平台证书列表
func Certificates(cfg *Config) (cert *Cert, err error) {
	method := "GET"
	path := "/v3/certificates"
	body := ""
	authorization, err := Authorization(cfg, method, path, body)
	if err != nil {
		return
	}
	resp, err := Request(cfg.ServiceUrl+path, method, authorization, body)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Println("获取证书列表:" + string(bytes))
	// 最新证书缓存到全局变量
	return
}

type Cert struct {
	PubKey        string
	SerialNo      string
	Algorithm     string
	EffectiveTime time.Time
	ExpireTime    time.Time
}
