package loader

import (
    "os"
    "strings"
    "image"
    "image/jpeg"
)

func SaveImageJPEG(img image.Image, name string) (image.Image) {
    fout, errCreate := os.Create(strings.Join([]string{name, "jpg"}, "."))
    if errCreate != nil {
        return nil
    }
    defer fout.Close()

	if jpeg.Encode(fout, img, &jpeg.Options{jpeg.DefaultQuality}) != nil {
		return nil
	}
    
    return img
}