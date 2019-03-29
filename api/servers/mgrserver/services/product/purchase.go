package product

import (
	"gsms/GSMSMerchant/api/modules/member"
	"gsms/GSMSMerchant/api/modules/product"
	"gsms/GSMSMerchant/api/modules/upchannel"
	"io"
	"os"
	"strconv"

	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

// ProductPurchaseHandler 对象
type ProductPurchaseHandler struct {
	container        component.IContainer
	productHandler   product.IProduct
	upchannelHandler upchannel.IUpChannel
}

//NewOrderPayHandler 实例化
func NewProductPurchaseHandler(container component.IContainer) *ProductPurchaseHandler {
	return &ProductPurchaseHandler{
		container:        container,
		productHandler:   product.NewProduct(container),
		upchannelHandler: upchannel.NewUpChannel(container),
	}
}

//QueryHandle handle
func (p *ProductPurchaseHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------商品采购列表查询----------------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check("pi", "ps", "start_time", "end_time"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.获取渠道账号")
	accounts, err := p.upchannelHandler.QueryAllUpChannelAccountAndIDByUser(member.GetMember(ctx))
	if err != nil {
		return
	}

	//accounts := "xpktm,grsXmXing"
	//accounts := fmt.Sprintf("'%s'", strings.Join(strings.Split(infos, ","), "','"))
	ctx.Log.Info("3.构建参数")
	input := map[string]interface{}{
		"pi":              ctx.Request.GetString("pi"),
		"ps":              ctx.Request.GetString("ps"),
		"start_time":      ctx.Request.GetString("start_time"),
		"end_time":        ctx.Request.GetString("end_time"),
		"need_logistics":  ctx.Request.GetString("need_logistics"),
		"down_channel_no": accounts,
		"card_no":         ctx.Request.GetString("card_no"),
	}

	ctx.Log.Info("4.执行操作")
	list, count, err := p.productHandler.QueryPurchaseList(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("5.查询下游渠道产品")
	productlist, err := p.productHandler.QueryDownProductListByChannelNos(accounts)
	if err != nil {
		return err
	}
	mergeProductFace(list, productlist)

	ctx.Log.Info("6. 返回数据")
	return map[string]interface{}{
		"count": count,
		"data":  list,
	}
}

func mergeProductFace(list, productlist db.QueryRows) {
	for k, v := range list {
		list[k]["per_price"] = ""
		list[k]["face"] = ""
		for _, v2 := range productlist {
			if v.GetString("down_product_no") == v2.GetString("product_no") {
				face, _ := strconv.ParseFloat(v2.GetString("face"), 64)
				discount, _ := strconv.ParseFloat(v2.GetString("deduct_discount"), 64)
				list[k]["per_price"] = face * discount
				list[k]["face"] = v2.GetString("face")
			}
		}

	}
}

func (p *ProductPurchaseHandler) CardsHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------商品采购卡号列表查询----------------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check("order_no", "pi", "ps"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	input := map[string]interface{}{
		"order_no": ctx.Request.GetString("order_no"),
		"pi":       ctx.Request.GetString("pi"),
		"ps":       ctx.Request.GetString("ps"),
	}

	ctx.Log.Info("3.执行操作")
	list, count, err := p.productHandler.QueryPurcharseCardNoList(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("4. 返回数据")
	return map[string]interface{}{
		"count": count,
		"data":  list,
	}
}

//QueryHandle handle
func (p *ProductPurchaseHandler) ExportHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------------商品采购导出卡----------------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check([]string{"order_no"}...); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	res, err := p.productHandler.QueryPurcharseExportCardNoList(ctx.Request.GetString("order_no"))
	if err != nil {
		return err
	}

	ctx.Log.Info("2.构造返回的excel", res)
	excelFile := excelize.NewFile()
	sheet1 := excelFile.NewSheet("Sheet1")
	excelFile.SetCellValue("Sheet1", "A1", "卡号")
	for index, item := range res {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%d", index+2), item.GetString("recharge_account_id"))
	}
	excelFile.SetColWidth("Sheet1", "A", "A", 40)
	excelFile.SetActiveSheet(sheet1)
	fileName := fmt.Sprintf("%s.xlsx", time.Now().Format("20060102150405"))
	err = excelFile.SaveAs(fmt.Sprintf("./%s", fileName))
	if err != nil {
		ctx.Log.Errorf("导出卡号,保存生成的文件异常,err:%+v", err)
		return fmt.Errorf("导出卡号异常")
	}

	response, err := ctx.Request.Http.GetResponse()
	if err != nil {
		ctx.Log.Errorf("导出卡号,GetResponse失败:err:%+v", err)
		return
	}
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Header().Set("Content-Disposition", "attachment;filename=卡号.xlsx")

	fileData, err := os.Open(fileName)
	if err != nil {
		ctx.Log.Errorf("导出卡号,打开生成的文件异常,err:%+v", err)
		return fmt.Errorf("导出卡号异常")
	}

	_, err = io.Copy(response, fileData)
	if err != nil {
		ctx.Log.Errorf("导出卡号,复制文件内容到response异常,err:%+v", err)
		return fmt.Errorf("导出卡号异常")
	}

	defer func() {
		fileData.Close()
		os.Remove(fmt.Sprintf("./%s", fileName))
	}()

	ctx.Log.Info("4. 返回数据")
	return
}
