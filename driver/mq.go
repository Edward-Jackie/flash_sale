package driver

import (
	"Flash_Sale/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var MQ *amqp.Connection

func InitMQ() {
	var err error
	MQ, err = amqp.Dial("amqp://guest:guest@106.12.140.51:5672/")
	if err != nil {
		logger.SLog.Error("RabbitMQ 连接失败！！！ err = ", zap.Error(err))
		defer MQ.Close()
	}

	logger.SLog.Info("RabbitMQ 连接成功！！！")
}
