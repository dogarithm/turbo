package turbo

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MakeRGBA(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	g := byte(0)
	b := byte(0)
	p := 0
	a := byte(255)
	for y := 0; y < height; y++ {
		r := byte(0)
		for x := 0; x < width; x++ {
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: a})
			r += 3
			b += 5
			p += 4
		}
		g += 1
	}
	return img
}

func TestCompress(t *testing.T) {
	w := 300
	h := 200
	raw1 := MakeRGBA(w, h)
	params := MakeCompressParams(PixelFormatRGBA, Sampling444, 90, 0)
	jpg, err := Compress(raw1, params)
	t.Logf("Encode return: %v, %v", len(jpg), err)
	raw2, err := Decompress(jpg)
	t.Logf("Decode return: %v x %v, %v, %v, %v", raw2.Bounds().Dx(), raw2.Bounds().Dy(), raw2.Stride, len(raw2.Pix), err)
	assert.Equal(t, w, raw2.Bounds().Dx(), "Width same")
	assert.Equal(t, h, raw2.Bounds().Dy(), "Height same")
	assert.Equal(t, raw1.Stride, raw2.Stride, "Stride same")
}
