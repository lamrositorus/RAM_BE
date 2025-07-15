-- +goose Up
CREATE TABLE estimasi_keuntungan (
    id SERIAL PRIMARY KEY,
    tanggal DATE NOT NULL DEFAULT CURRENT_DATE,
    sp_cair_pabrik INTEGER NOT NULL,
    harga_tbs_pabrik INTEGER NOT NULL,
    tonasi_sp_ram INTEGER NOT NULL,
    harga_tbs_beli_ram INTEGER NOT NULL,
    total_modal_beli INTEGER NOT NULL,
    estimasi_keuntungan INTEGER GENERATED ALWAYS AS ((sp_cair_pabrik * harga_tbs_pabrik) - total_modal_beli) STORED,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE estimasi_keuntungan;