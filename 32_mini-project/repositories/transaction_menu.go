package repositories

import (
	"myapp/config"
	"myapp/models"
)

func GetTransactionMenuByTransactionID(transactionID int) ([]*models.TransactionMenu, error) {
	var (
		db               = config.GetDB()
		transactionMenus []*models.TransactionMenu
	)

	if err := db.Where("transaction_id = ?", transactionID).Find(&transactionMenus).Error; err != nil {
		return nil, err
	}

	return transactionMenus, nil
}

func AddTransactionMenu(transactionID int, menuIDs []int) ([]*models.TransactionMenu, error) {
	var (
		db               = config.GetDB()
		transactionMenus []*models.TransactionMenu
	)

	for _, v := range menuIDs {
		transactionMenus = append(transactionMenus, &models.TransactionMenu{
			TransactionID: transactionID,
			MenuID:        v,
		})
	}

	if err := db.Create(&transactionMenus).Error; err != nil {
		return nil, err
	}

	return transactionMenus, nil
}
