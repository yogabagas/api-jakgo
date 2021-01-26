package controller

import (
	"context"
	"my-github/api-jakgo/domain/model"
	"my-github/api-jakgo/usecase/jakgo/interactor"
)

type JakGoImpl struct {
	jakGoInteractor interactor.JakGoService
	url, auth       string
}

type JakGoController interface {
	GetJakGoDetail(ctx context.Context) ([]*model.JakGo, error)
}

func NewJakGoController(jak interactor.JakGoService, url, auth string) JakGoController {
	return &JakGoImpl{
		jakGoInteractor: jak,
		url:             url,
		auth:            auth,
	}
}

func (j *JakGoImpl) GetJakGoDetail(ctx context.Context) (jakgo []*model.JakGo, err error) {

	var jak model.JakGo

	rsu, err := j.jakGoInteractor.JakGoRSU(ctx, j.url, j.auth)
	if err != nil {
		return nil, err
	}

	kel, err := j.jakGoInteractor.JakGoKelurahan(ctx, j.url, j.auth)
	if err != nil {
		return nil, err
	}


	for _, r := range rsu {
		jak.ID = r.ID
		jak.NamaRSU = r.NamaRSU
		jak.JenisRSU = r.JenisRSU
		jak.Location.Latitude = r.Location.Latitude
		jak.Location.Longitude=r.Location.Longitude
		jak.Alamat = r.Alamat
		jak.KodePos = r.KodePos
		jak.Telepon = r.Telepon
		jak.Faximile = r.Faximile
		jak.Website = r.Website
		jak.Email = r.Email
	}

	jakgo = []*model.JakGo {
		&model.JakGo{

		}
	}
}
