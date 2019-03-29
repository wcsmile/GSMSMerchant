package cachekey

import (
	"github.com/micro-plat/lib4go/transform"
	"github.com/micro-plat/lib4go/types"
)

const OrderMainPayLockCache = "recvcoupon.paylock.ordermain.@order_main_no"

func GetRealCacheKey(key string, data types.XMap) string {
	list := make([]interface{}, len(data)*2)
	idx := 0
	for k, v := range data {
		list[idx] = k
		list[idx+1] = types.GetString(v)
		idx += 2
	}

	newKey := transform.Translate(key, list...)
	return newKey
}
