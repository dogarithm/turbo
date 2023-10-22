package turbo

import (
	"image"
	"image/draw"
)

// Convert a Go image.Image into a turbo.Image
// If allowDeepClone is true, and the source image is type NRGBA or RGBA,
// then the resulting Image points directly to the pixel buffer of the source image.
func FromImage(src image.Image, allowDeepClone bool) *Image {
	dst := &Image{}
	switch v := src.(type) {
	case *image.RGBA:
		dst.Width, dst.Height, dst.Stride = v.Rect.Dx(), v.Rect.Dx(), v.Stride
		dst.Stride = v.Stride
		if allowDeepClone {
			dst.Pixels = v.Pix
		} else {
			dst.Pixels = make([]byte, dst.Stride*dst.Height)
			copy(dst.Pixels, v.Pix)
		}
		return dst
	case *image.NRGBA:
		dst.Width, dst.Height, dst.Stride = v.Rect.Dx(), v.Rect.Dx(), v.Stride
		if allowDeepClone {
			dst.Pixels = v.Pix
		} else {
			dst.Pixels = make([]byte, dst.Stride*dst.Height)
			copy(dst.Pixels, v.Pix)
		}
		return dst
	case *image.Gray:
		dst.Width, dst.Height, dst.Stride = v.Rect.Dx(), v.Rect.Dx(), v.Stride
		if allowDeepClone {
			dst.Pixels = v.Pix
		} else {
			dst.Pixels = make([]byte, dst.Stride*dst.Height)
			copy(dst.Pixels, v.Pix)
		}
		return dst
	default:
		nv := image.NewRGBA(src.Bounds())
		draw.Draw(nv, nv.Bounds(), src, image.Point{}, draw.Src)
		dst.Width, dst.Height, dst.Stride = nv.Rect.Dx(), nv.Rect.Dx(), nv.Stride
	}

	return dst
}
