package logic

import (
	"context"

	"acupofcoffee/api/internal/svc"
	"acupofcoffee/api/internal/types"
	"acupofcoffee/common/errorx"
	"acupofcoffee/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) GetUserInfo() (*types.UserInfoResponse, error) {
	userID, ok := l.ctx.Value("userId").(uint)
	if !ok {
		return nil, errorx.NewCodeError(401, "未登录")
	}

	var user model.User
	result := l.svcCtx.DB.First(&user, userID)
	if result.Error != nil {
		return nil, errorx.NewCodeError(404, "用户不存在")
	}

	return &types.UserInfoResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (l *UserLogic) UpdateUserInfo(req *types.UpdateUserRequest) error {
	userID, ok := l.ctx.Value("userId").(uint)
	if !ok {
		return errorx.NewCodeError(401, "未登录")
	}

	updates := make(map[string]interface{})
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}

	if len(updates) == 0 {
		return nil
	}

	result := l.svcCtx.DB.Model(&model.User{}).Where("id = ?", userID).Updates(updates)
	if result.Error != nil {
		l.Logger.Errorf("update user error: %v", result.Error)
		return errorx.NewDefaultError("更新失败")
	}

	return nil
}

