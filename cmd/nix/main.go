package main

import (
	"NIX-Education/internal/handler"
	"NIX-Education/internal/model"
	"NIX-Education/internal/service"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

const (
	NumberOfUsers = 10
)

func CreateDBPool() (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	var postgres service.PostgresService

	postgres.DB, err = CreateDBPool()
	if err != nil {
		log.Fatalln(err)
	}
	defer postgres.DB.Close()

	err = postgres.ClearDB()
	if err != nil {
		log.Fatalln(errors.New(err.Error()))
	}

	errs := make(chan error)

	for i := 1; i <= NumberOfUsers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			posts, err := handler.GetPostsByUserID(id)
			if err != nil {
				errs <- err
			}

			for _, post := range posts {
				wg.Add(1)

				go func(post model.Post) {
					defer wg.Done()

					err = postgres.WritePost(post)
					if err != nil {
						errs <- err
					}

					comments, err := handler.GetCommentsByPostID(post.Id)
					if err != nil {
						errs <- err
					}

					for _, comment := range comments {
						wg.Add(1)

						go func(comment model.Comment) {
							defer wg.Done()

							err = postgres.WriteComment(comment)
							if err != nil {
								errs <- err
							}
						}(comment)

					}
				}(post)
			}
		}(i)
	}

	var once sync.Once
	go func() {
		for e := range errs {
			log.Println(e)
		}

		once.Do(func() {
			close(errs)
		})
	}()

	wg.Wait()
}
