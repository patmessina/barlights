package types

type Color struct {
	Hex    string
	RGB    *RGB
	UInt32 uint32
	Name   string
}

type RGB struct {
	Red   uint8
	Blue  uint8
	Green uint8
}
