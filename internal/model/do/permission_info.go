// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure of table nv_permission_info for DAO operations like Where/Data.
type PermissionInfo struct {
	g.Meta    `orm:"table:nv_permission_info, do:true"`
	Id        interface{} //
	Name      interface{} // 权限名称
	Path      interface{} // 路径
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
