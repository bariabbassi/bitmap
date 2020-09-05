# 👾 Bitmap Distances
Counting the distance to the nearest white pixel in a bitmap

## How to run the project
```
go get -u github.com/bariabbassi/bitmap-distances
cd $GOPATH/src/github.com/bariabbassi/bitmap-distances
go install
bitmap-distances
```

## What the project does
- Read a test file
```
1
3 4
0001
0011
0110
```
- Parse the number of bitmaps from the first line
- Parse dimensions and pixels of all bitmaps
- Create a slice of white pixels
- Create a map of black pixels (a map lets you delete a pixel without iterating)
- **Count close distances**: white pixels get 0 as a distance and their 4 neighbours get 1 if they're black
```
0010
0100
1001
```
- **Count far distances**: count the minimum distances for the remaining black pixels
```
3210
2100
1001
```

## Testing
All functions have been unit tested. To run all tests, use the command.
```
go test
```

