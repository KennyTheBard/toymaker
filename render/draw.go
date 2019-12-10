package render

import (
	"image"
	"image/draw"
)

// RenderContext is used as a general wrapper for drawing methods
type RenderContext struct {
	img draw.Image
}

// NewRenderContext is used to create a new render context
func NewRenderContext(bounds image.Rectangle) RenderContext {
	return RenderContext{image.NewRGBA(bounds)}
}

// DrawImage draws the given image starting from given destination
func (c *RenderContext) DrawImage(src image.Image, origin image.Point) {
	bounds := c.img.Bounds()
	draw.Draw(c.img, bounds, src, origin.add(bounds.Min), draw.Over)
}

// DrawImageSection draws the given section of the image starting from given destination
func (c *RenderContext) DrawImageSection(src image.Image, section image.Rectangle, origin image.Point) {
	bounds := c.img.Bounds()
	draw.Draw(c.img, bounds, src.SubImage(section), origin.add(bounds.Min), draw.Over)
}

// Render returns the current image
func (c RenderContext) Render() image.Image {
	return img
}