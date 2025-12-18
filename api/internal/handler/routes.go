package handler

import (
	"net/http"

	"acupofcoffee/api/internal/middleware"
	"acupofcoffee/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, ctx *svc.ServiceContext) {
	corsMiddleware := middleware.NewCorsMiddleware()
	loggingMiddleware := middleware.NewLoggingMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(ctx.Config.Auth.AccessSecret)

	// 公开路由（无需认证）
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{corsMiddleware.Handle, loggingMiddleware.Handle},
			[]rest.Route{
				// 认证接口
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/auth/login",
					Handler: LoginHandler(ctx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/auth/register",
					Handler: RegisterHandler(ctx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/health",
					Handler: HealthHandler(ctx),
				},
				// 文章接口（开发阶段暂时公开）
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/articles",
					Handler: ListArticleHandler(ctx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/articles/:id",
					Handler: GetArticleHandler(ctx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/articles",
					Handler: CreateArticleHandler(ctx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/articles/:id",
					Handler: UpdateArticleHandler(ctx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/v1/articles/:id",
					Handler: DeleteArticleHandler(ctx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/articles/draft",
					Handler: SaveDraftHandler(ctx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/articles/:id/versions",
					Handler: GetVersionsHandler(ctx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/articles/:id/versions/:versionId/restore",
					Handler: RestoreVersionHandler(ctx),
				},
			}...,
		),
	)

	// 需要认证的路由
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{corsMiddleware.Handle, loggingMiddleware.Handle, authMiddleware.Handle},
			[]rest.Route{
				// 用户接口
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/user/info",
					Handler: GetUserInfoHandler(ctx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/user/info",
					Handler: UpdateUserInfoHandler(ctx),
				},
			}...,
		),
	)
}
