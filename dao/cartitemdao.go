package dao

import (
	"bookstore/common"
	"bookstore/model"
)

// 添加cartitem数据
func AddCartItem(ci *model.CartItem)error{
	db:=common.GetDB()
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"

	//执行
	_, err := db.Exec(sqlStr,ci.Count,ci.GetAmount(), ci.Book.ID, ci.CartID)
	if err != nil {
		return err
	}
	return nil
}

//根据图书的id和购物车id获取对应的购物项（用于判别该用户购物车是否已存在目标商品，如果已存在，则通过修改数量增加商品）
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	db:=common.GetDB()
	sqlStr := "select id,count,amount from cart_items where book_id=? and cart_id=?"
	//执行
	row := db.QueryRow(sqlStr, bookID, cartID)
	//创建cartitem变量，返回
	cartitem := &model.CartItem{}
	err := row.Scan(&cartitem.CartItemID, &cartitem.Count, &cartitem.Amount)
	if err != nil {
		return nil, err
	}
	//获取图书信息
	book, _ := GetBookByID(bookID)
	cartitem.Book = book

	return cartitem, err
}

//根据购物车id获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	db:=common.GetDB()
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id=?"

	//执行
	rows, err := db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}

	var cartitems []*model.CartItem
	for rows.Next() {
		var bookid string
		cartitem := &model.CartItem{}
		err := rows.Scan(&cartitem.CartItemID, &cartitem.Count, &cartitem.Amount, &bookid, &cartitem.CartID)
		if err != nil {
			continue
		}
		//通过bookid获取book信息
		book, _ := GetBookByID(bookid)
		//将book变量放到购物项中
		cartitem.Book = book
		cartitems = append(cartitems, cartitem)
	}
	return cartitems, err
}

//根据图书id和购物车id以及图书的数量更新购物项中的图书数量和金额
func UpdateBookCount(cartItem *model.CartItem) error {
	db:=common.GetDB()
	sqlStr := "update cart_items set count=?,amount=? where book_id=? and cart_id=?"

	_, err := db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//根据购物车ID删除购物项
func DeleteCartItemsByCartID(cartID string) error {
	db:=common.GetDB()
	sqlStr := "delete from cart_items where cart_id=?"

	_, err := db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

//根据购物项id删除购物项信息
func DeleteCartItemByID(cartItemID string) error {
	db:=common.GetDB()
	sqlStr := "delete from cart_items where id=?"
	//执行
	_, err := db.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil
}