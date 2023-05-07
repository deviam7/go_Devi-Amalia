package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"myapp/models"
	"myapp/repositories"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
)

func GetPayment(c echo.Context) error {
	payments, err := repositories.GetPayments()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payments)
}

func MakePayment(c echo.Context) error {
	var newPayment models.NewPayment

	err := c.Bind(&newPayment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	transaction, err := repositories.TransactionGetByID(newPayment.TransactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if newPayment.Pay < transaction.Total {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Jumlah pembayaran kurang",
		})
	}

	// Menghandle pembayaran
	var change int
	if newPayment.Type == "cash" {
		// Menghitung kembalian
		change = newPayment.Pay - transaction.Total
	} else if newPayment.Type == "credit_card" || newPayment.Type == "debit" || newPayment.Type == "qris" {
		if transaction.Total != newPayment.Pay {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Jumlah pembayaran harus sama untuk pembayaran dengan " + newPayment.Type,
			})
		}
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Jenis pembayaran tidak valid",
		})
	}

	_, err = repositories.AddPayment(newPayment, transaction.Total, change)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Pay":  newPayment.Pay,
		"Change": change,
	})
}

func MakeReceipt(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("payment_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid ID",
		})
	}

	payment, err := repositories.GetPaymentByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	transactionMenus, err := repositories.GetTransactionMenuByTransactionID(payment.TransactionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	menus, err := repositories.GetMenus()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	// Create PDF receipt
	h := 10.0
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, h, fmt.Sprintf("No: %d", payment.ID))
	pdf.Ln(12)
	h += 20
	pdf.Cell(40, h, payment.DateTime)
	pdf.Ln(12)
	h += 20
	pdf.Cell(40, h, "Receipt:")
	pdf.Ln(12)
	h += 20

	for _, transactionMenu := range transactionMenus {
		for _, menu := range menus {
			if menu.ID == transactionMenu.MenuID {
				h += 20
				pdf.Cell(40, h, "- "+menu.Name+fmt.Sprintf(" Rp. %d", menu.Price))
				pdf.Ln(10)
				break
			}
		}
	}

	h += 20
	pdf.Ln(10)

	pdf.Cell(40, h, "Total: "+fmt.Sprintf("Rp. %d", payment.Total))

	h += 20
	pdf.Ln(10)

	pdf.Cell(40, h, "Payment: "+fmt.Sprintf("Rp. %d", payment.Pay))

	h += 20
	pdf.Ln(10)

	pdf.Cell(40, h, "Change: "+fmt.Sprintf("Rp. %d", payment.Change))

	// Save PDF to file
	err = pdf.OutputFileAndClose("receipt.pdf")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.File("receipt.pdf")
}
