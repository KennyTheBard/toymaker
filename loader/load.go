package loader

import (
	"os"
	"image/draw"
    "image"
	"io/ioutil"

	freetype "github.com/golang/freetype"
	truetype "github.com/golang/freetype/truetype"
)

// LoadImage return the image loaded from the file at the given path
func LoadImage(path string) (image.Image, error) {
    file, errOpen := os.Open(path)
    if errOpen != nil {
        return nil, errOpen
    }
    defer file.Close()

    tempImg, _, errDecode := image.Decode(file)
    if errDecode != nil {
        return tempImg, errDecode
    }

    bounds := tempImg.Bounds()
    img := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
    draw.Draw(img, img.Bounds(), tempImg, bounds.Min, draw.Src)
    return img, nil
}

// LoadFont return the font loaded from the file at the given path
func LoadFont(path string) (*truetype.Font, error) {
    bytes, errRead := ioutil.ReadFile(path)
    if errRead != nil {
        return nil, errRead
    }

    font, errParse := freetype.ParseFont(bytes)

    return font, errParse
}