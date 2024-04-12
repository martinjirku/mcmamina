package models

type Sponsor struct {
	Img string
	Url string
}

func NewSponsor(url, img string) Sponsor {
	return Sponsor{img, url}
}
