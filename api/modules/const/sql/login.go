package sql

// QueryUserInfo 取得用户所有信息
const QueryUserInfo = `select 
						a.user_id,
						a.user_name,
						a.user_account,
						a.acc_type,
						a.ref_acc_id,
            a.password,
            to_char(a.last_login_time, 'yyyy-mm-dd hh24:mi:ss') last_login_time,
            a.last_login_ip
						from acc_user_info a where a.user_account = @user_account`

// UpdateUserLastTime 更新用户最后登录时间
const UpdateUserLastTime = `update acc_user_info a 
set 
a.last_login_time = sysdate,
a.last_login_ip = @last_login_ip
 where a.user_id =@user_id`

//QueryUserMenus 获取用户菜单信息
const QueryUserMenus = `select distinct s.menu_id id,
s.menu_name name,
s.parant_menu_id parent,
s.menu_level level_id,
s.icon,
s.path,
s.is_open,
to_char(s.create_time, 'yyyy-mm-dd hh24:mi:ss') create_time,
s.sort_id sortrank
from acc_user_role r
inner join acc_role_menu m on r.role_id = m.role_id and m.status = 0
inner join acc_menu_info s on s.menu_id = m.menu_id and s.status = 0
where r.user_id = @user_id
and r.status = 0
order by  s.menu_level asc, s.parant_menu_id asc, s.sort_id asc
`
const GetSysInfo = `
select t.sys_id,
       t.sys_logo,
       t.sys_name,
       t.default_pic,
       t.create_time,
       t.status,
       t.sys_copy,
       t.sys_themes
  from sys_config_info t
  where t.status=0
`

// QueryUserInfo 取得用户所有信息
const QueryUserInfoByID = `select 
						a.user_id,
						a.user_name,
						a.user_account,
						a.acc_type,
						a.ref_acc_id,
						a.password
						from acc_user_info a where a.user_id = @user_id`

const ChangePwd = `
update acc_user_info a set a.password = @password where a.user_id =@user_id
`
