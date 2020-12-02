package dao

import (
	"bookstore/common"
	"bookstore/model"
	"fmt"
)

func AddOrderItem(oi *model.OrderItem) error {
	db:=common.GetDB()
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	//执行
	_, err := db.Exec(sqlStr, oi.Count, oi.Amount, oi.Title, oi.Author, oi.Price, oi.ImgPath, oi.OrderID)
	if err != nil {
		return err
	}
	return nil
}

//GetOrderItemsByOrderID 根据订单号获取该订单的所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	db:=common.GetDB()
	sqlStr := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id=?"
	rows, err := db.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}

	//将数据库数据赋值给orderitem变量
	var orderItems []*model.OrderItem
	for rows.Next() {
		oi := &model.OrderItem{}
		err = rows.Scan(&oi.OrderItemID, &oi.Count, &oi.Amount, &oi.Title, &oi.Author, &oi.Price, &oi.ImgPath, &oi.OrderID)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		orderItems = append(orderItems, oi)
	}
	return orderItems, nil
}
