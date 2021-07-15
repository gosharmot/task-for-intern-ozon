package shortener

import (
	"github.com/speps/go-hashids/v2"
)

func GetShortUrl(longUrl string, urlLength int, urlID int) string {
	hd := hashids.NewData()
	hd.Salt = longUrl

	hd.MinLength = urlLength

	h, _ := hashids.NewWithData(hd)

	encodeUrl, _ := h.Encode([]int{urlID})

	return encodeUrl
}
