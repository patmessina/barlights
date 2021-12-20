package pkg

import (
	"barlights/types"
	"fmt"
	"testing"
)

func TestHextStrToUInt32(t *testing.T) {

	tcs := []struct {
		Input    string
		Expected uint32
	}{
		{"0xFFFFFF",
			16777215},
		{"0XFFFFFF",
			16777215},
		{"#FFFFFF",
			16777215},
		{"0x000000",
			0},
	}

	for _, c := range tcs {
		v, err := HexStrToUInt32(c.Input)
		if err != nil {
			t.Logf("Input: %v, Expected: %v, Got: %v, Error: %v\n",
				c.Input, c.Expected, v, err)
			t.Fail()
		}
		if v != c.Expected {
			t.Logf("Input: %v, Expected: %v, Got: %v\n",
				c.Input, c.Expected, v)
			t.Fail()
		}
	}

}

func TestUint32ToRGB(t *testing.T) {

	tcs := []struct {
		Input    uint32
		Expected types.RGB
	}{
		{16777215,
			types.RGB{255, 255, 255}},
		{6750207,
			types.RGB{102, 255, 255}},
		{0,
			types.RGB{0, 0, 0}},
	}

	for _, c := range tcs {
		v := Uint32ToRGB(c.Input)
		if v.Blue != c.Expected.Blue ||
			v.Green != c.Expected.Green ||
			v.Red != c.Expected.Red {
			exp := fmt.Sprintf("(r %v, g: %v, b: %v)",
				c.Expected.Red, c.Expected.Green, c.Expected.Blue)
			got := fmt.Sprintf("(r %v, g: %v, b: %v)",
				v.Red, v.Green, v.Blue)
			t.Logf("Input: %v, Expected: %v, Got: %v\n",
				c.Input, exp, got)
			t.Fail()
		}
	}

}
