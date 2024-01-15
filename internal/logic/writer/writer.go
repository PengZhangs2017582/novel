// Package backend -----------------------------
// @file      : writer.go
// @author    : Allen zhang
// @contact   : 364438619@qq.com
// @time      : 2024/1/10 14:49
// -------------------------------------------
package writer

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"

	"novel/internal/dao"
	"novel/internal/model"
	"novel/internal/model/do"
	"novel/internal/model/entity"
	"novel/internal/service"
)

// type sWriter struct{}
type (
	sWriter struct{}
)

func init() {
	service.RegisterWriter(New())
}
func New() *sWriter {
	return &sWriter{}
}

func (s *sWriter) CreateOrLogin(ctx context.Context, in model.WriterCreateOrLoginInput) (out *model.WriterCreateOrLoginOutput, err error) {
	// writerInfo := entity.Writer{}
	var writerInfo entity.Writer
	// writerInfo := new(entity.Writer)

	// check code
	codeCheck, codeErr := s.CheckPhoneVerifyCode(ctx, in.Phone, in.Code, 1)
	if codeErr != nil {
		return out, codeErr
	}
	if !codeCheck {
		return out, gerror.New(`验证码错误`)
	}

	// check account
	writerCheck, writerErr := s.CheckWriter(ctx, in.Phone)
	if writerErr != nil {
		return out, codeErr
	}
	if !writerCheck {
		// create new account
		_, insertErr := dao.Writer.Ctx(ctx).Data(do.Writer{
			Phone: in.Phone,
		}).Insert()
		if insertErr != nil {
			return out, gerror.New(`内部错误`)
		}
	}

	infoErr := dao.Writer.Ctx(ctx).Where(do.Writer{
		Phone: in.Phone,
	}).Scan(&writerInfo)
	if infoErr != nil {
		return out, infoErr
		// return out, gerror.New(`查找用户错误`)
	}

	// if writerInfo == nil {
	// 	return gerror.New(`Passport or Password not correct`)
	// }
	// fmt.Println("-------------------------")
	// fmt.Println(writerInfo)

	// err1 := service.Session().SetWriter(ctx, &writerInfo)
	// if err1 != nil {
	// 	return out, err1
	// }

	// service.BizCtx().SetUser(ctx, &model.ContextUser{
	// 	Id: writerInfo.Id,
	// })

	return &model.WriterCreateOrLoginOutput{Id: writerInfo.Id, PenName: writerInfo.PenName}, nil
}

func (s *sWriter) CheckPhoneVerifyCode(ctx context.Context, phone string, code, codeType int) (bool, error) {

	time1 := time.Now().Add(-10 * time.Minute)
	dateTime := time1.Format("2006-01-02 15:04:05")

	count, err := dao.PhoneVerify.Ctx(ctx).Where(do.PhoneVerify{
		Phone: phone,
		Code:  code,
		Type:  codeType,
	}).Where("created_at >=", dateTime).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (s *sWriter) CheckWriter(ctx context.Context, phone string) (bool, error) {
	count, err := dao.Writer.Ctx(ctx).Where(do.Writer{
		Phone: phone,
	}).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
