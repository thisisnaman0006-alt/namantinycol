package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// ==========================================
// Types & Structures
// ==========================================

type RGB struct {
	R float64
	G float64
	B float64
	A float64
}

type HSL struct {
	H float64
	S float64
	L float64
	A float64
}

type HSV struct {
	H float64
	S float64
	V float64
	A float64
}

type Options struct {
	Format       string
	GradientType bool
}

type WCAGOptions struct {
	Level                string // "AA" or "AAA"
	Size                 string // "small" or "large"
	IncludeFallbackColors bool
}

type TinyColor struct {
	r             float64
	g             float64
	b             float64
	a             float64
	roundA        float64
	format        string
	ok            bool
	originalInput interface{}
	gradientType  bool
}

// ==========================================
// Color Names Data Map
// ==========================================

var names = map[string]string{
	"aliceblue":            "f0f8ff",
	"antiquewhite":        "faebd7",
	"aqua":                "00ffff",
	"aquamarine":          "7fffd4",
	"azure":               "f0ffff",
	"beige":               "f5f5dc",
	"bisque":              "ffe4c4",
	"black":               "000000",
	"blanchedalmond":       "ffebcd",
	"blue":                "0000ff",
	"blueviolet":          "8a2be2",
	"brown":               "a52a2a",
	"burlywood":           "deb887",
	"burntsienna":         "ea7e5d",
	"cadetblue":           "5f9ea0",
	"chartreuse":          "7fff00",
	"chocolate":           "d2691e",
	"coral":               "ff7f50",
	"cornflowerblue":      "6495ed",
	"cornsilk":            "fff8dc",
	"crimson":             "dc143c",
	"cyan":                "00ffff",
	"darkblue":            "00008b",
	"darkcyan":            "008b8b",
	"darkgoldenrod":       "b8860b",
	"darkgray":            "a9a9a9",
	"darkgreen":           "006400",
	"darkgrey":            "a9a9a9",
	"darkkhaki":           "bdb76b",
	"darkmagenta":         "8b008b",
	"darkolivegreen":      "556b2f",
	"darkorange":          "ff8c00",
	"darkorchid":           "9932cc",
	"darkred":             "8b0000",
	"darksalmon":          "e9967a",
	"darkseagreen":        "8fbc8f",
	"darkslateblue":       "483d8b",
	"darkslategray":       "2f4f4f",
	"darkslategrey":       "2f4f4f",
	"darkturquoise":       "00ced1",
	"darkviolet":          "9400d3",
	"deeppink":            "ff1493",
	"deepskyblue":         "00bfff",
	"dimgray":             "696969",
	"dimgrey":             "696969",
	"dodgerblue":          "1e90ff",
	"firebrick":           "b22222",
	"floralwhite":         "fffaf0",
	"forestgreen":         "228b22",
	"fuchsia":             "ff00ff",
	"gainsboro":           "dcdcdc",
	"ghostwhite":          "f8f8ff",
	"gold":                "ffd700",
	"goldenrod":           "daa520",
	"gray":                "808080",
	"green":               "008000",
	"greenyellow":         "adff2f",
	"grey":                "808080",
	"honeydew":            "f0fff0",
	"hotpink":             "ff69b4",
	"indianred":           "cd5c5c",
	"indigo":              "4b0082",
	"ivory":               "fffff0",
	"khaki":               "f0e68c",
	"lavender":            "e6e6fa",
	"lavenderblush":       "fff0f5",
	"lawngreen":           "7cfc00",
	"lemonchiffon":        "fffacd",
	"lightblue":           "add8e6",
	"lightcoral":          "f08080",
	"lightcyan":           "e0ffff",
	"lightgoldenrodyellow": "fafad2",
	"lightgray":           "d3d3d3",
	"lightgreen":          "90ee90",
	"lightgrey":           "d3d3d3",
	"lightpink":           "ffb6c1",
	"lightsalmon":         "ffa07a",
	"lightseagreen":       "20b2aa",
	"lightskyblue":        "87cefa",
	"lightslategray":      "778899",
	"lightslategrey":      "778899",
	"lightsteelblue":      "b0c4de",
	"lightyellow":         "ffffe0",
	"lime":                "00ff00",
	"limegreen":           "32cd32",
	"linen":               "faf0e6",
	"magenta":             "ff00ff",
	"maroon":              "800000",
	"mediumaquamarine":    "66cdaa",
	"mediumblue":          "0000cd",
	"mediumorchid":        "ba55d3",
	"mediumpurple":        "9370db",
	"mediumseagreen":      "3cb371",
	"mediumslateblue":     "7b68ee",
	"mediumspringgreen":    "00fa9a",
	"mediumturquoise":     "48d1cc",
	"mediumvioletred":     "c71585",
	"midnightblue":        "191970",
	"mintcream":           "f5fffa",
	"mistyrose":           "ffe4e1",
	"moccasin":            "ffe4b5",
	"navajowhite":         "ffdead",
	"navy":                "000080",
	"oldlace":             "fdf5e6",
	"olive":               "808000",
	"olivedrab":           "6b8e23",
	"orange":              "ffa500",
	"orangered":           "ff4500",
	"orchid":              "da70d6",
	"palegoldenrod":       "eee8aa",
	"palegreen":           "98fb98",
	"paleturquoise":       "afeeee",
	"palevioletred":       "db7093",
	"papayawhip":          "ffefd5",
	"peachpuff":           "ffdab9",
	"peru":                "cd853f",
	"pink":                "ffc0cb",
	"plum":                "dda0dd",
	"powderblue":          "b0e0e6",
	"purple":              "800080",
	"rebeccapurple":       "663399",
	"red":                 "ff0000",
	"rosybrown":           "bc8f8f",
	"royalblue":           "4169e1",
	"saddlebrown":         "8b4513",
	"salmon":              "fa8072",
	"sandybrown":          "f4a460",
	"seagreen":            "2e8b57",
	"seashell":            "fff5ee",
	"sienna":              "a0522d",
	"silver":              "c0c0c0",
	"skyblue":             "87ceeb",
	"slateblue":           "6a5acd",
	"slategray":           "708090",
	"slategrey":           "708090",
	"snow":                "fffafa",
	"springgreen":         "00ff7f",
	"steelblue":           "4682b4",
	"tan":                 "d2b48c",
	"teal":                "008080",
	"thistle":             "d8bfd8",
	"tomato":              "ff6347",
	"turquoise":           "40e0d0",
	"violet":              "ee82ee",
	"wheat":               "f5deb3",
	"white":               "ffffff",
	"whitesmoke":          "f5f5f5",
	"yellow":              "ffff00",
	"yellowgreen":         "9acd32",
}

var hexNames = flipMap(names)

func flipMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

// ==========================================
// Regex Matchers
// ==========================================

const (
	cssInteger = `[-\+]?\d+%?`
	cssNumber  = `[-\+]?\d*\.\d+%?`
	cssUnit    = `(?:` + cssNumber + `)|(?:` + cssInteger + `)`
	permMatch3 = `[\s|\(]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)\s*\)?`
	permMatch4 = `[\s|\(]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)[,|\s]+(` + cssUnit + `)\s*\)?`
)

var (
	reRgb   = regexp.MustCompile(`rgb` + permMatch3)
	reRgba  = regexp.MustCompile(`rgba` + permMatch4)
	reHsl   = regexp.MustCompile(`hsl` + permMatch3)
	reHsla  = regexp.MustCompile(`hsla` + permMatch4)
	reHsv   = regexp.MustCompile(`hsv` + permMatch3)
	reHsva  = regexp.MustCompile(`hsva` + permMatch4)
	reHex3  = regexp.MustCompile(`^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$`)
	reHex6  = regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`)
	reHex4  = regexp.MustCompile(`^#?([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})([0-9a-fA-F]{1})$`)
	reHex8  = regexp.MustCompile(`^#?([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})([0-9a-fA-F]{2})$`)
	reUnit  = regexp.MustCompile(cssUnit)
)

// ==========================================
// Helper Utilities
// ==========================================

func boundAlpha(a interface{}) float64 {
	val := 1.0
	switch v := a.(type) {
	case float64:
		val = v
	case int:
		val = float64(v)
	case string:
		if parsed, err := strconv.ParseFloat(v, 64); err == nil {
			val = parsed
		}
	}
	if val < 0 || val > 1 {
		return 1.0
	}
	return val
}

func bound01(n interface{}, maxVal float64) float64 {
	var val float64
	strVal := fmt.Sprintf("%v", n)
	isPercent := strings.Contains(strVal, "%")

	if parsed, err := strconv.ParseFloat(strings.TrimSuffix(strVal, "%"), 64); err == nil {
		val = parsed
	}

	val = math.Min(maxVal, math.Max(0, val))

	if isPercent {
		val = float64(int(val*maxVal)) / 100.0
	}

	if math.Abs(val-maxVal) < 0.000001 {
		return 1.0
	}

	return math.Mod(val, maxVal) / maxVal
}

func clamp01(val float64) float64 {
	return math.Min(1, math.Max(0, val))
}

func parseIntFromHex(val string) float64 {
	i, _ := strconv.ParseInt(val, 16, 64)
	return float64(i)
}

func pad2(c string) string {
	if len(c) == 1 {
		return "0" + c
	}
	return c
}

func convertDecimalToHex(d float64) string {
	return fmt.Sprintf("%02x", int(math.Round(d*255.0)))
}

func convertHexToDecimal(h string) float64 {
	return parseIntFromHex(h) / 255.0
}

func convertToPercentage(n float64) string {
	if n <= 1 {
		return fmt.Sprintf("%g%%", n*100)
	}
	return fmt.Sprintf("%g", n)
}

// ==========================================
// Constructors & Parsing
// ==========================================

func NewTinyColor(input interface{}, opts ...Options) TinyColor {
	opt := Options{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	if tc, ok := input.(TinyColor); ok {
		return tc
	}

	res := inputToRGB(input)
	tc := TinyColor{
		originalInput: input,
		r:             res.R,
		g:             res.G,
		b:             res.B,
		a:             res.A,
		roundA:        math.Round(100*res.A) / 100,
		format:        res.Format,
		ok:            res.Ok,
		gradientType:  opt.GradientType,
	}

	if opt.Format != "" {
		tc.format = opt.Format
	}

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

func FromRatio(color RGB, opts ...Options) TinyColor {
	color.R = color.R * 100
	color.G = color.G * 100
	color.B = color.B * 100
	return NewTinyColor(color, opts...)
}

type internalRGB struct {
	RGB
	Ok     bool
	Format string
}

func inputToRGB(color interface{}) internalRGB {
	rgb := RGB{R: 0, G: 0, B: 0}
	a := 1.0
	ok := false
	format := ""

	switch v := color.(type) {
	case string:
		obj, fmtStr, isOk := stringInputToObject(v)
		if isOk {
			return parseParsedObject(obj, fmtStr)
		}
	case RGB:
		rgb = rgbToRgb(v.R, v.G, v.B)
		ok = true
		a = v.A
		format = "rgb"
	case HSL:
		rgb = hslToRgb(v.H, v.S, v.L)
		ok = true
		a = v.A
		format = "hsl"
	case HSV:
		rgb = hsvToRgb(v.H, v.S, v.V)
		ok = true
		a = v.A
		format = "hsv"
	}

	return internalRGB{
		RGB: RGB{
			R: math.Min(255, math.Max(rgb.R, 0)),
			G: math.Min(255, math.Max(rgb.G, 0)),
			B: math.Min(255, math.Max(rgb.B, 0)),
			A: boundAlpha(a),
		},
		Ok:     ok,
		Format: format,
	}
}

func parseParsedObject(obj map[string]interface{}, fmtStr string) internalRGB {
	rgb := RGB{}
	ok := true
	a := 1.0

	if val, exists := obj["a"]; exists {
		a = boundAlpha(val)
	}

	if r, okR := obj["r"]; okR {
		g, b := obj["g"], obj["b"]
		rgb = rgbToRgb(r, g, b)
	} else if h, okH := obj["h"]; okH {
		if s, okS := obj["s"]; okS {
			if v, okV := obj["v"]; okV {
				rgb = hsvToRgb(h, s, v)
			} else if l, okL := obj["l"]; okL {
				rgb = hslToRgb(h, s, l)
			}
		}
	}

	return internalRGB{
		RGB: RGB{
			R: math.Min(255, math.Max(rgb.R, 0)),
			G: math.Min(255, math.Max(rgb.G, 0)),
			B: math.Min(255, math.Max(rgb.B, 0)),
			A: boundAlpha(a),
		},
		Ok:     ok,
		Format: fmtStr,
	}
}

func stringInputToObject(color string) (map[string]interface{}, string, bool) {
	color = strings.ToLower(strings.TrimSpace(color))
	named := false

	if val, ok := names[color]; ok {
		color = val
		named = true
	} else if color == "transparent" {
		return map[string]interface{}{"r": 0, "g": 0, "b": 0, "a": 0.0}, "name", true
	}

	if match := reRgb.FindStringSubmatch(color); match != nil {
		return map[string]interface{}{"r": match[1], "g": match[2], "b": match[3]}, "rgb", true
	}
	if match := reRgba.FindStringSubmatch(color); match != nil {
		return map[string]interface{}{"r": match[1], "g": match[2], "b": match[3], "a": match[4]}, "rgb", true
	}
	if match := reHsl.FindStringSubmatch(color); match != nil {
		return map[string]interface{}{"h": match[1], "s": match[2], "l": match[3]}, "hsl", true
	}
	if match := reHsla.FindStringSubmatch(color); match != nil {
		return map[string]interface{}{"h": match[1], "s": match[2], "l": match[3], "a": match[4]}, "hsl", true
	}
	if match := reHsv.FindStringSubmatch(color); match != nil {
		return map[string]interface{}{"h": match[1], "s": match[2], "v": match[3]}, "hsv", true
	}
	if match := reHsva.FindStringSubmatch(color); match != nil {
		return map[string]interface{}{"h": match[1], "s": match[2], "v": match[3], "a": match[4]}, "hsv", true
	}
	if match := reHex8.FindStringSubmatch(color); match != nil {
		fmtStr := "hex8"
		if named {
			fmtStr = "name"
		}
		return map[string]interface{}{
			"r": parseIntFromHex(match[1]),
			"g": parseIntFromHex(match[2]),
			"b": parseIntFromHex(match[3]),
			"a": convertHexToDecimal(match[4]),
		}, fmtStr, true
	}
	if match := reHex6.FindStringSubmatch(color); match != nil {
		fmtStr := "hex"
		if named {
			fmtStr = "name"
		}
		return map[string]interface{}{
			"r": parseIntFromHex(match[1]),
			"g": parseIntFromHex(match[2]),
			"b": parseIntFromHex(match[3]),
		}, fmtStr, true
	}
	if match := reHex4.FindStringSubmatch(color); match != nil {
		fmtStr := "hex8"
		if named {
			fmtStr = "name"
		}
		return map[string]interface{}{
			"r": parseIntFromHex(match[1] + match[1]),
			"g": parseIntFromHex(match[2] + match[2]),
			"b": parseIntFromHex(match[3] + match[3]),
			"a": convertHexToDecimal(match[4] + match[4]),
		}, fmtStr, true
	}
	if match := reHex3.FindStringSubmatch(color); match != nil {
		fmtStr := "hex"
		if named {
			fmtStr = "name"
		}
		return map[string]interface{}{
			"r": parseIntFromHex(match[1] + match[1]),
			"g": parseIntFromHex(match[2] + match[2]),
			"b": parseIntFromHex(match[3] + match[3]),
		}, fmtStr, true
	}

	return nil, "", false
}

// ==========================================
// Conversions Algorithms
// ==========================================

func rgbToRgb(r, g, b interface{}) RGB {
	return RGB{
		R: bound01(r, 255) * 255,
		G: bound01(g, 255) * 255,
		B: bound01(b, 255) * 255,
	}
}

func rgbToHsl(r, g, b float64) HSL {
	r01 := bound01(r, 255)
	g01 := bound01(g, 255)
	b01 := bound01(b, 255)

	max := math.Max(r01, math.Max(g01, b01))
	min := math.Min(r01, math.Min(g01, b01))
	var h, s float64
	l := (max + min) / 2.0

	if max == min {
		h = 0
		s = 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r01:
			add := 0.0
			if g01 < b01 {
				add = 6.0
			}
			h = (g01-b01)/d + add
		case g01:
			h = (b01-r01)/d + 2.0
		case b01:
			h = (r01-g01)/d + 4.0
		}
		h /= 6.0
	}
	return HSL{H: h, S: s, L: l}
}

func hslToRgb(h, s, l interface{}) RGB {
	h01 := bound01(h, 360)
	s01 := bound01(s, 100)
	l01 := bound01(l, 100)

	hue2rgb := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		if t < 1.0/6.0 {
			return p + (q-p)*6.0*t
		}
		if t < 1.0/2.0 {
			return q
		}
		if t < 2.0/3.0 {
			return p + (q-p)*(2.0/3.0-t)*6.0
		}
		return p
	}

	var r, g, b float64
	if s01 == 0 {
		r, g, b = l01, l01, l01
	} else {
		var q float64
		if l01 < 0.5 {
			q = l01 * (1.0 + s01)
		} else {
			q = l01 + s01 - l01*s01
		}
		p := 2.0*l01 - q
		r = hue2rgb(p, q, h01+1.0/3.0)
		g = hue2rgb(p, q, h01)
		b = hue2rgb(p, q, h01-1.0/3.0)
	}

	return RGB{R: r * 255, G: g * 255, B: b * 255}
}

func rgbToHsv(r, g, b float64) HSV {
	r01 := bound01(r, 255)
	g01 := bound01(g, 255)
	b01 := bound01(b, 255)

	max := math.Max(r01, math.Max(g01, b01))
	min := math.Min(r01, math.Min(g01, b01))
	var h float64
	v := max
	d := max - min

	var s float64
	if max == 0 {
		s = 0
	} else {
		s = d / max
	}

	if max == min {
		h = 0
	} else {
		switch max {
		case r01:
			add := 0.0
			if g01 < b01 {
				add = 6.0
			}
			h = (g01-b01)/d + add
		case g01:
			h = (b01-r01)/d + 2.0
		case b01:
			h = (r01-g01)/d + 4.0
		}
		h /= 6.0
	}

	return HSV{H: h, S: s, V: v}
}

func hsvToRgb(h, s, v interface{}) RGB {
	h01 := bound01(h, 360) * 6.0
	s01 := bound01(s, 100)
	v01 := bound01(v, 100)

	i := math.Floor(h01)
	f := h01 - i
	p := v01 * (1.0 - s01)
	q := v01 * (1.0 - f*s01)
	t := v01 * (1.0 - (1.0-f)*s01)
	mod := int(i) % 6

	rList := []float64{v01, q, p, p, t, v01}
	gList := []float64{t, v01, v01, q, p, p}
	bList := []float64{p, p, t, v01, v01, q}

	return RGB{
		R: rList[mod] * 255,
		G: gList[mod] * 255,
		B: bList[mod] * 255,
	}
}

func rgbToHex(r, g, b float64, allow3Char bool) string {
	hex := []string{
		fmt.Sprintf("%02x", int(math.Round(r))),
		fmt.Sprintf("%02x", int(math.Round(g))),
		fmt.Sprintf("%02x", int(math.Round(b))),
	}

	if allow3Char && hex[0][0] == hex[0][1] && hex[1][0] == hex[1][1] && hex[2][0] == hex[2][1] {
		return string(hex[0][0]) + string(hex[1][0]) + string(hex[2][0])
	}
	return strings.Join(hex, "")
}

func rgbaToHex(r, g, b, a float64, allow4Char bool) string {
	hex := []string{
		fmt.Sprintf("%02x", int(math.Round(r))),
		fmt.Sprintf("%02x", int(math.Round(g))),
		fmt.Sprintf("%02x", int(math.Round(b))),
		fmt.Sprintf("%02x", int(math.Round(a*255))),
	}

	if allow4Char && hex[0][0] == hex[0][1] && hex[1][0] == hex[1][1] && hex[2][0] == hex[2][1] && hex[3][0] == hex[3][1] {
		return string(hex[0][0]) + string(hex[1][0]) + string(hex[2][0]) + string(hex[3][0])
	}
	return strings.Join(hex, "")
}

func rgbaToArgbHex(r, g, b, a float64) string {
	hex := []string{
		fmt.Sprintf("%02x", int(math.Round(a*255))),
		fmt.Sprintf("%02x", int(math.Round(r))),
		fmt.Sprintf("%02x", int(math.Round(g))),
		fmt.Sprintf("%02x", int(math.Round(b))),
	}
	return strings.Join(hex, "")
}

// ==========================================
// Instance Methods
// ==========================================

func (tc TinyColor) IsValid() bool         { return tc.ok }
func (tc TinyColor) GetOriginalInput() interface{} { return tc.originalInput }
func (tc TinyColor) GetFormat() string     { return tc.format }
func (tc TinyColor) GetAlpha() float64     { return tc.a }

func (tc *TinyColor) SetAlpha(val float64) *TinyColor {
	tc.a = boundAlpha(val)
	tc.roundA = math.Round(100*tc.a) / 100
	return tc
}

func (tc TinyColor) GetBrightness() float64 {
	rgb := tc.ToRgb()
	return (rgb.R*299 + rgb.G*587 + rgb.B*114) / 1000.0
}

func (tc TinyColor) IsDark() bool  { return tc.GetBrightness() < 128 }
func (tc TinyColor) IsLight() bool { return !tc.IsDark() }

func (tc TinyColor) GetLuminance() float64 {
	rgb := tc.ToRgb()
	RsRGB := rgb.R / 255.0
	GsRGB := rgb.G / 255.0
	BsRGB := rgb.B / 255.0

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

func (tc TinyColor) ToHsv() HSV {
	hsv := rgbToHsv(tc.r, tc.g, tc.b)
	return HSV{H: hsv.H * 360, S: hsv.S, V: hsv.V, A: tc.a}
}

func (tc TinyColor) ToHsvString() string {
	hsv := rgbToHsv(tc.r, tc.g, tc.b)
	h := math.Round(hsv.H * 360)
	s := math.Round(hsv.S * 100)
	v := math.Round(hsv.V * 100)
	if tc.a == 1 {
		return fmt.Sprintf("hsv(%g, %g%%, %g%%)", h, s, v)
	}
	return fmt.Sprintf("hsva(%g, %g%%, %g%%, %g)", h, s, v, tc.roundA)
}

func (tc TinyColor) ToHsl() HSL {
	hsl := rgbToHsl(tc.r, tc.g, tc.b)
	return HSL{H: hsl.H * 360, S: hsl.S, L: hsl.L, A: tc.a}
}

func (tc TinyColor) ToHslString() string {
	hsl := rgbToHsl(tc.r, tc.g, tc.b)
	h := math.Round(hsl.H * 360)
	s := math.Round(hsl.S * 100)
	l := math.Round(hsl.L * 100)
	if tc.a == 1 {
		return fmt.Sprintf("hsl(%g, %g%%, %g%%)", h, s, l)
	}
	return fmt.Sprintf("hsla(%g, %g%%, %g%%, %g)", h, s, l, tc.roundA)
}

func (tc TinyColor) ToHex(allow3Char ...bool) string {
	a3 := false
	if len(allow3Char) > 0 {
		a3 = allow3Char[0]
	}
	return rgbToHex(tc.r, tc.g, tc.b, a3)
}

func (tc TinyColor) ToHexString(allow3Char ...bool) string {
	return "#" + tc.ToHex(allow3Char...)
}

func (tc TinyColor) ToHex8(allow4Char ...bool) string {
	a4 := false
	if len(allow4Char) > 0 {
		a4 = allow4Char[0]
	}
	return rgbaToHex(tc.r, tc.g, tc.b, tc.a, a4)
}

func (tc TinyColor) ToHex8String(allow4Char ...bool) string {
	return "#" + tc.ToHex8(allow4Char...)
}

func (tc TinyColor) ToRgb() RGB {
	return RGB{
		R: math.Round(tc.r),
		G: math.Round(tc.g),
		B: math.Round(tc.b),
		A: tc.a,
	}
}

func (tc TinyColor) ToRgbString() string {
	if tc.a == 1 {
		return fmt.Sprintf("rgb(%g, %g, %g)", math.Round(tc.r), math.Round(tc.g), math.Round(tc.b))
	}
	return fmt.Sprintf("rgba(%g, %g, %g, %g)", math.Round(tc.r), math.Round(tc.g), math.Round(tc.b), tc.roundA)
}

func (tc TinyColor) ToPercentageRgbString() string {
	r := math.Round(bound01(tc.r, 255) * 100)
	g := math.Round(bound01(tc.g, 255) * 100)
	b := math.Round(bound01(tc.b, 255) * 100)
	if tc.a == 1 {
		return fmt.Sprintf("rgb(%g%%, %g%%, %g%%)", r, g, b)
	}
	return fmt.Sprintf("rgba(%g%%, %g%%, %g%%, %g)", r, g, b, tc.roundA)
}

func (tc TinyColor) ToName() interface{} {
	if tc.a == 0 {
		return "transparent"
	}
	if tc.a < 1 {
		return false
	}
	if val, ok := hexNames[rgbToHex(tc.r, tc.g, tc.b, true)]; ok {
		return val
	}
	return false
}

func (tc TinyColor) ToFilter(secondColor ...interface{}) string {
	hex8String := "#" + rgbaToArgbHex(tc.r, tc.g, tc.b, tc.a)
	secondHex8String := hex8String
	gradientType := ""
	if tc.gradientType {
		gradientType = "GradientType = 1, "
	}
	if len(secondColor) > 0 {
		s := NewTinyColor(secondColor[0])
		secondHex8String = "#" + rgbaToArgbHex(s.r, s.g, s.b, s.a)
	}
	return "progid:DXImageTransform.Microsoft.gradient(" + gradientType + "startColorstr=" + hex8String + ",endColorstr=" + secondHex8String + ")"
}

func (tc TinyColor) ToString(format ...string) string {
	fmtStr := tc.format
	formatSet := false
	if len(format) > 0 && format[0] != "" {
		fmtStr = format[0]
		formatSet = true
	}

	hasAlpha := tc.a < 1 && tc.a >= 0
	needsAlphaFormat := !formatSet && hasAlpha && (fmtStr == "hex" || fmtStr == "hex6" || fmtStr == "hex3" || fmtStr == "hex4" || fmtStr == "hex8" || fmtStr == "name")

	if needsAlphaFormat {
		if fmtStr == "name" && tc.a == 0 {
			if name, ok := tc.ToName().(string); ok {
				return name
			}
		}
		return tc.ToRgbString()
	}

	switch fmtStr {
	case "rgb":
		return tc.ToRgbString()
	case "prgb":
		return tc.ToPercentageRgbString()
	case "hex", "hex6":
		return tc.ToHexString()
	case "hex3":
		return tc.ToHexString(true)
	case "hex4":
		return tc.ToHex8String(true)
	case "hex8":
		return tc.ToHex8String()
	case "name":
		if name, ok := tc.ToName().(string); ok {
			return name
		}
	case "hsl":
		return tc.ToHslString()
	case "hsv":
		return tc.ToHsvString()
	}

	return tc.ToHexString()
}

func (tc TinyColor) Clone() TinyColor {
	return NewTinyColor(tc.ToString())
}

// ==========================================
// Color Modification Functions
// ==========================================

func (tc TinyColor) Lighten(amount ...float64) TinyColor {
	amt := 10.0
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := tc.ToHsl()
	hsl.L += amt / 100.0
	hsl.L = clamp01(hsl.L)
	return NewTinyColor(hsl)
}

func (tc TinyColor) Brighten(amount ...float64) TinyColor {
	amt := 10.0
	if len(amount) > 0 {
		amt = amount[0]
	}
	rgb := tc.ToRgb()
	rgb.R = math.Max(0, math.Min(255, rgb.R-math.Round(255*-(amt/100.0))))
	rgb.G = math.Max(0, math.Min(255, rgb.G-math.Round(255*-(amt/100.0))))
	rgb.B = math.Max(0, math.Min(255, rgb.B-math.Round(255*-(amt/100.0))))
	return NewTinyColor(rgb)
}

func (tc TinyColor) Darken(amount ...float64) TinyColor {
	amt := 10.0
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := tc.ToHsl()
	hsl.L -= amt / 100.0
	hsl.L = clamp01(hsl.L)
	return NewTinyColor(hsl)
}

func (tc TinyColor) Desaturate(amount ...float64) TinyColor {
	amt := 10.0
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := tc.ToHsl()
	hsl.S -= amt / 100.0
	hsl.S = clamp01(hsl.S)
	return NewTinyColor(hsl)
}

func (tc TinyColor) Saturate(amount ...float64) TinyColor {
	amt := 10.0
	if len(amount) > 0 {
		amt = amount[0]
	}
	hsl := tc.ToHsl()
	hsl.S += amt / 100.0
	hsl.S = clamp01(hsl.S)
	return NewTinyColor(hsl)
}

func (tc TinyColor) Greyscale() TinyColor {
	return tc.Desaturate(100)
}

func (tc TinyColor) Spin(amount float64) TinyColor {
	hsl := tc.ToHsl()
	hue := math.Mod(hsl.H+amount, 360)
	if hue < 0 {
		hue = 360 + hue
	}
	hsl.H = hue
	return NewTinyColor(hsl)
}

// ==========================================
// Color Combination Functions
// ==========================================

func (tc TinyColor) Complement() TinyColor {
	hsl := tc.ToHsl()
	hsl.H = math.Mod(hsl.H+180, 360)
	return NewTinyColor(hsl)
}

func Polyad(color interface{}, number int) []TinyColor {
	if number <= 0 {
		panic("Argument to polyad must be a positive number")
	}
	tc := NewTinyColor(color)
	hsl := tc.ToHsl()
	result := []TinyColor{tc}
	step := 360.0 / float64(number)

	for i := 1; i < number; i++ {
		result = append(result, NewTinyColor(HSL{
			H: math.Mod(hsl.H+float64(i)*step, 360),
			S: hsl.S,
			L: hsl.L,
			A: hsl.A,
		}))
	}
	return result
}

func (tc TinyColor) Triad() []TinyColor {
	return Polyad(tc, 3)
}

func (tc TinyColor) Tetrad() []TinyColor {
	return Polyad(tc, 4)
}

func (tc TinyColor) Splitcomplement() []TinyColor {
	hsl := tc.ToHsl()
	h := hsl.H
	return []TinyColor{
		tc,
		NewTinyColor(HSL{H: math.Mod(h+72, 360), S: hsl.S, L: hsl.L, A: hsl.A}),
		NewTinyColor(HSL{H: math.Mod(h+216, 360), S: hsl.S, L: hsl.L, A: hsl.A}),
	}
}

func (tc TinyColor) Analogous(resultsOpt ...int) []TinyColor {
	results := 6
	slices := 30
	if len(resultsOpt) > 0 {
		results = resultsOpt[0]
	}

	hsl := tc.ToHsl()
	part := 360.0 / float64(slices)
	ret := []TinyColor{tc}

	hsl.H = math.Mod(hsl.H-(part*float64(results>>1))+720, 360)
	for results > 1 {
		results--
		hsl.H = math.Mod(hsl.H+part, 360)
		ret = append(ret, NewTinyColor(hsl))
	}
	return ret
}

func (tc TinyColor) Monochromatic(resultsOpt ...int) []TinyColor {
	results := 6
	if len(resultsOpt) > 0 {
		results = resultsOpt[0]
	}

	hsv := tc.ToHsv()
	ret := []TinyColor{}
	modification := 1.0 / float64(results)

	for i := 0; i < results; i++ {
		ret = append(ret, NewTinyColor(HSV{
			H: hsv.H,
			S: hsv.S,
			V: hsv.V,
			A: hsv.A,
		}))
		hsv.V = math.Mod(hsv.V+modification, 1.0)
	}
	return ret
}

// ==========================================
// Static Utility Functions
// ==========================================

func Equals(color1, color2 interface{}) bool {
	if color1 == nil || color2 == nil {
		return false
	}
	return NewTinyColor(color1).ToRgbString() == NewTinyColor(color2).ToRgbString()
}

func Random() TinyColor {
	return FromRatio(RGB{
		R: rand.Float64(),
		G: rand.Float64(),
		B: rand.Float64(),
	})
}

func Mix(color1, color2 interface{}, amount ...float64) TinyColor {
	amt := 50.0
	if len(amount) > 0 {
		amt = amount[0]
	}
	rgb1 := NewTinyColor(color1).ToRgb()
	rgb2 := NewTinyColor(color2).ToRgb()

	p := amt / 100.0
	return NewTinyColor(RGB{
		R: (rgb2.R-rgb1.R)*p + rgb1.R,
		G: (rgb2.G-rgb1.G)*p + rgb1.G,
		B: (rgb2.B-rgb1.B)*p + rgb1.B,
		A: (rgb2.A-rgb1.A)*p + rgb1.A,
	})
}

func Readability(color1, color2 interface{}) float64 {
	c1 := NewTinyColor(color1)
	c2 := NewTinyColor(color2)
	return (math.Max(c1.GetLuminance(), c2.GetLuminance()) + 0.05) / (math.Min(c1.GetLuminance(), c2.GetLuminance()) + 0.05)
}

func IsReadable(color1, color2 interface{}, wcag ...WCAGOptions) bool {
	readability := Readability(color1, color2)
	opts := WCAGOptions{Level: "AA", Size: "small"}
	if len(wcag) > 0 {
		if wcag[0].Level != "" {
			opts.Level = strings.ToUpper(wcag[0].Level)
		}
		if wcag[0].Size != "" {
			opts.Size = strings.ToLower(wcag[0].Size)
		}
	}

	switch opts.Level + opts.Size {
	case "AAsmall", "AAAlarge":
		return readability >= 4.5
	case "AAlarge":
		return readability >= 3.0
	case "AAAsmall":
		return readability >= 7.0
	}
	return false
}

func MostReadable(baseColor interface{}, colorList []interface{}, args ...WCAGOptions) TinyColor {
	var bestColor TinyColor
	bestScore := 0.0
	opts := WCAGOptions{}
	if len(args) > 0 {
		opts = args[0]
	}

	for _, c := range colorList {
		r := Readability(baseColor, c)
		if r > bestScore {
			bestScore = r
			bestColor = NewTinyColor(c)
		}
	}

	if IsReadable(baseColor, bestColor, opts) || !opts.IncludeFallbackColors {
		return bestColor
	}

	opts.IncludeFallbackColors = false
	return MostReadable(baseColor, []interface{}{"#fff", "#000"}, opts)
}
