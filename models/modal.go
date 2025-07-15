package models

import (
	"database/sql"
)

type EstimasiModalRequest struct {
	ID            int `json:"id"`
	SisaModal     int `json:"sisa_modal"`
	TonasiGantung int `json:"tonasi_gantung"`
	HargaTbs      int `json:"harga_tbs_ram"`
	TotalModal    int `json:"total_modal"` // Optional: bisa ditambahkan
}

func CreateModal(db *sql.DB, m *EstimasiModalRequest) error {
	return db.QueryRow(
		"INSERT INTO modal (sisa_modal, tonasi_gantung, harga_tbs_ram) VALUES ($1, $2, $3) RETURNING id",
		m.SisaModal, m.TonasiGantung, m.HargaTbs,
	).Scan(&m.ID)
}

func GetModal(db *sql.DB) ([]EstimasiModalRequest, error) {
	rows, err := db.Query("SELECT id, sisa_modal, tonasi_gantung, harga_tbs_ram, total_modal FROM modal ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modal []EstimasiModalRequest
	for rows.Next() {
		var m EstimasiModalRequest
		if err := rows.Scan(&m.ID, &m.SisaModal, &m.TonasiGantung, &m.HargaTbs, &m.TotalModal); err != nil {
			return nil, err
		}
		modal = append(modal, m)
	}
	return modal, nil
}
