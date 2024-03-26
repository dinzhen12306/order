package logic

import (
	"context"
	"order/models/mysql"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderReq) (*order.CreateOrderResp, error) {
	orders := mysql.PbToOrder(in.Data)
	err := orders.Create(l.svcCtx)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResp{
		ID: int64(orders.ID),
	}, nil
}
