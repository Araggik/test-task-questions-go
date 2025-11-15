-- +goose Up
-- +goose StatementBegin
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    txt TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    question_id INT NOT NULL,
    user_id VARCHAR(36),
    txt TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

INSERT INTO questions (txt, created_at) VALUES 
('Как запустить проект?', '2025-11-11 20:57:00'),
('Работает ли сервис?', '2025-11-11 21:57:10'),
('Есть ли тесты?', '2025-11-15 18:57:00');

INSERT INTO answers (question_id, user_id, txt, created_at) VALUES 
(1, '1', 'Используй команду: docker-compose up -d', CURRENT_TIMESTAMP),
(1, '2', 'Не знаю', CURRENT_TIMESTAMP),
(2, '1', 'Да', CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS questions;
-- +goose StatementEnd
