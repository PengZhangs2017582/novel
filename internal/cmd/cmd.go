package cmd

import (
	"context"
	"strconv"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	apiBackend "novel/api/backend"
	"novel/internal/consts"
	"novel/internal/controller/backend"
	"novel/internal/dao"
	"novel/internal/model"
	"novel/internal/model/entity"
	"novel/internal/service"
	"novel/utility/response"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 认证接口
			// loginFunc := Login
			// 启动
			gfToken := &gtoken.GfToken{
				CacheMode:        consts.CacheModeRedis,
				ServerName:       "novel",
				LoginPath:        "/create-or-login",
				LoginBeforeFunc:  loginFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/backend/logout",
				AuthPaths:        g.SliceStr{"/backend/admin/info"},
				AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
				AuthAfterFunc:    authAfterFunc,
				MultiLogin:       true,
			}

			// s.Group("/backend", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(ghttp.MiddlewareHandlerResponse)
			// 	group.Group("/writer", func(group2 *ghttp.RouterGroup) {
			// 		group2.Bind(backend.New())
			// 	})
			// })

			// 管理后台路由组
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// group.Middleware(
				// 	service.Middleware().CORS,
				// 	service.Middleware().Ctx,
				// 	service.Middleware().ResponseHandler,
				// )

				// // gtoken中间件绑定
				// err := gfToken.Middleware(ctx, group)
				// if nil != nil {
				// 	panic(err)
				// }

				// 不需要登录的路由组绑定
				group.Bind(
					// backend.Login,               // 登录相关
					backend.Login.GetPhoneCode,  // 验证码
					backend.Login.CreateOrLogin, // 登录
				)
				// 需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						backend.Writer.Info, // 查询当前管理员信息
					)
				})
			})

			s.Run()
			return nil
		},
	}
)

func loginFunc(r *ghttp.Request) (string, interface{}) {
	phone := r.Get("phone").String()
	code := r.Get("code").Int()
	if phone == "" || code <= 0 {
		r.Response.WriteJson(gtoken.Fail("手机号或者验证码不能为空"))
		r.ExitAll()
	}

	ctx := context.TODO()
	writer, err := service.Writer().CreateOrLogin(ctx, model.WriterCreateOrLoginInput{
		Phone: phone,
		Code:  code,
	})

	if err != nil {
		r.Response.WriteJson(gtoken.Fail("手机号或者验证码错误"))
		r.ExitAll()
	}
	return consts.GtokenWriterPrefix + strconv.Itoa(int(writer.Id)), writer
}

// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		// 获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtokenWriterPrefix)
		// 根据id获得登录用户其他信息
		writerInfo := entity.Writer{}
		err := dao.Writer.Ctx(context.TODO()).WherePri(adminId).Scan(&writerInfo)
		if err != nil {
			return
		}
		// 通过角色查询权限
		// 先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{writerInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &apiBackend.CreateOrLoginRes{
			Type:        consts.TokenType,
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GTokenExpireIn, // 单位秒,
			IsAdmin:     writerInfo.IsAdmin,
			RoleIds:     writerInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 登录鉴权中间件for后台
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var writerInfo entity.Writer
	err := gconv.Struct(respData.GetString("data"), &writerInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	// todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxWriterId, writerInfo.Id)
	r.SetCtxVar(consts.CtxWriterName, writerInfo.PenName)
	r.SetCtxVar(consts.CtxWriterIsAdmin, writerInfo.IsAdmin)
	r.SetCtxVar(consts.CtxWriterRoleIds, writerInfo.RoleIds)
	r.Middleware.Next()
}
