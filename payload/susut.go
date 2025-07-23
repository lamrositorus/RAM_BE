package payload

type SusutTimbanganRequest struct {
	Tanggal      string  `json:"tanggal"`
	NomorPolisi  string  `json:"nomor_polisi"`
	NamaSupir    string  `json:"nama_supir"`
	SPPabrik     float64 `json:"sp_pabrik"`
	BuahPulangan float64 `json:"buah_pulangan"`
	SPRAM        float64 `json:"sp_ram"`
}