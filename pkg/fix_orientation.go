package instaResize

import (
	"fmt"
	imagego "image"
	"os"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

func FixOrientation(imgFile *os.File) imagego.Image {
	f := imgFile

	x, err := exif.Decode(f)

	if err != nil {
		fmt.Println(err)
	}

	tag, err := x.Get(exif.Orientation)

	if err != nil {
		fmt.Println(err)
	}
	// reset for image decode
	off, err := f.Seek(0, 0)

	if err != nil {
		fmt.Println(err)
	}

	if off != 0 {

	}

	img, _, err := imagego.Decode(f)

	if err != nil {
		fmt.Println(err)
	}

	if tag.Count == 1 && tag.Format() == tiff.IntVal {
		orientation, _ := tag.Int(0)

		switch orientation {
		case 3: // rotate 180
			img = imaging.Rotate180(img)
		case 6: // rotate 270
			img = imaging.Rotate270(img)
		case 8: //rotate 90
			img = imaging.Rotate90(img)
		}
	}

	return img
}
