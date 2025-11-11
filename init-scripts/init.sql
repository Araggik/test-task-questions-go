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

INSERT INTO questions (id, txt, created_at)
VALUES (1, 'Как запустить проект?', '2025-11-11 20:57:00');

