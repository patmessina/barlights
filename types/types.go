package types

type BarlightSettings struct {
	HexColors     []string `json:"colors,omitempty" binding:"dive,hexcolor"`
	Colors        []*Color `json:",omitempty"`
	Speed         int      `json:"speed,omitempty" binding:"omitempty,gte=0,lte=255"`
	Brightness    int      `json:"brightness" binding:"required,gte=0,lte=255"`
	BrightnessInc []*Color `json:",omitempty"`
}
