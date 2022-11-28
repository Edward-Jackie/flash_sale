package handler

import (
	"Flash_Sale/logger"
	daos "Flash_Sale/server/v1/daos/Mq"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"net/http"
)

func SendOrder(c *gin.Context) {
	qn := c.Param("queue")

	ch, q, err := daos.OrderQueue(qn)
	if err != nil {
		logger.SLog.Error(fmt.Sprintf("获取队列报错 %v 报错 , err = ", qn), zap.Error(err))
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "获取队列失败", "data": nil})
		return
	}

	body := "CherryNeko Test"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		logger.SLog.Error(fmt.Sprintf("发送队列报错 %v 报错 , err = ", qn), zap.Error(err))
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "发送失败", "data": nil})
		return
	}
	c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "发送成功", "data": nil})
}

func GetOrder(c *gin.Context) {
	qn := c.Param("queue")

	ch, q, err := daos.OrderQueue(qn)
	if err != nil {
		logger.SLog.Error(fmt.Sprintf("获取队列报错 %v 报错 , err = ", qn), zap.Error(err))
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "获取队列失败", "data": nil})
		return
	}

	msg, err := ch.Consume(
		q.Name,
		"CherryNeko Test Consumer",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.SLog.Error(fmt.Sprintf("获取队列信息 %v 报错 , err = ", qn), zap.Error(err))
		c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "获取失败", "data": nil})
		return
	}
	//tags := make([]string, 0)
	//bodies:= make([]string, 0)
	go func() {

		for info := range msg {
			logger.SLog.Info(fmt.Sprintf("获取到订单 %s Msg: %v", info.Body, info.ConsumerTag))
		}
	}()

	//c.JSONP(http.StatusOK, gin.H{"code" : 200 , "msg" : "获取成功", "data": gin.H{"body" : bodies, "tags": tags}})
	c.JSONP(http.StatusOK, gin.H{"code": 200, "msg": "获取成功", "data": gin.H{"body": "bodies", "tags": "tags"}})
}
