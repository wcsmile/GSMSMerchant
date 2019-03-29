package order

//QueryOrderInput 查询需要的参数
type QueryOrderInput struct {
	OrderNo       string `form:"order_no" json:"order_no"`
	Pi            string `form:"pi" json:"pi" `
	Ps            string `form:"ps" json:"ps" `
	PaymentStatus string `form:"payment_status" json:"payment_status" `
	AccType       string `form:"acc_type" json:"acc_type"`
	RefAccID      string `form:"ref_acc_id" json:"ref_acc_id"`
	ChannelNo     string `form:"channel_no" json:"channel_no"`
}
