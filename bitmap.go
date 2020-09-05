package main

import "fmt"

// Bitmap TODO
type Bitmap struct {
	Width       int
	Length      int
	Pixels      [][]bool
	WhitePixels [][2]int
	BlackPixels map[string]bool
	Distances   [][]int
}

func (bitmap *Bitmap) String() string {
	output := ""
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			if bitmap.Pixels[i][j] {
				output += "1 "
			} else {
				output += "0 "
			}
		}
		output += "\n"
	}
	return output
}

// Output shows distances in a readable format
func (bitmap *Bitmap) Output() string {
	output := ""
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			output += fmt.Sprintf("%d ", bitmap.Distances[i][j])
		}
		output += "\n"
	}
	return output
}

// CountCloseDistances fills 0 and 1 distances on and around white pixels
func (bitmap *Bitmap) CountCloseDistances() {
	for _, pixel := range bitmap.WhitePixels {
		i := pixel[0]
		j := pixel[1]

		// the 4 neighbours of a white pixel have a distance of 1
		if j != 0 {
			bitmap.Distances[i][j-1] = 1
			delete(bitmap.BlackPixels, fmt.Sprintf("%d %d", i, j-1))
		}
		if j != bitmap.Width-1 {
			bitmap.Distances[i][j+1] = 1
			delete(bitmap.BlackPixels, fmt.Sprintf("%d %d", i, j+1))
		}
		if i != 0 {
			bitmap.Distances[i-1][j] = 1
			delete(bitmap.BlackPixels, fmt.Sprintf("%d %d", i-1, j))
		}
		if i != bitmap.Length-1 {
			bitmap.Distances[i+1][j] = 1
			delete(bitmap.BlackPixels, fmt.Sprintf("%d %d", i+1, j))
		}
	}
	// the distance of a white pixel is 0
	for _, pixel := range bitmap.WhitePixels {
		bitmap.Distances[pixel[0]][pixel[1]] = 0
	}
}

// CountFarDistances counts the distances of remaining black pixels not close to a white pixel
func (bitmap *Bitmap) CountFarDistances() {
	for key := range bitmap.BlackPixels {
		iBlack, jBlack, _ := parseBlackPixel(key)

		// just a big value that will be over writen by the 1st distance
		minDistance := maxDimension * 2
		for _, pixel := range bitmap.WhitePixels {
			iWhite := pixel[0]
			jWhite := pixel[1]

			// there isn't an abs function in golang
			iDistance := iBlack - iWhite
			if iDistance < 0 {
				iDistance = -iDistance
			}
			jDistance := jBlack - jWhite
			if jDistance < 0 {
				jDistance = -jDistance
			}
			distance := iDistance + jDistance

			// we pick the closest distance
			if distance < minDistance {
				minDistance = distance
			}
		}
		bitmap.Distances[iBlack][jBlack] = minDistance
	}
}
