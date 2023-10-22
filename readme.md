# Turbo (TurboJPEG Go wrapper)

This is a very thin wrapper around turbojpeg.

Why not use https://github.com/pixiv/go-libjpeg ?

It's buggy and generates invalid output (probably due to incorrect use of more complex API).

### How to use

```go
import "github.com/bmharper/turbo"

func compressImage(width, height int, rgba []byte) {
	raw := turbo.Image{
		Width: width,
		Height: height,
		Stride: width * 4,
		RGBA: rgba,
	}
	params := turbo.MakeCompressParams(turbo.PixelFormatRGBA, turbo.Sampling420, 35, 0)
	jpg, err := turbo.Compress(&raw, params)
}

func decompressImage(jpg []byte) (*Image, error) {
	return turbo.Decompress(jpg)
}
```
