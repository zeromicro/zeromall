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
	/*
		TODO X :注册流程
			1. 校验参数
			2. 校验账号是否存在
			3. 根据注册方式, 执行不同的注册逻辑:
				3.1 手机号注册: 生成验证码, 发送短信, 保存验证码, 保存账号, 返回账号信息, 返回token,
				3.2 邮箱注册: 生成验证码, 发送邮件, 保存验证码, 保存账号
				3.3 用户名+密码注册: 保存账号, 返回账号信息, 返回token
	*/

	var resp = new(pb.UserRegisterResp)

	// query if exist
	authUnique := in.Username
	l.Logger.Debugf("Register req=%+v", in)
	ret, err := l.svcCtx.MAuthn.FindOneByAuthUnique(l.ctx, authUnique)
	l.Logger.Debugf("Register query account: ret=%+v, err=%v", ret, err)

	// existed, return error
	if ret != nil {
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
