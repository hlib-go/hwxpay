package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 通知应答
func NotifyResponse(w http.ResponseWriter, err error) {
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code": "FAIL","message": "%s"}`, err.Error())))
		return
	}
	w.Write([]byte(`{"code": "SUCCESS","message": "成功"}`))
}

// 微信支付与退款通知验签
func NotifyVerifySign(cfg *Config, h *http.Header, body string) (err error) {
	signature := h.Get("Wechatpay-Signature")
	serial := h.Get("Wechatpay-Serial")
	timestamp := h.Get("Wechatpay-Timestamp")
	nonce := h.Get("Wechatpay-Nonce")

	pubKey, err := cfg.GetWxPubKey(serial)
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

type NotifyContent struct {
	Id           string                 `json:"id"`            //通知的唯一ID
	CreateTime   string                 `json:"create_time"`   //通知创建的时间
	EventType    string                 `json:"event_type"`    //通知的类型，支付成功通知的类型为TRANSACTION.SUCCESS
	ResourceType string                 `json:"resource_type"` //通知的资源数据类型，支付成功通知为encrypt-resource
	Summary      string                 `json:"summary"`       //回调摘要 示例值：支付成功
	Resource     *NotifyContentResource `json:"resource"`
}

type NotifyContentResource struct {
	Algorithm       string `json:"algorithm"`       //对开启结果数据进行加密的加密算法，目前只支持AEAD_AES_256_GCM
	Ciphertext      string `json:"ciphertext"`      //Base64编码后的开启/停用结果数据密文
	Associated_data string `json:"associated_data"` //附加数据
	Original_type   string `json:"original_type"`   //原始回调类型，为transaction
	Nonce           string `json:"nonce"`           //加密使用的随机串
}

// 解析通知内容
func NotifyDecrypt(conf *Config, data string, result interface{}) (err error) {
	var nc *NotifyContent
	err = json.Unmarshal([]byte(data), &nc)
	if err != nil {
		return
	}
	// 解密
	plaintext, err := AesGcmDecrypt(nc.Resource.Ciphertext, nc.Resource.Nonce, nc.Resource.Associated_data, conf.V3Secret)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(plaintext), &result)
	if err != nil {
		return
	}
	return
}
