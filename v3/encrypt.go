package v3

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// 微信APIv3 签名验签、回调数据解密、敏感数据加密解密

// RsaSignWithSha256 商户私钥签名
func RsaSignWithSha256(data, priKey string) (string, error) {
	keyBytes := []byte(priKey)
	dataBytes := []byte(data)
	h := sha256.New()
	h.Write(dataBytes)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", errors.New("rsa private key error")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// RsaVeryWithSha256 平台公钥验签
func RsaVeryWithSha256(data, signature, pubKey string) (bool, error) {
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return false, errors.New("rsa private key error")
	}
	oldSign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	pk, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	hashed := sha256.Sum256([]byte(data))
	err = rsa.VerifyPKCS1v15(pk.(*rsa.PublicKey), crypto.SHA256, hashed[:], oldSign)
	if err != nil {
		return false, err
	}
	return true, nil
}

// AesGcmDecrypt 证书和回调报文解密
func AesGcmDecrypt(ciphertext string, nonce, additionalData []byte, v3Secret string) (plaintext string, err error) {
	key := []byte(v3Secret) //key是APIv3密钥，长度32位，由管理员在商户平台上自行设置的
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	aesGcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return
	}
	cipherData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	plainData, err := aesGcm.Open(nil, nonce, cipherData, additionalData)
	if err != nil {
		return
	}
	plaintext = string(plainData)
	fmt.Println("plaintext: ", plaintext)
	return
}

// RsaEncrypt 敏感信息加密
func RsaEncrypt(plaintext []byte, pub *rsa.PublicKey) (ciphertext string, err error) {
	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pub, plaintext, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}
	ciphertext = base64.StdEncoding.EncodeToString(cipherdata)
	fmt.Printf("Ciphertext: %s\n", ciphertext)
	return
}

// RsaDecrypt 敏感信息解密
func RsaDecrypt(ciphertext string, priv *rsa.PrivateKey) (plaintext []byte, err error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	plaintext, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, cipherdata, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	return
}
