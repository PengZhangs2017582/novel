// Package backend -----------------------------
// @file      : login.go
// @author    : Allen zhang
// @contact   : 364438619@qq.com
// @time      : 2024/1/10 14:07
// -------------------------------------------
package backend

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"novel/api/backend"
	"novel/internal/model"
	"novel/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func New() *cLogin {
	return &cLogin{}
}

func (c *cLogin) CreateOrLogin(ctx context.Context, req *backend.CreateOrLoginReq) (res *backend.CreateLoginRes, err error) {
	input := model.WriterCreateOrLoginInput{}
	err = gconv.Scan(req, &input)
	if err != nil {
		return nil, err
	}

	writer, err := service.Writer().CreateOrLogin(ctx, model.WriterCreateOrLoginInput{
		Phone: req.Phone,
		Code:  req.Code,
	})

	if err != nil {
		return nil, err
	}
	return &backend.CreateLoginRes{Id: writer.Id, PenName: writer.PenName}, nil
	// return &backend.CreateOrLoginRes{Id: writer.Id}, nil
}

func (c *cLogin) GetPhoneCode(ctx context.Context, req *backend.GetCodeReq) (res *backend.GetCodeRes, err error) {
	input := model.GetCodeInput{}
	err = gconv.Scan(req, &input)
	if err != nil {
		return nil, err
	}
	code, err := service.PhoneVerify().GenerateCode(ctx, req.Phone, req.CodeType)
	if err != nil {
		return nil, err
	}
	return &backend.GetCodeRes{Code: code.Code}, nil
}
