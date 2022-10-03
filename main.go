package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"time"
	"strconv"
)

func main() {
	// create a 400x400 2d array and initialize all values to 0
	var arr [400][400]int

	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			arr[i][j] = 0
		}
	}

	total := 1_370_000_000
	rate := 940

	var years int

	//get input for number of years from argv as type int
	if len(os.Args) > 1 {
		years, _ = strconv.Atoi(os.Args[1])
	} else {
		years = 1
	}

	if years == -1 {
		os.Exit(0)
	}

	start := time.Now()

	points := 0
	// for every day in the number of years subtract the rate from the total and for every time that 8562 is subtracted from the total incriment points
	for i := 0; i < years*365; i++ {
		total -= rate
		if total%8562 == 0 {
			points++
		}
	}

	// for every point in the points variable add 1 to the 2d array at the next location
	for i := 0; i < points; i++ {
		arr[i%400][i/400] = 1
	}

	// create basic 400x400 image
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))

	//set pixels of image on 2d array where 1 is white and 0 is black and edit image accordingly
	for i := 0; i < 400; i++ {
		for j := 0; j < 400; j++ {
			if arr[i][j] == 1 {
				img.Set(i, j, color.White)
			} else {
				img.Set(i, j, color.Black)
			}
		}
	}

	// create file
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	// encode image to file
	png.Encode(f, img)

	fmt.Println("Done! : ", time.Since(start))
}
