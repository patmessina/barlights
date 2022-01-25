package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

var (
	LightOptions *ws2811.Option
	done         chan bool
	errch        chan error
)

func Start(lightOptions *ws2811.Option, port int) error {

	LightOptions = lightOptions
	done = make(chan bool, 1)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("health", healthCheck)
		v1.POST("solid", setSolid)
		v1.POST("cycle", setCycle)
		v1.POST("off", setOff)
	}

	// TODO: Make sure port is within the correct range
	router.Run(":" + strconv.Itoa(port))

	return nil
}
