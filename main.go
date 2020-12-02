package main

import (
	"bookstore/common"
	"github.com/gin-gonic/gin"
)

func main(){
	common.InitDB()
	r:=gin.Default()
	r=CollectRouter(r)
	r.Run()
}