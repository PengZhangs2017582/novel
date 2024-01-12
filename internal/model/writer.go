// Package model -----------------------------
// @file      : login.go
// @author    : Allen zhang
// @contact   : 364438619@qq.com
// @time      : 2024/1/10 14:40
// -------------------------------------------
package model

type WriterCreateOrLoginInput struct {
	Phone string
	Code  int
}

type WriterCreateOrLoginOutput struct {
	Id uint
}

// type GetCodeInput struct {
// 	Phone string
// }
//
// type GetCodeOutput struct {
// 	Code int
// }
