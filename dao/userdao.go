package dao

import (
	"bookstore/common"
	"bookstore/model"
	"golang.org/x/crypto/bcrypt"
)

// CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func GetUserByUsernameAndPassword(username,password string)(*model.User,error){
	db:=common.GetDB()
	sqlStr:="select id,username,password,email from users where username=?"

	row:=db.QueryRow(sqlStr,username,password)
	user:=model.User{}
	// 创建一个user变量
	err:=row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	if err!=nil{
		return nil,err
	}
	// 判断返回的Username是否正确
	if user.Username!=username{
		return nil,nil
	}
	// 判断密码是否正确
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		return nil,err
	}
	return &user,nil
}

// IfUserExist 根据用户名查看用户存在
func IfUserExist(username string)bool{
	DB:=common.GetDB()
	sqlStr:="select id,username,password,email from users where username=?"
	row:=DB.QueryRow(sqlStr,username)

	user:=model.User{}
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)

	// 判断返回的Username是否正确
	if user.Username==username{
		return true
	}

	return false
}

// CreatUser 创建用户信息插入数据库
func CreateUser(user model.User)error{
	db:=common.GetDB()
	sqlStr:="Insert into users(username,password,email) values(?,?,?)"

	stmt,err:=db.Prepare(sqlStr)
	if err!=nil{
		return err
	}

	_,err=stmt.Exec(user.Username,user.Password,user.Email)
	if err!=nil{
		return err
	}
	return err
}