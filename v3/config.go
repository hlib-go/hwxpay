package v3

// 微信支付V3接口参数配置项
type Config struct {
	ServiceUrl string // 微信接口服务地址
	Mchid      string // 微信支付商户号
	V3Secret   string // 微信V3秘钥 ，用于证书与回调报文解密 AEAD_AES_256_GCM
	PrivateKey string // 微信支付商户私钥
	SerialNo   string // 微信商户证书序列号
}

// WxPublicKey 微信平台公钥，通过接口获取
func (c *Config) WxPublicKey(serial string) string {
	// serial 证书序号，用于判断当前证书是否时最新的，如果不是则更新
	return ""
}
