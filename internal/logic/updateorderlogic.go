package logic

import (
	"context"
	"order/models/mysql"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *order.UpdateOrderReq) (*order.UpdateOrderResp, error) {
	orders := mysql.PbToOrder(in.Data)
	err := orders.Update(l.svcCtx)
	if err != nil {
		return nil, err
	}
	return &order.UpdateOrderResp{
		Data: mysql.OrderToPb(orders),
	}, nil
}
