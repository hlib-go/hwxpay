package main

import (
	"encoding/json"
	"github.com/hlib-go/hwxpay/v3"
	"io/ioutil"
	"testing"
)

func Config() *v3.Config {
	// cert 在当前目录/cert/放微信证书工具下载的私钥
	wxpkBytes, err := ioutil.ReadFile("/run/wxmch-1579635811/apiclient_key.pem")
	if err != nil {
		panic(err)
	}

	// 以下配置信息，从微信支付商家中心获取
	cfg, err := v3.NewConfig("",
		"wx239c521c61221a8a", // 公众号的appid，需要与mchid有绑定关系
		"1579635811",
		"u943wfsdfwerwergrthydrrtwerwefd2",
		"78E20F4064DDBE72537C194DE0C089CBF2D14447",
		string(wxpkBytes))
	if err != nil {
		panic(err)
	}
	return cfg
}

func TestConfig(t *testing.T) {
	cfg := Config()
	bytes, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	t.Log(string(bytes))
}
