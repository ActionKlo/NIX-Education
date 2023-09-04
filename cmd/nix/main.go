package main

import (
	nix "NIX-Education"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	//ctx := context.Background()

	ch := make(chan string)
	var wg sync.WaitGroup

	// Цикл для получения всех постов и сохранению их в файлы (согласно заданию)
	for i := 1; i <= 2; i++ { // if i = 100 -> server sent GOAWAY and closed the connection
		wg.Add(1)

		//  уже выглядит дико =\
		//var post nix.Post
		var posts []nix.Posts

		//url := "https://jsonplaceholder.typicode.com/posts?userId=" + strconv.Itoa(i)
		//posts.GetFromJP(url)
		//for _, item := range posts {
		//	item.GetFromJP(url)
		//}
		//posts[0].GetFromJP(url)

		//nix.GetFromJP(url, posts)

		fmt.Println(posts)
		//go nix.SaveInFile(i, ch, &wg)
	}

	//for i := 1; i <= 10; i++ {
	//	wg.Add(1)
	//
	//	go nix.InsertPosts(i, ctx, dbPool, &wg)
	//}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		if _, ok := <-ch; !ok {
			break
		}
	}
}
