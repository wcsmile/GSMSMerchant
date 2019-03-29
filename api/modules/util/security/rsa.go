package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RsaDecrypt RSA解密
// privateKey 解密时候用到的秘钥
func RsaDecrypt(ciphertext string, privateKey string) (string, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509 ParsePKCS1PrivateKey err:%v", err)
	}
	input, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 StdEncoding DecodeString err:%v", err)
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, input)
	if err != nil {
		return "", fmt.Errorf("rsa DecryptPKCS1v15 err:%v", err)
	}

	return string(data), nil

}

func RsaEncrypt(originalData, publicKey string) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	data, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(originalData))
	if err != nil {
		return "", fmt.Errorf("rsa EncryptPKCS1v15 err:%v", err)
	}
	ouput := base64.StdEncoding.EncodeToString(data)

	return ouput, nil
}
