// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermissionInfo is the golang structure of table nv_role_permission_info for DAO operations like Where/Data.
type RolePermissionInfo struct {
	g.Meta       `orm:"table:nv_role_permission_info, do:true"`
	Id           interface{} //
	RoleId       interface{} // 角色id
	PermissionId interface{} // 权限id
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
