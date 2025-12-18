package model

// User 用户模型
type User struct {
	BaseModel
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"-"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Nickname string `gorm:"type:varchar(50)" json:"nickname"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`
	Phone    string `gorm:"type:varchar(20);index" json:"phone"`
	Status   int8   `gorm:"type:tinyint;default:1;comment:状态 1:正常 0:禁用" json:"status"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}

// IsActive 判断用户是否激活
func (u *User) IsActive() bool {
	return u.Status == 1
}
