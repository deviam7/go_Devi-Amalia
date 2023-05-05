package controllers

import (
	"net/http"

	"myapp/models"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func GetTransactions(c echo.Context) error {
	transactions, err := repositories.GetTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	for _, transaction := range transactions {
		TransactionMenus, err := repositories.GetTransactionMenuByTransactionID(transaction.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		}

		for _, transactionMenu := range TransactionMenus {
			menu, err := repositories.GetMenuByID(transactionMenu.MenuID)
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"error": err.Error(),
				})
			}

			transaction.Menus = append(transaction.Menus, menu)
		}
	}

	return c.JSON(http.StatusOK, transactions)
}

func AddTransaction(c echo.Context) error {
	var newTransaction models.NewTransaction

	err := c.Bind(&newTransaction)
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

	// Calculate total amount
	var total int
	for _, item := range newTransaction.Items {
		for _, menu := range menus {
			if item == menu.ID {
				total += menu.Price
				break
			}
		}
	}

	transaction, err := repositories.AddTransaction(total)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = repositories.AddTransactionMenu(transaction.ID, newTransaction.Items)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}
