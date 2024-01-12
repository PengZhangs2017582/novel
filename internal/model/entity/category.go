// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure for table category.
type Category struct {
	Id        uint        `json:"id"        description:""`
	Name      string      `json:"name"      description:"分类名称"`
	ParentId  uint        `json:"parentId"  description:"父级id"`
	CateType  uint        `json:"cateType"  description:"分类类型：1-主题，2-角色，3-情节"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
