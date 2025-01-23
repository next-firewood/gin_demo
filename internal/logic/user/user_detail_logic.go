package user

import (
	"context"
	"gin_demo/internal/api"
	"gin_demo/internal/api/user"
	"gin_demo/internal/svc"
	"github.com/pkg/errors"
)

type UserDetailLogic struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewUserDetailLogic(context context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Ctx:    context,
		SvcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *api.UuidForm) (resp *user.DetailResp, err error) {
	if req.Uuid == "" {
		return resp, errors.New("没有传入参数")
	}

	return &user.DetailResp{Uuid: req.Uuid, Name: "test"}, err
}
