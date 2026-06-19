package themes

import (
	"image"
	"os"
	"path"
)

func LoadImage(s string) image.Image {
	imgPath := path.Join(ThemePath, s)
	infile, err := os.Open(imgPath)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, _, err := image.Decode(infile)
	if err != nil {
		panic(err)
	}
	return src
}
