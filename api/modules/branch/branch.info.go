package branch

//QueryBranchInput 查询需要的参数
type QueryBranchInput struct {
	Status   string `form:"status" json:"status"`
	Pi       string `form:"pi" json:"pi" `
	Ps       string `form:"ps" json:"ps" `
	BranchID string `form:"branch_id" json:"branch_id" `
	AccType  string `form:"acc_type" json:"acc_type"`
	RefAccID string `form:"ref_acc_id" json:"ref_acc_id"`
}

//EditStatusInput 编辑门店状态需要的参数
type EditStatusInput struct {
	Status   string `form:"status" json:"status"`
	BranchID string `form:"branch_id" json:"branch_id" valid:"required"`
	AccType  string `form:"acc_type" json:"acc_type"`
	RefAccID string `form:"ref_acc_id" json:"ref_acc_id"`
}

//AddBranchInput 添加门店信息
type AddBranchInput struct {
	BranchName  string `form:"branch_name" json:"branch_name" valid:"required"`
	Contactname string `form:"contact_name" json:"contact_name" valid:"required"`
	ContactTel  string `form:"contact_tel" json:"contact_tel" valid:"required"`
	Province    string `form:"province" json:"province" valid:"required"`
	City        string `form:"city" json:"city" valid:"required"`
	District    string `form:"district" json:"district" valid:"required"`
	Address     string `form:"address" json:"address" valid:"required"`
	AccType     string `form:"acc_type" json:"acc_type"`
	RefAccID    string `form:"ref_acc_id" json:"ref_acc_id"`
}

//UpdateBranchInput 更新门店信息
type UpdateBranchInput struct {
	BranchID    string `form:"branch_id" json:"branch_id" valid:"required"`
	BranchName  string `form:"branch_name" json:"branch_name" valid:"required"`
	Contactname string `form:"contact_name" json:"contact_name" valid:"required"`
	ContactTel  string `form:"contact_tel" json:"contact_tel" valid:"required"`
	Province    string `form:"province" json:"province" valid:"required"`
	City        string `form:"city" json:"city" valid:"required"`
	District    string `form:"district" json:"district" valid:"required"`
	Address     string `form:"address" json:"address" valid:"required"`
	AccType     string `form:"acc_type" json:"acc_type"`
	RefAccID    string `form:"ref_acc_id" json:"ref_acc_id"`
}

//DetailInput 详细需要的参数
type DetailInput struct {
	BranchID string `form:"branch_id" json:"branch_id" valid:"required"`
	AccType  string `form:"acc_type" json:"acc_type"`
	RefAccID string `form:"ref_acc_id" json:"ref_acc_id"`
}
