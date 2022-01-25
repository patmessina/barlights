package api

const ()

type Flash struct {
	Colors     []string `json:"colors" binding:"required,hexcolor"`
	Speed      int      `json:"speed" binding:"required,gte=10,lte=255"`
	Brightness int      `json:"brightness" binding:"required,gte=0,lte=255"`
	Fade       bool     `json:"fade" binding:"required,boolean"`
}
