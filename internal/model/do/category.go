// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure of table nv_category for DAO operations like Where/Data.
type Category struct {
	g.Meta    `orm:"table:nv_category, do:true"`
	Id        interface{} //
	Name      interface{} // 分类名称
	ParentId  interface{} // 父级id
	CateType  interface{} // 分类类型：1-主题，2-角色，3-情节
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
