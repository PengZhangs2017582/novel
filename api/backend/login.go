package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// verify code
type GetCodeReq struct {
	g.Meta   `path:"/get-code" method:"post" tags:"获取验证码" summary:"获取验证码"`
	Phone    string `json:"phone" v:"required|phone" description:"手机号"`
	CodeType int    `json:"codeType" v:"required|in:1,2#类型不合法" description:"验证码类型"`
}
type GetCodeRes struct {
	Code int
}

// createOrLogin
type CreateOrLoginReq struct {
	g.Meta `path:"/create-or-login" method:"post" tags:"注册或登录" summary:"注册或者登录"`
	Phone  string `json:"phone" v:"required|phone" description:"手机号"`
	Code   int    `json:"code" v:"required|size:6" description:"验证码"`
}

type CreateOrLoginRes struct {
	Id uint
}
