package daos

import (
	"Flash_Sale/driver"
	"Flash_Sale/logger"
	"Flash_Sale/models"
	"go.uber.org/zap"
)

func InsertUser(u *models.User) bool {
	err := driver.DB.Create(u).Error
	if err != nil {
		logger.SLog.Error("err = ", zap.Error(err))
		return false
	}
	return true
}

func UpdateUser(u *models.User) {

}

func SelectUserById(id int64) (*models.User, error) {
	u := new(models.User)
	err := driver.DB.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func SelectUserByName(n string) (*models.User, error) {
	u := new(models.User)
	err := driver.DB.Where("name = ?", n).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
