package types

type Flash struct {
	HexColors     []string `json:"colors" binding:"required,hexcolor"`
	Colors        []*Color `json:",omitempty"`
	Speed         int      `json:"speed" binding:"required,gte=10,lte=255"`
	Brightness    int      `json:"brightness" binding:"required,gte=0,lte=255"`
	BrightnessInc int      `json:"brightnessInc" binding:"required,gte=0,lte=30"`
}
