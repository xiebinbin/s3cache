package imagepro

import (
	"image"
)

type Options struct {
	Format   string
	Flip     interface{}
	Rotate   interface{}
	X        int
	Y        int
	Mode     string
	Width    int
	Pixelate interface{}
	Height   int
	Quality  interface{}
}
type ImagePro struct {
	image     image.Image
	Lossless  bool
	Exact     bool
	extension string
	opts      *Options
}
