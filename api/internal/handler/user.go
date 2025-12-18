package handler

import (
	"net/http"

	"acupofcoffee/api/internal/logic"
	"acupofcoffee/api/internal/svc"
	"acupofcoffee/api/internal/types"
	"acupofcoffee/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserLogic(r.Context(), ctx)
		resp, err := l.GetUserInfo()
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func UpdateUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), ctx)
		err := l.UpdateUserInfo(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, nil)
	}
}
