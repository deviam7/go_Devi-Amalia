package repositories

import (
	"myapp/config"
	"myapp/models"
	"time"
)

func GetPayments() ([]*models.Payment, error) {
	var (
		db       = config.GetDB()
		payments []*models.Payment
	)

	if err := db.Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

func AddPayment(newPayment models.NewPayment, total, change int) (*models.Payment, error) {
	var (
		db      = config.GetDB()
		payment = models.Payment{
			TransactionID: newPayment.TransactionID,
			Total:         total,
			Type:          newPayment.Type,
			Pay:           newPayment.Pay,
			Change:        change,
			DateTime:      time.Now().String(),
		}
	)

	if err := db.Create(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func GetPaymentByID(id int) (*models.Payment, error) {
	var (
		db      = config.GetDB()
		payment models.Payment
	)

	if err := db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}
