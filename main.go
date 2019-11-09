package main

import (
	"github.com/h00s/bmp180-server/sensors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	bmp180 := sensors.NewBMP180(r)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", bmp180.GetValues)

	e.Logger.Fatal(e.Start(":1702"))
}
