package backcard

// 银行卡认证类型  0-待认证； 1-认证成功；2-认证失败； 3-认证中
const (
	ToBeCertified    = 0 //待认证
	CertifiedSuccess = 1 //认证成功
	CertifiedFailed  = 2 //认证失败
	OnCertifying     = 3 //认证中
)
