package order

import (
	"gsms/GSMSMerchant/api/modules/member"
	"gsms/GSMSMerchant/api/modules/order"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

//OrderManager 对象
type OrderManager struct {
	container component.IContainer
	order     order.IOrder
}

//NewOrderManagerHandler 实例化
func NewOrderManagerHandler(container component.IContainer) *OrderManager {
	return &OrderManager{
		container: container,
		order:     order.NewOrder(container),
	}
}

//QueryHandle handle
func (u *OrderManager) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询订单信息 开始--------")
	var inputData order.QueryOrderInput

	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	inputData.AccType = types.GetString(member.GetMember(ctx).AccType)
	inputData.RefAccID = types.GetString(member.GetMember(ctx).RefAccID)

	ctx.Log.Info("2.执行操作")
	data, count, err := u.order.Query(&inputData)
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询订单信息 结束--------")

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"count": count,
		"data":  data,
	}

}

//QueryChannelHandle 查询渠道
func (u *OrderManager) QueryChannelHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询渠道 开始--------")

	ctx.Log.Info("1.执行操作")
	data, err := u.order.QueryChannel()
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询订单来源 结束--------")

	ctx.Log.Info("2. 返回数据")
	return map[string]interface{}{
		"data": data,
	}

}
