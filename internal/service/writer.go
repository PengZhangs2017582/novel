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
	IWriter interface {
		CreateOrLogin(ctx context.Context, in model.WriterCreateOrLoginInput) (out *model.WriterCreateOrLoginOutput, err error)
		CheckPhoneVerifyCode(ctx context.Context, phone string, code, codeType int) (bool, error)
		CheckWriter(ctx context.Context, phone string) (bool, error)
	}
)

var (
	localWriter IWriter
)

func Writer() IWriter {
	if localWriter == nil {
		panic("implement not found for interface IWriter, forgot register?")
	}
	return localWriter
}

func RegisterWriter(i IWriter) {
	localWriter = i
}
