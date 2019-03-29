package security

import (
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/net"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"
)

//Sign 签名
func Sign(key string, input map[string]interface{}) (sign string) {
	values := net.NewValues()
	for k, v := range input {
		vv := types.GetString(v)
		if vv == "" {
			continue
		}
		if strings.Contains(strings.ToLower(vv), "e+") {
			values.Set(k, fmt.Sprintf("%d", types.GetInt64(v)))
		} else {
			values.Set(k, types.GetString(v))
		}
	}

	values.Sort()
	raw := values.Join("=", "&")
	//raw = key + raw + key
	fmt.Println("SignRaw:", raw)
	raw = raw + key
	return md5.Encrypt(raw)
}

//Verify 验证签名
func Verify(key string, input map[string]interface{}) bool {
	outSign := types.GetString(input["sign"])
	delete(input, "sign")
	innerSign := Sign(key, input)

	fmt.Println("outSign:", outSign)
	fmt.Println("innerSign:", innerSign)
	return strings.EqualFold(outSign, innerSign)
}
