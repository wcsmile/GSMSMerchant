package util

import "github.com/micro-plat/lib4go/utility"

func NonceStr() string {
	return utility.GetGUID()
}
