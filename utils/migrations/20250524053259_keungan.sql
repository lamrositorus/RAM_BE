-- +goose Up
CREATE TABLE keuangan (
    id SERIAL PRIMARY KEY,
    tanggal DATE NOT NULL DEFAULT CURRENT_DATE,
    deskripsi TEXT NOT NULL,
    nominal INTEGER NOT NULL CHECK (nominal > 0),
    tipe VARCHAR(15) NOT NULL CHECK (tipe IN ('pemasukan', 'pengeluaran')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_keuangan_tanggal ON keuangan(tanggal);
CREATE INDEX idx_keuangan_tipe ON keuangan(tipe);

-- +goose Down
DROP TABLE keuangan;