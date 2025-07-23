package payload

type EstimasiKeuntunganRequest struct {
	SpCairPabrik   int `json:"sp_cair_pabrik"`
	HargaTbsPabrik int `json:"harga_tbs_pabrik"`
	TonasiSpRam    int `json:"tonasi_sp_ram"`
	HargaTbsBeliRam int `json:"harga_tbs_beli_ram"`
}