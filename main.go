package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	bmp180 := i2c.NewBMP180Driver(r, i2c.WithBus(1), i2c.WithAddress(0x77))
	bmp180.Start()
	for {
		fmt.Println(bmp180.Temperature())
		time.Sleep(2 * time.Second)
	}
}
