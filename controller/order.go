package controller

import (
	"bookstore/common"
	"bookstore/dao"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 结算订单
func Checkout(ctx *gin.Context){
	// 获取session
	session,_:=dao.IsLogin(ctx)
	userID:=session.UserID
	cart,_:=dao.GetCartByUserID(userID)
	// 创建订单号
	orderID:=common.CreateUUID()
	//创建生成订单的时间
	timeNow:=time.Now().Format("2006-01-02 15:04:05")
	// 创建订单
	order:=&model.Order{
		OrderID: orderID,
		CreateTime: timeNow,
		TotalCount: cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:0,
		UserID: int64(userID),
	}
	// 将订单添加入库，更新订单信息，更新订单项信息，更新图书信息（数量）
	dao.AddOrderA(order,cart)

	// 删除购物车
	dao.DeleteCartByCartID(cart.CartID)
	// 将订单号也传出
	session.Order=order

	ctx.HTML(http.StatusOK,"checkout.html",session)
}

// 订单管理界面获取所有订单
func GetOrders(ctx *gin.Context){
	orders,_:=dao.GetOrders()
	ctx.HTML(http.StatusOK,"order_manager.html",orders)
}

// 订单详情，查看订单项信息
func GetOrderInfo(ctx *gin.Context){
	orderID:=ctx.Query("orderId")
	orderItems,_:=dao.GetOrderItemsByOrderID(orderID)
	ctx.HTML(http.StatusOK,"order_info.html",orderItems)
}

// 通过userId获取我的订单
func GetMyOrders(ctx *gin.Context){
	// 获取session，取出userId，查询订单信息并返回
	session,_:=dao.IsLogin(ctx)
	userID:=session.UserID
	orders,_:=dao.GetMyOrders(userID)
	session.Orders=orders

	ctx.HTML(http.StatusOK,"order.html",session)
}

// 发货
func SendOrder(ctx *gin.Context){
	// 获取订单id麻将订单的state更新为1
	orderID:=ctx.Query("orderId")
	dao.UpdateOrderState(orderID,1)
	// 刷新订单页面信息
	GetOrders(ctx)
}

// 收货
func TakeOrder(ctx *gin.Context){
	orderID:=ctx.Query("orderId")
	dao.UpdateOrderState(orderID,2)
	GetMyOrders(ctx)
}
