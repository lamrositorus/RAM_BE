-- +goose Up
CREATE TABLE modal (
    id SERIAL PRIMARY KEY,
    tanggal DATE NOT NULL DEFAULT CURRENT_DATE,
    sisa_modal INTEGER NOT NULL,
    tonasi_gantung INTEGER NOT NULL,
    harga_tbs_ram INTEGER NOT NULL,
    total_modal INTEGER GENERATED ALWAYS AS (sisa_modal + (tonasi_gantung * harga_tbs_ram)) STORED,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE modal;