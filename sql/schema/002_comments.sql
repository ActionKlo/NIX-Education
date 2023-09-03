-- +goose Up
CREATE TABLE comments (
    id INT,
    postId INT,
    name TEXT,
    email TEXT,
    body TEXT,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE comments;