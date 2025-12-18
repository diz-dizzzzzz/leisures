package model

// Article 文章模型
type Article struct {
	BaseModel
	Title      string `gorm:"type:varchar(255);not null;index" json:"title"`
	Content    string `gorm:"type:longtext" json:"content"`     // 富文本内容（JSON字符串）
	ContentRaw string `gorm:"type:longtext" json:"contentRaw"`  // 纯文本内容（用于搜索）
	Cover      string `gorm:"type:varchar(500)" json:"cover"`   // 封面图
	Summary    string `gorm:"type:varchar(500)" json:"summary"` // 摘要
	AuthorID   uint   `gorm:"index;not null" json:"authorId"`   // 作者ID
	Author     *User  `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Status     int8   `gorm:"type:tinyint;default:0;index" json:"status"` // 0:草稿 1:已发布 2:已归档
	Version    int    `gorm:"default:1" json:"version"`                   // 版本号
	ViewCount  int64  `gorm:"default:0" json:"viewCount"`                 // 浏览量
	LikeCount  int64  `gorm:"default:0" json:"likeCount"`                 // 点赞数
}

func (Article) TableName() string {
	return "articles"
}

// ArticleVersion 文章版本历史
type ArticleVersion struct {
	BaseModel
	ArticleID uint   `gorm:"index;not null" json:"articleId"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Content   string `gorm:"type:longtext" json:"content"`
	Version   int    `gorm:"not null" json:"version"`
	Remark    string `gorm:"type:varchar(255)" json:"remark"` // 版本备注
}

func (ArticleVersion) TableName() string {
	return "article_versions"
}

// ArticleDraft 文章草稿（实时保存）
type ArticleDraft struct {
	BaseModel
	ArticleID uint   `gorm:"uniqueIndex;not null" json:"articleId"` // 0 表示新文章
	UserID    uint   `gorm:"index;not null" json:"userId"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Content   string `gorm:"type:longtext" json:"content"`
}

func (ArticleDraft) TableName() string {
	return "article_drafts"
}

// ArticleStatus 文章状态常量
const (
	ArticleStatusDraft     int8 = 0 // 草稿
	ArticleStatusPublished int8 = 1 // 已发布
	ArticleStatusArchived  int8 = 2 // 已归档
)
