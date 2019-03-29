package dictionary

import (
	"gsms/GSMSMerchant/api/modules/dictionary"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// DictionaryHandler DictionaryHandler对象
type DictionaryHandler struct {
	container component.IContainer
	handler   dictionary.IDictionary
}

// NewDictionaryHandler 构建DictionaryHandler
func NewDictionaryHandler(container component.IContainer) (u *DictionaryHandler) {
	return &DictionaryHandler{
		container: container,
		handler:   dictionary.NewDictionary(container),
	}
}

// Handle handle
func (c *DictionaryHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------获取所有字典数据--------------------")

	ctx.Log.Info("1. 获取所有字典数据")
	list, err := c.handler.QueryAll()
	if err != nil {
		return err
	}

	ctx.Log.Info("2. 返回结果")
	return map[string]interface{}{
		"list": list,
	}
}

