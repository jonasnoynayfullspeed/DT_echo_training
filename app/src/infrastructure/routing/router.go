package routing

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func Init() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// set template
	SetTemplate(e)

	// set routing
	SetRouting(e)

	//Error page
	e.HTTPErrorHandler = errorPageHandler

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}

func errorPageHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	renderErr := c.Render(code, "error.html", he)
	//_ = c.JSONPretty(http.StatusOK, he.Message, "  ")
	//log.Print(code)
	if err != nil {
		log.Print(renderErr)
	}
	return
}
