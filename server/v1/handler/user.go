package handler

import (
	"Flash_Sale/server/v1/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	type user struct {
		Name string `json:"name" form:"name"`
		Age  int    `json:"age" form:"age"`
	}

	u := new(user)
	err := c.Bind(&u)
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "参数错误"})
		return
	}

	if !service.AddNewUser(u.Name, u.Age) {
		c.JSONP(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "注册失败"})
		return
	}
	c.JSONP(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "注册成功"})
}

func FindUserByName(c *gin.Context) {
	n := c.Param("name")
	u, err := service.FindUserByName(n)
	if err != nil {
		c.JSONP(http.StatusOK, gin.H{"code": 200, "data": nil, "msg": "查询失败"})
		return
	}

	c.JSONP(http.StatusOK, gin.H{"code": 200, "data": u, "msg": "查询成功"})
}
