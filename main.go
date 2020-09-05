package main

import (
	"fmt"
	"log"
)

const maxBitmapsNumber = 1000
const maxDimension = 182

func main() {
	// Read file
	fileLines, err := readFile("./testfiles/default_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileLines)
	fmt.Println(len(fileLines))

	// Parse the number of bitmaps
	bitmapsNumber, err := parseBitmapsNumber(fileLines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitmapsNumber)
	fmt.Println()

	// Parse all bitmaps
	bitmaps := make([]Bitmap, bitmapsNumber)
	bitmapsCount := 0 // the i in this while loop
	index := 1        // index of the 1st line of a bitmap
	for index < len(fileLines) && bitmapsCount < bitmapsNumber {
		bitmaps[bitmapsCount], err = parseBitmap(index, fileLines)
		if err != nil {
			log.Fatal(err)
		}

		index += bitmaps[bitmapsCount].Length + 1
		bitmapsCount++
	}

	// Print Bitmap
	for _, bitmap := range bitmaps {
		fmt.Println(bitmap.String())
		fmt.Println(bitmap.WhitePixels)
		fmt.Println(bitmap.BlackPixels)
	}
}
