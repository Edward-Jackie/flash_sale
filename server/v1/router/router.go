package router

import (
	"Flash_Sale/server/v1/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "请求成功", "data": gin.H{"msg": "主人，我是喵喵"}})
	})

	q := r.Group("/order")
	{
		q.GET("/send/:queue", handler.SendOrder)
		q.GET("/get/:queue", handler.GetOrder)
	}

	r.Run(":8087")
}
