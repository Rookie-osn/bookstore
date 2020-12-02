package main

import(
	"bookstore/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func CollectRouter(r *gin.Engine)*gin.Engine{
	staticDir,_:=os.Getwd()
	r.Static("/static",staticDir+"/views/static")
	r.Static("/pages",staticDir+"/views/pages")
	// 加载html文件
	r.LoadHTMLGlob("views/pages/**/*")
	// 主页
	r.GET("/",controller.GetPageBooksByPrice)
	r.GET("/main",controller.GetPageBooksByPrice)

	// 用户
	// 处理登陆
	r.POST("/login",controller.Login)
	// 处理注销
	r.GET("/logout",controller.Logout)
	// 处理注册
	r.POST("/regist",controller.Regist)

	// 图书
	// 获取分页图书
	r.GET("/getPageBooks",controller.GetPageBooksByPageNo)
	// 根据价格筛选图书
	r.GET("/getPageBooksByPrice",controller.GetPageBooksByPrice)
	// 添加或更新图书
	r.GET("/toUpdateBookPage",controller.ToUpdateBookPage)
	// 删除图书
	r.GET("/deleteBook",controller.DeleteBook)

	
	// 更新或添加图书信息
	r.POST("/updateOrAddBook",controller.UpdateOrAddBook)

	// 购物车
	// 添加图书到购物车中
	r.POST("/addBook2Cart",controller.AddBookToCart)
	// 获取购物车信息
	r.GET("/getCartInfo",controller.GetCartInfo)
	// 清空购物车
	r.GET("/deleteCart",controller.DeleteCart)
	// 删除购物项
	r.GET("/deleteCartItem",controller.DeleteCartItem)
	// 更新购物项
	r.GET("/updateCartItem",controller.UpdateCartItem)

	// 订单
	// 订单结算
	r.GET("/checkout",controller.Checkout)
	// 通过订单id获取订单项信息
	r.GET("/getOrderInfo",controller.GetOrderInfo)
	// 通过用户id获取我的订单
	r.GET("/getMyOrders",controller.GetMyOrders)
	// 发货
	r.GET("/sendOrder",controller.SendOrder)
	// 确认发货
	r.GET("/takeOrder",controller.TakeOrder)

	// 获取所有订单信息
	r.GET("/getOrders",controller.GetOrders)

	// 通过Ajax请求验证用户名是否可用
	r.POST("/checkUserName",controller.IfUserExist)

	return r
}