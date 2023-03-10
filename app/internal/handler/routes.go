// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	captcha "zero-scaffold/app/internal/handler/captcha"
	"zero-scaffold/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha/:id",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)
}
