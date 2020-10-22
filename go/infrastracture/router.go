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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	prefectureController := controller.NewPrefectureController(NewSqlHandler())
	cityController := controller.NewCityController(NewSqlHandler())
	userController := controller.NewUserController(NewSqlHandler())
	restaurantController := controller.NewRestaurantController(NewSqlHandler())
	commentController := controller.NewCommentController(NewSqlHandler())
	goodController := controller.NewGoodController(NewSqlHandler())
	favoriteController := controller.NewFavoriteController(NewSqlHandler())

	e.GET("/", hello)

	e.GET("/prefectures", func(c echo.Context) error { return prefectureController.Index(c) })
	e.GET("/prefectures/:id", func(c echo.Context) error { return prefectureController.Show(c) })
	e.POST("/prefectures", func(c echo.Context) error { return prefectureController.Create(c) })
	e.PUT("/prefectures/:id", func(c echo.Context) error { return prefectureController.Save(c) })
	e.DELETE("/prefectures/:id", func(c echo.Context) error { return prefectureController.Delete(c) })

	e.GET("/cities", func(c echo.Context) error { return cityController.Index(c) })
	e.GET("/cities/:id", func(c echo.Context) error { return cityController.Show(c) })
	e.POST("/cities", func(c echo.Context) error { return cityController.Create(c) })
	e.PUT("/cities/:id", func(c echo.Context) error { return cityController.Save(c) })
	e.DELETE("/cities/:id", func(c echo.Context) error { return cityController.Delete(c) })

	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })

	e.GET("/restaurants", func(c echo.Context) error { return restaurantController.Index(c) })
	e.GET("/restaurants/:id", func(c echo.Context) error { return restaurantController.Show(c) })
	e.POST("/restaurants", func(c echo.Context) error { return restaurantController.Create(c) })
	e.PUT("/restaurants/:id", func(c echo.Context) error { return restaurantController.Save(c) })
	e.DELETE("/restaurants/:id", func(c echo.Context) error { return restaurantController.Delete(c) })

	e.GET("/comments", func(c echo.Context) error { return commentController.Index(c) })
	e.GET("/comments/:id", func(c echo.Context) error { return commentController.Show(c) })
	e.POST("/comments", func(c echo.Context) error { return commentController.Create(c) })
	e.PUT("/comments/:id", func(c echo.Context) error { return commentController.Save(c) })
	e.DELETE("/comments/:id", func(c echo.Context) error { return commentController.Delete(c) })

	e.GET("/goods", func(c echo.Context) error { return goodController.Index(c) })
	e.GET("/goods/:id", func(c echo.Context) error { return goodController.Show(c) })
	e.POST("/goods", func(c echo.Context) error { return goodController.Create(c) })
	e.PUT("/goods/:id", func(c echo.Context) error { return goodController.Save(c) })
	e.DELETE("/goods/:id", func(c echo.Context) error { return goodController.Delete(c) })

	e.GET("/favorites", func(c echo.Context) error { return favoriteController.Index(c) })
	e.GET("/favorites/:id", func(c echo.Context) error { return favoriteController.Show(c) })
	e.POST("/favorites", func(c echo.Context) error { return favoriteController.Create(c) })
	e.PUT("/favorites/:id", func(c echo.Context) error { return favoriteController.Save(c) })
	e.DELETE("/favorites/:id", func(c echo.Context) error { return favoriteController.Delete(c) })

	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
