package usecase

import (
	"errors"

	"github.com/alterra-academy/psd-go-course/internal/entity"
	"github.com/alterra-academy/psd-go-course/internal/infra"
)

type UsecaseHandler struct {
	repo   PhoneNumberRepository
	input  InputHandler
	output OutputHandler
}

func NewUsecaseHandler(repo PhoneNumberRepository, input InputHandler, output OutputHandler) *UsecaseHandler {
	return &UsecaseHandler{
		repo:   repo,
		input:  input,
		output: output,
	}
}
