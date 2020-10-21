package infrastracture

import (
	"net/http"

	"github.com/Khmer495/social-creators-teamG-server/go/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// IniitApiServer is
func InitApiServer() {
	e := echo.New()

	userController := controller.NewUserController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", hello)
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
