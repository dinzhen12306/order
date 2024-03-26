package logic

import (
	"context"
	"order/models/mysql"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersLogic {
	return &GetOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrdersLogic) GetOrders(in *order.GetOrdersReq) (*order.GetOrdersResp, error) {
	orders := mysql.NewOrder()
	which := make(map[string]interface{}, 0)
	for k, v := range in.Where {
		which[k] = v
	}
	gets, err := orders.Gets(l.svcCtx, which)
	if err != nil {
		return nil, err
	}
	data := make([]*order.OrderInfo, 0)
	for _, v := range gets {
		data = append(data, mysql.OrderToPb(v))
	}
	return &order.GetOrdersResp{
		Data: data,
	}, nil
}
