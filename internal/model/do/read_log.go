// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReadLog is the golang structure of table nv_read_log for DAO operations like Where/Data.
type ReadLog struct {
	g.Meta       `orm:"table:nv_read_log, do:true"`
	Id           interface{} //
	ShortStoryId interface{} //
	OperateType  interface{} // 操作类型：1-阅读，2-加入书架
	OperateTime  *gtime.Time // 操作时间
	ReadDuration interface{} // 阅读时长，秒为单位
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
