package driver

import (
	log "Flash_Sale/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb() {
	var err error
	//新增Gorm配置

	//DB, err = gorm.Open("mysql", "root:123456@tcp(106.12.140.51:9777)/cherry_neko?charset=utf8&parseTime=true&loc=Local")
	//DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/neko?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		log.SLog.Error("数据库连接报错, err = ", zap.Error(err))
	}
	////打开SQL打印
	//DB.LogMode(true)
	////启用配置
	//DB.SetLogger(gorm.Logger{logger.GormLog})
	log.SLog.Info("数据库连接成功！！！")
}

func InitDB() {
	var err error
	DB, err = gorm.Open(
		mysql.Open("root:123456@tcp(106.12.140.51:9777)/cherry_neko?charset=utf8&parseTime=true&loc=Local"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		log.SLog.Error("数据连接出错 err = ", zap.Error(err))
		panic(err)
		return
	}

	log.SLog.Info("数据库连接成功")
}
