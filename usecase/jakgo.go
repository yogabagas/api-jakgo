package usecase

import (
	sv "my-github/api-jakgo/service"
	"net/url"
	"strings"
)

const (
	rsu = "/v1/rumahsakitumum"
	kel = "/v1/kelurahan"
)

type JakGoImpl struct {
	urlRSU, urlKel string
	auth           string
}

func NewJakGoService(auth string, urls ...string) sv.JakGoService {
	jkt := &JakGoImpl{
		auth: auth,
	}

	for _, u := range urls {

		p, _ := url.Parse(u)
		if strings.Contains(p.Path, kel) {
			jkt.urlKel = u
		}
		if strings.Contains(p.Path, rsu) {
			jkt.urlRSU = u
		}
	}
	return jkt
}
