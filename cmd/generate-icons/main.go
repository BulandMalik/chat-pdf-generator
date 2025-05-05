package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"path/filepath"
)

func main() {
	// Create images directory if it doesn't exist
	os.MkdirAll("../../images", 0755)

	// Generate check icon (green checkmark)
	generateCheckIcon()

	// Generate close icon (red X)
	generateCloseIcon()

	// Generate book icon (blue book)
	generateBookIcon()

	// Generate target icon (orange target)
	generateTargetIcon()

	// Generate rocket icon (purple rocket)
	generateRocketIcon()

	fmt.Println("Icons generated successfully in the images directory")
}

func generateCheckIcon() {
	// Create a 32x32 image with transparent background
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with transparent color
	transparent := color.RGBA{0, 0, 0, 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{transparent}, image.Point{}, draw.Src)

	// Draw a green checkmark
	green := color.RGBA{0, 255, 0, 255}

	// Draw the checkmark
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			// Checkmark shape
			if (i >= 5 && i <= 12 && j >= 8 && j <= 14) || // Vertical part
				(i >= 8 && i <= 24 && j >= 14 && j <= 20) { // Diagonal part
				img.Set(i, j, green)
			}
		}
	}

	// Save the image
	saveImage(img, "check.png")
}

func generateCloseIcon() {
	// Create a 32x32 image with transparent background
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with transparent color
	transparent := color.RGBA{0, 0, 0, 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{transparent}, image.Point{}, draw.Src)

	// Draw a red X
	red := color.RGBA{255, 0, 0, 255}

	// Draw the X
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			// X shape
			if (i >= 5 && i <= 12 && j >= 5 && j <= 12) || // Top-left to bottom-right
				(i >= 20 && i <= 27 && j >= 5 && j <= 12) || // Top-right to bottom-left
				(i >= 5 && i <= 12 && j >= 20 && j <= 27) ||
				(i >= 20 && i <= 27 && j >= 20 && j <= 27) {
				img.Set(i, j, red)
			}
		}
	}

	// Save the image
	saveImage(img, "close.png")
}

func generateBookIcon() {
	// Create a 32x32 image with transparent background
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with transparent color
	transparent := color.RGBA{0, 0, 0, 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{transparent}, image.Point{}, draw.Src)

	// Draw a blue book
	blue := color.RGBA{0, 0, 255, 255}

	// Draw the book
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			// Book shape
			if i >= 5 && i <= 27 && j >= 5 && j <= 27 {
				img.Set(i, j, blue)
			}
		}
	}

	// Save the image
	saveImage(img, "book.png")
}

func generateTargetIcon() {
	// Create a 32x32 image with transparent background
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with transparent color
	transparent := color.RGBA{0, 0, 0, 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{transparent}, image.Point{}, draw.Src)

	// Draw an orange target
	orange := color.RGBA{255, 165, 0, 255}

	// Draw the target
	centerX, centerY := 16, 16
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			// Calculate distance from center
			dx := float64(i - centerX)
			dy := float64(j - centerY)
			distance := math.Sqrt(dx*dx + dy*dy)

			// Draw concentric circles
			if distance <= 12 && distance >= 8 {
				img.Set(i, j, orange)
			}
		}
	}

	// Save the image
	saveImage(img, "target.png")
}

func generateRocketIcon() {
	// Create a 32x32 image with transparent background
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	// Fill with transparent color
	transparent := color.RGBA{0, 0, 0, 0}
	draw.Draw(img, img.Bounds(), &image.Uniform{transparent}, image.Point{}, draw.Src)

	// Draw a purple rocket
	purple := color.RGBA{128, 0, 128, 255}

	// Draw the rocket
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			// Rocket shape
			if i >= 8 && i <= 24 && j >= 5 && j <= 27 {
				img.Set(i, j, purple)
			}
		}
	}

	// Save the image
	saveImage(img, "rocket.png")
}

func saveImage(img image.Image, filename string) {
	f, err := os.Create(filepath.Join("../../images", filename))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
