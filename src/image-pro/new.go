package imagepro

import (
	"github.com/fishtailstudio/imgo"
)

func New(src string, opts *Options) *ImagePro {
	i := &ImagePro{}
	i.Lossless = false
	i.Exact = false
	img := imgo.Load(src)
	i.image = img.ToImage()
	if opts.Quality == nil {
		opts.Quality = 85
	}
	i.extension = img.Extension()
	i.opts = opts
	return i
}
