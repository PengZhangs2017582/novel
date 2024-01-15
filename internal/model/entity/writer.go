// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Writer is the golang structure for table writer.
type Writer struct {
	Id           uint        `json:"id"           description:""`
	PenName      string      `json:"penName"      description:"笔名"`
	Phone        string      `json:"phone"        description:"手机号"`
	Password     string      `json:"password"     description:"密码"`
	Introduction string      `json:"introduction" description:"简介"`
	Avatar       string      `json:"avatar"       description:"头像"`
	Qq           string      `json:"qq"           description:""`
	RealName     string      `json:"realName"     description:"真名"`
	IdNo         string      `json:"idNo"         description:"身份证号"`
	IsAdmin      uint        `json:"isAdmin"      description:"是否是管理员"`
	RoleIds      string      `json:"roleIds"      description:"角色id"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    description:""`
}
