package models

import "database/sql"

type EstimasiModalRequest struct {
	ID         int `json:"id"`
	SisaModal  int `json:"sisa_modal"`
	TonasiGantung int `json:"tonasi_gantung"`
	HargaTbs   int `json:"harga_tbs_ram"`
	TotalModal int `json:"total_modal"`
}

func CreateModal(db *sql.DB, m *EstimasiModalRequest) error {
	m.TotalModal = m.SisaModal + (m.TonasiGantung * m.HargaTbs)
	return db.QueryRow(
		"INSERT INTO modal (sisa_modal, tonasi_gantung, harga_tbs_ram, total_modal) VALUES ($1, $2, $3, $4) RETURNING id",
		m.SisaModal, m.TonasiGantung, m.HargaTbs, m.TotalModal,
	).Scan(&m.ID)
}

func GetModal(db *sql.DB) ([]EstimasiModalRequest, error) {
	rows, err := db.Query("SELECT id, sisa_modal, tonasi_gantung, harga_tbs_ram, total_modal FROM modal ORDER BY id DESC")
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