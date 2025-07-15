-- +goose Up
CREATE TABLE susut_timbangan (
    id SERIAL PRIMARY KEY,
    tanggal DATE NOT NULL DEFAULT CURRENT_DATE,
    nomor_polisi VARCHAR(20) NOT NULL,
    nama_supir VARCHAR(100) NOT NULL,
    jam_masuk_pabrik TIME NOT NULL,
    sp_pabrik DECIMAL(10,2) NOT NULL CHECK (sp_pabrik > 0),
    buah_pulangan DECIMAL(10,2) NOT NULL DEFAULT 0,
    jam_keluar_ram TIME NOT NULL,
    sp_ram DECIMAL(10,2) NOT NULL CHECK (sp_ram > 0),
    selisih DECIMAL(10,2) GENERATED ALWAYS AS ((sp_pabrik + buah_pulangan) - sp_ram) STORED,
    status VARCHAR(10) GENERATED ALWAYS AS (
        CASE 
            WHEN (sp_pabrik + buah_pulangan) > sp_ram THEN 'kelebihan'
            WHEN (sp_pabrik + buah_pulangan) < sp_ram THEN 'susut'
            ELSE 'normal'
        END
    ) STORED,
    persentase DECIMAL(5,2) GENERATED ALWAYS AS (
        CASE
            WHEN (sp_pabrik + buah_pulangan) = 0 THEN 0
            ELSE ABS((sp_pabrik + buah_pulangan) - sp_ram) / (sp_pabrik + buah_pulangan) * 100
        END
    ) STORED,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT valid_jam CHECK (jam_keluar_ram > jam_masuk_pabrik)
);

CREATE INDEX idx_susut_tanggal ON susut_timbangan(tanggal);
CREATE INDEX idx_susut_nopol ON susut_timbangan(nomor_polisi);

-- +goose Down
DROP TABLE susut_timbangan;