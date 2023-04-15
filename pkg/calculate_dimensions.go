package instaResize

import (
	"image"
)

func CalculateDimensions(img image.Image, aspectRatioType int) (int, int) {
	aspectRatioVertical := 0.8
	aspectRatioHorizontal := 1.91

	bounds := img.Bounds()

	imgWidth := float64(bounds.Dx())
	imgHeight := float64(bounds.Dy())
	reqWidth := int(imgWidth)
	reqHeight := int(imgHeight)

	// Logic based on current image's aspect ratio
	aspectRatio := aspectRatioVertical

	if aspectRatioType == 2 {
		aspectRatio = aspectRatioHorizontal
	}

	if imgHeight > imgWidth {
		reqWidth = int(imgHeight * aspectRatio)
	} else if aspectRatioType == 2 {
		reqWidth = int(imgHeight * aspectRatio)
		reqHeight = int(imgHeight)

	} else {
		reqHeight = int(imgWidth / aspectRatio)
	}

	return reqWidth, reqHeight
}
