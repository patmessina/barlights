package pkg

var (
	CycleColors = map[string][]uint32{
		// comes from the following pallet
		// https://lospec.com/palette-list/mr-colors
		"default": {
			uint32(0x027baa),
			uint32(0x149098),
			uint32(0x5ac68e),
			uint32(0xc0180a),
			uint32(0xdb3a05),
			uint32(0xec5903),
			uint32(0xf98d00),
			uint32(0x000882),
			uint32(0x007ba3),
			uint32(0x0dd4b4),
			uint32(0xf0658a),
			uint32(0xf95cde),
			uint32(0xec407a),
		},
		"fire": {
			uint32(0xff0000),
			uint32(0xffff00),
		},
		"police": {
			uint32(0x0006b1),
			uint32(0xffff00),
		},
	}

	DefaultSolidColors = map[string]uint32{
		"off":    uint32(0x000000),
		"blue":   uint32(0x21abcD),
		"yellow": uint32(0xffe135),
		"green":  uint32(0x7fff00),
		"red":    uint32(0xc41e3a),
		"white":  uint32(0xf0f8ff),
		"purple": uint32(0x8b008b),
		"pink":   uint32(0xde3163),
		"orange": uint32(0xff4f00),
	}
)
