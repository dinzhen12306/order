package logic

import (
	"context"
	"order/internal/svc"
	"order/models/mysql"
	"order/order"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderByOrderNOLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderByOrderNOLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderByOrderNOLogic {
	return &UpdateOrderByOrderNOLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderByOrderNOLogic) UpdateOrderByOrderNO(in *order.UpdateOrderByOrderNOReq) (*order.UpdateOrderByOrderNOResp, error) {
	orders := mysql.PbToOrder(in.Data)
	now := time.Now()
	orders.PaymentTime = &now
	err := orders.UpdateByOrderNO(l.svcCtx)
	if err != nil {
		return nil, err
	}
	return &order.UpdateOrderByOrderNOResp{
		Data: mysql.OrderToPb(orders),
	}, nil
}
