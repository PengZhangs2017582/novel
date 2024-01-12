// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"novel/internal/model"
)

type (
	IPhoneVerify interface {
		GenerateCode(ctx context.Context, phone string, codeType int) (out *model.GetCodeOutput, err error)
	}
)

var (
	localPhoneVerify IPhoneVerify
)

func PhoneVerify() IPhoneVerify {
	if localPhoneVerify == nil {
		panic("implement not found for interface IPhoneVerify, forgot register?")
	}
	return localPhoneVerify
}

func RegisterPhoneVerify(i IPhoneVerify) {
	localPhoneVerify = i
}
