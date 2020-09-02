package main

// Bitmap TODO
type Bitmap [][]bool

func (bitmap Bitmap) String() string {
	output := ""
	for i := 0; i < len(bitmap); i++ {
		for j := 0; j < len(bitmap[0]); j++ {
			if bitmap[i][j] {
				output += "1 "
			} else {
				output += "0 "
			}
		}
		output += "\n"
	}
	return output
}
