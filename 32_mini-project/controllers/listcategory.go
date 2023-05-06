package controllers

import (
    "github.com/labstack/echo/v4"
    "myapp/repositories"
)

func GetListCategory(c echo.Context) error {
    repo := repositories.NewListCategoryRepository("dbConnection")
    categories, err := repo.GetAll()
    if err != nil {
        return c.JSON(500, map[string]string{
            "message": "Internal Server Error",
        })
    }
    return c.JSON(200, categories)
}
