package logic

import (
	"context"
	"order/models/mysql"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *order.GetOrderReq) (*order.GetOrderResp, error) {
	orders := mysql.NewOrder()
	which := make(map[string]interface{}, 0)
	for k, v := range in.Where {
		which[k] = v
	}
	err := orders.Get(l.svcCtx, which)
	if err != nil {
		return nil, err
	}
	return &order.GetOrderResp{
		Data: mysql.OrderToPb(orders),
	}, nil
}
