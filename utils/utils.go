package utils

import (
	"strconv"
	"strings"
)

func HexStrToInt(strHex string) (uint32, error) {

	var c uint32

	strHex = strings.Replace(strHex, "0x", "", -1)
	strHex = strings.Replace(strHex, "0X", "", -1)
	strHex = strings.Replace(strHex, "#", "", -1)

	_c, err := strconv.ParseInt(strHex, 16, 32)
	if err != nil {
		return 0, err
	}
	c = uint32(_c)

	return c, nil
}
