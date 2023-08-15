package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllData(c echo.Context) error {
	result, err := models.FetchAllData()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StoreMahasiswa(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telephone := c.FormValue("telephone")

	result, err := models.StoreMahasiswa(nama, alamat, telephone)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateMahasiswa(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telephone := c.FormValue("telephone")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	result, err := models.UpdateMahasiswa(conv_id, nama, alamat, telephone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteMahasiswa(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	result, err := models.DeleteMahasiswa(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, result)
}
