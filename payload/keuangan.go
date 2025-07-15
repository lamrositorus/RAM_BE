package payload

type KeuanganRequest struct {
	Tanggal   string `json:"tanggal"`
	Deskripsi string `json:"deskripsi"`
	Nominal   int    `json:"nominal"`
	Tipe      string `json:"tipe"`
}
