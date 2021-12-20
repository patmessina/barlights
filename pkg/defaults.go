package pkg

var (
	OFF = uint32(0x000000)

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

	DefaultColors = map[string]uint32{
		"red":        uint32(0xff0000),
		"orange":     uint32(0xff8000),
		"yellow":     uint32(0xffff00),
		"lime":       uint32(0x80ff00),
		"green":      uint32(0x00ff00),
		"greenblue":  uint32(0x00ff80),
		"lightblue":  uint32(0x00ffff),
		"blue":       uint32(0x0080ff),
		"darkblue":   uint32(0x0000ff),
		"purple":     uint32(0x7f00ff),
		"purplepink": uint32(0xff00ff),
		"pink":       uint32(0xff007f),
		"white":      uint32(0xffffff),
	}
)
