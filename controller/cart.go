package controller

import (
	"bookstore/common"
	"bookstore/dao"
	"bookstore/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加书本到购物车
func AddBookToCart(ctx *gin.Context){
	// 判断是否登录
	if session,ok:=dao.IsLogin(ctx);ok{
		bookId:=ctx.PostForm("bookId")
		book,_:=dao.GetBookByID(bookId)
		// 获取用户id
		userID:=session.UserID
		// 判断数据库中是否有当前用户的购物车
		cart,_:=dao.GetCartByUserID(userID)
		if cart!=nil{
			// 当前用户有购物车，此时判断购物车里是否有这个购物项
			cartItem,_:=dao.GetCartItemByBookIDAndCartID(bookId,cart.CartID)
			if cartItem!=nil{
				// 如果不为空，则购物车里已有该购物项，此时只需数据加1
				// 获取购物车切片中的所有购物项，遍历找到匹配项
				cartItems:=cart.CartItems
				for _,v:=range cartItems{
					if v.Book.ID==cartItem.Book.ID{
						// 将购物车的count+1
						v.Count++
						// 更新数据库中该购物项的图书数量并更新数据库中购物车的总数量和总价格
						dao.UpdateBookCount(v,cart)
					}
				}
			}else{
				// 如果为空，则无该购物项，此时添加购物项信息到数据库
				newCartItem:=&model.CartItem{
					Book:book,
					Count:1,
					CartID: cart.CartID,
				}
				// 将购物项添加到切片中
				cart.CartItems=append(cart.CartItems,newCartItem)
				// 将新购物项加入到数据库并更新数据库中购物车的总数量和总价格
				dao.AddCartItem(newCartItem,cart)
			}
			//// 更新数据库中购物车的总数量和总价格
			//dao.UpdateCart(cart)
		}else{
			// 当前用户没有购物车，创建一个购物车添加到数据库
			cartID:=common.CreateUUID()
			cart:=&model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			// 创建购物车中的购物项组并将该购物项加入购物项组中
			var cartItems []*model.CartItem
			cartItem:=&model.CartItem{
				Book:book,
				Count:1,
				CartID: cartID,
			}
			cartItems=append(cartItems,cartItem)
			// 将购物项组赋值给购物车
			cart.CartItems=cartItems
			// 将购物车添加到数据库
			dao.AddCart(cart)
		}

		// 将新增的图书名返回客户端
		ctx.String(http.StatusOK,fmt.Sprintf("您刚刚将%s添加到购物车中", book.Title))

	}else{
		// 未登录
		ctx.String(http.StatusOK,"请先登录！")
	}
}

// 获取购物车信息
func GetCartInfo(ctx *gin.Context){
	// 判断是否登录
	session,ok:=dao.IsLogin(ctx)
	if !ok{
		ctx.HTML(http.StatusUnprocessableEntity,"login.html",nil)
		return
	}

	// 获取用户id
	userID:=session.UserID
	// 从数据库中获取当前用户的购物车
	cart,_:=dao.GetCartByUserID(userID)

	if cart!=nil{
		// 设置用户名
		cart.UserName=session.UserName
		cart.UserID=session.UserID

		ctx.HTML(http.StatusOK,"cart.html",cart)
	}else{
		cart:=&model.Cart{}
		cart.UserName=session.UserName
		cart.UserID=session.UserID
		ctx.HTML(http.StatusOK,"cart.html",cart)
	}
}

// 清空购物车
func DeleteCart(ctx *gin.Context){
	// 获取购物车id
	cartID:=ctx.Query("cartId")
	// 清空购物车
	dao.DeleteCartByCartID(cartID)
	//再次查询购物车信息
	GetCartInfo(ctx)
}

// 删除购物项
func DeleteCartItem(ctx *gin.Context){
	// 获取要删除的购物项id
	cartItemID:=ctx.Query("cartItemId")
	iCartItemID,_:=strconv.ParseInt(cartItemID,10,64)
	//获取session
	session,_:=dao.IsLogin(ctx)
	userID:=session.UserID
	// 获取用户的购物车
	cart,_:=dao.GetCartByUserID(userID)
	// 遍历购物车内的所有购物项
	cartItems:=cart.CartItems
	for k,v:=range cartItems{
		if v.CartItemID==iCartItemID{
			// 从购物项组中删除当前购物项
			cartItems=append(cartItems[:k],cartItems[k+1:]...)
			cart.CartItems=cartItems
			// 从数据库中删除当前购物项信息
			dao.DeleteCartItemByID(cartItemID,cart)
		}
	}

	// 再次查询购物车信息
	GetCartInfo(ctx)
}

// 更新购物项
func UpdateCartItem(ctx *gin.Context){
	// 获取要更新的购物项id
	cartItemID:=ctx.Query("cartItemId")
	iCartItemID,_:=strconv.ParseInt(cartItemID,10,64)
	// 获取要更新的图书数量
	bookCount:=ctx.Query("bookCount")
	ibookCount,_:=strconv.ParseInt(bookCount,10,64)
	//获取session
	session,_:=dao.IsLogin(ctx)
	userID:=session.UserID
	// 获取用户的购物车
	cart,_:=dao.GetCartByUserID(userID)
	// 遍历购物车内的所有购物项
	cartItems:=cart.CartItems
	for _,v:=range cartItems{
		if v.CartItemID==iCartItemID{
			// 将图书数量赋值给购物项
			v.Count=ibookCount
			// 从数据库中更新购物项信息并更新购物车信息
			dao.UpdateBookCount(v,cart)
		}
	}
	// 再次查询购物车信息
	GetCartInfo(ctx)
}