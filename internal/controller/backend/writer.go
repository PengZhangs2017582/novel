// Package backend -----------------------------
// @file      : writer.go
// @author    : Allen zhang
// @contact   : 364438619@qq.com
// @time      : 2024/1/15 15:15
// -------------------------------------------
package backend

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"novel/api/backend"
	"novel/internal/consts"
)

// Writer 内容管理
var Writer = cWriter{}

type cWriter struct{}

// gtoken 版本返回结果
func (c *cWriter) Info(ctx context.Context, req *backend.WriterGetInfoReq) (res *backend.WriterGetInfoRes, err error) {
	return &backend.WriterGetInfoRes{
		Id:      gconv.Int(ctx.Value(consts.CtxWriterId)),
		PenName: gconv.String(ctx.Value(consts.CtxWriterName)),
		IsAdmin: gconv.Int(ctx.Value(consts.CtxWriterIsAdmin)),
		RoleIds: gconv.String(ctx.Value(consts.CtxWriterRoleIds)),
	}, err
}
