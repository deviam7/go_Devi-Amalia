package controllers

import (
	"DEVI-AMALIA/config"
	"DEVI-AMALIA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	var (
		DB    = config.GetDB()
		users []models.User
	)

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil mendapatkan semua pengguna",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id pengguna tidak valid")
	}

	var (
		DB   = config.GetDB()
		user models.User
	)

	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "pengguna tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil mendapatkan pengguna dengan id",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var (
		DB   = config.GetDB()
		user = models.User{}
	)
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil membuat pengguna baru",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id pengguna tidak valid")
	}

	var (
		DB   = config.GetDB()
		user models.User
	)

	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "pengguna tidak ditemukan")
	}

	if err := DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil menghapus pengguna",
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id pengguna tidak valid")
	}

	var (
		DB   = config.GetDB()
		user models.User
	)

	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "pengguna tidak ditemukan")
	}

	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil memperbarui pengguna",
		"user":    user,
	})

}
