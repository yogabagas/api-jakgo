package model

type JakGo struct {
	ID        int       `json:"id"`
	NamaRSU   string    `json:"nama_rsu"`
	JenisRSU  string    `json:"jenis_rsu"`
	Location  Location  `json:"location"`
	Alamat    string    `json:"alamat"`
	KodePos   int       `json:"kode_pos"`
	Telepon   []string  `json:"telepon"`
	Faximile  []string  `json:"faximile"`
	Website   string    `json:"website"`
	Email     string    `json:"email"`
	Kelurahan Kelurahan `json:"kelurahan"`
	Kecamatan Kecamatan `json:"kecamatan"`
	Kota      Kota      `json:"kota"`
}

type RSUResponse struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
	Data   []RSU  `json:"data"`
}

type RSU struct {
	ID            int      `json:"id"`
	NamaRSU       string   `json:"nama_rsu"`
	JenisRSU      string   `json:"jenis_rsu"`
	Location      Location `json:"location"`
	KodePos       int      `json:"kode_pos"`
	Telepon       []string `json:"telepon"`
	Faximile      []string `json:"faximile"`
	Website       string   `json:"website"`
	Email         string   `json:"email"`
	KodeKota      int      `json:"kode_kota"`
	KodeKecamatan int      `json:"kode_kecamatan"`
	KodeKelurahan int      `json:"kode_kelurahan"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
}

type Kota struct {
	KodeKota int    `json:"kode"`
	Nama     string `json:"nama"`
}

type Kecamatan struct {
	KodeKecamatan int    `json:"kode"`
	Nama          string `json:"nama"`
}

type Kelurahan struct {
	KodeKelurahan int    `json:"kode"`
	Nama          string `json:"nama"`
}

type KelurahanResponse struct {
	Status string          `json:"status"`
	Count  int             `json:"count"`
	Data   []KelurahanData `json:"data"`
}

type KelurahanData struct {
	KodeProvinsi  int    `json:"kode_provinsi"`
	NamaProvinsi  string `json:"nama_provinsi"`
	KodeKota      int    `json:"kode_kota"`
	NamaKota      string `json:"nama_kota"`
	KodeKecamatan int    `json:"kode_kecamatan"`
	NamaKecamatan string `json:"nama_kecamatan"`
	KodeKelurahan int    `json:"kode_kelurahan"`
	NamaKelurahan string `json:"nama_kelurahan"`
}

type Location struct {
	Alamat    string  `json:"alamat"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
