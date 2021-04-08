package v3

import (
	"crypto/rsa"
	"sync"
)

// 微信支付 API SDK v3  接口调用规则文档：https://wechatpay-api.gitbook.io/wechatpay-api-v3/

// NewConfig 新建配置
func NewConfig(serviceUrl, appid, mchid, v3secret, serialNo, priKey string) (cfg *Config, err error) {
	if serviceUrl == "" {
		serviceUrl = "https://api.mch.weixin.qq.com"
	}
	cfg = &Config{
		ServiceUrl: serviceUrl,
		Appid:      appid,
		Mchid:      mchid,
		V3Secret:   v3secret,
		SerialNo:   serialNo,
		PrivateKey: priKey,
	}
	return
}

// Config 微信支付V3接口参数配置项
type Config struct {
	ServiceUrl string `json:"serviceUrl"` // 微信接口服务地址
	Appid      string `json:"appid"`      // 微信公众号appid，需要与商户有绑定关系
	Mchid      string `json:"mchid"`      // 微信支付商户号
	V3Secret   string `json:"v3Secret"`   // 微信V3秘钥 ，用于证书与回调报文解密 AEAD_AES_256_GCM
	SerialNo   string `json:"serialNo"`   // 微信商户证书序列号
	PrivateKey string `json:"privateKey"` // 微信支付商户私钥,PEM格式

	mchPriKey  *rsa.PrivateKey `json:"-"` // 微信支付商户私钥
	wxSerialNo string          `json:"-"` // 微信平台证书序号,
	wpk        sync.Map        `json:"-"` //微信平台证书序号与公钥，使用接口自动更新 ,k=序号 v=证书
}

// 商户私钥
func (c *Config) GetMchPriKey() (*rsa.PrivateKey, error) {
	if c.mchPriKey != nil {
		return c.mchPriKey, nil
	}
	mchPriKey, err := PrivateKeyPemParse(c.PrivateKey)
	if err != nil {
		return nil, err
	}
	c.mchPriKey = mchPriKey
	return c.mchPriKey, nil
}

// 微信公钥
func (c *Config) GetWxPubKey(wxSerial string) (pub *rsa.PublicKey, err error) {
	value, ok := c.wpk.Load(wxSerial)
	if ok {
		pub = value.(*rsa.PublicKey)
		return
	}
	cert, err := Certificates(c)
	if err != nil {
		return
	}
	c.wxSerialNo = cert.SerialNo
	c.wpk.Store(cert.SerialNo, cert.PublicKey)
	pub = cert.PublicKey
	return
}

func (c *Config) SetWxSerialNo() {
	cert, err := Certificates(c)
	if err != nil {
		return
	}
	c.wxSerialNo = cert.SerialNo
	c.wpk.Store(cert.SerialNo, cert.PublicKey)
}
