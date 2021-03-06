package util

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"
)

//MakeSign .
func MakeSign(mReq map[string]interface{}, secret string) (sign string) {
	fmt.Println("签名计算, API KEY:", secret)
	//1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	//2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//3, 在键值对的最后加上key=secret
	if secret != "" {
		signStrings = signStrings + "key=" + secret
	}

	//4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

//CreateCertClient CreateCertClient
func CreateCertClient(certFile string, keyFile string) (client *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       tlsConfig,
		},
		Timeout: 5 * time.Second,
	}, nil
}
