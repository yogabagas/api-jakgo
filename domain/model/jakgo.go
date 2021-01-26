package model

type JakGo struct {
	ID        int       `json:"id"`
	NamaRSU   string    `json:"nama_rsu"`
	JenisRSU  string    `json:"jenis_rsu"`
	Location  Location  `json:"location"`
	Alamat    string    `json:"alamat"`
	KodePos   string    `json:"kode_pos"`
	Telepon   []string  `json:"telepon"`
	Faximile  []string  `json:"faximile"`
	Website   string    `json:"website"`
	Email     string    `json:"email"`
	Kelurahan Kelurahan `json:"kelurahan"`
	Kecamatan Kecamatan `json:"kecamatan"`
	Kota      Kota      `json:"kota"`
}

type RSU struct {
	ID            int      `json:"id"`
	NamaRSU       string   `json:"nama_rsu"`
	JenisRSU      string   `json:"jenis_rsu"`
	Location      Location `json:"location"`
	Alamat        string   `json:"alamat"`
	KodePos       string   `json:"kode_pos"`
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

type Location struct {
	Alamat    string  `json:"alamat"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
