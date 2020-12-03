package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 获取页码筛选分页的图书信息（首页）
func GetPageBooksByPageNo(ctx *gin.Context){
	// 获取参数
	pageNo:=ctx.Query("pageNo")
	if pageNo==""{
		pageNo="1"
	}

	// 获取分页图书信息
	page,_:=dao.GetPageBooksByPageNo(pageNo)
	ctx.HTML(http.StatusOK,"index.html",page)
}

// 获取价格筛选分页的图书信息
func GetPageBooksByPrice(ctx *gin.Context){
	// 获取传入的页码
	pageNo:=ctx.Query("pageNo")
	if pageNo==""{
		pageNo="1"
	}

	minPrice:=ctx.Query("min")
	maxPrice:=ctx.Query("max")

	var page *model.Page
	if minPrice==""&& maxPrice==""{
		// 调用bookdao中带分页的图书的函数
		page,_=dao.GetPageBooksByPageNo(pageNo)
	}else{
		// 调用bookdao中带价格过滤分页的图书的函数
		page,_=dao.GetPageBooksByPrice(pageNo,minPrice,maxPrice)
		// 将传入的价格赋值给page
		page.MinPrice=minPrice
		page.MaxPrice=maxPrice
	}

	// 查询是否已登录，返回用户名和bool类型值
	if session,ok:=dao.IsLogin(ctx);ok{
		// 已登录
		page.IsLogin=true
		page.Username=session.UserName
	}

	ctx.HTML(http.StatusOK,"index.html",page)
}

// 删除图书
func DeleteBook(ctx *gin.Context){
	bookId:=ctx.Query("bookId")
	dao.DeleteBook(bookId)
	// 调用GetPageBooks函数刷新页面
	GetPageBooksByPageNo(ctx)
}

// 跳转到更新或添加图书的页面
func ToUpdateBookPage(ctx *gin.Context){
	// 获取要更新的图书ID
	bookId:=ctx.Query("bookId")
	book,_:=dao.GetBookByID(bookId)

	// 判断是添加还是更新
	if bookId!="0"{
		// 代表图书存在，则更新
		ctx.HTML(http.StatusOK,"book_edit.html",book)
	}else{
		ctx.HTML(http.StatusOK,"book_edit.html",nil)
	}
}

// 更新图书信息
func UpdateOrAddBook(ctx *gin.Context){
	// 获取图书信息
	bookId:=ctx.PostForm("bookId")
	title:=ctx.PostForm("title")
	author:=ctx.PostForm("author")
	price:=ctx.PostForm("price")
	sales:=ctx.PostForm("sales")
	stock:=ctx.PostForm("stock")
	// 转换部分数据格式
	ibookId,_:=strconv.ParseInt(bookId,10,0)
	fprice,_:=strconv.ParseFloat(price,64)
	isales,_:=strconv.ParseInt(sales,10,0)
	istock,_:=strconv.ParseInt(stock,10,0)
	// 创建Book
	book:=&model.Book{
		ID:int(ibookId),
		Title: title,
		Author: author,
		Price: fprice,
		Sales: int(isales),
		Stock: int(istock),
		ImgPath: "/static/img/default.jpg",
	}

	// 判断是更新还是添加
	if book.ID>0{
		// 表示有ID传入，是更新
		dao.UpdateBook(book)
	}else{
		dao.AddBook(book)
	}
	// 调用GetBooks函数刷新
	GetPageBooksByPageNo(ctx)
}