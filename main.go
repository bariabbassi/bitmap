package main

import (
	"fmt"
	"log"
)

const maxBitmapsNumber = 1000
const maxDimension = 182

func main() {
	// Read file
	fileLines, err := readFile("./testfiles/big_file.txt") // there are other files are in ./testfiles
	if err != nil {
		log.Fatal(err)
	}

	// Parse the number of bitmaps
	bitmapsNumber, err := parseBitmapsNumber(fileLines[0])
	if err != nil {
		log.Fatal(err)
	}
	bitmaps := make([]Bitmap, bitmapsNumber)

	// Parse all bitmaps
	index := 1        // index of the 1st line of a bitmap. it jumps by bitmap length
	bitmapsCount := 0 // the i in this while loop
	for index < len(fileLines) && bitmapsCount < bitmapsNumber {
		bitmaps[bitmapsCount], err = parseBitmap(index, fileLines)
		if err != nil {
			log.Fatal(err)
		}

		index += bitmaps[bitmapsCount].Length + 1
		bitmapsCount++
	}

	// Print bitmaps
	fmt.Println("# Bitmaps:")
	for _, bitmap := range bitmaps {
		fmt.Println(bitmap.String())
	}

	// Count distances
	for _, bitmap := range bitmaps {
		bitmap.CountCloseDistances()
		bitmap.CountFarDistances()
	}

	// Print distances
	fmt.Println("# Distances:")
	for _, bitmap := range bitmaps {
		fmt.Println(bitmap.Output())
	}
}
