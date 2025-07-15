package payload

type EstimasiModalRequest struct {
	SisaModal     int `json:"sisa_modal"`
	TonasiGantung int `json:"tonasi_gantung"`
	HargaTbs      int `json:"harga_tbs_ram"`
}
