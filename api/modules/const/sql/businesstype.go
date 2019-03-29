package sql

const QueryAllSysBusinessType = `
select
	business_type,
	business_name,
	product_line,
	need_user_active,
	remark,
	status,
	need_recharge
from
	sys_business_type
where
	status=0
`
