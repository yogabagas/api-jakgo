package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	model "my-github/api-jakgo/domain/model"
	"net/http"
)

func JakGoRSU(ctx context.Context, url, auth string) (rsu model.RSUResponse, err error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		return rsu, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return rsu, err
	}

	json.Unmarshal([]byte(body), &rsu)

	return
}

func JakGoKelurahan(ctx context.Context, url, auth string) (kel model.KelurahanResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		return kel, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return kel, err
	}

	json.Unmarshal([]byte(body), &kel)

	return
}

func JakGo(w http.ResponseWriter, r *http.Request) {
	rsu, err := JakGoRSU(context.Background(), "http://api.jakarta.go.id/v1/rumahsakitumum", "LdT23Q9rv8g9bVf8v/fQYsyIcuD14svaYL6Bi8f9uGhLBVlHA3ybTFjjqe+cQO8k")
	if err != nil {
		return
	}

	// fmt.Println(rsu)
	// fmt.Println("RSU", rsu)

	kel, err := JakGoKelurahan(context.Background(), "http://api.jakarta.go.id/v1/kelurahan", "LdT23Q9rv8g9bVf8v/fQYsyIcuD14svaYL6Bi8f9uGhLBVlHA3ybTFjjqe+cQO8k")
	if err != nil {
		return
	}

	// fmt.Println(kel)
	// log.Println(kel)

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

	err = json.NewEncoder(w).Encode(jaks)
	if err != nil {
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1", JakGo)
	http.ListenAndServe(":5000", mux)
	log.Println("tes..")
}
