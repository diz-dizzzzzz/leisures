package types

// ============== 用户相关 ==============

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Nickname string `json:"nickname" validate:"max=50"`
}

type UserInfoResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname" validate:"max=50"`
	Avatar   string `json:"avatar" validate:"max=255"`
}

// ============== 分页相关 ==============

type PageRequest struct {
	Page     int `json:"page" form:"page,optional"`
	PageSize int `json:"pageSize" form:"pageSize,optional"`
}

type PageResponse struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
}

// ============== 通用 ==============

type IDRequest struct {
	ID uint `json:"id" path:"id"`
}
