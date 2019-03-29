package branch

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
)

// IBranch 方法
type IBranch interface {
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

// Branch 对象
type Branch struct {
	c  component.IContainer
	db IDbBranch
}

// NewBranch 实例化
func NewBranch(c component.IContainer) *Branch {
	return &Branch{
		c:  c,
		db: NewDbBranch(c),
	}
}

// QueryAll 查询门店信息
func (p *Branch) QueryAll(inputData *QueryBranchInput) (data db.QueryRows, count int, err error) {
	return p.db.QueryAll(inputData)
}

// EditStatus 编辑门店状态
func (p *Branch) EditStatus(inputData *EditStatusInput) (err error) {
	return p.db.EditStatus(inputData)
}

// AddBranchInfo 添加门店信息
func (p *Branch) AddBranchInfo(inputData *AddBranchInput) (err error) {
	return p.db.AddBranchInfo(inputData)
}

// QueryProvince 查询省信息
func (p *Branch) QueryProvince() (data db.QueryRows, err error) {
	return p.db.QueryProvince()
}

// QueryCityByProvince 查询市信息
func (p *Branch) QueryCityByProvince(province string) (data db.QueryRows, err error) {
	return p.db.QueryCityByProvince(province)
}

// QueryDistrictByCity 查询区信息
func (p *Branch) QueryDistrictByCity(city string) (data db.QueryRows, err error) {
	return p.db.QueryDistrictByCity(city)
}

// QueryDetailInfo 查询门店详细信息
func (p *Branch) QueryDetailInfo(inputData *DetailInput) (data db.QueryRows, err error) {
	return p.db.QueryDetailInfo(inputData)
}

// UpdateBranchInfo 更新门店信息
func (p *Branch) UpdateBranchInfo(inputData *UpdateBranchInput) (err error) {
	return p.db.UpdateBranchInfo(inputData)
}

func (p *Branch) GetAllBranchesByUser(channelNos string) (data db.QueryRows, err error) {
	return p.db.GetAllBranchesByUser(channelNos)
}

func (p *Branch) GetBranchIDsByName(name string) (branchIDs string, err error) {
	return p.db.GetBranchIDsByName(name)
}

func (p *Branch) GetAllBranchesByChannelID(AccType, RefAccID string) (data db.QueryRows, err error) {
	return p.db.GetAllBranchesByChannelID(AccType, RefAccID)
}
