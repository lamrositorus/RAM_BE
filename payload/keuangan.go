package payload

type KeuanganRequest struct {
	Deskripsi string `json:"deskripsi"`
	Nominal   int    `json:"nominal"`
	Tipe      string `json:"tipe"`
}