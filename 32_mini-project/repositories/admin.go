package repositories

import (
	"fmt"
	"myapp/config"
	"myapp/models"

	"gorm.io/gorm"
)

func GetAdmin(username, password string) (*models.Admin, error) {
	var (
		db    = config.GetDB()
		admin models.Admin
	)

	if err := db.Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}

		return nil, err
	}

	if admin.Password != password {
		return nil, fmt.Errorf("password not match")
	}

	return &admin, nil
}
