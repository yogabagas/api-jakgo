package controller

import (
	"context"
	"encoding/json"
	"my-github/api-jakgo/domain/model"
	"my-github/api-jakgo/service"
	"net/http"
)

type JakGoImpl struct {
	jakGo service.JakGoService
}

type JakGoController interface {
	JakGo(w http.ResponseWriter, r *http.Request)
}

func NewJakGoController(sv service.JakGoService) JakGoController {
	return &JakGoImpl{
		jakGo: sv,
	}
}

func (j *JakGoImpl) JakGo(w http.ResponseWriter, r *http.Request) {

	var jr model.JakGoResponse

	rsu, err := j.jakGo.JakGoRSU(context.Background())
	if err != nil {
		jr.Status = "error"
	}

	kel, err := j.jakGo.JakGoKelurahan(context.Background())
	if err != nil {
		jr.Status = "error"
	}

	namaKec := make(map[int]string)
	namaKel := make(map[int]string)
	namaKota := make(map[int]string)
	for _, k := range kel.Data {
		namaKel[k.KodeKelurahan] = k.NamaKelurahan
		namaKec[k.KodeKecamatan] = k.NamaKecamatan
		namaKota[k.KodeKota] = k.NamaKota
	}

	var jaks []model.JakGo
	var jak model.JakGo
	for _, r := range rsu.Data {
		jak.ID = r.ID
		jak.NamaRSU = r.NamaRSU
		jak.JenisRSU = r.JenisRSU
		jak.Location.Latitude = r.Location.Latitude
		jak.Location.Longitude = r.Location.Longitude
		jak.Alamat = r.Location.Alamat
		jak.KodePos = r.KodePos
		jak.Telepon = r.Telepon
		jak.Faximile = r.Faximile
		jak.Website = r.Website
		jak.Email = r.Email
		jak.Kelurahan.KodeKelurahan = r.KodeKelurahan
		jak.Kecamatan.KodeKecamatan = r.KodeKecamatan
		jak.Kota.KodeKota = r.KodeKota

		if v, ok := namaKel[r.KodeKelurahan]; ok {
			jak.Kelurahan.Nama = v
		}
		if w, ok := namaKec[r.KodeKecamatan]; ok {
			jak.Kecamatan.Nama = w
		}
		if x, ok := namaKota[r.KodeKota]; ok {
			jak.Kota.Nama = x
		}

		jaks = append(jaks, jak)
	}

	jr.Status = "success"
	jr.Count = len(jaks)
	jr.Data = jaks

	err = json.NewEncoder(w).Encode(jr)
	if err != nil {
		return
	}
}
