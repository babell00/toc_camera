package camera

import (
	"image"
)

type Camera struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	MJpegUrl    string `json:"-"`
	PublicImageUrl string `json:"public_image_url"`
	PublicStreamUrl string `json:"-"`
	UrlPath   string `json:"url_path"`
	Image  image.Image `json:"-"`
	Online bool `json:"online"`
}
