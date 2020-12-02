package model

type CartItem struct{
	CartItemID int64 //购物项ID
	Book *Book //购物项中的图书信息
	Count int64 // 购物项的图书信息
	Amount float64 // 购物项中的图书金额
	CartID string // 当前购物项所属购物车ID
}

// 获取购物项金额
func (ci *CartItem)GetAmount()float64{
	price:=ci.Book.Price
	return float64(ci.Count)*price
}