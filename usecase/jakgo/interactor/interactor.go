package interactor

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"my-github/api-jakgo/domain/model"
	"my-github/api-jakgo/usecase/jakgo/presenter"
	"net/http"
)

type JakGoImpl struct {
	presenter presenter.JakGo
}

type Option func(j *JakGoImpl)

type JakGoService interface {
	JakGoRSU(ctx context.Context, link, auth string) ([]*model.RSU, error)
	JakGoKelurahan(ctx context.Context, link, auth string) ([]*model.Kelurahan, error)
}

// func LinkAPI(rsu, kec, auth string) Option {
// 	return func(j *JakGoImpl) {
// 		j.apiRSU = rsu
// 		j.apiKecamatan = kec
// 		j.authorization = auth
// 	}
// }

func NewJakGoService(p presenter.JakGo, options ...Option) JakGoService {
	jak := &JakGoImpl{
		presenter: p,
	}

	for _, o := range options {
		o(jak)
	}

	return jak
}

func (j *JakGoImpl) JakGoRSU(ctx context.Context, link, auth string) (rsu []*model.RSU, err error) {

	cli := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &rsu)
	if err != nil {
		return nil, err
	}

	return
}

func (j *JakGoImpl) JakGoKelurahan(ctx context.Context, link, auth string) (kel []*model.Kelurahan, err error) {
	cli := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", auth)

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &kel)
	if err != nil {
		return nil, err
	}

	return
}
