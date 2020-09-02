package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if len(fileData) < 11 {
		return nil, errors.New("File contains too litle data")
	}

	fileLines := strings.Split(string(fileData), "\n")
	if len(fileData) < 3 {
		return nil, errors.New("File should contain at least 3 lines")
	}

	return fileLines, nil
}

func parseBitmapsNumber(firstLine string) (int, error) {
	bitmapsNumber, err := strconv.Atoi(firstLine)
	if err != nil {
		return 0, err
	}
	if bitmapsNumber < 1 || bitmapsNumber > 1000 {
		return 0, errors.New("The number of bitmaps is out of range")
	}

	return bitmapsNumber, nil
}

func parseDimensions(firstLine string) (int, int, error) {
	dimensions := strings.Split(firstLine, " ")

	if len(dimensions) != 2 {
		return 0, 0, errors.New("Dimensions line should contain 2 numbers: m and n")
	}

	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, err
	}
	if width < 1 {
		return 0, 0, errors.New("m is out of range")
	}

	length, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return 0, 0, err
	}
	if length < 1 {
		return 0, 0, errors.New("n is out of range")
	}

	return width, length, nil
}

func parseBitmap(m int, n int, index int, fileLines []string) (Bitmap, error) {
	bitmap := make(Bitmap, m)
	for i := range bitmap {
		bitmap[i] = make([]bool, n)
	}
	for i := index; i < index+m; i++ {
		row := strings.Split(fileLines[i], "")
		for j := 0; j < n; j++ {
			if row[j] == "0" {
				bitmap[i-index][j] = false
			} else if row[j] == "1" {
				bitmap[i-index][j] = true
			} else {
				return nil, errors.New("A bitmap can only contain ones and zeros")
			}
		}
	}
	return bitmap, nil
}
