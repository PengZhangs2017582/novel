// Package phone_verify -----------------------------
// @file      : phone_verify.go
// @author    : Allen zhang
// @contact   : 364438619@qq.com
// @time      : 2024/1/11 15:17
// -------------------------------------------
package phone_verify

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/grand"

	"novel/internal/dao"
	"novel/internal/model"
	"novel/internal/model/do"
	"novel/internal/service"
)

type sPhoneVerify struct{}

func init() {
	service.RegisterPhoneVerify(New())
}
func New() *sPhoneVerify {
	return &sPhoneVerify{}
}

func (s *sPhoneVerify) GenerateCode(ctx context.Context, phone string, codeType int) (out *model.GetCodeOutput, err error) {
	// once within a minute
	time1 := time.Now().Add(-1 * time.Minute)
	dateTime := time1.Format("2006-01-02 15:04:05")
	// record := entity.PhoneVerify{}
	count, recordErr := dao.PhoneVerify.Ctx(ctx).Where(do.PhoneVerify{
		Phone: phone,
		Type:  codeType,
	}).Where("created_at >=", dateTime).Count()
	// fmt.Println(dateTime, count)
	if recordErr != nil {
		return nil, recordErr
	}
	if count > 0 {
		return nil, gerror.New(`一分钟内只能发送一次验证码`)
	}

	// generate new code
	code := grand.N(100000, 999999)
	_, insertErr := dao.PhoneVerify.Ctx(ctx).Data(do.PhoneVerify{
		Phone: phone,
		Code:  code,
		Type:  codeType,
	}).Insert()
	if insertErr != nil {
		return nil, gerror.New(`内部错误`)
	}
	return &model.GetCodeOutput{Code: code}, nil
}
