package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	instaResize "github.com/raiyan24r/go-insta-img-resize/pkg"
)

func main() {

	fmt.Println("Enter your full folder directory:")
	var directory string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	directory = scanner.Text()

	fmt.Println("Enter image orientation number. \n1 for vertical\n2 for horizontal:")

	// scanner.Scan() // use `for scanner.Scan()` to keep reading
	typeString := scanner.Text()
	fmt.Scan(&typeString)

	aspectRatioType, err := strconv.Atoi(typeString)

	if err != nil {
		fmt.Println("Error with orientation number")
		return
	}

	if (aspectRatioType != 1) && (aspectRatioType != 2) {
		aspectRatioType = 1
	}
	// aspectRatioType := 1
	// directory := `G:\Goa 23\upload\dd`
	file, _ := os.Open(directory)

	defer file.Close()
	fileNames, _ := file.Readdirnames(0)

	for _, fileName := range fileNames {
		filePath := filepath.Join(directory, fileName)

		fileInfo, _ := os.Stat(filePath)
		if fileInfo.IsDir() {
			continue
		}

		fmt.Println(fileName)
		os.Mkdir(filepath.Join(filepath.Dir(filePath), "insta-resized"), 007)
		exportFileName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".png"
		outputDirectory := filepath.Join(filepath.Dir(filePath), "insta-resized", exportFileName)

		existingImageFile := instaResize.Parse(filePath)
		defer existingImageFile.Close()

		img := instaResize.FixOrientation(existingImageFile)

		reqWidth, reqHeight := instaResize.CalculateDimensions(img, aspectRatioType)

		blankWhite := instaResize.CreateWhiteBg(reqWidth, reqHeight)

		resizedImg := instaResize.ResizeForInsta(blankWhite, img, aspectRatioType)

		instaResize.ExportResizedImage(resizedImg, outputDirectory)
	}
}
