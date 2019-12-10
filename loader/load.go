package loader

import (
	"os"
	"image"
	"io/ioutil"

	freetype "github.com/golang/freetype"
	truetype "github.com/golang/freetype/truetype"
)

func LoadImage(path string) (image.Image, error) {
    file, errOpen := os.Open(path)
    if errOpen != nil {
        return nil, errOpen
    }
    defer file.Close()

    img, _, errDecode := image.Decode(file)

    return img, errDecode
}

func LoadFont(path string) (*truetype.Font, error) {
    bytes, errRead := ioutil.ReadFile(path)
    if errRead != nil {
        return nil, errRead
    }

    font, errParse := freetype.ParseFont(bytes)

    return font, errParse
}