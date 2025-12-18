package logic

import (
	"context"
	"time"

	"acupofcoffee/api/internal/svc"
	"acupofcoffee/api/internal/types"
	"acupofcoffee/common/errorx"
	"acupofcoffee/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Create 创建文章
func (l *ArticleLogic) Create(req *types.CreateArticleRequest) (*types.ArticleResponse, error) {
	userID, ok := l.ctx.Value("userId").(uint)
	if !ok {
		userID = 1 // 开发模式：默认用户ID
	}

	article := model.Article{
		Title:      req.Title,
		Content:    req.Content,
		ContentRaw: req.Content, // 简化：直接使用 content 作为搜索内容
		Cover:      req.Cover,
		Summary:    req.Summary,
		AuthorID:   userID,
		Status:     req.Status,
		Version:    1,
	}

	if err := l.svcCtx.DB.Create(&article).Error; err != nil {
		l.Logger.Errorf("create article error: %v", err)
		return nil, errorx.NewDefaultError("创建文章失败")
	}

	// 删除草稿（如果存在）
	l.svcCtx.DB.Where("article_id = 0 AND user_id = ?", userID).Delete(&model.ArticleDraft{})

	return l.articleToResponse(&article), nil
}

// Update 更新文章（带版本控制）
func (l *ArticleLogic) Update(req *types.UpdateArticleRequest) (*types.ArticleResponse, error) {
	userID, _ := l.ctx.Value("userId").(uint)
	// 开发模式：不检查权限

	var article model.Article
	if err := l.svcCtx.DB.First(&article, req.ID).Error; err != nil {
		return nil, errorx.NewNotFoundError("文章不存在")
	}

	_ = userID // 暂时忽略用户ID检查

	// 使用事务保存版本历史和更新文章
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 保存当前版本到历史
		version := model.ArticleVersion{
			ArticleID: article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Version:   article.Version,
			Remark:    req.Remark,
		}
		if err := tx.Create(&version).Error; err != nil {
			return err
		}

		// 2. 更新文章
		updates := map[string]interface{}{
			"version": article.Version + 1,
		}
		if req.Title != "" {
			updates["title"] = req.Title
		}
		if req.Content != "" {
			updates["content"] = req.Content
			updates["content_raw"] = req.Content
		}
		if req.Cover != "" {
			updates["cover"] = req.Cover
		}
		if req.Summary != "" {
			updates["summary"] = req.Summary
		}
		if req.Status != 0 {
			updates["status"] = req.Status
		}

		return tx.Model(&article).Updates(updates).Error
	})

	if err != nil {
		l.Logger.Errorf("update article error: %v", err)
		return nil, errorx.NewDefaultError("更新文章失败")
	}

	// 重新查询更新后的文章
	l.svcCtx.DB.First(&article, req.ID)
	return l.articleToResponse(&article), nil
}

// Get 获取文章详情
func (l *ArticleLogic) Get(id uint) (*types.ArticleResponse, error) {
	var article model.Article
	if err := l.svcCtx.DB.Preload("Author").First(&article, id).Error; err != nil {
		return nil, errorx.NewNotFoundError("文章不存在")
	}

	// 增加浏览量
	l.svcCtx.DB.Model(&article).UpdateColumn("view_count", gorm.Expr("view_count + 1"))

	return l.articleToResponse(&article), nil
}

// List 文章列表
func (l *ArticleLogic) List(req *types.ArticleListRequest) (*types.ArticleListResponse, error) {
	var articles []model.Article
	var total int64

	query := l.svcCtx.DB.Model(&model.Article{})

	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.AuthorID > 0 {
		query = query.Where("author_id = ?", req.AuthorID)
	}
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		query = query.Where("title LIKE ? OR content_raw LIKE ?", keyword, keyword)
	}

	query.Count(&total)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	offset := (req.Page - 1) * req.PageSize
	if err := query.Preload("Author").
		Order("created_at DESC").
		Offset(offset).
		Limit(req.PageSize).
		Find(&articles).Error; err != nil {
		return nil, errorx.NewDefaultError("获取文章列表失败")
	}

	list := make([]*types.ArticleResponse, len(articles))
	for i, article := range articles {
		list[i] = l.articleToResponse(&article)
	}

	return &types.ArticleListResponse{
		PageResponse: types.PageResponse{
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
			List:     list,
		},
	}, nil
}

// SaveDraft 保存草稿（实时自动保存）
func (l *ArticleLogic) SaveDraft(req *types.SaveDraftRequest) (*types.SaveDraftResponse, error) {
	userID, ok := l.ctx.Value("userId").(uint)
	if !ok {
		userID = 1 // 开发模式：默认用户ID
	}

	draft := model.ArticleDraft{
		ArticleID: req.ArticleID,
		UserID:    userID,
		Title:     req.Title,
		Content:   req.Content,
	}

	// 使用 Upsert（存在则更新，不存在则创建）
	err := l.svcCtx.DB.Where("article_id = ? AND user_id = ?", req.ArticleID, userID).
		Assign(model.ArticleDraft{
			Title:   req.Title,
			Content: req.Content,
		}).
		FirstOrCreate(&draft).Error

	if err != nil {
		l.Logger.Errorf("save draft error: %v", err)
		return nil, errorx.NewDefaultError("保存草稿失败")
	}

	return &types.SaveDraftResponse{
		DraftID: draft.ID,
		SavedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

// GetVersions 获取版本历史
func (l *ArticleLogic) GetVersions(articleID uint) ([]*types.ArticleVersionResponse, error) {
	var versions []model.ArticleVersion
	if err := l.svcCtx.DB.Where("article_id = ?", articleID).
		Order("version DESC").
		Find(&versions).Error; err != nil {
		return nil, errorx.NewDefaultError("获取版本历史失败")
	}

	list := make([]*types.ArticleVersionResponse, len(versions))
	for i, v := range versions {
		list[i] = &types.ArticleVersionResponse{
			ID:        v.ID,
			ArticleID: v.ArticleID,
			Title:     v.Title,
			Content:   v.Content,
			Version:   v.Version,
			Remark:    v.Remark,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return list, nil
}

// RestoreVersion 恢复到指定版本
func (l *ArticleLogic) RestoreVersion(articleID, versionID uint) (*types.ArticleResponse, error) {
	// 开发模式：不检查权限

	var article model.Article
	if err := l.svcCtx.DB.First(&article, articleID).Error; err != nil {
		return nil, errorx.NewNotFoundError("文章不存在")
	}

	var version model.ArticleVersion
	if err := l.svcCtx.DB.First(&version, versionID).Error; err != nil {
		return nil, errorx.NewNotFoundError("版本不存在")
	}

	// 更新文章为历史版本的内容
	req := &types.UpdateArticleRequest{
		ID:      articleID,
		Title:   version.Title,
		Content: version.Content,
		Remark:  "恢复到版本 " + string(rune(version.Version)),
	}

	return l.Update(req)
}

// Delete 删除文章
func (l *ArticleLogic) Delete(id uint) error {
	// 开发模式：不检查权限

	var article model.Article
	if err := l.svcCtx.DB.First(&article, id).Error; err != nil {
		return errorx.NewNotFoundError("文章不存在")
	}

	// 软删除
	if err := l.svcCtx.DB.Delete(&article).Error; err != nil {
		l.Logger.Errorf("delete article error: %v", err)
		return errorx.NewDefaultError("删除文章失败")
	}

	return nil
}

func (l *ArticleLogic) articleToResponse(article *model.Article) *types.ArticleResponse {
	resp := &types.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Cover:     article.Cover,
		Summary:   article.Summary,
		AuthorID:  article.AuthorID,
		Status:    article.Status,
		Version:   article.Version,
		ViewCount: article.ViewCount,
		LikeCount: article.LikeCount,
		CreatedAt: article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if article.Author != nil {
		resp.AuthorName = article.Author.Nickname
		if resp.AuthorName == "" {
			resp.AuthorName = article.Author.Username
		}
	}

	return resp
}
