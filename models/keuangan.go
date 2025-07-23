package models

import (
	"database/sql"
	"time"
)

type Keuangan struct {
	ID        int       `json:"id"`
	Tanggal   time.Time `json:"tanggal"`
	Deskripsi string    `json:"deskripsi"`
	Nominal   int       `json:"nominal"`
	Tipe      string    `json:"tipe"`
}

func CreateKeuangan(db *sql.DB, k *Keuangan) error {
	k.Tanggal = time.Now()
	return db.QueryRow(
		"INSERT INTO keuangan (tanggal, deskripsi, nominal, tipe) VALUES ($1, $2, $3, $4) RETURNING id",
		k.Tanggal, k.Deskripsi, k.Nominal, k.Tipe,
	).Scan(&k.ID)
}

func GetKeuangan(db *sql.DB) ([]Keuangan, error) {
	rows, err := db.Query("SELECT id, tanggal, deskripsi, nominal, tipe FROM keuangan ORDER BY tanggal DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keuangan []Keuangan
	for rows.Next() {
		var k Keuangan
		if err := rows.Scan(&k.ID, &k.Tanggal, &k.Deskripsi, &k.Nominal, &k.Tipe); err != nil {
			return nil, err
		}
		keuangan = append(keuangan, k)
	}
	return keuangan, nil
}