package branch

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/errorcode"
	"gsms/GSMSMerchant/api/modules/const/sql"
	"gsms/GSMSMerchant/api/modules/const/usertype"
	"gsms/GSMSMerchant/api/modules/member"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//IDbBranch 方法
type IDbBranch interface {
	QueryAll(inputData *QueryBranchInput) (data db.QueryRows, count int, err error)
	EditStatus(inputData *EditStatusInput) (err error)
	AddBranchInfo(inputData *AddBranchInput) (err error)
	QueryProvince() (data db.QueryRows, err error)
	QueryCityByProvince(province string) (data db.QueryRows, err error)
	QueryDistrictByCity(city string) (data db.QueryRows, err error)
	QueryDetailInfo(inputData *DetailInput) (data db.QueryRows, err error)
	UpdateBranchInfo(inputData *UpdateBranchInput) (err error)

	GetAllBranchesByUser(channelNos string) (data db.QueryRows, err error)

	GetBranchIDsByName(name string) (branchIDs string, err error)
	GetAllBranchesByChannelID(AccType, RefAccID string) (data db.QueryRows, err error)
}

// DbBranch 对象
type DbBranch struct {
	c component.IContainer
}

// NewDbBranch 实例化
func NewDbBranch(c component.IContainer) *DbBranch {
	return &DbBranch{
		c: c,
	}
}

// QueryAll 查询门店信息
func (d *DbBranch) QueryAll(inputData *QueryBranchInput) (data db.QueryRows, count int, err error) {
	db := d.c.GetRegularDB("industry_marketing")
	var rowsdate []types.XMap
	var totals int

	params := map[string]interface{}{
		"status":    inputData.Status,
		"pi":        inputData.Pi,
		"ps":        inputData.Ps,
		"branch_id": inputData.BranchID,
	}

	user := member.LoginAdminState{
		AccType:  inputData.AccType,
		RefAccID: inputData.RefAccID,
	}
	fmt.Println("user:", user)

	if strings.EqualFold(inputData.AccType, usertype.Agent) { //代理
		params["agent_id"] = inputData.RefAccID
		total, sq, sa, err := db.Scalar(sql.QueryAgentBranchInfoCount, params)

		if err != nil || types.GetInt(total) < 0 {
			return nil, 0, fmt.Errorf("获取查询代理门店信息条数发生错误(err:%v),sql:%s,输入参数:%v", err, sq, sa)
		}

		// 判断返回的数量是否为空
		count = types.GetInt(total, 0)
		if count <= 0 {
			return nil, count, nil
		}

		datas, q, a, err := db.Query(sql.QueryAgentBranchInfo, params)
		if err != nil {
			return nil, 0, fmt.Errorf("获取查询代理门店信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		rowsdate = datas
		totals = types.GetInt(total)

	} else if strings.EqualFold(inputData.AccType, usertype.Channel) { //渠道
		params["channel_ids"] = inputData.RefAccID
		total, sq, sa, err := db.Scalar(sql.QueryChannelBranchInfoCount, params)
		if err != nil || types.GetInt(total) < 0 {
			return nil, 0, fmt.Errorf("获取查询渠道门店信息条数发生错误(err:%v),sql:%s,输入参数:%v", err, sq, sa)
		}

		// 判断返回的数量是否为空
		count = types.GetInt(total, 0)
		if count <= 0 {
			return nil, count, nil
		}

		datas, q, a, err := db.Query(sql.QueryChannelBranchInfo, params)
		if err != nil {
			return nil, 0, fmt.Errorf("获取查询渠道门店信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		rowsdate = datas
		totals = types.GetInt(total)
	}
	return rowsdate, totals, nil
}

// EditStatus 编辑门店状态
func (d *DbBranch) EditStatus(inputData *EditStatusInput) (err error) {
	db := d.c.GetRegularDB("industry_marketing")

	params := map[string]interface{}{
		"status":    inputData.Status,
		"branch_id": inputData.BranchID,
	}

	if strings.EqualFold(inputData.AccType, usertype.Agent) { //代理

		total, q, a, err := db.Execute(sql.EditAgentStatusInfo, params)
		if err != nil || total <= 0 {
			return fmt.Errorf("编辑代理门店状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}

	} else if strings.EqualFold(inputData.AccType, usertype.Channel) { //渠道

		total, q, a, err := db.Execute(sql.EditChannelStatusInfo, params)
		if err != nil || total <= 0 {
			return fmt.Errorf("编辑渠道门店状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	return nil
}

// AddBranchInfo 添加门店信息
func (d *DbBranch) AddBranchInfo(inputData *AddBranchInput) (err error) {
	// 打开事物
	db := d.c.GetRegularDB("industry_marketing")
	dbTrans, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("添加门店信息,开启事务失败:err:%+v", err)
		return
	}

	params := map[string]interface{}{
		"branch_name":  inputData.BranchName,
		"contact_name": inputData.Contactname,
		"contact_tel":  inputData.ContactTel,
		"province":     inputData.Province,
		"city":         inputData.City,
		"district":     inputData.District,
		"address":      inputData.Address,
		"channel_id":   inputData.RefAccID,
	}

	if inputData.AccType == usertype.Agent { //代理

		// 当前代理下面是否存在相同门店
		count, q1, a1, err1 := dbTrans.Scalar(sql.IsExitAgentBranchNameInfo, map[string]interface{}{
			"branch_name": inputData.BranchName,
			"agent_id":    inputData.RefAccID,
		})
		if err1 != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加代理门店信息发生错误(err:%v),sql:%s,输入参数:%v", err1, q1, a1)
		}
		if types.GetInt(count) > 0 {
			dbTrans.Rollback()
			return context.NewError(errorcode.HTTPErrorBranchNameError, fmt.Errorf("门店名称已经存在")) // 返回给前端的error
		}

		total, q, a, err := db.Execute(sql.AddAgentBranchInfo, params)
		if err != nil || total <= 0 {
			dbTrans.Rollback()
			return fmt.Errorf("添加代理门店信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}

	} else if inputData.AccType == usertype.Channel { //渠道
		// 当前代理商渠道下面是否存在相同门店
		count, q1, a1, err1 := dbTrans.Scalar(sql.IsExitsChannelBranchNameInfo, map[string]interface{}{
			"branch_name":   inputData.BranchName,
			"up_channel_id": inputData.RefAccID,
		})
		if err1 != nil {
			dbTrans.Rollback()
			return fmt.Errorf("添加渠道门店信息发生错误(err:%v),sql:%s,输入参数:%v", err1, q1, a1)
		}
		if types.GetInt(count) > 0 {
			dbTrans.Rollback()
			return context.NewError(errorcode.HTTPErrorBranchNameError, fmt.Errorf("门店名称已经存在")) // 返回给前端的error
		}

		total, q, a, err := db.Execute(sql.AddChannelBranchInfo, params)
		if err != nil || total <= 0 {
			dbTrans.Rollback()
			return fmt.Errorf("渠道账户添加门店信息发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}

	dbTrans.Commit()
	return nil
}

// QueryDetailInfo 查询门店详细信息
func (d *DbBranch) QueryDetailInfo(inputData *DetailInput) (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("industry_marketing")
	var datarows []types.XMap

	if strings.EqualFold(inputData.AccType, usertype.Agent) { //代理

		data, q, a, err := db.Query(sql.QueryAgentDetailInfo, map[string]interface{}{
			"branch_id": inputData.BranchID,
		})
		if err != nil {
			return nil, fmt.Errorf("查询代理门店详细信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		if len(data) == 0 {
			return nil, nil
		}
		datarows = data

	} else if strings.EqualFold(inputData.AccType, usertype.Channel) { //渠道
		data, q, a, err := db.Query(sql.QueryChannelDetailInfo, map[string]interface{}{
			"branch_id": inputData.BranchID,
		})
		if err != nil {
			return nil, fmt.Errorf("查询渠道门店详细信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		if len(data) == 0 {
			return nil, nil
		}
		datarows = data
	}

	return datarows, nil
}

// UpdateBranchInfo 更新门店信息
func (d *DbBranch) UpdateBranchInfo(inputData *UpdateBranchInput) (err error) {

	// 打开事物
	db := d.c.GetRegularDB("industry_marketing")
	dbTrans, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("更新门店信息,开启事务失败:err:%+v", err)
		return
	}
	params := map[string]interface{}{
		"branch_id":    inputData.BranchID,
		"contact_name": inputData.Contactname,
		"contact_tel":  inputData.ContactTel,
		"province":     inputData.Province,
		"city":         inputData.City,
		"district":     inputData.District,
		"address":      inputData.Address,
	}

	if strings.EqualFold(inputData.AccType, usertype.Agent) { //代理
		total, q, a, err := dbTrans.Execute(sql.UpdateAgentBranchInfo, params)
		if err != nil || total <= 0 {
			dbTrans.Rollback()
			return fmt.Errorf("编辑代理门店状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	} else if strings.EqualFold(inputData.AccType, usertype.Channel) { //渠道

		total, q, a, err := dbTrans.Execute(sql.UpdateChannelBranchInfo, params)
		if err != nil || total <= 0 {
			dbTrans.Rollback()
			return fmt.Errorf("编辑渠道门店状态发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
		}
	}
	dbTrans.Commit()
	return nil
}

// GetAllBranchesByChannelID GetAllBranchesByChannelID
func (d *DbBranch) GetAllBranchesByChannelID(AccType, RefAccID string) (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("industry_marketing")
	var rowsdate []types.XMap

	if strings.EqualFold(AccType, usertype.Agent) { //代理
		data, q, a, err := db.Query(sql.GetAllAgentBranchesByChannelID, map[string]interface{}{
			"channel_id": RefAccID,
		})

		if err != nil {
			return nil, fmt.Errorf("查询代理用户所有的门店信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		rowsdate = data
	} else if strings.EqualFold(AccType, usertype.Channel) { //渠道
		data, q, a, err := db.Query(sql.GetAllChannelBranchesByChannelID, map[string]interface{}{
			"up_channel_id": RefAccID,
		})

		if err != nil {
			return nil, fmt.Errorf("查询渠道用户所有的门店信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
		}
		rowsdate = data
	}

	return rowsdate, nil
}

// QueryProvince 查询省信息
func (d *DbBranch) QueryProvince() (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("industry_marketing")

	data, q, a, err := db.Query(sql.QueryProvince, nil)
	if err != nil {
		return nil, fmt.Errorf("查询省信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	if len(data) == 0 {
		return nil, nil
	}
	return data, nil
}

// QueryCityByProvince 查询市信息
func (d *DbBranch) QueryCityByProvince(province string) (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("industry_marketing")

	data, q, a, err := db.Query(sql.QueryCityByProvince, map[string]interface{}{
		"province": province,
	})
	if err != nil {
		return nil, fmt.Errorf("查询市信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	if len(data) == 0 {
		return nil, nil
	}

	return data, nil
}

// QueryDistrictByCity 查询区信息
func (d *DbBranch) QueryDistrictByCity(city string) (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("industry_marketing")

	data, q, a, err := db.Query(sql.QueryDistrictByCity, map[string]interface{}{
		"city": city,
	})
	if err != nil {
		return nil, fmt.Errorf("查询区信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	if len(data) == 0 {
		return nil, nil
	}
	return data, nil
}

// GetAllBranchesByUser查询所有的门店信息
func (d *DbBranch) GetAllBranchesByUser(channelNos string) (data db.QueryRows, err error) {
	db := d.c.GetRegularDB("industry_marketing")
	data, q, a, err := db.Query(sql.GetAllBranchesByUser, map[string]interface{}{
		"up_channel_id": channelNos,
	})
	// "channel_id": user.RefAccID,
	// "agent_id":   user.RefAccID,
	// "agent":      agent,
	// "channel":    channel,

	if err != nil {
		return nil, fmt.Errorf("查询用户所有的门店信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, nil
}

func (d *DbBranch) GetBranchIDsByName(name string) (branchIDs string, err error) {
	db := d.c.GetRegularDB("industry_marketing")

	data, q, a, err := db.Scalar(sql.GetBranchIDsByName, map[string]interface{}{
		"branch_name": name,
	})
	if err != nil {
		return "", fmt.Errorf("查询所有的门店信息发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return types.GetString(data), nil
}
