package logic

import (
	"context"
	"time"

	"acupofcoffee/api/internal/svc"
	"acupofcoffee/api/internal/types"
	"acupofcoffee/common/errorx"
	"acupofcoffee/common/utils"
	"acupofcoffee/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	var user model.User
	result := l.svcCtx.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		return nil, errorx.NewCodeError(401, "用户名或密码错误")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errorx.NewCodeError(401, "用户名或密码错误")
	}

	// 生成 JWT Token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := l.generateToken(user.ID, now, accessExpire)
	if err != nil {
		l.Logger.Errorf("generate token error: %v", err)
		return nil, errorx.NewDefaultError("登录失败")
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *AuthLogic) Register(req *types.RegisterRequest) error {
	// 检查用户名是否已存在
	var count int64
	l.svcCtx.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return errorx.NewCodeError(400, "用户名已存在")
	}

	// 检查邮箱是否已存在
	l.svcCtx.DB.Model(&model.User{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		return errorx.NewCodeError(400, "邮箱已被注册")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		l.Logger.Errorf("hash password error: %v", err)
		return errorx.NewDefaultError("注册失败")
	}

	// 创建用户
	user := model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Nickname: req.Nickname,
	}

	if result := l.svcCtx.DB.Create(&user); result.Error != nil {
		l.Logger.Errorf("create user error: %v", result.Error)
		return errorx.NewDefaultError("注册失败")
	}

	return nil
}

func (l *AuthLogic) generateToken(userID uint, iat, seconds int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"iat":    iat,
		"exp":    iat + seconds,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}

