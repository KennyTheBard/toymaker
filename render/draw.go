package render

import (
	"image"
	"image/draw"
)

type RenderContext struct {
	img draw.Image
}

func (c *RenderContext) DrawImage(src image.Image, origin image.Point) {
	bounds := c.img.Bounds()
	draw.Draw(c.img, bounds, src, origin.add(bounds.Min), draw.Over)
}

func (c *RenderContext) DrawImageSection(src image.Image, section image.Rectangle, origin image.Point) {
	bounds := c.img.Bounds()
	draw.Draw(c.img, bounds, src.SubImage(section), origin.add(bounds.Min), draw.Over)
}

func (c RenderContext) Render() image.Image {
	return img
}