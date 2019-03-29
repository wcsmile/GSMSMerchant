package common

import "github.com/micro-plat/lib4go/types"

type RequestResponse struct {
	Response map[string]interface{} `json:"response"`
}

func (r *RequestResponse) Get(name string) string {
	return types.GetString(r.Response[name])
}
