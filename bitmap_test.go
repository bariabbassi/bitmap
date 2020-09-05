package main

import (
	"testing"
)

/////////////////////// Count Close Distances ///////////////////////

func TestCountCloseDistances(t *testing.T) {
	expectedDistances := [][]int{
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
	}

	index := 1
	fileLines := []string{"1", "3 4", "0001", "0011", "0110"}
	bitmap, err := parseBitmap(index, fileLines)
	if err != nil {
		t.Error(err)
	}

	bitmap.CountCloseDistances()
	if len(bitmap.Distances) != len(expectedDistances) {
		t.Error("Bitmap distances length is not correct")
	}
	if len(bitmap.Distances[0]) != len(expectedDistances[0]) {
		t.Error("Bitmap distances width is not correct")
	}
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			if bitmap.Distances[i][j] != expectedDistances[i][j] {
				t.Error("Bitmap close distances not filled correctly")
			}
		}
	}
}

/////////////////////// Count Far Distances ///////////////////////

func TestCountFarDistances(t *testing.T) {
	expectedDistances := [][]int{
		{3, 2, 1, 0},
		{2, 1, 0, 0},
		{1, 0, 0, 1},
	}

	index := 1
	fileLines := []string{"1", "3 4", "0001", "0011", "0110"}
	bitmap, err := parseBitmap(index, fileLines)
	if err != nil {
		t.Error(err)
	}

	bitmap.CountFarDistances()
	if len(bitmap.Distances) != len(expectedDistances) {
		t.Error("Bitmap distances length is not correct")
	}
	if len(bitmap.Distances[0]) != len(expectedDistances[0]) {
		t.Error("Bitmap distances width is not correct")
	}
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			if bitmap.Distances[i][j] != expectedDistances[i][j] {
				t.Error("Bitmap far distances not filled correctly")
			}
		}
	}
}

/////////////////////// Count Close & Far Distances ///////////////////////

func TestCountCloseFarDistances(t *testing.T) {
	expectedCloseDistances := [][]int{
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
	}
	expectedFarDistances := [][]int{
		{3, 2, 1, 0},
		{2, 1, 0, 0},
		{1, 0, 0, 1},
	}

	index := 1
	fileLines := []string{"1", "3 4", "0001", "0011", "0110"}
	bitmap, err := parseBitmap(index, fileLines)
	if err != nil {
		t.Error(err)
	}

	bitmap.CountCloseDistances()
	if len(bitmap.Distances) != len(expectedCloseDistances) {
		t.Error("Bitmap close distances length is not correct")
	}
	if len(bitmap.Distances[0]) != len(expectedCloseDistances[0]) {
		t.Error("Bitmap close distances width is not correct")
	}
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			if bitmap.Distances[i][j] != expectedCloseDistances[i][j] {
				t.Error("Bitmap close distances not filled correctly")
			}
		}
	}
	bitmap.CountFarDistances()
	if len(bitmap.Distances) != len(expectedFarDistances) {
		t.Error("Bitmap far distances length is not correct")
	}
	if len(bitmap.Distances[0]) != len(expectedFarDistances[0]) {
		t.Error("Bitmap far distances width is not correct")
	}
	for i := 0; i < bitmap.Length; i++ {
		for j := 0; j < bitmap.Width; j++ {
			if bitmap.Distances[i][j] != expectedFarDistances[i][j] {
				t.Error("Bitmap far distances not filled correctly")
			}
		}
	}
}
