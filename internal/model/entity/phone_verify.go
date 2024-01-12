// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PhoneVerify is the golang structure for table phone_verify.
type PhoneVerify struct {
	Id        uint64      `json:"id"        description:""`
	Phone     string      `json:"phone"     description:"手机号"`
	Code      string      `json:"code"      description:"验证码"`
	Type      uint        `json:"type"      description:"类型：1-注册/登录，2-找回密码"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}
