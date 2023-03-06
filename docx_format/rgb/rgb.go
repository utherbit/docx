package rgb

import (
	"fmt"
	"strconv"
	"strings"
)

type ColorRGB struct {
	R uint8
	G uint8
	B uint8
}

func (c ColorRGB) ToHEX() string {
	return strings.ToUpper(fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B))
}

func HexToRGB(hex string) (ColorRGB, error) {
	var rgb ColorRGB
	values, err := strconv.ParseUint(string(hex), 16, 32)

	if err != nil {
		return ColorRGB{}, err
	}

	rgb.R = uint8(values >> 16)
	rgb.G = uint8((values >> 8) & 0xFF)
	rgb.B = uint8(values & 0xFF)

	return rgb, nil
}
