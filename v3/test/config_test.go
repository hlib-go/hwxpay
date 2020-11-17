package main

import (
	"encoding/json"
	hwxpay "github.com/hlib-go/hwxpay/v3"
	"io/ioutil"
	"testing"
)

func Config() *hwxpay.Config {
	// cert 在当前目录/cert/放微信证书工具下载的私钥 ,用于测试
	wxpkBytes, err := ioutil.ReadFile("./cert/apiclient_key.pem")
	if err != nil {
		panic(err)
	}

	// 以下配置信息，从微信支付商家中心获取
	cfg, err := hwxpay.NewConfig("",
		"wx239c521c61221a8a", // 公众号的appid，需要与mchid有绑定关系
		"1579444051",
		"himktfsdfwerwergrthydrrtwerwefd2",
		"3AF43C0A75C61C12CE6423400BB4978D7B41FC0A", string(wxpkBytes))
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
