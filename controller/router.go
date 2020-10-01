package controller

import (
	"github.com/andersonlira/item-api/config"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//MapRoutes for the endpoints which the API listens for
func MapRoutes(e *echo.Echo) {
	g := e.Group("/item-api/v1")
	if config.Values.UsePrometheus {
		p := prometheus.NewPrometheus("echo", nil)
		p.Use(e)
	}
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	g.GET("/item", GetItemList)
	g.GET("/item/:id", GetItemByID)
	g.POST("/item", SaveItem)
	g.PUT("/item/:id", UpdateItem)
	g.DELETE("/item/:id", DeleteItem)
	g.GET("/health", CheckHealth)
	g.GET("/info", GetInfo)
	g.PATCH("/item/:id/price/lowest", LowestPrice)
	g.PATCH("/item/:id/price/highest", HighestPrice)
	g.PATCH("/item/:id/price/last", LastPrice)
}
