# bap

## テーブル一覧

| 名前 | カラム一覧 | コメント | タイプ |
| ---- | ------- | ------- | ---- |
| [public.banks](public.banks.md) | 6 | 全体休日設定テーブル | BASE TABLE |
| [public.menu_masters](public.menu_masters.md) | 7 | メニューマスターテーブル | BASE TABLE |

## ER図

```mermaid
erDiagram

"public.menu_masters" }o--|| "public.banks" : "FOREIGN KEY (bank_id) REFERENCES banks(id) ON DELETE CASCADE"

"public.banks" {
  character_4_ id
  varchar_255_ name
  timestamp_without_time_zone created_at
  varchar_255_ created_by
  timestamp_without_time_zone updated_at
  varchar_255_ updated_by
}
"public.menu_masters" {
  integer id
  character_4_ bank_id FK
  varchar_255_ name
  timestamp_without_time_zone created_at
  varchar_255_ created_by
  timestamp_without_time_zone updated_at
  varchar_255_ updated_by
}
```

---

> Generated by [tbls](https://github.com/k1LoW/tbls)