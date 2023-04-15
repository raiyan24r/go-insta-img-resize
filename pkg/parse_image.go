package instaResize

import (
	"os"
)

func Parse(filePath string) *os.File {
	imageFile, _ := os.Open(filePath)
	return imageFile
}
