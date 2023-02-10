package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-scaffold/app/internal/logic/captcha"
	"zero-scaffold/app/internal/svc"
	"zero-scaffold/app/internal/types"
	"zero-scaffold/common/errorx"
	"zero-scaffold/common/response"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDPathReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errorx.NewCodeError(errorx.ERR_PARAMS, err.Error()))
			return
		}

		l := captcha.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha(&req)
		response.Response(w, resp, err)
	}
}
