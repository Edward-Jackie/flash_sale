package start

import (
	"Flash_Sale/driver"
	"Flash_Sale/logger"
	"Flash_Sale/router"
)

func Init() {
	logger.Init("./log/Flash_Sale.log")
	driver.InitDb()
	router.InitRouter()
}
