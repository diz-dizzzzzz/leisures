package handler

import (
	"net/http"
	"strconv"

	"acupofcoffee/api/internal/logic"
	"acupofcoffee/api/internal/svc"
	"acupofcoffee/api/internal/types"
	"acupofcoffee/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateArticleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.Create(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func UpdateArticleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.Update(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func GetArticleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.Get(req.ID)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func ListArticleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := types.ArticleListRequest{
			Page:     1,
			PageSize: 10,
		}

		// 手动解析 query 参数
		query := r.URL.Query()
		if v := query.Get("page"); v != "" {
			if page, err := strconv.Atoi(v); err == nil {
				req.Page = page
			}
		}
		if v := query.Get("pageSize"); v != "" {
			if pageSize, err := strconv.Atoi(v); err == nil {
				req.PageSize = pageSize
			}
		}
		if v := query.Get("status"); v != "" {
			if status, err := strconv.Atoi(v); err == nil {
				s := int8(status)
				req.Status = &s
			}
		}
		if v := query.Get("keyword"); v != "" {
			req.Keyword = v
		}
		if v := query.Get("authorId"); v != "" {
			if authorId, err := strconv.Atoi(v); err == nil {
				req.AuthorID = uint(authorId)
			}
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.List(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func DeleteArticleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		if err := l.Delete(req.ID); err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, nil)
	}
}

func SaveDraftHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveDraftRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.SaveDraft(&req)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

func GetVersionsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.GetVersions(req.ID)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}

type RestoreVersionRequest struct {
	ID        uint `path:"id"`
	VersionID uint `path:"versionId"`
}

func RestoreVersionHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RestoreVersionRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(w, err)
			return
		}

		l := logic.NewArticleLogic(r.Context(), ctx)
		resp, err := l.RestoreVersion(req.ID, req.VersionID)
		if err != nil {
			response.Error(w, err)
			return
		}

		response.Success(w, resp)
	}
}
