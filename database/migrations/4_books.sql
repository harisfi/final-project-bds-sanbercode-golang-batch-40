-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(256) NOT NULL,
    author VARCHAR(256),
    year_published INT,
    publisher_id INT NOT NULL,
    subject_id INT NOT NULL,
    is_borrowed BOOLEAN NOT NULL DEFAULT('0'),
    borrowed_by INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_publisher FOREIGN KEY(publisher_id) REFERENCES publishers(id),
    CONSTRAINT fk_subject FOREIGN KEY(subject_id) REFERENCES subjects(id),
    CONSTRAINT fk_borrower FOREIGN KEY(borrowed_by) REFERENCES users(id)
)

-- +migrate StatementEnd