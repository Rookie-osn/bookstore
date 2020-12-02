package dao

import (
	"bookstore/common"
	"bookstore/model"
	"strconv"
)

//查找所有图书
func GetBooks() ([]*model.Book, error) {
	db:=common.GetDB()
	//连接数据库查询所有图书信息
	sqlStr := "Select id,title,author,price,sales,stock,img_path from books"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	//创建一个切片保存图书
	var books []*model.Book
	//遍历数据库查询到的数据
	for rows.Next() {
		//创建一个book结构体
		book := &model.Book{}
		//将数据赋值到变量中
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			continue
		}
		//将读取到的图书扩展到切片中
		books = append(books, book)
	}
	return books,err
}

//添加图书
func AddBook(book *model.Book) error {
	db:=common.GetDB()
	//数据库语句
	sqlStr := "Insert into books (title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"

	_, err := db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//删除图书
func DeleteBook(bookId string) error {
	db:=common.GetDB()
	//数据库语句
	sqlStr := "Delete from books where id=?"

	_, err := db.Exec(sqlStr, bookId)
	if err != nil {
		return err
	}
	return nil
}

//通过ID搜索book并返回book信息
func GetBookByID(bookId string) (*model.Book, error) {
	db:=common.GetDB()
	//数据库语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	//执行
	row :=db.QueryRow(sqlStr, bookId)

	//创建一个model.Book变量
	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

	if err != nil {
		return nil, err
	}
	return book, nil
}

//更新book信息
func UpdateBook(book *model.Book) error {
	db:=common.GetDB()
	sqlStr := "Update books set title=?,author=?,price=?,sales=?,stock=? where id=?"

	//执行
	_, err := db.Exec(sqlStr, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ID)
	if err != nil {
		return err
	}
	return nil
}

//传入页码，获取页码对应页上的图书信息
func GetPageBooksByPageNo(pageNo string) (*model.Page, error) {
	db:=common.GetDB()
	//将传入的页码转码
	ipageNo, _ := strconv.ParseInt(pageNo, 10, 0)
	//获取图书总数量
	sqlStr := "select count(*) from books"
	//创建总记录量变量
	var totalRecord int64
	//执行
	row := db.QueryRow(sqlStr)
	row.Scan(&totalRecord)

	//设置每页显示的图书量
	var pageSize int64
	pageSize = 4

	//计算总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取目标页的图书信息
	sqlStr = "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	//执行
	rows, err := db.Query(sqlStr, (ipageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	//创建books变量切片，接收查询的图书信息
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      ipageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, err
}


//传入页码、价格范围，获取页码对应页上的图书信息
func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*model.Page, error) {
	db:=common.GetDB()
	//将传入的页码转码
	ipageNo, _ := strconv.ParseInt(pageNo, 10, 0)

	//获取价格过滤后的图书数量
	sqlStr := "select count(*) from books where price between ? and ?"
	//创建总记录量变量
	var totalRecord int64
	//执行
	row := db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)

	//设置每页显示的图书量
	var pageSize int64
	pageSize = 4

	//计算总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取目标页的图书信息
	sqlStr = "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows, err := db.Query(sqlStr, minPrice, maxPrice, (ipageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	//创建books变量切片，接收查询的图书信息
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}

	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      ipageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, err
}
