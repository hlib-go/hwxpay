package v3

import (
	"errors"
	"net/http"
)

// 微信支付与退款通知验签
func NotifyVerifySign(cfg *Config, h *http.Header, body string) (err error) {
	signature := h.Get("Wechatpay-Signature")
	serial := h.Get("Wechatpay-Serial")
	timestamp := h.Get("Wechatpay-Timestamp")
	nonce := h.Get("Wechatpay-Nonce")

	pubKey, err := GetWxPublicKey(cfg, serial)
	if err != nil {
		return
	}
	ok, err := V3SignVery(signature, timestamp, nonce, body, pubKey)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("WX_NOTIFY:签名校验失败")
		return
	}
	return
}

// 对通知密文解密
