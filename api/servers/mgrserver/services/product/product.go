package product

import (
	"gsms/GSMSMerchant/api/modules/card"
	"gsms/GSMSMerchant/api/modules/member"
	"gsms/GSMSMerchant/api/modules/product"
	"gsms/GSMSMerchant/api/modules/upchannel"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

// ProductHandler 对象
type ProductHandler struct {
	container        component.IContainer
	productHandler   product.IProduct
	upchannelHandler upchannel.IUpChannel
	cardHandler      card.ICard
}

//NewOrderPayHandler 实例化
func NewProductHandler(container component.IContainer) *ProductHandler {
	return &ProductHandler{
		container:        container,
		productHandler:   product.NewProduct(container),
		upchannelHandler: upchannel.NewUpChannel(container),
		cardHandler:      card.NewCard(container),
	}
}

//QueryHandle handle
func (p *ProductHandler) CardQueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------芯片卡商品列表查询----------------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check("pi", "ps", "start_time", "end_time"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.获取渠道账号")
	accounts, err := p.upchannelHandler.QueryAllUpChannelAccountAndIDByUser(member.GetMember(ctx))
	if err != nil {
		return err
	}
	// accounts := "xpktm,grsXmXing"
	ctx.Log.Info("3.构建参数", accounts)
	input := map[string]interface{}{
		"pi":                  ctx.Request.GetString("pi"),
		"ps":                  ctx.Request.GetString("ps"),
		"start_time":          ctx.Request.GetString("start_time"),
		"end_time":            ctx.Request.GetString("end_time"),
		"recharge_account_id": ctx.Request.GetString("card_no"),
		"order_source":        ctx.Request.GetString("order_source"),
		"down_channel_no":     accounts,
	}

	ctx.Log.Info("4.执行操作", input)
	list, count, err := p.productHandler.QuerydCardList(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("5.查询卡片列表")
	temp := []string{}
	for _, v := range list {
		temp = append(temp, v.GetString("recharge_account_id"))
	}
	cardNos := strings.Join(temp, ",")
	cardList, err := p.cardHandler.QueryCardInfoByCardNo(cardNos)
	if err != nil {
		return err
	}
	ctx.Log.Info("6.查询下游渠道产品")
	productlist, err := p.productHandler.QueryDownProductListByChannelNos(accounts)
	if err != nil {
		return err
	}
	mergeCardFace(list, productlist, cardList)

	ctx.Log.Info("7. 返回数据")
	return map[string]interface{}{
		"count": count,
		"data":  list,
	}
}

func mergeCardFace(list, productlist, cardList db.QueryRows) {
	for k, v := range list {
		list[k]["face"] = ""
		list[k]["has_first_recharge"] = ""
		list[k]["status"] = ""
		for _, v2 := range productlist {
			if v.GetString("down_product_no") == v2.GetString("product_no") {
				list[k]["face"] = v2.GetString("face")
			}
		}
		for _, v3 := range cardList {
			if v.GetString("recharge_account_id") == v3.GetString("card_no") {
				list[k]["has_first_recharge"] = v3.GetString("has_first_recharge")
				list[k]["status"] = v3.GetString("status")
			}
		}

	}
}

//QueryHandle handle
func (p *ProductHandler) RechargeQueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------充值商品列表查询----------------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check("pi", "ps"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.获取渠道账号")
	accounts, err := p.upchannelHandler.QueryAllUpChannelAccountAndIDByUser(member.GetMember(ctx))
	if err != nil {
		return err
	}
	// accounts := "xpktm,grsXmXing"
	ctx.Log.Info("3.构建参数", accounts)
	input := map[string]interface{}{
		"pi":              ctx.Request.GetString("pi"),
		"ps":              ctx.Request.GetString("ps"),
		"down_channel_no": accounts,
	}

	ctx.Log.Info("4.执行操作", input)
	list, count, err := p.productHandler.QuerydRechargeList(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("5. 返回数据")
	return map[string]interface{}{
		"count": count,
		"data":  list,
	}
}
