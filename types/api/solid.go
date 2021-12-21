package api

type Solid struct {
	Hex        string `json:"hex" binding:"required,hexcolor"`
	Brightness int    `json:"brightness" binding:"required,gte=0,lte=255"`
}
