// Package model -----------------------------
// @file      : phone_verify.go
// @author    : Allen zhang
// @contact   : 364438619@qq.com
// @time      : 2024/1/11 15:30
// -------------------------------------------
package model

type GetCodeInput struct {
	Phone    string
	CodeType int
}

type GetCodeOutput struct {
	Code int
}
