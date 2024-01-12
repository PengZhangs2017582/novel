// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PhoneVerify is the golang structure of table nv_phone_verify for DAO operations like Where/Data.
type PhoneVerify struct {
	g.Meta    `orm:"table:nv_phone_verify, do:true"`
	Id        interface{} //
	Phone     interface{} // 手机号
	Code      interface{} // 验证码
	Type      interface{} // 类型：1-注册/登录，2-找回密码
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
