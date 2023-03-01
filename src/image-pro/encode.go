package imagepro

import (
	"bytes"
	"github.com/chai2010/webp"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image/jpeg"
	"image/png"
	"log"
)

func (i *ImagePro) Encode() ([]byte, error) {
	extension := i.extension
	var buff bytes.Buffer
	var err error
	if extension == "png" {
		err = png.Encode(&buff, i.image)
	} else if extension == "jpg" || extension == "jpeg" {
		quality := 0
		if i.opts.Quality != nil {
			quality = i.opts.Quality.(int)
		}
		if quality > 0 {
			err = jpeg.Encode(&buff, i.image, &jpeg.Options{Quality: quality})
		} else {
			err = jpeg.Encode(&buff, i.image, &jpeg.Options{Quality: 100})
		}
	} else if extension == "tiff" {
		err = tiff.Encode(&buff, i.image, &tiff.Options{Compression: tiff.Deflate, Predictor: true})
	} else if extension == "bmp" {
		err = bmp.Encode(&buff, i.image)
	} else if extension == "webp" {
		err := webp.Encode(&buff, i.image, &webp.Options{
			Quality:  float32(i.opts.Quality.(int)),
			Lossless: i.Lossless,
			Exact:    i.Exact,
		})
		if err != nil {
			log.Panic(err)
		}
	}
	return buff.Bytes(), err
}
