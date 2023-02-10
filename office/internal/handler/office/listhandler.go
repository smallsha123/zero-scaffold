package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-scaffold/common/errorx"
	"zero-scaffold/common/response"
	"zero-scaffold/office/internal/logic/office"
	"zero-scaffold/office/internal/svc"
	"zero-scaffold/office/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OfficeSearch
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errorx.NewCodeError(errorx.ERR_PARAMS, err.Error()))
			return
		}

		l := office.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		response.Response(w, resp, err)

	}
}
