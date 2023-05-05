package repositories

import (
	"myapp/config"
	"myapp/models"
)

func GetCategories() ([]*models.Category, error) {
	var (
		db         = config.GetDB()
		categories []*models.Category
	)

	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
