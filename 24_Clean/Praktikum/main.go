package main

import (
	"fmt"
	"os"

	"github.com/alterra-academy/psd-go-course/internal/infra"
	"github.com/alterra-academy/psd-go-course/internal/input"
	"github.com/alterra-academy/psd-go-course/internal/output"
	"github.com/alterra-academy/psd-go-course/internal/usecase"
)

func main() {
	inputHandler := input.NewInputHandler(os.Stdin, os.Stdout)
	outputHandler := output.NewOutputHandler(os.Stdout)
	repo := infra.NewPhoneNumberRepository()
	usecaseHandler := usecase.NewUsecaseHandler(repo, inputHandler, outputHandler)

	outputHandler.WriteString("Masukkan nomor telepon:")
	phoneNumber, err := inputHandler.GetPhoneNumber()
	if err != nil {
		outputHandler.WriteString(err.Error())
		return
	}

	result, err := usecaseHandler.ValidatePhoneNumber(phoneNumber)
	if err != nil {
		outputHandler.WriteString(err.Error())
		return
	}

	outputHandler.WriteString(fmt.Sprintf("Nomor telepon %s valid\n", result))
}
