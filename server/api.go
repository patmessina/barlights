package server

import (
	"barlights/pkg"
	"barlights/types/api"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"health": "ok"})
}

func setSolid(c *gin.Context) {
	var body api.Solid
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	color, err := pkg.NewColorFromHex(body.Hex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.WithFields(
		log.Fields{
			"function":   "setSolid",
			"color":      color.Hex,
			"brightness": LightOptions.Channels[0].Brightness,
		},
	).Debug("setting lights")

	LightOptions.Channels[0].Brightness = body.Brightness
	err = pkg.Solid(LightOptions, color)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func setOff(c *gin.Context) {
	color, err := pkg.NewColorFromHex("000000")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.WithFields(
		log.Fields{
			"function":   "setSolid",
			"color":      color.Hex,
			"brightness": LightOptions.Channels[0].Brightness,
		},
	).Debug("turning off lights")

	LightOptions.Channels[0].Brightness = 0
	err = pkg.Solid(LightOptions, color)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
