package main

import (
	"github.com/gin-gonic/gin"
	"gocasbin/lib"
)

func main() {
	r:=gin.New()
	r.Use(lib.Middlewares()...)
	r.GET("/depts", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":"部门列表"})
	})
	r.GET("/depts/:id", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":"部门详细"})
	})
	r.POST("/depts", func(context *gin.Context) {
		context.JSON(200,gin.H{"result":"批量修改部门列表"})
	})
	r.Run(":8080")
}
