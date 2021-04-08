package v3

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*var (
	_emap = map[int]error{
		204: errors.New("204微信处理成功，无返回Body"),
		400: errors.New("400微信协议或者参数非法"),
		401: errors.New("401微信签名验证失败"),
		403: errors.New("403微信权限异常"),
		404: errors.New("404微信请求的资源不存在"),
		429: errors.New("429微信请求超过频率限制"),
		500: errors.New("500微信系统错误"),
		502: errors.New("502微信服务下线，暂时不可用"),
		503: errors.New("503微信服务不可用，过载保护"),
	}
)*/

func GET(cfg *Config, path string, i interface{}, o interface{}) (err error) {
	err = Call(cfg, path, "GET", i, o)
	return
}

func POST(cfg *Config, path string, i interface{}, o interface{}) (err error) {
	err = Call(cfg, path, "POST", i, o)
	return
}

// Call 调用接口方法
func Call(cfg *Config, path, method string, i interface{}, o interface{}) (err error) {
	var (
		reqBody string
		resBody string
	)
	defer func() {
		log.Println("微信请求接口：", cfg.ServiceUrl+path)
		log.Println("微信请求报文：", reqBody)
		log.Println("微信响应报文：", resBody)
		if err != nil {
			log.Println("微信响应错误：", err.Error())
		}
	}()

	if i != nil {
		reqBytes, err := json.Marshal(i)
		if err != nil {
			return err
		}
		reqBody = string(reqBytes)
	}
	authorization, err := Authorization(cfg, method, path, reqBody)
	if err != nil {
		return
	}
	resp, err := Request(cfg, path, method, authorization, reqBody)
	if err != nil {
		return
	}
	resBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	resBody = string(resBytes)

	// HTTP 返回204，处理成功，应答无内容
	if resp.StatusCode == 204 {
		return
	}

	// HTTP 返回非200，直接返回错误
	if resp.StatusCode != 200 {
		var eres *ErrResponse
		json.Unmarshal(resBytes, &eres)
		err = errors.New(eres.Code + ":" + eres.Message)
		return
	}

	//requestId := resp.Header.Get("Request-ID")
	signature := resp.Header.Get("Wechatpay-Signature")
	serial := resp.Header.Get("Wechatpay-Serial")
	timestamp := resp.Header.Get("Wechatpay-Timestamp")
	nonce := resp.Header.Get("Wechatpay-Nonce")

	pubKey, err := cfg.GetWxPubKey(serial)
	if err != nil {
		return
	}
	ok, err := V3SignVery(signature, timestamp, nonce, resBody, pubKey)
	if err != nil {
		return
	}
	if !ok {
		return errors.New("签名校验失败")
	}
	err = json.Unmarshal(resBytes, o)
	if err != nil {
		return
	}
	return
}

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Request 发送接口请求
func Request(cfg *Config, path, method, authorization, body string) (resp *http.Response, err error) {
	request, err := http.NewRequest(method, cfg.ServiceUrl+path, strings.NewReader(body))
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", "APIv3;golang")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", authorization)
	request.Header.Set("Wechatpay-Serial", cfg.wxSerialNo) // 存在敏感字段加密时必填，其它场景可选
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	return
}

// Authorization 拼接权限验证字符串
func Authorization(cfg *Config, method, path, body string) (authorization string, err error) {
	authType := "WECHATPAY2-SHA256-RSA2048" //固定字符串
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonceStr := RandomString(32)
	priKey, err := cfg.GetMchPriKey()
	if err != nil {
		return
	}
	signature, err := V3Sign(method, path, body, timestamp, nonceStr, priKey)
	if err != nil {
		return
	}
	authorization = fmt.Sprintf(`%s mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`,
		authType, cfg.Mchid, nonceStr, signature, timestamp, cfg.SerialNo)
	return
}

// RandomString 生成随机字符串
func RandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
