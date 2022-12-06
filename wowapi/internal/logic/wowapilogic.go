package logic

import (
	"context"

	"wow_server/wowapi/internal/svc"
	"wow_server/wowapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WowapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWowapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WowapiLogic {
	return &WowapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WowapiLogic) Wowapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
