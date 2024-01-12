// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Writer is the golang structure of table nv_writer for DAO operations like Where/Data.
type Writer struct {
	g.Meta       `orm:"table:nv_writer, do:true"`
	Id           interface{} //
	PenName      interface{} // 笔名
	Phone        interface{} // 手机号
	Password     interface{} // 密码
	Introduction interface{} // 简介
	Avatar       interface{} // 头像
	Qq           interface{} //
	RealName     interface{} // 真名
	IdNo         interface{} // 身份证号
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
