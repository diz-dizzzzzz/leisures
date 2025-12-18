package handler

import (
	"net/http"

	"acupofcoffee/api/internal/logic"
	"acupofcoffee/api/internal/svc"
	"acupofcoffee/api/internal/types"
	"acupofcoffee/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewAuthLogic(r.Context(), ctx)
		resp, err := l.Login(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewAuthLogic(r.Context(), ctx)
		err := l.Register(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, nil)
	}
}
