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
	defer rows.Close()

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
	defer rows.Close()

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

// 添加订单
func AddOrderA(order *model.Order,cart *model.Cart) error {
	db:=common.GetDB()
	tx,_:=db.Begin()
	// 添加订单数据到orders表
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := tx.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 遍历订单项，加入orderitem表
	cis:=cart.CartItems
	for _,v:=range cis{
		orderItem:=&model.OrderItem{
			Count: v.Count,
			Amount: v.Amount,
			Title: v.Book.Title,
			Author: v.Book.Author,
			Price: v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: order.OrderID,
		}
		// 更新购物项信息入库
		sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
		//执行
		_, err := tx.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title,orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
		if err != nil {
			tx.Rollback()
			return err
		}

		// 将图书库存和销量更新
		book:=v.Book
		book.Sales=book.Sales+int(v.Count)
		book.Stock=book.Stock-int(v.Count)
		// 更新图书信息入库
		sqlStr = "Update books set title=?,author=?,price=?,sales=?,stock=? where id=?"

		//执行
		_, err = tx.Exec(sqlStr, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
