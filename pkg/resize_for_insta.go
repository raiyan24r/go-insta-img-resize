package instaResize

import (
	"image"
	"image/draw"
)

func ResizeForInsta(blankWhite *image.Alpha, imgToResize image.Image, aspectRatioType int) *image.RGBA {

	sp2 := image.Point{(blankWhite.Bounds().Dx() - imgToResize.Bounds().Dx()) / 2, 0}

	if imgToResize.Bounds().Dx() > imgToResize.Bounds().Dy() && aspectRatioType != 2 {
		sp2 = image.Point{0, (blankWhite.Bounds().Dy() - imgToResize.Bounds().Dy()) / 2}
	}

	if aspectRatioType == 2 {

	}

	//new rectangle for the second image for bounds
	r2 := image.Rectangle{sp2, sp2.Add(imgToResize.Bounds().Size())}

	r := image.Rectangle{image.Point{0, 0}, blankWhite.Bounds().Max}
	rgba := image.NewRGBA(r)

	draw.Draw(rgba, blankWhite.Bounds(), blankWhite, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, imgToResize, image.Point{0, 0}, draw.Src)

	return rgba
}
