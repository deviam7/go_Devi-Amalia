package repositories

import (
	"myapp/config"
	"myapp/models"
	"time"
)

func GetTransactions() ([]*models.Transaction, error) {
	var (
		db           = config.GetDB()
		transactions []*models.Transaction
	)

	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func AddTransaction(total int) (*models.Transaction, error) {
	var (
		db          = config.GetDB()
		transaction = models.Transaction{
			Total:    total,
			DateTime: time.Now().Format("2006-01-02 15:04:05"),
		}
	)

	if err := db.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func TransactionGetByID(id int) (*models.Transaction, error) {
	var (
		db          = config.GetDB()
		transaction models.Transaction
	)

	if err := db.Where("id = ?", id).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
