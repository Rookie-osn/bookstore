package model

type OrderItem struct {
	OrderItemID int64   //订单项的id
	Count       int64   //订单项中的图书数量
	Amount      float64 //订单项中的图书金额
	Title       string  //订单项中图书的书名
	Author      string  //订单项中图书的作者
	Price       float64 //订单项中图书的价格
	ImgPath     string  //订单项中图书封面
	OrderID     string  //订单项所属订单ID
}
