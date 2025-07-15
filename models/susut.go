package models

import (
	"database/sql"
	"time"
)

type SusutTimbangan struct {
	ID             int     `json:"id"`
	Tanggal        string  `json:"tanggal"`
	NomorPolisi    string  `json:"nomor_polisi"`
	NamaSupir      string  `json:"nama_supir"`
	JamMasukPabrik string  `json:"jam_masuk_pabrik"`
	SPPabrik       float64 `json:"sp_pabrik"`
	BuahPulangan   float64 `json:"buah_pulangan"`
	JamKeluarRAM   string  `json:"jam_keluar_ram"`
	SPRAM          float64 `json:"sp_ram"`
	Selisih        float64 `json:"selisih"`
	Status         string  `json:"status"`
	Persentase     float64 `json:"persentase"`
}

func CreateSusutTimbangan(db *sql.DB, s *SusutTimbangan) error {
	s.Tanggal = time.Now().Format("2006-01-02")

	err := db.QueryRow(
		`INSERT INTO susut_timbangan 
		(tanggal, nomor_polisi, nama_supir, jam_masuk_pabrik, sp_pabrik, 
		buah_pulangan, jam_keluar_ram, sp_ram) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING id, selisih, status, persentase`,
		s.Tanggal, s.NomorPolisi, s.NamaSupir, s.JamMasukPabrik, 
		s.SPPabrik, s.BuahPulangan, s.JamKeluarRAM, s.SPRAM,
	).Scan(&s.ID, &s.Selisih, &s.Status, &s.Persentase)

	return err
}

func GetSusutTimbangan(db *sql.DB) ([]SusutTimbangan, error) {
	rows, err := db.Query(`
		SELECT id, tanggal, nomor_polisi, nama_supir, jam_masuk_pabrik, 
		sp_pabrik, jam_keluar_ram, sp_ram, selisih, status, persentase
		FROM susut_timbangan
		ORDER BY tanggal DESC, jam_masuk_pabrik DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []SusutTimbangan
	for rows.Next() {
		var s SusutTimbangan
		err := rows.Scan(
			&s.ID, &s.Tanggal, &s.NomorPolisi, &s.NamaSupir, &s.JamMasukPabrik,
			&s.SPPabrik, &s.JamKeluarRAM, &s.SPRAM, &s.Selisih, &s.Status, &s.Persentase,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}
	return results, nil
}
