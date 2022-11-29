package router

import (
	"Flash_Sale/logger"
	"Flash_Sale/server/v1/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	r := gin.Default()
	//重置颜色
	gin.DisableConsoleColor()
	//启用日志文件打印
	config := gin.LoggerConfig{
		Output: logger.GinLog,
	}
	//启用配置
	withConfig := gin.LoggerWithConfig(config)
	r.Use(withConfig)

	r.GET("/hello", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "请求成功", "data": gin.H{"msg": "主人，我是喵喵"}})
	})

	q := r.Group("/order")
	{
		q.GET("/send/:queue", handler.SendOrder)
		q.GET("/get/:queue", handler.GetOrder)
	}

	u := r.Group("/user")
	{
		u.POST("/add", handler.RegisterUser)
		u.GET("/:name/select", handler.FindUserByName)
	}
	r.Run(":8087")
}
