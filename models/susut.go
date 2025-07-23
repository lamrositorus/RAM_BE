package models

import (
	"database/sql"
	"fmt"
	"time"
)

type SusutTimbangan struct {
	ID             int     `json:"id"`
	Tanggal        string  `json:"tanggal"`
	NomorPolisi    string  `json:"nomor_polisi"`
	NamaSupir      string  `json:"nama_supir"`
	SPPabrik       float64 `json:"sp_pabrik"`
	BuahPulangan   float64 `json:"buah_pulangan"`
	SPRAM          float64 `json:"sp_ram"`
	Selisih        float64 `json:"selisih"`
	Status         string  `json:"status"`
	Persentase     float64 `json:"persentase"`
}

func CreateSusutTimbangan(db *sql.DB, s *SusutTimbangan) error {
if s.Tanggal == "" {
    return fmt.Errorf("field tanggal wajib diisi dan harus format yyyy-MM-dd")
}

if _, err := time.Parse("2006-01-02", s.Tanggal); err != nil {
    return fmt.Errorf("format tanggal salah, harus yyyy-MM-dd: %v", err)
}

	// Validasi format tanggal
	if _, err := time.Parse("2006-01-02", s.Tanggal); err != nil {
		return fmt.Errorf("format tanggal salah, harus yyyy-MM-dd: %v", err)
	}

	err := db.QueryRow(
		`INSERT INTO susut_timbangan 
		(tanggal, nomor_polisi, nama_supir, sp_pabrik, 
		buah_pulangan, sp_ram) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, selisih, status, persentase`,
		s.Tanggal, s.NomorPolisi, s.NamaSupir, s.SPPabrik, 
		s.BuahPulangan, s.SPRAM,
	).Scan(&s.ID, &s.Selisih, &s.Status, &s.Persentase)

	return err
}

func GetSusutTimbangan(db *sql.DB) ([]SusutTimbangan, error) {
	rows, err := db.Query(`
		SELECT id, tanggal, nomor_polisi, nama_supir, sp_pabrik, 
		buah_pulangan, sp_ram, selisih, status, persentase
		FROM susut_timbangan
		ORDER BY tanggal DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []SusutTimbangan
	for rows.Next() {
		var s SusutTimbangan
		err := rows.Scan(
			&s.ID, &s.Tanggal, &s.NomorPolisi, &s.NamaSupir, &s.SPPabrik,
			&s.BuahPulangan, &s.SPRAM, &s.Selisih, &s.Status, &s.Persentase,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, s)
	}
	return results, nil
}