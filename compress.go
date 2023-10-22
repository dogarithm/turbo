package turbo

/*
#cgo LDFLAGS: -lturbojpeg
#include <turbojpeg.h>
*/
import "C"

import (
	"image"
	"image/draw"
	"unsafe"
)

type CompressParams struct {
	PixelFormat PixelFormat
	Sampling    Sampling
	Quality     int // 1 .. 100
	Flags       Flags
}

func MakeCompressParams(pixelFormat PixelFormat, sampling Sampling, quality int, flags Flags) CompressParams {
	return CompressParams{
		PixelFormat: pixelFormat,
		Sampling:    sampling,
		Quality:     quality,
		Flags:       flags,
	}
}

func Compress(src image.Image, params CompressParams) ([]byte, error) {
	var width, height, stride int
	var pix []byte

	switch v := src.(type) {
	case *image.RGBA:
		width, height, stride, pix = v.Rect.Dx(), v.Rect.Dy(), v.Stride, v.Pix
	case *image.NRGBA:
		width, height, stride, pix = v.Rect.Dx(), v.Rect.Dy(), v.Stride, v.Pix
	case *image.Gray:
		width, height, stride, pix = v.Rect.Dx(), v.Rect.Dy(), v.Stride, v.Pix
	default:
		nv := image.NewRGBA(src.Bounds())
		draw.Draw(nv, nv.Bounds(), src, image.Point{}, draw.Src)
		width, height, stride, pix = nv.Rect.Dx(), nv.Rect.Dy(), nv.Stride, nv.Pix
	}

	encoder := C.tjInitCompress()
	defer C.tjDestroy(encoder)

	var outBuf *C.uchar
	var outBufSize C.ulong

	res := C.tjCompress2(encoder, (*C.uchar)(&pix[0]), C.int(width), C.int(stride), C.int(height), C.int(params.PixelFormat),
		&outBuf, &outBufSize, C.int(params.Sampling), C.int(params.Quality), C.int(params.Flags))

	err := makeError(encoder, res)

	if err != nil {
		C.tjFree(outBuf)
		return nil, err
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(outBuf)), outBufSize), nil
}
