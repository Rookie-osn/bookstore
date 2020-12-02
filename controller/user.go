package controller

import (
	"bookstore/common"
	"bookstore/dao"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

// 注册
func Regist(ctx *gin.Context){
	// 获取参数
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")
	email:=ctx.PostForm("email")

	// 从数据库查看该用户是否已存在
	if ok:=dao.IfUserExist(username);!ok{
		// 密码加密
		hasedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
		if err!=nil{
			ctx.HTML(http.StatusInternalServerError,"regist.html","注册失败，请重试！")
			return
		}
		//用户不存在，注册用户并注册成功
		user:=model.User{
			Username: username,
			Password: string(hasedPassword),
			Email: email,
		}
		dao.CreateUser(user)
		ctx.HTML(http.StatusOK,"regist_success.html","")
	}else{
		// 用户存在，返回
		ctx.HTML(http.StatusUnprocessableEntity,"regist.html","用户名已存在！")
	}
}


// 登录
func Login(ctx *gin.Context){
	// 防止重复登录，判断是否已登录
	_,flag:=dao.IsLogin(ctx)
	if flag{
		// 直接传回图书页面
		GetPageBooksByPrice(ctx)
		return
	}

	// 获取参数
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")

	// 验证数据
	if username==""||password==""{
		ctx.HTML(http.StatusUnprocessableEntity,"login.html",gin.H{"data":"请填写用户名或密码！"})
		return
	}
	// 从数据库请求数据
	user,err:=dao.GetUserByUsernameAndPassword(username,password)
	if err!=nil{
		ctx.HTML(http.StatusUnprocessableEntity,"login.html",gin.H{"data":"用户名或密码错误！"})
		return
	}

	if user!=nil{
		// 创建uuid
		uuid:=common.CreateUUID()
		// 创建session
		ss:=&model.Session{
			SessionID: uuid,
			UserName: user.Username,
			UserID: user.ID,
		}
		// 添加入库
		dao.AddSession(ss)
		// 发送cookie给客户端
		ctx.SetCookie("user",uuid,3600,"/","localhost",false,true)

		ctx.HTML(http.StatusOK,"login_success.html",gin.H{"data":*user})
	}else{
		// 用户不存在
		ctx.HTML(http.StatusUnprocessableEntity,"login.html","用户名或密码不正确！")
	}
}

// 注销
func Logout(ctx *gin.Context){
	// 获取cookie
	cookie,_:=ctx.Request.Cookie("user")
	// 通过cookie的sessionid删除session
	if cookie!=nil{
		cookieValue:=cookie.Value
		// 删除session
		dao.DeleteSessionBySessionID(cookieValue)
		// 将cookie清空
		cookie.MaxAge=-1
		ctx.SetCookie("user","",-1,"/","localhost",false,true)
	}
	// 回到首页
	GetPageBooksByPrice(ctx)
}


// 检查用户是否已存在，返回给Ajax请求
func IfUserExist(ctx *gin.Context){
	username:=ctx.PostForm("username")

	if ok:=dao.IfUserExist(username);!ok{
		ctx.String(http.StatusUnprocessableEntity,"<font style='color:green'>用户名可用！</font>")
	}
}