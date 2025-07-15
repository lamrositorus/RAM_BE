package models

import (
	"database/sql"
	"fmt"
	"time"
)

type EstimasiKeuntungan struct {
    ID                   int    `json:"id"`
    Tanggal              string `json:"tanggal"`
    SpCairPabrik         int    `json:"sp_cair_pabrik"`
    HargaTbsPabrik       int    `json:"harga_tbs_pabrik"`
    TonasiSpRam          int    `json:"tonasi_sp_ram"`
    HargaTbsBeliRam      int    `json:"harga_tbs_beli_ram"`
    TotalModalBeli       int    `json:"total_modal_beli"`
    EstimasiKeuntungan   int    `json:"estimasi_keuntungan"`
}

func CalculateProfit(spCairPabrik, hargaTbsPabrik, totalModalBeli int) int {
	return (spCairPabrik * hargaTbsPabrik) - totalModalBeli
}

func CreateEstimasiKeuntungan(db *sql.DB, e *EstimasiKeuntungan) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	e.TotalModalBeli = e.TonasiSpRam * e.HargaTbsBeliRam
	e.Tanggal = time.Now().Format("2006-01-02")

	err = tx.QueryRow(
		`INSERT INTO estimasi_keuntungan 
		(tanggal, sp_cair_pabrik, harga_tbs_pabrik, tonasi_sp_ram, harga_tbs_beli_ram, total_modal_beli) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		e.Tanggal, e.SpCairPabrik, e.HargaTbsPabrik, e.TonasiSpRam, e.HargaTbsBeliRam, e.TotalModalBeli,
	).Scan(&e.ID)
	if err != nil {
		return err
	}

	e.EstimasiKeuntungan = CalculateProfit(e.SpCairPabrik, e.HargaTbsPabrik, e.TotalModalBeli)

	keuangan := Keuangan{
		Tanggal:   e.Tanggal,
		Deskripsi: fmt.Sprintf("Estimasi keuntungan tonasi SP pabrik %d kg x  @%d", e.SpCairPabrik, e.HargaTbsPabrik),
		Nominal:   e.EstimasiKeuntungan,
		Tipe:      "pemasukan",
	}

	err = tx.QueryRow(
		`INSERT INTO keuangan 
		(tanggal, deskripsi, nominal, tipe) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		keuangan.Tanggal, keuangan.Deskripsi, keuangan.Nominal, keuangan.Tipe,
	).Scan(&keuangan.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func GetEstimasiKeuntungan(db *sql.DB) ([]EstimasiKeuntungan, error) {
    rows, err := db.Query(`
        SELECT id, tanggal, sp_cair_pabrik, harga_tbs_pabrik, tonasi_sp_ram, harga_tbs_beli_ram, total_modal_beli, estimasi_keuntungan 
        FROM estimasi_keuntungan 
        ORDER BY tanggal ASC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var estimasi []EstimasiKeuntungan
    for rows.Next() {
        var e EstimasiKeuntungan
        if err := rows.Scan(
            &e.ID,
            &e.Tanggal,
            &e.SpCairPabrik,
            &e.HargaTbsPabrik,
            &e.TonasiSpRam,
            &e.HargaTbsBeliRam,
            &e.TotalModalBeli,
            &e.EstimasiKeuntungan,
        ); err != nil {
            return nil, err
        }
        estimasi = append(estimasi, e)
    }
    return estimasi, nil
}