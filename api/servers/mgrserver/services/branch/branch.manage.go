package branch

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/branch"
	"gsms/GSMSMerchant/api/modules/member"
	"gsms/GSMSMerchant/api/modules/upchannel"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

//BranchManager 对象
type BranchManager struct {
	container        component.IContainer
	branch           branch.IBranch
	upchannelHandler upchannel.IUpChannel
}

//NewBranchManagerHandler 实例化
func NewBranchManagerHandler(container component.IContainer) *BranchManager {
	return &BranchManager{
		container:        container,
		branch:           branch.NewBranch(container),
		upchannelHandler: upchannel.NewUpChannel(container),
	}
}

//QueryAllHandle handle
func (u *BranchManager) QueryAllHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询门店信息 开始--------")

	var inputData branch.QueryBranchInput

	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	inputData.AccType = types.GetString(member.GetMember(ctx).AccType)
	inputData.RefAccID = types.GetString(member.GetMember(ctx).RefAccID)

	ctx.Log.Info("2.执行操作")
	data, count, err := u.branch.QueryAll(&inputData)
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询门店信息 结束--------")

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"count": count,
		"data":  data,
	}

}

//EditStatusHandle handle
func (u *BranchManager) EditStatusHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------编辑门店状态 开始--------")

	var inputData branch.EditStatusInput

	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	inputData.AccType = types.GetString(member.GetMember(ctx).AccType)
	inputData.RefAccID = types.GetString(member.GetMember(ctx).RefAccID)

	ctx.Log.Info("2.执行操作")
	err := u.branch.EditStatus(&inputData)
	if err != nil {
		return err
	}

	ctx.Log.Info("--------编辑门店状态 结束--------")

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"data": "SUCCESS",
	}

}

//AddBranchHandle 添加门店信息
func (u *BranchManager) AddBranchHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加门店信息 开始--------")

	var inputData branch.AddBranchInput

	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	inputData.AccType = types.GetString(member.GetMember(ctx).AccType)
	inputData.RefAccID = types.GetString(member.GetMember(ctx).RefAccID)

	fmt.Println("AccType:", inputData.AccType)
	ctx.Log.Info("2.执行操作")
	err := u.branch.AddBranchInfo(&inputData)
	if err != nil {
		return err
	}

	ctx.Log.Info("--------添加门店信息 结束--------")

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"data": "SUCCESS",
	}
}

//QueryProvinceHandle handle
func (u *BranchManager) QueryProvinceHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询省信息 开始--------")

	ctx.Log.Info("1.执行操作")

	data, err := u.branch.QueryProvince()
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询省信息 结束--------")
	ctx.Log.Info("2. 返回数据")
	return map[string]interface{}{
		"data": data,
	}
}

//QueryCityByProvinceHandle handle
func (u *BranchManager) QueryCityByProvinceHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询市信息 开始--------")
	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Check(); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.branch.QueryCityByProvince(ctx.Request.GetString("province"))
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询市信息 结束--------")
	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"data": data,
	}
}

//QueryDistrictByCityHandle handle
func (u *BranchManager) QueryDistrictByCityHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询区信息 开始--------")

	ctx.Log.Info("1.执行操作")
	data, err := u.branch.QueryDistrictByCity(ctx.Request.GetString("city"))
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询区信息 结束--------")
	ctx.Log.Info("2. 返回数据")
	return map[string]interface{}{
		"data": data,
	}
}

//QueryDetailInfoHandle handle
func (u *BranchManager) QueryDetailInfoHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询门店详细信息 开始--------")

	var inputData branch.DetailInput

	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	inputData.AccType = types.GetString(member.GetMember(ctx).AccType)
	inputData.RefAccID = types.GetString(member.GetMember(ctx).RefAccID)

	ctx.Log.Info("2.执行操作")
	data, err := u.branch.QueryDetailInfo(&inputData)
	if err != nil {
		return err
	}

	ctx.Log.Info("--------查询门店详细信息 结束--------")
	ctx.Log.Info("3. 返回数据")
	return data
}

//UpdateBranchInfoHandle handle
func (u *BranchManager) UpdateBranchInfoHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新门店信息 开始--------")

	var inputData branch.UpdateBranchInput

	ctx.Log.Info("1.校验参数")
	if err := ctx.Request.Bind(&inputData); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	inputData.AccType = types.GetString(member.GetMember(ctx).AccType)
	inputData.RefAccID = types.GetString(member.GetMember(ctx).RefAccID)

	ctx.Log.Info("2.执行操作")
	err := u.branch.UpdateBranchInfo(&inputData)
	if err != nil {
		return err
	}

	ctx.Log.Info("--------更新门店信息 结束--------")

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"data": "SUCCESS",
	}

}

func (u *BranchManager) GetAllBranchHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询所有的门店信息--------")
	ctx.Log.Info("1.获取渠道账号")
	accounts, err := u.upchannelHandler.QueryAllUpChannelAccountAndIDByUser(member.GetMember(ctx))
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.branch.GetAllBranchesByUser(accounts)
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 返回数据")
	return map[string]interface{}{
		"data": data,
	}
}

// GetAllBranchByAccIDHandle GetAllBranchByAccIDHandle
func (u *BranchManager) GetAllBranchByAccIDHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------查询所有的门店信息--------")

	ctx.Log.Info("1.执行操作")
	data, err := u.branch.GetAllBranchesByChannelID(types.GetString(member.GetMember(ctx).AccType), types.GetString(member.GetMember(ctx).RefAccID))

	if err != nil {
		return err
	}

	ctx.Log.Info("2. 返回数据")
	return map[string]interface{}{
		"data": data,
	}
}
