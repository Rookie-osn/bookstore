package dao

import (
	"bookstore/common"
	"bookstore/model"
	"github.com/gin-gonic/gin"
)

//添加session
func AddSession(session *model.Session) error {
	db:=common.GetDB()
	sqlStr := "insert into sessions values(?,?,?)"

	_, err := db.Exec(sqlStr, session.SessionID, session.UserName, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

//删除session
func DeleteSessionBySessionID(sessionID string) error {
	db:=common.GetDB()
	sqlStr := "delete from sessions where session_id=?"

	_, err := db.Exec(sqlStr, sessionID)
	if err != nil {
		return err
	}
	return nil
}

//通过id从数据库获取session
func GetSessionBySessionID(sessionID string) (*model.Session, error) {
	db:=common.GetDB()
	sqlStr := "select session_id,username,user_id from sessions where session_id=?"

	row := db.QueryRow(sqlStr, sessionID)
	//创建变量接收
	sess := &model.Session{}
	err := row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, err
}

//判断是否登录
func IsLogin(ctx *gin.Context) (*model.Session, bool) {
	cookie, _ := ctx.Request.Cookie("user") //获取cookie
	//判断是否有cookie，没有则未登录
	if cookie != nil {
		//提取session
		cookieValue := cookie.Value
		session, _ := GetSessionBySessionID(cookieValue)
		//判断是否有返回的数据
		if session.UserID > 0 {
			return session, true
		}
	}
	return nil, false
}

