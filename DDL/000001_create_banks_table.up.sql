CREATE TABLE IF NOT EXISTS banks
(
    id         CHAR(4) PRIMARY KEY NOT NULL UNIQUE, -- 統一金融機関コードは3桁で先頭に0がくることがあるためCHARで
    name       VARCHAR(255)        NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255)        NOT NULL,
    updated_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(255)        NOT NULL
);

COMMENT ON TABLE banks IS '全体休日設定テーブル';
COMMENT ON COLUMN banks.id IS 'ID';
COMMENT ON COLUMN banks.name IS '銀行名';
COMMENT ON COLUMN banks.created_at IS '登録日時';
COMMENT ON COLUMN banks.created_by IS '登録者';
COMMENT ON COLUMN banks.updated_at IS '更新日時';
COMMENT ON COLUMN banks.updated_by IS '更新者';