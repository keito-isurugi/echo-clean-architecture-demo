TRUNCATE banks RESTART IDENTITY CASCADE;
INSERT INTO banks (id, name, created_by, updated_by)
VALUES ('0158', '京都銀行', 'system', 'system'),
       ('0188', '沖縄銀行', 'system', 'system');

TRUNCATE TABLE menu_masters RESTART IDENTITY CASCADE;
INSERT INTO menu_masters (bank_id, name, created_by, updated_by)
VALUES ('0158', '融資相談', 'system', 'system'),
       ('0158', '各種手続き', 'system', 'system'),
       ('0158', '資金運用', 'system', 'system'),
       ('0158', '相続', 'system', 'system');

TRUNCATE TABLE todos RESTART IDENTITY CASCADE;
INSERT INTO todos (title, content, created_by, updated_by)
VALUES ('テストToDo1', 'これはテストToDo1です！', 'system', 'system'),
       ('テストToDo2', 'これはテストToDo2です！', 'system', 'system'),
       ('テストToDo3', 'これはテストToDo3です！', 'system', 'system'),
       ('テストToDo4', 'これはテストToDo4です！', 'system', 'system');