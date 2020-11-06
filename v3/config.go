package v3

import (
	"crypto/rsa"
)

// 微信支付 API SDK v3  接口调用规则文档：https://wechatpay-api.gitbook.io/wechatpay-api-v3/

// Config 微信支付V3接口参数配置项
type Config struct {
	ServiceUrl string          // 微信接口服务地址
	Appid      string          // 微信公众号appid，需要与商户有绑定关系
	Mchid      string          // 微信支付商户号
	V3Secret   string          // 微信V3秘钥 ，用于证书与回调报文解密 AEAD_AES_256_GCM
	SerialNo   string          // 微信商户证书序列号
	PrivateKey *rsa.PrivateKey // 微信支付商户私钥

	//私有属性：微信平台证书序号与公钥，使用接口自动更新,微信要求12小时内更新一次
	wxSerialNo  string
	wxPublicKey *rsa.PublicKey
}

// NewConfig 新建配置
func NewConfig(serviceUrl, appid, mchid, v3secret, serialNo, priKey string) (cfg *Config, err error) {
	if serviceUrl == "" {
		serviceUrl = "https://api.mch.weixin.qq.com"
	}
	privateKey, err := PrivateKeyPemParse(priKey)
	if err != nil {
		return
	}
	cfg = &Config{
		ServiceUrl: serviceUrl,
		Appid:      appid,
		Mchid:      mchid,
		V3Secret:   v3secret,
		SerialNo:   serialNo,
		PrivateKey: privateKey,
	}
	// 创建Config对象时，加载微信平台公钥
	err = cfg.LoadWxPublicKey()
	if err != nil {
		return
	}
	return
}

// WxPublicKey 微信平台最新公钥，通过接口获取
func (c *Config) WxPublicKey(wxSerial string) (pub *rsa.PublicKey, err error) {
	// serial 接口响应的证书序号，用于判断当前证书是否最新的，如果不是则更新
	if c.SerialNo == wxSerial {
		pub = c.wxPublicKey
		return
	}
	// 缓存不存在时，接口查询微信公钥
	err = c.LoadWxPublicKey()
	if err != nil {
		return
	}
	pub = c.wxPublicKey
	return
}

func (c *Config) LoadWxPublicKey() (err error) {
	cert, err := Certificates(c)
	if err != nil {
		return
	}
	c.wxSerialNo = cert.SerialNo
	c.wxPublicKey = cert.PublicKey
	return
}
