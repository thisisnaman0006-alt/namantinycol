# TinyColor (Go)

## Go color tooling

TinyColor is a small, fast library for color manipulation and conversion written in Go (Golang). It allows many forms of input while providing color conversions and other color utility functions. It has zero external dependencies.

## Installation

You can install `tinycolor` using `go get`:

```bash
go get [github.com/yourusername/tinycolor](https://github.com/yourusername/tinycolor)

import "[github.com/yourusername/tinycolor](https://github.com/yourusername/tinycolor)"

color := tinycolor.New("red")

tinycolor.New("#000")
tinycolor.New("000")
tinycolor.New("#369C")
tinycolor.New("369C")
tinycolor.New("#f0f0f6")
tinycolor.New("f0f0f6")
tinycolor.New("#f0f0f688")
tinycolor.New("f0f0f688")

tinycolor.New("rgb (255, 0, 0)")
tinycolor.New("rgb 255 0 0")
tinycolor.New("rgba (255, 0, 0, .5)")
tinycolor.New(tinycolor.RGB{R: 255, G: 0, B: 0})
tinycolor.FromRatio(tinycolor.RGB{R: 1, G: 0, B: 0})
tinycolor.FromRatio(tinycolor.RGB{R: 0.5, G: 0.5, B: 0.5})

tinycolor.New("hsl(0, 100%, 50%)")
tinycolor.New("hsla(0, 100%, 50%, .5)")
tinycolor.New("hsl 0 1.0 0.5")
tinycolor.New(tinycolor.HSL{H: 0, S: 1, L: 0.5})
tinycolor.FromRatio(tinycolor.HSL{H: 1, S: 0, L: 0})
tinycolor.FromRatio(tinycolor.HSL{H: 0.5, S: 0.5, L: 0.5})

tinycolor.New("hsv(0, 100%, 100%)")
tinycolor.New("hsva(0, 100%, 100%, .5)")
tinycolor.New("hsv 0 1 1")
tinycolor.New(tinycolor.HSV{H: 0, S: 100, V: 100})
tinycolor.FromRatio(tinycolor.HSV{H: 1, S: 0, V: 0})
tinycolor.FromRatio(tinycolor.HSV{H: 0.5, S: 0.5, V: 0.5})

tinycolor.New("RED")
tinycolor.New("blanchedalmond")
tinycolor.New("darkblue")

tinycolor.New(tinycolor.RGB{R: 255, G: 0, B: 0})
tinycolor.New(tinycolor.RGB{R: 255, G: 0, B: 0, A: 0.5})
tinycolor.New(tinycolor.HSL{H: 0, S: 100, L: 50})
tinycolor.New(tinycolor.HSV{H: 0, S: 100, V: 100})

tinycolor.New(tinycolor.RGB{R: 255, G: 0, B: 0})
tinycolor.New(tinycolor.RGB{R: 255, G: 0, B: 0, A: 0.5})
tinycolor.New(tinycolor.HSL{H: 0, S: 100, L: 50})
tinycolor.New(tinycolor.HSV{H: 0, S: 100, V: 100})

color := tinycolor.New("red")
color.GetFormat() // "name"

color = tinycolor.New(tinycolor.RGB{R: 255, G: 255, B: 255})
color.GetFormat() // "rgb"

color := tinycolor.New("red")
color.GetOriginalInput() // "red"

color1 := tinycolor.New("red")
color1.IsValid() // true
color1.ToHexString() // "#ff0000"

color2 := tinycolor.New("not a color")
color2.IsValid() // false
color2.ToString("") // "#000000"

color1 := tinycolor.New("#fff")
color1.GetBrightness() // 255.0

color2 := tinycolor.New("#000")
color2.GetBrightness() // 0.0

color := tinycolor.New("#fff")
color.IsLight() // true
color.IsDark()  // false

color := tinycolor.New("#fff")
color.GetLuminance() // 1.0

color := tinycolor.New("red")
color.GetAlpha() // 1.0
color.SetAlpha(0.5)
color.GetAlpha() // 0.5
color.ToRgbString() // "rgba(255, 0, 0, 0.5)"

color := tinycolor.New("red")
color.ToHsv()       // tinycolor.HSV{H: 0, S: 1, V: 1, A: 1}
color.ToHsvString() // "hsv(0, 100%, 100%)"

color := tinycolor.New("red")
color.ToHsl()       // tinycolor.HSL{H: 0, S: 1, L: 0.5, A: 1}
color.ToHslString() // "hsl(0, 100%, 50%)"

color := tinycolor.New("red")
color.ToHex()        // "ff0000"
color.ToHexString()  // "#ff0000"
color.ToHex8()       // "ff0000ff"
color.ToHex8String() // "#ff0000ff"

color := tinycolor.New("red")
color.ToRgb()       // tinycolor.RGB{R: 255, G: 0, B: 0, A: 1}
color.ToRgbString() // "rgb(255, 0, 0)"

color := tinycolor.New("red")
color.ToPercentageRgb()       // tinycolor.PercentageRGB{R: "100%", G: "0%", B: "0%", A: 1}
color.ToPercentageRgbString() // "rgb(100%, 0%, 0%)"

color := tinycolor.New("red")
color.ToName() // "red"

color1 := tinycolor.New("red")
color1.ToString("")      // "red"
color1.ToString("hsv")  // "hsv(0, 100%, 100%)"

tinycolor.New("red").Lighten(10).Desaturate(10).ToHexString() // "#f53d3d"

tinycolor.New("#f00").Lighten(10).ToString("") // "#ff3333"
tinycolor.New("#f00").Darken(10).ToString("")  // "#cc0000"
tinycolor.New("#f00").Brighten(10).ToString("") // "#ff1919"

tinycolor.New("#f00").Saturate(10).ToString("")
tinycolor.New("#f00").Desaturate(10).ToString("")
tinycolor.New("#f00").Greyscale().ToString("") // "#808080"

tinycolor.New("#f00").Spin(180).ToString("") // "#00ffff"

colors := tinycolor.New("#f00").Analogous(6, 30)
colors = tinycolor.New("#f00").Monochromatic(6)
colors = tinycolor.New("#f00").Splitcomplement()
colors = tinycolor.New("#f00").Triad()
colors = tinycolor.New("#f00").Tetrad()
comp := tinycolor.New("#f00").Complement()

tinycolor.Equals(color1, color2)
tinycolor.Mix(color1, color2, 50)
tinycolor.Random()

tinycolor.Equals(color1, color2)
tinycolor.Mix(color1, color2, 50)
tinycolor.Random()

c1 := tinycolor.New("#000")
c2 := tinycolor.New("#fff")

tinycolor.Readability(c1, c2) // 21.0

// Readability check
options := tinycolor.ReadabilityOptions{Level: "AA", Size: "small"}
tinycolor.IsReadable(c1, c2, options) // true

// Most Readable
list := []*tinycolor.Color{tinycolor.New("#f00"), tinycolor.New("#0f0")}
best := tinycolor.MostReadable(c1, list, tinycolor.MostReadableOptions{IncludeFallbackColors: true})

color1 := tinycolor.New("#F00")
color2 := color1.Clone()
color2.SetAlpha(0.5)

color1.ToString("") // "#ff0000"
color2.ToString("") // "rgba(255, 0, 0, 0.5)"

