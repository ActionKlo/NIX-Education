package main

import (
	nix "NIX-Education/internal/service"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"sync"
)

func createDBPool() (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil {
		log.Fatalln("Connect Config err:", err)
		return nil, err
	}

	return dbPool, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	dbPool, err := createDBPool()
	if err != nil {
		log.Fatalln("createDBPool err:", err)
	}
	defer dbPool.Close()

	ctx := context.Background()

	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		go nix.SaveInFile(i, &wg)
	}

	var post nix.Post
	var comment nix.Comment

	_, err = dbPool.Exec(ctx, "TRUNCATE posts, comments")
	if err != nil {
		log.Fatalln(err)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			urlPost := "https://jsonplaceholder.typicode.com/posts?userId=" + strconv.Itoa(i)

			posts := post.ReadFromJP(urlPost).([]nix.Post)

			for _, post := range posts {
				wg.Add(1)

				go func(post nix.Post) {
					defer wg.Done()

					post.WriteToDB(ctx, dbPool, &wg)

					urlComm := "https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(post.Id)

					comments := comment.ReadFromJP(urlComm).([]nix.Comment)

					for _, comment := range comments {
						wg.Add(1)

						go func(comment nix.Comment) {
							defer wg.Done()
							comment.WriteToDB(ctx, dbPool, &wg)
						}(comment)

					}
				}(post)
			}
		}(i)
	}

	wg.Wait()
}
