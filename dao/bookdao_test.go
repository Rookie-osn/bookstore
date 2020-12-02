package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestRun(t *testing.T){
	//testBook(t)
	//testAddCartItem(t)
	testAddCartItem(t)
}

func testBook(t *testing.T){
	page:=&model.Page{}
	var err error
	page,err=GetPageBooksByPageNo("3")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(page)
	for _,v:=range page.Books{
		fmt.Println(v)
	}
}

// 添加到购物项
func testAddCartItem(t *testing.T) {
	book:=&model.Book{
		ID: 3,
	}
	ci:=&model.CartItem{
		Book:book,
		Count:1,
		CartID: "3333333333333333",
	}
	err:=AddCartItem(ci)
	fmt.Println(err)
}

// 添加订单项
func testAddOrderItem(t *testing.T){
	oi:=&model.OrderItem{
		Count: 2,
		Amount: 12.2,
		Title: "虚拟小说",
		Author: "虚拟作者",
		Price: 6.1,
		ImgPath: "default.jpg",
		OrderID:"xxxxxxxxxxx",
	}

	err:=AddOrderItem(oi)
	if err!=nil{
		fmt.Println(err)
	}
}