package nix

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5/pgxpool"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Post struct {
	Id, UserId  int
	Title, Body string
}

type Comment struct {
	Id, PostId        int
	Name, Email, Body string
}

// можно ли не прокидывать сюда контекст а отправить его сразу в getCommentsByPostId()?
// точно можно, но не знаю как
func getPostsByUserId(userId int, ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup) []Post {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=" + strconv.Itoa(userId))
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

	// по идее тут вставить цикл в котором будет go GET /comments?postId=(5)
	// insertComments ?????
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go InsertComments(posts[i].Id, ctx, dbPool, wg)
	}

	return posts
}

func getCommentsByPostId(postId int, ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup) []Comment {
	res, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(postId))
	if err != nil {
		log.Fatalln("Error get comments:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("io.ReadAll commnets body err:", err)
	}

	var comments []Comment

	err = json.Unmarshal([]byte(body), &comments)
	if err != nil {
		log.Fatalln("Unmarshal err:", err)
	}

	return comments
}

func InsertPosts(userId int, ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := dbPool.Exec(ctx, "TRUNCATE posts")
	if err != nil {
		log.Fatalln("TRUNCATE posts error: ", err)
	}

	posts := getPostsByUserId(userId, ctx, dbPool, wg)

	query := "INSERT INTO posts VALUES ($1, $2, $3, $4)"

	for i := 0; i < 10; i++ {
		_, err := dbPool.Exec(ctx, query, posts[i].Id, posts[i].UserId, posts[i].Title, posts[i].Body)
		if err != nil {
			log.Fatalln("Exec err:", err)
		}
	}
}

// !TODO есть баг, был случай когда отсутствовало 2 записи: 271 и 296, то есть первые записи для 55 и 60 юзера
// бывает и больше =\
func InsertComments(postId int, ctx context.Context, dbPool *pgxpool.Pool, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := dbPool.Exec(ctx, "TRUNCATE comments")
	if err != nil {
		log.Fatalln("TRUNCATE comments error: ", err)
	}

	comments := getCommentsByPostId(postId, ctx, dbPool, wg)

	query := "INSERT INTO comments VALUES ($1, $2, $3, $4, $5)"

	for i := 0; i < 5; i++ {
		_, err := dbPool.Exec(ctx, query, comments[i].Id, comments[i].PostId, comments[i].Name, comments[i].Email, comments[i].Body)
		if err != nil {
			log.Fatalln("Exec err:", err)
		}
	}
}
