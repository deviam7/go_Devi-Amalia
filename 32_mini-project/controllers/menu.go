package controllers

import (
	"net/http"
	"strconv"

	"myapp/models"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func GetMenu(c echo.Context) error {
	menus, err := repositories.GetMenus()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, menus)
}

func GetMenuByCategory(c echo.Context) error {
	categories, err := repositories.GetCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	for _, v := range categories {
		menus, err := repositories.GetMenuByCategoryID(v.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		}

		v.Menus = menus
	}

	return c.JSON(http.StatusOK, categories)
}

func AddMenu(c echo.Context) error {
	var newMenu models.NewMenu

	err := c.Bind(&newMenu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = repositories.AddMenu(newMenu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Menu berhasil ditambahkan",
	})
}

func EditMenu(c echo.Context) error {
	var newMenu models.NewMenu

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid ID",
		})
	}

	err = c.Bind(&newMenu)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	_, err = repositories.EditMenu(id, newMenu)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Menu berhasil diubah",
	})
}
