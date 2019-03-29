package businesstype

import (
	"gsms/GSMSMerchant/api/modules/businesstype"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// BusinessTypeHandler 对象
type BusinessTypeHandler struct {
	container component.IContainer
	btHandler businesstype.IBusinessType
}

//NewOrderPayHandler 实例化
func NewBusinessTypeHandler(container component.IContainer) *BusinessTypeHandler {
	return &BusinessTypeHandler{
		container: container,
		btHandler: businesstype.NewBusinessType(container),
	}
}

//QueryHandle handle
func (p *BusinessTypeHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------系统业务类型查询----------------")
	ctx.Log.Info("1.校验参数")
	ctx.Log.Info("2.执行操作")
	list, err := p.btHandler.QueryAllBusnessType()
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"data": list,
	}
}
