package render

import (
	"image"
	"image/draw"
)

// Context is used as a general wrapper for drawing methods
type Context struct {
	img draw.Image
}

// NewContext is used to create a new render context
func NewContext(bounds image.Rectangle) Context {
	return Context{image.NewRGBA(bounds)}
}

// DrawImage draws the given image starting from given destination
func (c *Context) DrawImage(src image.Image, origin image.Point) {
	bounds := c.img.Bounds()
	draw.Draw(c.img, bounds, src, origin.add(bounds.Min), draw.Over)
}

// DrawImageSection draws the given section of the image starting from given destination
func (c *Context) DrawImageSection(src image.Image, section image.Rectangle, origin image.Point) {
	bounds := c.img.Bounds()
	draw.Draw(c.img, bounds, src.SubImage(section), origin.add(bounds.Min), draw.Over)
}

// Render returns the current image
func (c Context) Render() image.Image {
	return img
}