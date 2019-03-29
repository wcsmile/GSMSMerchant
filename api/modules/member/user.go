package member

import (
	"fmt"
	"gsms/GSMSMerchant/api/modules/const/errorcode"
	"gsms/GSMSMerchant/api/modules/const/sql"
	"net/http"
	"strings"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/security/md5"
	"github.com/micro-plat/lib4go/types"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

type Menu struct {
	ID       string `json:"menu_id"`
	Name     string `json:"name_name"`
	Level    string `json:"menu_level"`
	IsOpen   string `json:"is_open"`
	Icon     string `json:"icon"`
	Parent   string `json:"parant_menu_id"`
	Path     string `json:"path"`
	Sortrank string `json:"sort_id"`
	Children []Menu `json:"children,omitempty"`
}

type IUser interface {
	Login(userAccount, password, clientIP string) (rows db.QueryRows, err error)
	GetMenu(userID string) (datas []map[string]interface{}, err error)

	// GetUserInfo(username, ident string) (*LoginUserInfo, error)
	ChangePwd(userid, passwordOld, password string) (err error)
	GetSysInfo() (map[string]interface{}, error)
}

type User struct {
	c    component.IContainer
	http *http.Client
}

func NewUser(c component.IContainer) *User {
	return &User{
		c:    c,
		http: &http.Client{},
	}
}

//Login 登录
func (u *User) Login(userAccount, password, clientIP string) (rows db.QueryRows, err error) {
	db := u.c.GetRegularDB("industry_marketing")

	// 获取用户信息
	data, q, p, err := db.Query(sql.QueryUserInfo, map[string]interface{}{
		"user_account": userAccount,
	})
	if err != nil {
		err = fmt.Errorf("获取用户信息错误:user_account:%+v,(err:%v) \nq:%v \na:%v", userAccount, err, q, p)

		return nil, err
	}

	if len(data) == 0 {
		err := context.NewError(errorcode.HTTPErrorUserAccountError, fmt.Errorf("用户账号不存在")) // 406 返回给前端的error
		return nil, err
	}

	if !strings.EqualFold(types.GetString(data[0]["password"]), md5.Encrypt(password)) {
		err := context.NewError(errorcode.HTTPErrorUserPasswordError, fmt.Errorf("用户密码错误")) // 407 返回给前端的error
		return nil, err
	}

	// 更新当前用户的最后登录时间
	count, q1, p1, err1 := db.Execute(sql.UpdateUserLastTime, map[string]interface{}{
		"user_id":       data[0]["user_id"],
		"last_login_ip": clientIP,
	})
	if err1 != nil || count <= 0 {
		return nil, fmt.Errorf("更新用户最后登录时间和IP发生错误(err:%v),sql:%s,输入参数:%v,", err, q1, p1)
	}
	return data, err
}

// GetMenu 获取用户系统的菜单信息
func (u *User) GetMenu(userID string) (datas []map[string]interface{}, err error) {
	db := u.c.GetRegularDB("industry_marketing")

	// 查询当前用户角色菜单
	data, _, _, err := db.Query(sql.QueryUserMenus, map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return nil, err
	}
	menuDatas := []map[string]interface{}{}
	result := menuConstruct(data, menuDatas, 0, 1)
	return result, nil
}

func menuConstruct(datas db.QueryRows, menuDatas []map[string]interface{}, ID, levelID int) (result []map[string]interface{}) {

	for _, row := range datas {
		// 一级菜单
		if row.GetInt("parent") == ID && row.GetInt("level_id") == levelID {

			// 三级菜单 没有children
			if row.GetInt("parent") == ID && row.GetInt("level_id") != 3 {
				// 传入一个空的children
				var children []map[string]interface{}
				row["children"] = menuConstruct(datas, children, row.GetInt("id"), levelID+1)
			}

			// 放入每一成菜单的数据
			menuDatas = append(menuDatas, row)

			// 递归返回的的2,3级children，或1级所有数据
			result = menuDatas

		}
	}
	return
}

//changePwd
func (u *User) ChangePwd(userid, passwordOld, password string) (err error) {
	db := u.c.GetRegularDB("industry_marketing")

	// 获取用户信息
	data, q, p, err := db.Query(sql.QueryUserInfoByID, map[string]interface{}{
		"user_id": userid,
	})
	if err != nil {
		err = fmt.Errorf("根据用户ID获取用户信息错误:user_id:%+v,(err:%v) \nq:%v \na:%v", userid, err, q, p)
		return err
	}
	if len(data) == 0 {
		err := context.NewError(errorcode.HTTPErrorUserAccountError, fmt.Errorf("用户账号不存在")) // 406 返回给前端的error
		return err
	}

	if !strings.EqualFold(types.GetString(data[0]["password"]), md5.Encrypt(passwordOld)) {
		err := context.NewError(errorcode.HTTPErrorOldPWDError, fmt.Errorf("用户密码错误")) // 705 返回给前端的error
		return err
	}

	//密码加密
	passwordEnctypt := md5.Encrypt(password)

	// 获取用户信息
	count, q, p, err := db.Execute(sql.ChangePwd, map[string]interface{}{
		"user_id":  userid,
		"password": passwordEnctypt,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("修改用户登录密码发生错误(err:%v),sql:%s,输入参数:%v,", err, q, p)
	}
	return nil
}

func (u *User) GetSysInfo() (map[string]interface{}, error) {
	db := u.c.GetRegularDB("industry_marketing")
	// 获取用户信息
	data, q, p, err := db.Query(sql.GetSysInfo, map[string]interface{}{})
	if err != nil {
		err = fmt.Errorf("获取系统信息错误:(err:%v) \nq:%v \na:%v", err, q, p)
		return nil, err
	}
	if data.IsEmpty() {
		err := context.NewError(errorcode.HTTPErrorSysInfoError, fmt.Errorf("系统信息不存在")) // 703 返回给前端的error
		return nil, err
	}
	return data.Get(0), err
}
