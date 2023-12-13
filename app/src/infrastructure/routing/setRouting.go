package routing

import (
	"app/src/infrastructure/sqlhandler"
	"app/src/interfaces/controllers"

	"github.com/labstack/echo/v4"
)

// このファイルにはリクエストのルーティング処理を実装する

func SetRouting(e *echo.Echo) {

	controller := controllers.NewController(sqlhandler.NewSqlHandler())

	e.GET("/", controller.Index)
	e.GET("/blog/list", controller.ShowBlogListPage)
	e.GET("/blog/show/:id", controller.ShowBlogDetailsPage)
	e.GET("/blog/edit/:id", controller.EditBlogDetailsPage)
	e.POST("/blog/edit/:id", controller.EditBlogPost)
	e.DELETE("/blog/delete/:id", controller.DeleteBlogPost)
	e.GET("/blog/create", controller.ShowBlogCreatePage)
	e.POST("/blog/add", controller.AddBlogPost)

	e.Static("assets", "./static")
}
