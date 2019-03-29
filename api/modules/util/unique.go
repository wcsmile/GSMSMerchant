package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/micro-plat/lib4go/security/md5"
)

//0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
const str = "0123456789abcdefghijklmnopqrstuvwxyz"
const staticKey = "6789abcdefghijklmnop"

func BuildUnique(key string, length int) string {
	if length <= 0 {
		return ""
	}
	key = staticKey + key
	code := RandomStr(length)
	now := time.Now()
	year := now.Year()
	mon := int(now.Month())
	day := now.Day()

	yearVal := str[year%len(str)]
	monVal := str[mon]
	dayVal := str[day]

	newCode := string(yearVal) + string(monVal) + string(dayVal) + code

	validCode := calcValidateCode(newCode, key)
	//println("code:", code)
	//println("validCode:", validCode)
	return newCode + validCode
}

//ValidateUnique ValidateUnique
func ValidateUnique(code, key string) bool {
	key = staticKey + key

	val := code[0 : len(code)-1]
	vc := code[len(code)-1 : len(code)]
	nvc := calcValidateCode(val, key)
	return strings.EqualFold(vc, nvc)
}

func calcValidateCode(val, key string) string {
	md5Val := md5.Encrypt(val + key)
	//fmt.Println("len(md5Val):", len(md5Val))
	idx := val[0] % byte(len(md5Val))
	//fmt.Println("idx:", idx)
	vc := md5Val[idx]
	//fmt.Println("vc:", vc)
	return string(vc)
}

//RandomStr 随机生成字符串
func RandomStr(length int) string {

	bytes := []byte(str)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func OriginStr(compStr string, randStrLen int) (str string, err error) {
	if len(compStr) < randStrLen {
		return "", fmt.Errorf("随机串解析出错，解析串：%s", compStr)
	}
	str = compStr[0:(len(compStr) - randStrLen - 1)]
	return str, nil
}
