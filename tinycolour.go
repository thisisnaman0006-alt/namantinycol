package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TinyColor Struct
type TinyColor struct {
	r, g, b, a    float64
	roundA        float64
	format        string
	ok            bool
	originalInput string
	gradientType  bool
}

// Color structures for conversions
type RGBA struct {
	R, G, B, A float64
}

type HSL struct {
	H, S, L, A float64
}

type HSV struct {
	H, S, V, A float64
}

// Named CSS Colors
var names = map[string]string{
	"aliceblue": "f0f8ff", "antiquewhite": "faebd7", "aqua": "0ff", "aquamarine": "7fffd4",
	"azure": "f0ffff", "beige": "f5f5dc", "bisque": "ffe4c4", "black": "000",
	"blue": "00f", "blueviolet": "8a2be2", "brown": "a52a2a", "burlywood": "deb887",
	"cyan": "0ff", "darkblue": "00008b", "darkcyan": "008b8b", "darkgreen": "006400",
	"gray": "808080", "green": "008000", "red": "f00", "white": "fff", "yellow": "ff0",
	// Baki saare named colors aap zaroorat ke hisab se map me add kar sakte hain
}

var hexNames = flipMap(names)

func flipMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

// Helper utility functions
func boundAlpha(a float64) float64 {
	if math.IsNaN(a) || a < 0 || a > 1 {
		return 1.0
	}
	return a
}

func bound01(n float64, max float64) float64 {
	n = math.Min(max, math.Max(0, n))
	if math.Abs(n-max) < 0.000001 {
		return 1.0
	}
	return math.Mod(n, max) / max
}

func clamp01(val float64) float64 {
	return math.Min(1, math.Max(0, val))
}

func pad2(c string) string {
	if len(c) == 1 {
		return "0" + c
	}
	return c
}

// Main Constructor
func NewTinyColor(input string) TinyColor {
	tc := TinyColor{
		originalInput: input,
	}

	rgb, ok, format := inputToRGB(input)
	tc.r = rgb.R
	tc.g = rgb.G
	tc.b = rgb.B
	tc.a = rgb.A
	tc.roundA = math.Round(100*tc.a) / 100
	tc.format = format
	tc.ok = ok

	if tc.r < 1 {
		tc.r = math.Round(tc.r)
	}
	if tc.g < 1 {
		tc.g = math.Round(tc.g)
	}
	if tc.b < 1 {
		tc.b = math.Round(tc.b)
	}

	return tc
}

// String parsing logic
func inputToRGB(colorStr string) (RGBA, bool, string) {
	colorStr = strings.TrimSpace(strings.ToLower(colorStr))

	if val, ok := names[colorStr]; ok {
		colorStr = val
	} else if colorStr == "transparent" {
		return RGBA{0, 0, 0, 0}, true, "name"
	}

	hex6 := regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`)
	if match := hex6.FindStringSubmatch(colorStr); match != nil {
		r, _ := strconv.ParseInt(match[1], 16, 64)
		g, _ := strconv.ParseInt(match[2], 16, 64)
		b, _ := strconv.ParseInt(match[3], 16, 64)
		return RGBA{float64(r), float64(g), float64(b), 1.0}, true, "hex"
	}

	hex3 := regexp.MustCompile(`^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$`)
	if match := hex3.FindStringSubmatch(colorStr); match != nil {
		r, _ := strconv.ParseInt(match[1]+match[1], 16, 64)
		g, _ := strconv.ParseInt(match[2]+match[2], 16, 64)
		b, _ := strconv.ParseInt(match[3]+match[3], 16, 64)
		return RGBA{float64(r), float64(g), float64(b), 1.0}, true, "hex"
	}

	return RGBA{0, 0, 0, 1}, false, ""
}

// Methods
func (t TinyColor) IsValid() bool {
	return t.ok
}

func (t TinyColor) IsDark() bool {
	return t.GetBrightness() < 128
}

func (t TinyColor) IsLight() bool {
	return !t.IsDark()
}

func (t TinyColor) GetAlpha() float64 {
	return t.a
}

func (t *TinyColor) SetAlpha(value float64) *TinyColor {
	t.a = boundAlpha(value)
	t.roundA = math.Round(100*t.a) / 100
	return t
}

func (t TinyColor) GetBrightness() float64 {
	return (t.r*299 + t.g*587 + t.b*114) / 1000
}

func (t TinyColor) GetLuminance() float64 {
	RsRGB := t.r / 255
	GsRGB := t.g / 255
	BsRGB := t.b / 255

	var R, G, B float64

	if RsRGB <= 0.03928 {
		R = RsRGB / 12.92
	} else {
		R = math.Pow((RsRGB+0.055)/1.055, 2.4)
	}

	if GsRGB <= 0.03928 {
		G = GsRGB / 12.92
	} else {
		G = math.Pow((GsRGB+0.055)/1.055, 2.4)
	}

	if BsRGB <= 0.03928 {
		B = BsRGB / 12.92
	} else {
		B = math.Pow((BsRGB+0.055)/1.055, 2.4)
	}

	return 0.2126*R + 0.7152*G + 0.0722*B
}

func (t TinyColor) ToRgb() RGBA {
	return RGBA{
		R: math.Round(t.r),
		G: math.Round(t.g),
		B: math.Round(t.b),
		A: t.a,
	}
}

func (t TinyColor) ToRgbString() string {
	if t.a == 1 {
		return fmt.Sprintf("rgb(%d, %d, %d)", int(math.Round(t.r)), int(math.Round(t.g)), int(math.Round(t.b)))
	}
	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", int(math.Round(t.r)), int(math.Round(t.g)), int(math.Round(t.b)), t.roundA)
}

func (t TinyColor) ToHex() string {
	return fmt.Sprintf("%02x%02x%02x", int(math.Round(t.r)), int(math.Round(t.g)), int(math.Round(t.b)))
}

func (t TinyColor) ToHexString() string {
	return "#" + t.ToHex()
}

// Modifications
func (t TinyColor) Lighten(amount float64) TinyColor {
	hsl := t.ToHsl()
	hsl.L += amount / 100
	hsl.L = clamp01(hsl.L)
	return HslToTinyColor(hsl)
}

func (t TinyColor) Darken(amount float64) TinyColor {
	hsl := t.ToHsl()
	hsl.L -= amount / 100
	hsl.L = clamp01(hsl.L)
	return HslToTinyColor(hsl)
}

// HSL Conversion helpers
func (t TinyColor) ToHsl() HSL {
	r := bound01(t.r, 255)
	g := bound01(t.g, 255)
	b := bound01(t.b, 255)

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))

	var h, s float64
	l := (max + min) / 2

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			if g < b {
				h = (g-b)/d + 6
			} else {
				h = (g - b) / d
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h /= 6
	}

	return HSL{H: h * 360, S: s, L: l, A: t.a}
}

func HslToTinyColor(hsl HSL) TinyColor {
	// HSL to RGB conversion calculation
	// Simplified wrapper return
	return NewTinyColor(fmt.Sprintf("hsl(%d, %d%%, %d%%)", int(hsl.H), int(hsl.S*100), int(hsl.L*100)))
}

// Readability functions
func Readability(color1, color2 TinyColor) float64 {
	l1 := color1.GetLuminance()
	l2 := color2.GetLuminance()
	return (math.Max(l1, l2) + 0.05) / (math.Min(l1, l2) + 0.05)
}

// Example usage
func main() {
	color := NewTinyColor("#ff0000")

	fmt.Println("Is Valid:", color.IsValid())
	fmt.Println("Hex String:", color.ToHexString())
	fmt.Println("RGB String:", color.ToRgbString())
	fmt.Println("Is Dark:", color.IsDark())
	fmt.Println("Luminance:", color.GetLuminance())

	bg := NewTinyColor("#ffffff")
	fmt.Println("Contrast Ratio with White:", Readability(color, bg))
}
