package sensors

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

// BMP180 contains struct for BMP180 sensor
type BMP180 struct {
	driver      *i2c.BMP180Driver
	Temperature float64 `json:"temperature"`
	Pressure    float64 `json:"pressure"`
}

// NewBMP180 creates struct connected to Adaptor
func NewBMP180(a *raspi.Adaptor) *BMP180 {
	b := &BMP180{driver: i2c.NewBMP180Driver(a, i2c.WithBus(1), i2c.WithAddress(0x77))}
	b.driver.Start()
	go b.start()
	return b
}

// GetValues returns values in JSON format
func (b *BMP180) GetValues(c echo.Context) error {
	return c.JSON(http.StatusOK, b)
}

func (b *BMP180) start() {
	for {
		temperature, err := b.driver.Temperature()
		if err == nil {
			b.Temperature = float64(temperature)
		}

		pressure, err := b.driver.Pressure()
		if err == nil {
			b.Pressure = float64(pressure)
		}
		time.Sleep(1 * time.Second)
	}
}
