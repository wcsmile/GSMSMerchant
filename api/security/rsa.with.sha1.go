package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

func RsaSignWithSha1(ciphertext string, privateKey string) (result string, err error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		err = errors.New("private key error")
		return
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("x509 ParsePKCS1PrivateKey err:%v", err)
		return
	}
	t := sha1.New()
	//fmt.Println("RsaSignWithSha1:", ciphertext)
	io.WriteString(t, ciphertext)
	digest := t.Sum(nil)
	data, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, digest)
	if err != nil {
		return
	}
	result = base64.StdEncoding.EncodeToString(data)
	return

}

func RsaVerifyWithSha1(originalData, signdata, publicKey string) (isMatch bool, err error) {

	//步骤1，加载RSA的公钥
	block, _ := pem.Decode([]byte(publicKey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err = fmt.Errorf("RsaVerifyWithSha1.1.:%+v", err)
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	data, _ := base64.StdEncoding.DecodeString(signdata)

	t := sha1.New()
	io.WriteString(t, originalData)
	digest := t.Sum(nil)
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	if err != nil {
		err = fmt.Errorf("RsaVerifyWithSha1.2.:%+v", err)
		return
	}
	isMatch = true
	return
}

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
