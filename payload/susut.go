package payload

type SusutTimbanganRequest struct {
	NomorPolisi    string  `json:"nomor_polisi"`
	NamaSupir      string  `json:"nama_supir"`
	JamMasukPabrik string  `json:"jam_masuk_pabrik"`
	SPPabrik       float64 `json:"sp_pabrik"`
	BuahPulangan   float64 `json:"buah_pulangan"`
	JamKeluarRAM   string  `json:"jam_keluar_ram"`
	SPRAM          float64 `json:"sp_ram"`
}
