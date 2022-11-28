package start

import (
	"Flash_Sale/driver"
	"Flash_Sale/logger"
	"Flash_Sale/server/v1/router"
)

func Init() {
	logger.Init("./log/Flash_Sale.log")
	Banner()
	driver.InitDb()
	driver.InitMQ()
	router.InitRouter()
}
