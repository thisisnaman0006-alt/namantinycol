package main

import (
	"math"
	"testing"
)

// ==========================================
// Helper Assertion Functions
// ==========================================

func assertEquals(t *testing.T, got, expected interface{}, msg string) {
	t.Helper()
	if got != expected {
		t.Errorf("%s | Expected: %v, Got: %v", msg, expected, got)
	}
}

func assertTrue(t *testing.T, condition bool, msg string) {
	t.Helper()
	if !condition {
		t.Errorf("Assertion failed: %s", msg)
	}
}

func assertFalse(t *testing.T, condition bool, msg string) {
	t.Helper()
	if condition {
		t.Errorf("Assertion failed (expected false): %s", msg)
	}
}

func assertFloatAlmostEqual(t *testing.T, got, expected, tolerance float64, msg string) {
	t.Helper()
	if math.Abs(got-expected) > tolerance {
		t.Errorf("%s | Expected approx: %f, Got: %f", msg, expected, got)
	}
}

// ==========================================
// 1. Initialization & Basic Tests
// ==========================================

func TestTinyColorInitialization(t *testing.T) {
	tc := NewTinyColor("red")
	assertTrue(t, tc.IsValid(), "tinycolor should be instantiated")
	assertEquals(t, tc.ToHexString(), "#ff0000", "tinycolor hex output")
}

func TestOriginalInput(t *testing.T) {
	colorRgbUp := "RGB(39, 39, 39)"
	colorRgbLow := "rgb(39, 39, 39)"
	colorRgbMix := "RgB(39, 39, 39)"

	assertEquals(t, NewTinyColor(colorRgbLow).originalInput, colorRgbLow, "original lowercase input")
	assertEquals(t, NewTinyColor(colorRgbUp).originalInput, colorRgbUp, "original uppercase input")
	assertEquals(t, NewTinyColor(colorRgbMix).originalInput, colorRgbMix, "original mixed input")
	assertEquals(t, NewTinyColor("").originalInput, "", "empty string input")
}

func TestCloningColor(t *testing.T) {
	originalColor := NewTinyColor("red")
	originalRgbStr := originalColor.ToRgbString()

	// Go me struct assignment se value copy hoti hai
	clonedColor := originalColor
	assertEquals(t, clonedColor.ToRgbString(), originalColor.ToRgbString(), "cloned color is identical")

	clonedColor.SetAlpha(0.5)
	assertTrue(t, clonedColor.ToRgbString() != originalColor.ToRgbString(), "cloned color changes independently")
	assertEquals(t, originalRgbStr, originalColor.ToRgbString(), "original color remains unchanged")
}

// ==========================================
// 2. Parsing & Conversion Tests
// ==========================================

func TestHexParsing(t *testing.T) {
	assertEquals(t, NewTinyColor("rgb(255, 0, 0)").ToHexString(), "#ff0000", "hex output from rgb")
	assertEquals(t, NewTinyColor("rgb(255, 0, 0)").ToHex(), "ff0000", "hex raw output from rgb")
	assertEquals(t, NewTinyColor("#f00").ToHexString(), "#ff0000", "3-digit hex parsing")
	assertEquals(t, NewTinyColor("f00").ToHexString(), "#ff0000", "3-digit hex without # prefix")
	assertEquals(t, NewTinyColor("#ff0000").ToHexString(), "#ff0000", "6-digit hex parsing")
	assertEquals(t, NewTinyColor("ff0000").ToHexString(), "#ff0000", "6-digit hex without # prefix")
}

func TestInvalidParsing(t *testing.T) {
	invalidInputs := []string{
		"this is not a color",
		"#red",
		"  #red",
		"##123456",
		"  ##123456",
		"rgb 255 0 0",
		"hsl(1000, 100%, 100%)",
	}

	for _, input := range invalidInputs {
		invalidColor := NewTinyColor(input)
		assertFalse(t, invalidColor.IsValid(), "Invalid input test: "+input)
	}
}

func TestNamedColors(t *testing.T) {
	namedColors := map[string]string{
		"aliceblue":    "f0f8ff",
		"antiquewhite": "faebd7",
		"aqua":         "00ffff",
		"aquamarine":   "7fffd4",
		"azure":        "f0ffff",
		"beige":        "f5f5dc",
		"bisque":       "ffe4c4",
		"black":        "000000",
		"blue":         "0000ff",
		"brown":        "a52a2a",
		"cyan":         "00ffff",
		"gray":         "808080",
		"green":        "008000",
		"magenta":      "ff00ff",
		"orange":       "ffa500",
		"red":          "ff0000",
		"white":        "ffffff",
		"yellow":       "ffff00",
	}

	for name, hex := range namedColors {
		assertEquals(t, NewTinyColor(name).ToHex(), hex, "Named color: "+name)
	}
}

// ==========================================
// 3. Alpha / Transparency Tests
// ==========================================

func TestAlphaHandling(t *testing.T) {
	c1 := NewTinyColor("rgba(255, 0, 0, 0.5)")
	assertEquals(t, c1.GetAlpha(), 0.5, "alpha parsing")
	assertEquals(t, c1.ToRgbString(), "rgba(255, 0, 0, 0.5)", "rgba string output")

	c2 := NewTinyColor("#ff0000")
	assertEquals(t, c2.GetAlpha(), 1.0, "default alpha")
	c2.SetAlpha(0.25)
	assertEquals(t, c2.GetAlpha(), 0.25, "set alpha")
}

// ==========================================
// 4. Brightness & Luminance Tests
// ==========================================

func TestGetBrightnessAndLuminance(t *testing.T) {
	assertEquals(t, NewTinyColor("#000000").GetBrightness(), 0.0, "black brightness")
	assertEquals(t, NewTinyColor("#ffffff").GetBrightness(), 255.0, "white brightness")

	assertFloatAlmostEqual(t, NewTinyColor("#000000").GetLuminance(), 0.0, 0.001, "black luminance")
	assertFloatAlmostEqual(t, NewTinyColor("#ffffff").GetLuminance(), 1.0, 0.001, "white luminance")
}

func TestIsDarkAndIsLight(t *testing.T) {
	assertTrue(t, NewTinyColor("#000000").IsDark(), "#000000 is dark")
	assertTrue(t, NewTinyColor("#333333").IsDark(), "#333333 is dark")
	assertFalse(t, NewTinyColor("#ffffff").IsDark(), "#ffffff is not dark")

	assertTrue(t, NewTinyColor("#ffffff").IsLight(), "#ffffff is light")
	assertTrue(t, NewTinyColor("#cccccc").IsLight(), "#cccccc is light")
	assertFalse(t, NewTinyColor("#000000").IsLight(), "#000000 is not light")
}

// ==========================================
// 5. Modifications (Lighten, Darken, Spin)
// ==========================================

func TestModifications(t *testing.T) {
	// Lighten / Darken
	red := NewTinyColor("red")
	assertEquals(t, red.Lighten(20).ToHexString(), "#ff6666", "lighten 20%")
	assertEquals(t, red.Darken(20).ToHexString(), "#990000", "darken 20%")

	// Greyscale
	assertEquals(t, NewTinyColor("#ff0000").Greyscale().ToHexString(), "#808080", "greyscale red")

	// Spin (Hue rotation)
	assertEquals(t, NewTinyColor("red").Spin(180).ToHexString(), "#00ffff", "spin 180 deg")
	assertEquals(t, NewTinyColor("red").Spin(-180).ToHexString(), "#00ffff", "spin -180 deg")
}

// ==========================================
// 6. Readability & Contrast
// ==========================================

func TestReadability(t *testing.T) {
	black := NewTinyColor("#000000")
	white := NewTinyColor("#ffffff")

	assertFloatAlmostEqual(t, Readability(black, black), 1.0, 0.01, "black on black")
	assertFloatAlmostEqual(t, Readability(black, white), 21.0, 0.01, "black on white")
}
