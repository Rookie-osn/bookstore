package dao

import (
	"bookstore/common"
	"bookstore/model"
	"fmt"
)

// 添加订单
func AddOrder(order *model.Order) error {
	db:=common.GetDB()
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

//获取所有订单
func GetOrders() ([]*model.Order, error) {
	db:=common.GetDB()
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err :=db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	//将数据库查到的数据放到切片中
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err = rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		orders = append(orders, order)
	}
	return orders, err
}

//根据用户id获取订单信息
func GetMyOrders(UserID int) ([]*model.Order, error) {
	db:=common.GetDB()
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id=?"
	rows, err :=db.Query(sqlStr, UserID)
	if err != nil {
		return nil, err
	}

	//将数据库查到的数据放到切片中
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err = rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		orders = append(orders, order)
	}
	return orders, err
}

//更新订单状态
func UpdateOrderState(orderID string, state int64) error {
	db:=common.GetDB()
	sqlStr := "update orders set state=? where id=?"

	_, err :=db.Exec(sqlStr, state, orderID)
	if err != nil {
		return err
	}
	return nil
}
