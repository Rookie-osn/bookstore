package dao

import (
	"bookstore/common"
	"bookstore/model"
)

//添加购物车
func AddCart(cart *model.Cart) error {
	db:=common.GetDB()
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"

	//执行
	_, err :=db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//将购物车内的购物项添加进购物项表中
	for _, v := range cart.CartItems {
		AddCartItem(v)
	}

	return nil
}

//通过用户id查询数据库获取购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	db:=common.GetDB()
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"

	//执行
	row := db.QueryRow(sqlStr, userID)

	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	//获取当前购物车所有的购物项
	cartitems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartitems
	return cart, nil
}

//更新购物车中的图书总数量和总价格
func UpdateCart(cart *model.Cart) error {
	db:=common.GetDB()
	sqlStr := "update carts set total_count=?,total_amount=? where id=?"

	_, err := db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

//通过购物车id删除购物车数据
func DeleteCartByCartID(cartID string) error {
	db:=common.GetDB()
	tx,_:=db.Begin()
	//先删除购物项，才能删除购物车
	sqlStr := "delete from cart_items where cart_id=?"

	_, err := tx.Exec(sqlStr, cartID)
	if err != nil {
		return err
		tx.Rollback()
	}

	sqlStr = "delete from carts where id=?"
	_, err = tx.Exec(sqlStr, cartID)
	if err != nil {
		return err
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
