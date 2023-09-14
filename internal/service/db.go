package service

import (
	"NIX-Education/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	WritePost(post model.Post) error
	WriteComment(comment model.Comment) error
	ClearDB() error
}

type PostgresService struct {
	DB *pgxpool.Pool
}

func (s *PostgresService) WritePost(p model.Post) error {
	query := "INSERT INTO posts VALUES ($1, $2, $3, $4)"
	_, err := s.DB.Exec(context.Background(), query, p.Id, p.UserId, p.Title, p.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresService) WriteComment(c model.Comment) error {
	query := "INSERT INTO comments VALUES ($1, $2, $3, $4, $5)"
	_, err := s.DB.Exec(context.Background(), query, c.Id, c.PostId, c.Name, c.Email, c.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresService) ClearDB() error {
	_, err := s.DB.Exec(context.Background(), "TRUNCATE posts, comments")
	if err != nil {
		return err
	}

	return nil
}
