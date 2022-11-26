package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "主人好，我是喵喵"})
	})

	r.Run(":8087")
}
