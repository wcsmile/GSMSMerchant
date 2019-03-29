package capitalrecord

import (
	"gsms/GSMSMerchant/api/modules/capitalrecord"
	"gsms/GSMSMerchant/api/modules/member"
	"gsms/GSMSMerchant/api/modules/upchannel"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// CapitalRecordHandler 对象
type CapitalRecordHandler struct {
	container        component.IContainer
	recordHandler    capitalrecord.ICapitalrecord
	upchannelHandler upchannel.IUpChannel
	fieldsReceive    []string
}

//NewOrderPayHandler 实例化
func NewCapitalRecordHandler(container component.IContainer) *CapitalRecordHandler {
	return &CapitalRecordHandler{
		container:        container,
		recordHandler:    capitalrecord.NewCapitalrecord(container),
		upchannelHandler: upchannel.NewUpChannel(container),
		fieldsReceive:    []string{"pi", "ps"},
	}
}

//QueryHandle handle
func (p *CapitalRecordHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------资金（佣金）流水列表查询----------------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check(p.fieldsReceive...); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	userInfo := member.GetMember(ctx)

	input := map[string]interface{}{
		"pi":            ctx.Request.GetString("pi"),
		"ps":            ctx.Request.GetString("ps"),
		"start_time":    ctx.Request.GetString("start_time"),
		"end_time":      ctx.Request.GetString("end_time"),
		"business_type": ctx.Request.GetString("record_type"),
		"record_no":     ctx.Request.GetString("record_no"),
	}

	ctx.Log.Info("2.执行操作", input, userInfo)
	list, count, err := p.recordHandler.QueryRecordList(userInfo, input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"count": count,
		"list":  list,
	}
}
