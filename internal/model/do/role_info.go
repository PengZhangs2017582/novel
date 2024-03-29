// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleInfo is the golang structure of table nv_role_info for DAO operations like Where/Data.
type RoleInfo struct {
	g.Meta    `orm:"table:nv_role_info, do:true"`
	Id        interface{} //
	Name      interface{} // 角色名称
	Desc      interface{} // 描述
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
