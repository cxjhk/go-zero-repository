package logic

import (
	"context"

	"go-zero-repository/user/rpc/internal/svc"
	"go-zero-repository/user/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *rpc.Request) (*rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &rpc.Response{}, nil
}
