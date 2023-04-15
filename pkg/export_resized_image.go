package instaResize

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func ExportResizedImage(resizedImage *image.RGBA, directory string) {
	file, _ := os.Create(directory)
	fmt.Println(directory)
	png.Encode(file, resizedImage)
}
