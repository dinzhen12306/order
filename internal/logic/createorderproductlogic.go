package logic

import (
	"context"
	"order/models/mysql"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderProductLogic {
	return &CreateOrderProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderProductLogic) CreateOrderProduct(in *order.CreateOrderProductReq) (*order.CreateOrderProductResp, error) {
	orders := mysql.NewOrderProduct()
	orderProducts := make([]*mysql.OrderProduct, 0)
	for _, v := range in.Data {
		orderProducts = append(orderProducts, mysql.PbToOrderProduct(v))
	}
	products, err := orders.CreateOrderProducts(l.svcCtx, orderProducts)
	if err != nil {
		return nil, err
	}
	data := make([]*order.OrderProductInfo, 0)
	for _, v := range products {
		data = append(data, mysql.OrderProductToPb(v))
	}
	return &order.CreateOrderProductResp{
		Data: data,
	}, nil
}
