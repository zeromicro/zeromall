package logic

import (
	"context"
	"user/authn/proto/model"
	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 用户注册:
func (l *RegisterLogic) Register(in *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	// todo: add your logic here and delete this line

	var resp = new(pb.UserRegisterResp)

	// query if exist
	authUnique := in.Username
	l.Logger.Debugf("Register req=%+v", in)
	ret, err := l.svcCtx.MAuthn.FindOneByAuthUnique(l.ctx, authUnique)
	l.Logger.Debugf("Register query account: ret=%+v, err=%v", ret, err)
	if err != nil {
		// logger
		l.Logger.Errorf("FindOneByAuthUnique err=%v", err)
		return nil, err
	}

	// existed, return error
	if ret.AuthUnique != "" {
		// 重复
		resp.Status = "account repeated" // 0: success, -1: fail
		return resp, err
	}

	// not exist, insert

	var data = new(model.UserAuthn)
	data.From(in)
	result, err := l.svcCtx.MAuthn.Insert(l.ctx, data)
	if err != nil {
		// logger
		l.Logger.Errorf("Insert err=%v", err)
		return nil, err
	}
	l.Logger.Debugf("Register insert account: ret=%+v, err=%v", result, err)

	return resp, nil
}
