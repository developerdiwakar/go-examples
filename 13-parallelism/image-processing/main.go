package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

func processBlock(img *image.RGBA, startX, startY, endX, endY int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Process the image block here
	for x := startX; x < endX; x++ {
		for y := startY; y < endY; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Apply your image processing logic here
			// For example, grayscale conversion:
			gray := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
			img.Set(x, y, color.Gray{gray})
		}
	}
}

func main() {
	// Load image
	file, err := os.Open("lion.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	rgbaImg, ok := img.(*image.RGBA) // Checking image color type using type assertion
	if !ok {
		// If img is not an *image.RGBA, try converting it
		bounds = img.Bounds()
		rgbaImg = image.NewRGBA(bounds)
		draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)
	}

	// Divide image into blocks
	blockWidth := 100
	blockHeight := 100
	numBlocksX := (bounds.Max.X - bounds.Min.X) / blockWidth
	numBlocksY := (bounds.Max.Y - bounds.Min.Y) / blockHeight

	var wg sync.WaitGroup
	wg.Add(numBlocksX * numBlocksY)

	// Process blocks concurrently
	for i := 0; i < numBlocksX; i++ {
		for j := 0; j < numBlocksY; j++ {
			startX := bounds.Min.X + i*blockWidth
			startY := bounds.Min.Y + j*blockHeight
			endX := startX + blockWidth
			endY := startY + blockHeight
			go processBlock(rgbaImg, startX, startY, endX, endY, &wg)
		}
	}

	wg.Wait()

	// Save processed image
	outFile, err := os.Create("shera.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	// Encode writes the Image img to outFile in PNG format. Any Image may be encoded,
	// but images that are not image.NRGBA might be encoded lossily.
	if err = png.Encode(outFile, img); err != nil {
		panic(err)
	}
}
