package main

import (
	"fmt"
	"log"
)

func main() {
	// Read file
	fileLines, err := readFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileLines)
	fmt.Println(len(fileLines))

	bitmapsNumber, err := parseBitmapsNumber(fileLines[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bitmapsNumber)
	fmt.Println()

	bitmaps := make([]Bitmap, bitmapsNumber)
	bitmapsCount := 0
	dimensionsIndex := 1
	for dimensionsIndex < len(fileLines) && bitmapsCount < bitmapsNumber {
		// fmt.Println(dimensionsIndex)
		// fmt.Println(fileLines[dimensionsIndex])
		m, n, err := parseDimensions(fileLines[dimensionsIndex])
		if err != nil {
			log.Fatal(err)
		}
		bitmapIndex := dimensionsIndex + 1
		// fmt.Println("-", m, n, bitmapIndex)
		bitmaps[bitmapsCount], _ = parseBitmap(m, n, bitmapIndex, fileLines)

		dimensionsIndex += m + 1
		bitmapsCount++
	}
	for i := 0; i < bitmapsNumber; i++ {
		fmt.Println(bitmaps[i])
	}
}
