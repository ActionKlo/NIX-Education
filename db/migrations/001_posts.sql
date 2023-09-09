-- +goose Up
CREATE TABLE posts (
    id INT,
    userId INT,
    title TEXT,
    body TEXT,
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE posts;