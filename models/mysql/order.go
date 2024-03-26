package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"order/internal/config"
	"order/internal/svc"
	pb "order/order"
	"time"
)

type Order struct {
	gorm.Model
	UserId        int64      `gorm:"int;comment('用户id')"`
	ProductId     int64      `gorm:"int;comment('商品id')"`
	AddressId     int64      `gorm:"int;comment('用户地址')"`
	OrderNo       string     `gorm:"index;varchar(50);comment('订单编号')"`
	Title         string     `gorm:"varchar(255);comment('商品信息')"`
	TotalAmount   string     `gorm:"varchar(255);comment('总金额')"`
	PaidAmount    string     `gorm:"varchar(255);comment('实际支付金额')"`
	Status        int64      `gorm:"int;comment('0待支付、1已支付、2已发货、3已完成、4已取消')"`
	PaymentMethod string     `gorm:"varchar(30);comment('支付方式')"`
	OrderNotes    string     `gorm:"varchar(255);comment('订单备注')"`
	PaymentTime   *time.Time `gorm:"comment('支付时间')"`
}

type OrderProduct struct {
	gorm.Model
	OrderId         int64  `gorm:"int;comment('订单id')"`
	ValueOneId      int64  `gorm:"int;comment('规格1')"`
	ValueTwlId      int64  `gorm:"int;comment('规格2')"`
	ValueThreeId    int64  `gorm:"int;comment('规格3')"`
	ProductPrice    string `gorm:"decimal(10,2);comment('商品价格')"`
	ProductQuantity int64  `gorm:"int;comment('商品数量')"`
}

func withDB(c config.Config, fun func(db *gorm.DB) error) error {
	open, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.MysqlConf.Username, c.MysqlConf.Password, c.MysqlConf.Host, c.MysqlConf.Port, c.MysqlConf.DatabasesName)), &gorm.Config{})
	if err != nil {
		return err
	}
	defer func() {
		db, _ := open.DB()
		db.Close()
	}()
	return fun(open)
}

func Sync(c config.Config) error {
	log.Println(c.MysqlConf)
	return withDB(c, func(db *gorm.DB) error {
		return db.AutoMigrate(new(Order), new(OrderProduct))
	})
}

func NewOrder() *Order {
	return new(Order)
}

func NewOrderProduct() *OrderProduct {
	return new(OrderProduct)
}

func (o *Order) Create(c *svc.ServiceContext) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		o.PaymentTime = nil
		return db.Create(o).Error
	})
}

func (o *Order) Update(c *svc.ServiceContext) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Where("id = ?", o.ID).Updates(o).Error
	})
}

func (o *Order) UpdateByOrderNO(c *svc.ServiceContext) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Where("order_no = ?", o.OrderNo).Updates(o).Error
	})
}

func (o *Order) Get(c *svc.ServiceContext, where map[string]interface{}) error {
	return withDB(c.Config, func(db *gorm.DB) error {
		return db.Where(where).First(o).Error
	})
}

func (o *Order) Gets(c *svc.ServiceContext, where map[string]interface{}) ([]*Order, error) {
	orders := make([]*Order, 0)
	err := withDB(c.Config, func(db *gorm.DB) error {
		return db.Where(where).Find(&orders).Error
	})
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (p *OrderProduct) CreateOrderProducts(c *svc.ServiceContext, info []*OrderProduct) ([]*OrderProduct, error) {
	err := withDB(c.Config, func(db *gorm.DB) error {
		return db.Create(info).Error
	})
	if err != nil {
		return nil, err
	}
	return info, err
}

func OrderToPb(order *Order) *pb.OrderInfo {
	payTime := ""
	if order.PaymentTime == nil {
		payTime = "nil"
	} else {
		payTime = order.PaymentTime.String()
	}
	return &pb.OrderInfo{
		ID:            int64(order.ID),
		UserID:        order.UserId,
		ProductID:     order.ProductId,
		AddressID:     order.AddressId,
		OrderNO:       order.OrderNo,
		Title:         order.Title,
		TotalAmount:   order.TotalAmount,
		PaidAmount:    order.PaidAmount,
		Status:        order.Status,
		PaymentMethod: order.PaymentMethod,
		OrderNotes:    order.OrderNotes,
		PaymentTime:   payTime,
	}
}
func PbToOrder(info *pb.OrderInfo) *Order {
	return &Order{
		Model:         gorm.Model{ID: uint(info.ID)},
		UserId:        info.UserID,
		ProductId:     info.ProductID,
		AddressId:     info.AddressID,
		OrderNo:       info.OrderNO,
		Title:         info.Title,
		TotalAmount:   info.TotalAmount,
		PaidAmount:    info.PaidAmount,
		Status:        info.Status,
		PaymentMethod: info.PaymentMethod,
		OrderNotes:    info.OrderNotes,
	}
}

func PbToOrderProduct(info *pb.OrderProductInfo) *OrderProduct {
	return &OrderProduct{
		Model:           gorm.Model{ID: uint(info.ID)},
		OrderId:         info.OrderID,
		ValueOneId:      info.ValueOneID,
		ValueTwlId:      info.ValueTwoID,
		ValueThreeId:    info.ValueThreeID,
		ProductPrice:    info.ProductPrice,
		ProductQuantity: info.ProductQuantity,
	}
}
func OrderProductToPb(product *OrderProduct) *pb.OrderProductInfo {
	return &pb.OrderProductInfo{
		ID:              int64(product.ID),
		OrderID:         product.OrderId,
		ValueOneID:      product.ValueOneId,
		ValueTwoID:      product.ValueTwlId,
		ValueThreeID:    product.ValueThreeId,
		ProductPrice:    product.ProductPrice,
		ProductQuantity: product.ProductQuantity,
	}
}
