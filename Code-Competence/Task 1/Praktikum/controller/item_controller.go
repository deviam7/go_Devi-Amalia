package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"model"
)

type ItemController struct{}

func (ic *ItemController) AddItem(c echo.Context) error {
	var newItem model.Item
	err := json.NewDecoder(c.Request().Body).Decode(&newItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newItem.ID = len(model.Items) + 1
	model.Items = append(model.Items, newItem)

	return c.JSON(http.StatusCreated, newItem)
}

func (ic *ItemController) UpdateItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var updatedItem model.Item
	err = json.NewDecoder(c.Request().Body).Decode(&updatedItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for i, item := range model.Items {
		if item.ID == id {
			updatedItem.ID = id
			model.Items[i] = updatedItem

			return c.JSON(http.StatusOK, updatedItem)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
}

func (ic *ItemController) DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for i, item := range model.Items {
		if item.ID == id {
			model.Items = append(model.Items[:i], model.Items[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
}

func (ic *ItemController) GetAllItems(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Items)
}

func (ic *ItemController) GetItemByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for _, item := range model.Items {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Item not found"})
}

func (ic *ItemController) GetItemsByCategory(c echo.Context) error {
	category := c.Param("category_id")

	var result []model.Item
	for _, item := range model.Items {
		if item.Category == category {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No items found in this category"})
	}

	return c.JSON(http.StatusOK, result)
}

	func (ic *ItemController) GetItemsByCategory(c echo.Context) error {
	category := c.Param("category_id")

	var result []model.Item
	for _, item := range model.Items {
		if item.Category == category {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No items found in this category"})
	}

	return c.JSON(http.StatusOK, result)
}

	func (ic *ItemController) SearchItems(c echo.Context) error {
		query := c.QueryParam("keyword")
	
		var result []model.Item
		for _, item := range model.Items {
			if item.Name == query {
				result = append(result, item)
			}
		}
	
		if len(result) == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "No items found with this name"})
		}
	
		return c.JSON(http.StatusOK, result)
	}
	