package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// readFile reads a file and returns a slice of strings, 1 for every line
func readFile(filePath string) ([]string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// You need 7 bytes for a 1 pixel bitmap
	if len(fileData) < 7 {
		return nil, errors.New("File contains too litle data")
	}

	fileLines := strings.Split(string(fileData), "\n")
	// You need 3 need lines for a 1 pixel bitmap
	if len(fileData) < 3 {
		return nil, errors.New("File should contain at least 3 lines")
	}

	return fileLines, nil
}

// parseBitmapsNumber parses the number of bitmaps from the first line's string
func parseBitmapsNumber(firstLine string) (int, error) {
	bitmapsNumber, err := strconv.Atoi(firstLine)
	if err != nil {
		return 0, err
	}
	if bitmapsNumber < 1 || bitmapsNumber > maxBitmapsNumber {
		return 0, errors.New("The number of bitmaps is out of range")
	}

	return bitmapsNumber, nil
}

// parseDimensions parses the length and width of a bitmap from a string
func parseDimensions(firstLine string) (int, int, error) {
	dimensions := strings.Split(firstLine, " ")

	if len(dimensions) != 2 {
		return 0, 0, errors.New("Dimensions line should contain 2 numbers: length and width")
	}

	length, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, err
	}
	if length < 1 || length > maxDimension {
		return 0, 0, errors.New("bitmap length is out of range")
	}

	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return 0, 0, err
	}
	if width < 1 || width > maxDimension {
		return 0, 0, errors.New("bitmap width is out of range")
	}

	return length, width, nil
}

// parseBitmap creates a bitmap from a slice of strings by starting at the right index
func parseBitmap(index int, fileLines []string) (Bitmap, error) {
	var bitmap Bitmap
	var err error
	// parse length and width at index
	bitmap.Length, bitmap.Width, err = parseDimensions(fileLines[index])
	if err != nil {
		return bitmap, err
	}
	// go from dimensions index to bitmap index
	index++

	// initiate pixels and blackPixels
	bitmap.Pixels = make([][]bool, bitmap.Length)
	for i := range bitmap.Pixels {
		bitmap.Pixels[i] = make([]bool, bitmap.Width)
	}
	bitmap.BlackPixels = make(map[string]bool)

	// fill in pixels, withePixels, and blackPixels
	for i := index; i < index+bitmap.Length; i++ {
		// row is a list of pixels
		row := strings.Split(fileLines[i], "")
		for j := 0; j < bitmap.Width; j++ {
			if row[j] == "1" {
				bitmap.Pixels[i-index][j] = true
				bitmap.WhitePixels = append(bitmap.WhitePixels, [2]int{i - index, j})

			} else if row[j] == "0" {
				bitmap.Pixels[i-index][j] = false
				bitmap.BlackPixels[fmt.Sprintf("%d %d", i-index, j)] = true
			} else {
				return bitmap, errors.New("A bitmap can only contain ones and zeros")
			}
		}
	}

	if len(bitmap.WhitePixels) < 1 {
		return bitmap, errors.New("A bitmap should contain at least 1 white pixel")
	}

	// initiate distances
	bitmap.Distances = make([][]int, bitmap.Length)
	for i := range bitmap.Distances {
		bitmap.Distances[i] = make([]int, bitmap.Width)
	}

	return bitmap, nil
}

// parseBlackPixel converts black pixel's string key to 2 ints i and j
func parseBlackPixel(key string) (int, int, error) {
	blackPixel := strings.Split(key, " ")

	if len(blackPixel) != 2 {
		return 0, 0, errors.New("Black pixel key should contain 2 numbers: i and j")
	}

	i, err := strconv.Atoi(blackPixel[0])
	if err != nil {
		return 0, 0, err
	}
	if i < 0 || i > maxDimension-1 {
		return 0, 0, errors.New("Black pixel's i is out of range")
	}

	j, err := strconv.Atoi(blackPixel[1])
	if err != nil {
		return 0, 0, err
	}
	if j < 0 || i > maxDimension {
		return 0, 0, errors.New("Black pixel's j is out of range")
	}

	return i, j, nil
}
