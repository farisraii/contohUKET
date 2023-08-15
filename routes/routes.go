package routes

import (
	"myapp/controllers"
	"myapp/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Faris!")
	})

	e.GET("/mahasiswa", controllers.FetchAllData, middleware.IsAuthenticated)
	e.POST("/mahasiswa", controllers.StoreMahasiswa)
	e.PUT("/mahasiswa", controllers.UpdateMahasiswa)
	e.DELETE("/mahasiswa", controllers.DeleteMahasiswa)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("test-struct", controllers.TestStructValidation)

	return e

}
