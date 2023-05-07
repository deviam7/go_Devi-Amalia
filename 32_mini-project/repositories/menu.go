package repositories

import (
	"fmt"
	"myapp/config"
	"myapp/models"

	"gorm.io/gorm"
)

func GetMenus() ([]*models.Menu, error) {
	var (
		db    = config.GetDB()
		menus []*models.Menu
	)

	if err := db.Find(&menus).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("menu not found")
		}

		return nil, err
	}

	return menus, nil
}

func AddMenu(newMenu models.NewMenu) (*models.Menu, error) {
	var (
		db   = config.GetDB()
		menu = models.Menu{
			Name:        newMenu.Name,
			CategoryID:  newMenu.CategoryID,
			Description: newMenu.Description,
			Price:       newMenu.Price,
		}
	)

	if err := db.Create(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

func EditMenu(id int, newMenu models.NewMenu) (*models.Menu, error) {
	var (
		db   = config.GetDB()
		menu = models.Menu{
			ID:          id,
			Name:        newMenu.Name,
			CategoryID:  newMenu.CategoryID,
			Description: newMenu.Description,
			Price:       newMenu.Price,
		}
	)

	if err := db.Updates(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

func GetMenuByCategoryID(categoryID int) ([]*models.Menu, error) {
	var (
		db    = config.GetDB()
		menus []*models.Menu
	)

	if err := db.Where("category_id = ?", categoryID).Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}


func GetMenuByID(id int) (*models.Menu, error) {
	var (
		db   = config.GetDB()
		menu models.Menu
	)

	if err := db.Where("id = ?", id).First(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}
