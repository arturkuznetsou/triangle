package main
import (
	"image/color"
	"strings"
	"strconv"
)

func ParseHex(hex string) color.RGBA {
	hex = strings.ToLower(hex)
	var r, g, b int64
	r, g, b = 255, 255, 255

	switch len(hex) {
	case 6:
		r, _ = strconv.ParseInt(string(hex[0:2]), 16, 64)
		g, _ = strconv.ParseInt(string(hex[2:4]), 16, 64)
		b, _ = strconv.ParseInt(string(hex[4:6]), 16, 64)
		print(r," ", g," ", b, "\n")
	case 3:
		r, _ = strconv.ParseInt(string(hex[0]), 16, 64)
		g, _ = strconv.ParseInt(string(hex[1]), 16, 64)
		b, _ = strconv.ParseInt(string(hex[2]), 16, 64)
		r *= 16
		g *= 16
		b *= 16
	default:
		println("Invalid color. Using black.")
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
