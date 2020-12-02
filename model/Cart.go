package model

type Cart struct{
	CartID string  //购物车ID
	CartItems []*CartItem // 购物项切片
	TotalCount int64 // 购物车中图书数量
	TotalAmount float64 // 购物车中图书总金额
	UserID int // 购物车所属用户ID
	UserName string // 购物车所属的用户名
}

// GetTotalCount 获取购物车中图书总数量
func (c *Cart)GetTotalCount()int64{
	var tc int64
	for _,v:=range c.CartItems{
		tc+=v.Count
	}
	return tc
}

// 获取购物车中图书总金额
func (c *Cart)GetTotalAmount() float64{
	var ta float64
	for _,v:=range c.CartItems{
		ta+=v.GetAmount()
	}
	return ta
}
