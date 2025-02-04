package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"zhihu/application/applet/internal/svc"
	"zhihu/application/applet/internal/types"
	"zhihu/application/user/user"
	"zhihu/pkg/ecode"
	"zhihu/pkg/validator"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	if err = validator.Struct(req); err != nil {
		err = ecode.RequestErr
		return
	}
	userId, err := l.ctx.Value(types.UserIdKey).(json.Number).Int64()
	if err != nil {
		logx.Errorf("l.ctx.Value(%s) error: %v", types.UserIdKey, err)
		return nil, err
	}
	if userId == 0 {
		return &types.InfoResponse{}, nil
	}
	idResp, err := l.svcCtx.UserRpc.FindById(l.ctx, &user.FindByIdRequest{UserId: userId})
	if err != nil {
		logx.Errorf("UserRpc->FindById userId: %d error: %v", userId, err)
		return nil, err
	}
	if idResp.UserId <= 0 {
		return nil, nil
	}
	return &types.InfoResponse{
		UserId:   idResp.UserId,
		Username: idResp.Username,
		Avatar:   idResp.Avatar,
	}, nil
}
