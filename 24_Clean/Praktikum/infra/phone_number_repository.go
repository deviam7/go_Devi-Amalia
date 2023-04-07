package infra

import (
	"errors"

	"github.com/alterra-academy/psd-go-course/internal/entity"
)

type PhoneNumberRepository struct {
}

func NewPhoneNumberRepository() *PhoneNumberRepository {
	return &PhoneNumberRepository{}
}

func (r *PhoneNumberRepository) SavePhoneNumber(phoneNumber *entity.PhoneNumber) error {
	// menyimpan nomor telepon ke database
	return errors.New("gagal menyimpan nomor telepon ke database")
}
