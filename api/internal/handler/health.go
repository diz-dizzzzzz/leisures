package handler

import (
	"net/http"

	"acupofcoffee/api/internal/svc"
	"acupofcoffee/common/response"
)

func HealthHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, map[string]string{
			"status":  "ok",
			"service": ctx.Config.Name,
		})
	}
}
