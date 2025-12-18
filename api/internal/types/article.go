package types

// ============== 文章相关 ==============

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"` // JSON 字符串格式的富文本内容
	Cover   string `json:"cover,optional"`
	Summary string `json:"summary,optional"`
	Status  int8   `json:"status,optional"` // 0:草稿 1:发布
}

type UpdateArticleRequest struct {
	ID      uint   `json:"id" path:"id"`
	Title   string `json:"title,optional"`
	Content string `json:"content,optional"` // JSON 字符串格式
	Cover   string `json:"cover,optional"`
	Summary string `json:"summary,optional"`
	Status  int8   `json:"status,optional"`
	Remark  string `json:"remark,optional"` // 版本备注
}

type ArticleResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"` // JSON 字符串
	Cover      string `json:"cover"`
	Summary    string `json:"summary"`
	AuthorID   uint   `json:"authorId"`
	AuthorName string `json:"authorName"`
	Status     int8   `json:"status"`
	Version    int    `json:"version"`
	ViewCount  int64  `json:"viewCount"`
	LikeCount  int64  `json:"likeCount"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type ArticleListRequest struct {
	Page     int    `json:"page" form:"page,optional"`
	PageSize int    `json:"pageSize" form:"pageSize,optional"`
	Status   *int8  `json:"status" form:"status,optional"`
	Keyword  string `json:"keyword" form:"keyword,optional"`
	AuthorID uint   `json:"authorId" form:"authorId,optional"`
}

type ArticleListResponse struct {
	PageResponse
}

// ============== 实时保存草稿 ==============

type SaveDraftRequest struct {
	ArticleID uint   `json:"articleId,optional"` // 0 表示新文章
	Title     string `json:"title,optional"`
	Content   string `json:"content,optional"`
}

type SaveDraftResponse struct {
	DraftID uint   `json:"draftId"`
	SavedAt string `json:"savedAt"`
}

// ============== 版本历史 ==============

type ArticleVersionListRequest struct {
	ArticleID uint `json:"articleId" path:"articleId" validate:"required"`
}

type ArticleVersionResponse struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"articleId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Version   int    `json:"version"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
}

// ============== WebSocket 实时同步 ==============

type ArticleSyncMessage struct {
	Type      string          `json:"type"` // edit/cursor/presence
	ArticleID uint            `json:"articleId"`
	UserID    uint            `json:"userId"`
	UserName  string          `json:"userName"`
	Content   string          `json:"content,omitempty"`
	Cursor    *CursorPosition `json:"cursor,omitempty"`
	Timestamp int64           `json:"timestamp"`
}

type CursorPosition struct {
	Index  int `json:"index"`
	Length int `json:"length"`
}
