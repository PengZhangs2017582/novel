// // Package session -----------------------------
// // @file      : session.go
// // @author    : Allen zhang
// // @contact   : 364438619@qq.com
// // @time      : 2024/1/10 15:51
// // -------------------------------------------
package session

import (
	"context"
	"fmt"

	"novel/internal/model/entity"
	"novel/internal/service"
)

// type (
// 	sSession struct{}
// )

type sSession struct{}

const (
	sessionKeyWriter = "SessionKeyWriter" // 用户信息存放在Session中的Key
)

func init() {
	service.RegisterSession(New())
}
func New() *sSession {
	return &sSession{}
}

func (s *sSession) SetWriter(ctx context.Context, writer *entity.Writer) error {
	// fmt.Println(writer)
	// fmt.Println(ctx)
	// fmt.Println(sessionKeyWriter)
	fmt.Printf("writer: %p \n", &writer)
	fmt.Printf("ctx: %p \n", &ctx)
	// fmt.Printf("sessionKeyWriter: %p \n", &sessionKeyWriter)
	fmt.Println(*writer, &writer, writer)
	return service.BizCtx().Get(ctx).Session.Set(sessionKeyWriter, writer)
}
