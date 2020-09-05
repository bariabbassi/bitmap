package main

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
