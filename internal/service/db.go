package service

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"io"
	"log"
	"net/http"
	"sync"
)

type Reader interface {
	ReadFromJP(url string) interface{}
}

type Writer interface {
	WriteToDB(ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup)
}

type Post struct {
	Id, UserId  int
	Title, Body string
}

type Comment struct {
	Id, PostId        int
	Name, Email, Body string
}

func (p *Post) ReadFromJP(url string) interface{} {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error get posts:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("io.ReadAll body err:", err)
	}

	var posts []Post
	err = json.Unmarshal([]byte(body), &posts)
	if err != nil {
		log.Fatalln("Unmarshal err:", err)
	}

	return posts
}

func (p *Post) WriteToDB(ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup) {
	_, err := dbPool.Exec(ctx, "INSERT INTO posts VALUES ($1, $2, $3, $4)", p.Id, p.UserId, p.Title, p.Body)
	if err != nil {
		log.Fatalln("Insert posts error: ", err, p.Id, p.UserId)
	}
}

func (c *Comment) ReadFromJP(url string) interface{} {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error get comments:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("io.ReadAll body err:", err)
	}

	var comments []Comment
	err = json.Unmarshal([]byte(body), &comments)
	if err != nil {
		log.Fatalln("Unmarshal err:", err)
	}

	return comments
}

func (c *Comment) WriteToDB(ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup) {
	_, err := dbPool.Exec(ctx, "INSERT INTO comments VALUES ($1, $2, $3, $4, $5)", c.Id, c.PostId, c.Name, c.Email, c.Body)
	if err != nil {
		log.Fatalln("Insert posts error: ", err, c.Id, c.PostId)
	}
}
