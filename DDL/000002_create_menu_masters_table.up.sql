CREATE TABLE IF NOT EXISTS menu_masters
(
    id         SERIAL PRIMARY KEY NOT NULL,
    bank_id    CHAR(4)            NOT NULL REFERENCES banks (id) ON DELETE CASCADE,
    name       VARCHAR(255)       NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)        NOT NULL,
    updated_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255)        NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_menu_masters_bank_id ON menu_masters (bank_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_menu_masters_bank_id_name ON menu_masters (bank_id, name);

COMMENT ON TABLE menu_masters IS 'メニューマスターテーブル';
COMMENT ON COLUMN menu_masters.id IS 'ID';
COMMENT ON COLUMN menu_masters.bank_id IS '銀行ID';
COMMENT ON COLUMN menu_masters.name IS '名前';
COMMENT ON COLUMN menu_masters.created_at IS '登録日時';
COMMENT ON COLUMN menu_masters.created_by IS '登録者';
COMMENT ON COLUMN menu_masters.updated_at IS '更新日時';
COMMENT ON COLUMN menu_masters.updated_by IS '更新者';
