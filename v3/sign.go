package v3

import (
	"log"
)

// V3Sign 计算请求报文签名
func V3Sign(method, path, body, timestamp, nonceStr, priKey string) (sign string, err error) {
	targetStr := method + "\n" + path + "\n" + timestamp + "\n" + nonceStr + "\n" + body + "\n"
	log.Println("签名原始字符串：\n" + targetStr)
	sign, err = RsaSignWithSha256(targetStr, priKey)
	log.Println("签名结果字符串：" + sign)
	return
}

// V3SignVery 验签响应报文签名
func V3SignVery(signature, time, nonce, body string, pubKey string) (ok bool, err error) {
	data := time + "\n" + nonce + "\n" + body + "\n"
	return RsaVeryWithSha256(data, signature, pubKey)
}
