// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReadLog is the golang structure for table read_log.
type ReadLog struct {
	Id           uint64      `json:"id"           description:""`
	ShortStoryId int         `json:"shortStoryId" description:""`
	OperateType  int         `json:"operateType"  description:"操作类型：1-阅读，2-加入书架"`
	OperateTime  *gtime.Time `json:"operateTime"  description:"操作时间"`
	ReadDuration uint        `json:"readDuration" description:"阅读时长，秒为单位"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    description:""`
}
