package driver

import (
	"Flash_Sale/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var DB *gorm.DB

func InitDb() {
	var err error
	//新增Gorm配置

	DB, err = gorm.Open("mysql", "root:123456@tcp(106.12.140.51:9777)/cherry_neko?charset=utf8&parseTime=true&loc=Local")
	//DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/neko?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		logger.SLog.Error("数据库连接报错, err = ", zap.Error(err))
	}
	//打开SQL打印
	DB.LogMode(true)
	//启用配置
	DB.SetLogger(gorm.Logger{logger.GormLog})
	logger.SLog.Info("数据库连接成功！！！")
}
