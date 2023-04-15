package instaResize

import (
	"image"
	"image/color"
	"image/draw"
)

func CreateWhiteBg(reqWidth, reqHeight int) *image.Alpha {

	// drawing the white background
	upLeft := image.Point{0, 0}
	lowRight := image.Point{reqWidth, reqHeight}
	bgColor := color.RGBA{100, 200, 200, 0xff}
	blankWhite := image.NewAlpha(image.Rectangle{upLeft, lowRight})

	draw.Draw(blankWhite, blankWhite.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	return blankWhite
}
