package main

import (
	"reflect"
	"testing"
)

/////////////////////// Read File ///////////////////////

func TestReadDefaultFile(t *testing.T) {
	expected := []string{"1", "3 4", "0001", "0011", "0110"}

	fileLines, err := readFile("./testfiles/default_file.txt")
	if err != nil {
		t.Error(err)
	}
	if len(fileLines) != len(expected) {
		t.Error("File was not read intierly")
	}
	for i, line := range fileLines {
		if line != expected[i] {
			t.Error("File read incorrectly")
		}
	}
}

func TestRead1PixelFile(t *testing.T) {
	expected := []string{"1", "1 1", "1"}

	fileLines, err := readFile("./testfiles/1_pixel_file.txt")
	if err != nil {
		t.Error(err)
	}
	if len(fileLines) != len(expected) {
		t.Error("File was not read intierly")
	}
	for i, line := range fileLines {
		if line != expected[i] {
			t.Error("File read incorrectly")
		}
	}
}

func TestReadEmptyFile(t *testing.T) {
	_, err := readFile("./testfiles/empty_file.txt")
	if err.Error() != "File contains too litle data" {
		t.Error(err)
	}
}

func TestRead1LineFile(t *testing.T) {
	_, err := readFile("./testfiles/1_line_file.txt")
	if err.Error() != "File contains too litle data" {
		t.Error(err)
	}
}

func TestRead2LineFile(t *testing.T) {
	_, err := readFile("./testfiles/2_lines_file.txt")
	if err.Error() != "File contains too litle data" {
		t.Error(err)
	}
}

/////////////////////// Parse Bitmaps Number ///////////////////////

func TestParseBitmapsNumber1(t *testing.T) {
	expectedNumber := 1

	number, err := parseBitmapsNumber("1")
	if err != nil {
		t.Error(err)
	}
	if number != expectedNumber {
		t.Error("Bitmaps number parsed incorrectly")
	}
}

func TestParseBitmapsNumber0(t *testing.T) {
	_, err := parseBitmapsNumber("0")
	if err.Error() != "The number of bitmaps is out of range" {
		t.Error("0 was not considered out of range")
	}
}

func TestParseBitmapsNumber1000(t *testing.T) {
	_, err := parseBitmapsNumber("1001")
	if err.Error() != "The number of bitmaps is out of range" {
		t.Error("0 was not considered out of range")
	}
}

func TestParseWrongBitmapsNumber1(t *testing.T) {
	_, err := parseBitmapsNumber("2 ")
	if err == nil {
		t.Error(err)
	}
}

func TestParseWrongBitmapsNumber2(t *testing.T) {
	_, err := parseBitmapsNumber("2 3")
	if err == nil {
		t.Error(err)
	}
}

func TestParseWrongBitmapsNumber3(t *testing.T) {
	_, err := parseBitmapsNumber("a")
	if err == nil {
		t.Error(err)
	}
}

/////////////////////// Parse Dimensions ///////////////////////

func TestParseDefaultDimensions1(t *testing.T) {
	expectedLength, expectedWidth := 3, 4

	length, width, err := parseDimensions("3 4")
	if err != nil {
		t.Error(err)
	}
	if width != expectedWidth || length != expectedLength {
		t.Error("Bitmap dimensions parsed incorrectly")
	}
}

func TestParseDefaultDimensions2(t *testing.T) {
	expectedLength, expectedWidth := 1, 181

	length, width, err := parseDimensions("1 181")
	if err != nil {
		t.Error(err)
	}
	if width != expectedWidth || length != expectedLength {
		t.Error("Bitmap dimensions parsed incorrectly")
	}
}

func TestParseEmptyDimensions(t *testing.T) {
	_, _, err := parseDimensions("")
	if err.Error() != "Dimensions line should contain 2 numbers: length and width" {
		t.Error(err)
	}

}

func TestParse1Dimensions(t *testing.T) {
	_, _, err := parseDimensions("10")
	if err.Error() != "Dimensions line should contain 2 numbers: length and width" {
		t.Error(err)
	}
}

func TestParse3Dimensions(t *testing.T) {
	_, _, err := parseDimensions("3 4 5")
	if err.Error() != "Dimensions line should contain 2 numbers: length and width" {
		t.Error(err)
	}
}

func TestParseAlphaDimensions(t *testing.T) {
	_, _, err := parseDimensions("27 a")
	if err.Error() != `strconv.Atoi: parsing "a": invalid syntax` {
		t.Error(err)
	}
}

func TestParseSpacedDimensions(t *testing.T) {
	_, _, err := parseDimensions("36  5")
	if err.Error() != "Dimensions line should contain 2 numbers: length and width" {
		t.Error(err)
	}
}

func TestParseNegativeDimensions(t *testing.T) {
	_, _, err := parseDimensions("40 -2")
	if err.Error() != "bitmap width is out of range" {
		t.Error(err)
	}
}

func TestParse0Dimensions(t *testing.T) {
	_, _, err := parseDimensions("0 98")
	if err.Error() != "bitmap length is out of range" {
		t.Error(err)
	}
}

func TestParse183Dimensions(t *testing.T) {
	_, _, err := parseDimensions("183 7")
	if err.Error() != "bitmap length is out of range" {
		t.Error(err)
	}
}

/////////////////////// Parse Bitmap ///////////////////////

func TestParseDefaultBitmap(t *testing.T) {
	expectedPixels := [][]bool{
		{false, false, false, true},
		{false, false, true, true},
		{false, true, true, false},
	}
	expectedWhitePixels := [][2]int{{0, 3}, {1, 2}, {1, 3}, {2, 1}, {2, 2}}
	expectedBlackPixels := map[string]bool{
		"0 0": true,
		"0 1": true,
		"0 2": true,
		"1 0": true,
		"1 1": true,
		"2 0": true,
		"2 3": true,
	}

	index := 1
	fileLines := []string{"1", "3 4", "0001", "0011", "0110"}
	bitmap, err := parseBitmap(index, fileLines)
	if err != nil {
		t.Error(err)
	}

	if bitmap.Length != len(expectedPixels) || len(bitmap.Pixels) != len(expectedPixels) {
		t.Error("Bitmap pixels length is not correct")
	}
	if bitmap.Width != len(expectedPixels[0]) || len(bitmap.Pixels[0]) != len(expectedPixels[0]) {
		t.Error("Bitmap pixels width is not correct")
	}
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			if bitmap.Pixels[i][j] != expectedPixels[i][j] {
				t.Error("Bitmap pixels read incorrectly")
			}
		}
	}

	if len(bitmap.WhitePixels) != len(expectedWhitePixels) {
		t.Error("White pixels read incorrectly")
	}
	for i, pixel := range bitmap.WhitePixels {
		if pixel != bitmap.WhitePixels[i] {
			t.Error("White pixels read incorrectly")
		}
	}

	if !(reflect.DeepEqual(bitmap.BlackPixels, expectedBlackPixels)) {
		t.Error("Black pixels read incorrectly")
	}
}

func TestParseWrongBitmap(t *testing.T) {
	index := 1
	fileLines := []string{"1", "3 4", "0001", "00*1", "0110"}
	_, err := parseBitmap(index, fileLines)
	if err.Error() != "A bitmap can only contain ones and zeros" {
		t.Error(err)
	}
}

func TestParseBlackBitmap(t *testing.T) {
	index := 1
	fileLines := []string{"1", "3 4", "0000", "0000", "0000"}
	_, err := parseBitmap(index, fileLines)
	if err.Error() != "A bitmap should contain at least 1 white pixel" {
		t.Error(err)
	}
}

/////////////////////// Parse Black Pixel ///////////////////////

func TestParseDefaultBlackPixel(t *testing.T) {
	expectedI, expectedJ := 3, 4

	i, j, err := parseBlackPixel("3 4")
	if err != nil {
		t.Error(err)
	}
	if j != expectedJ || i != expectedI {
		t.Error("Black pixel parsed incorrectly")
	}
}

func TestParseDefaultBlackPixel2(t *testing.T) {
	expectedI, expectedJ := 1, 180

	i, j, err := parseBlackPixel("1 180")
	if err != nil {
		t.Error(err)
	}
	if j != expectedJ || i != expectedI {
		t.Error("Black pixel parsed incorrectly")
	}
}

func TestParseEmptyBlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("")
	if err.Error() != "Black pixel key should contain 2 numbers: i and j" {
		t.Error(err)
	}

}

func TestParse1BlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("10")
	if err.Error() != "Black pixel key should contain 2 numbers: i and j" {
		t.Error(err)
	}
}

func TestParse3BlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("3 4 5")
	if err.Error() != "Black pixel key should contain 2 numbers: i and j" {
		t.Error(err)
	}
}

func TestParseAlphaBlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("27 a")
	if err.Error() != `strconv.Atoi: parsing "a": invalid syntax` {
		t.Error(err)
	}
}

func TestParseSpacedBlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("36  5")
	if err.Error() != "Black pixel key should contain 2 numbers: i and j" {
		t.Error(err)
	}
}

func TestParseNegativeBlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("40 -2")
	if err.Error() != "Black pixel's j is out of range" {
		t.Error(err)
	}
}

func TestParse183BlackPixel(t *testing.T) {
	_, _, err := parseBlackPixel("182 7")
	if err.Error() != "Black pixel's i is out of range" {
		t.Error(err)
	}
}
