package v3test

import (
	"encoding/json"
	hwxpay "hwxpay/v3"
	"io/ioutil"
	"testing"
)

func Config() *hwxpay.Config {
	wxpkBytes, err := ioutil.ReadFile("./cert/apiclient_key.pem")
	if err != nil {
		panic(err)
	}

	cfg, err := hwxpay.NewConfig("",
		"wx239c521c61221a8a",
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
