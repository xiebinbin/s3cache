package imagepro

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/fishtailstudio/imgo"
	"log"
)

func (i *ImagePro) Transform() *ImagePro {
	img := imgo.Load(i.image)
	if i.extension == "png" {
		i.Lossless = true
		i.Exact = true
	}
	if i.opts.Flip != nil {
		img = img.Flip(imgo.FlipType(i.opts.Flip.(int)))
	}
	if i.opts.Rotate != nil {
		img = img.Rotate(i.opts.Flip.(int))
	}
	if i.opts.Pixelate != nil {
		img = img.Pixelate(i.opts.Pixelate.(int))
	}
	switch i.opts.Mode {
	case "crop":
		img = img.Crop(i.opts.X, i.opts.Y, i.opts.Width, i.opts.Height)
		break
	case "resize":
		img = img.Resize(i.opts.Width, i.opts.Height)
		break
	case "thumbnail":
		img = img.Thumbnail(i.opts.Width, i.opts.Height)
		break
	default:
		break
	}
	if i.opts.Format == "webp" {
		var buff bytes.Buffer
		err := webp.Encode(&buff, img.ToImage(), &webp.Options{
			Lossless: i.Lossless,
			Exact:    i.Exact,
			Quality:  float32(i.opts.Quality.(int)),
		})
		if err != nil {
			log.Panic(err)
		}

		webpImg, err := webp.Decode(&buff)
		if err != nil {
			log.Println("webp编码失败")
			log.Panic(err)
		}
		i.image = webpImg
		i.extension = "webp"
	} else {
		i.image = img.ToImage()
	}
	return i
}
