package pkg

import (
	"barlights/types"
	"math/rand"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// NewColorFromHex Given a Hex value create Color
func NewColorFromHex(hex string) (*types.Color, error) {

	v, err := HexStrToUInt32(hex)
	if err != nil {
		return nil, err
	}
	c := &types.Color{
		UInt32: v,
		Hex:    strings.ToLower(hex),
		RGB:    Uint32ToRGB(v),
	}

	log.WithFields(log.Fields{
		"hex":       c.Hex,
		"uint32":    c.UInt32,
		"RGB.Red":   c.RGB.Red,
		"RGB.Green": c.RGB.Green,
		"RGB.Blue":  c.RGB.Blue,
	}).Debug("Created color from Hex")

	return c, nil
}

// NewColorFromHex gets a new Color type given a hex value
func NewColorFromUInt32(v uint32) (*types.Color, error) {
	c := &types.Color{
		UInt32: v,
		Hex:    strconv.FormatInt(int64(v), 16),
	}

	c.RGB = Uint32ToRGB(v)

	log.WithFields(log.Fields{
		"hex":       c.Hex,
		"uint32":    c.UInt32,
		"RGB.Red":   c.RGB.Red,
		"RGB.Green": c.RGB.Green,
		"RGB.Blue":  c.RGB.Blue,
	}).Debug("Created color from uint32")

	return c, nil
}

// NewColorFromRGB gets a new Color type given RGB values
func NewColorFromRGB(r uint8, g uint8, b uint8) (*types.Color, error) {
	// ensure rgb values are between the correct values [0,255]

	rgb := &types.RGB{
		Red:   r,
		Green: g,
		Blue:  b,
	}

	// convert to int
	v := int(r)
	v = (v << 8) + int(g)
	v = (v << 8) + int(b)

	c := &types.Color{
		UInt32: uint32(v),
		RGB:    rgb,
		Hex:    strconv.FormatInt(int64(v), 16),
	}

	log.WithFields(log.Fields{
		"hex":       c.Hex,
		"uint32":    c.UInt32,
		"RGB.Red":   c.RGB.Red,
		"RGB.Green": c.RGB.Green,
		"RGB.Blue":  c.RGB.Blue,
	}).Debug("Created color from RGB")

	return c, nil
}

// Uint32ToRGB convert int to rgb
func Uint32ToRGB(v uint32) *types.RGB {

	// Get RGB Values
	rgb := &types.RGB{
		Red:   uint8(v >> 16),
		Green: uint8(v >> 8),
		Blue:  uint8(v),
	}

	return rgb

}

// HexStrToUInt32 converts 'hex' strings to uint32
func HexStrToUInt32(strHex string) (uint32, error) {

	var c uint32

	strHex = strings.Replace(strHex, "0x", "", -1)
	strHex = strings.Replace(strHex, "0X", "", -1)
	strHex = strings.Replace(strHex, "#", "", -1)
	strHex = strings.ToLower(strHex)

	_c, err := strconv.ParseInt(strHex, 16, 32)
	if err != nil {
		return 0, err
	}
	c = uint32(_c)

	return c, nil
}

// RandColor given a list grab a random color from it
func RandColor(colors []*types.Color) *types.Color {
	i := rand.Intn(len(colors))
	return colors[i]
}
