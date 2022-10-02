package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
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

	//get input for number of years
	var years int
	fmt.Print("Enter a number of years: ")
	fmt.Scanln(&years)

	points := 0
	// for every day in the number of years subtract the rate from the total and for every time that 400 is subtracted from the total incriment points
	for i := 0; i < years*365; i++ {
		total -= rate
		if total%400 == 0 {
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

	fmt.Println("Done!")
}
