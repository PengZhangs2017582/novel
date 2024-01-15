package backend

import (
	"github.com/gogf/gf/v2/frame/g"

	"novel/internal/model/entity"
)

// verify code
type GetCodeReq struct {
	g.Meta   `path:"/get-phone-code" method:"post" tags:"获取验证码" summary:"获取验证码"`
	Phone    string `json:"phone" v:"required|phone" description:"手机号"`
	CodeType int    `json:"codeType" v:"required|in:1,2#类型不合法" description:"验证码类型"`
}

type GetCodeRes struct {
	Code int
}

// createOrLogin
type CreateOrLoginReq struct {
	// g.Meta `path:"/create-or-login" method:"post" tags:"注册或登录" summary:"注册或者登录"`
	Phone string `json:"phone" v:"required|phone" description:"手机号"`
	Code  int    `json:"code" v:"required|size:6" description:"验证码"`
}

type CreateLoginRes struct {
	Id      uint
	PenName string
}

// for gtoken
type CreateOrLoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     uint                    `json:"is_admin"`    // 是否超管
	RoleIds     string                  `json:"role_ids"`    // 角色
	Permissions []entity.PermissionInfo `json:"permissions"` // 权限列表
}
