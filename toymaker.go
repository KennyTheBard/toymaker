package main

import (
	"image"
	"fmt"
	"image/draw"

	freetype "github.com/golang/freetype"

	loader "./loader"
)

func main() {
	context := freetype.NewContext()
	img, _ := loader.LoadImage("images/test.jpg")
	font, _ := loader.LoadFont("fonts/OldLondon.ttf")
	pt := freetype.Pt(0, 0)

	context.SetDst(img.(draw.Image))
	context.SetClip(image.Rect(0, 0, 100, 500))
	context.SetFont(font)
	context.SetFontSize(25)
	context.DrawString("This is a long ass text designed to test text wrapping", pt)

	loader.SaveImageJPEG(img, "output")
}
