package nix

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Interfacer interface {
	GetFromJP(url string) interface{}
	GetFromDB(query string, ctx context.Context, DBPool *pgxpool.Pool)
}

type Post struct {
	Id, UserId  int
	Title, Body string
}

type Comment struct {
	Id, PostId        int
	Name, Email, Body string
}

type Posts []Post

func (p Post) GetFromJP(url string) interface{} {
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

func (p Post) GetFromDB(query string, ctx context.Context, DBPool *pgxpool.Pool) {

}

//func (c Comment) GetFromJP(url string) []Comments {
//	res, err := http.Get(url)
//	if err != nil {
//		log.Fatalln("Error get comments:", err)
//	}
//
//	body, err := io.ReadAll(res.Body)
//	if err != nil {
//		log.Fatalln("io.ReadAll body err:", err)
//	}
//
//	var comments []Comment
//	err = json.Unmarshal([]byte(body), &comments)
//	if err != nil {
//		log.Fatalln("Unmarshal err:", err)
//	}
//
//	return comments
//}
//
//func (c Comment) GetFromDB(query string, ctx context.Context, DBPool *pgxpool.Pool) {
//
//}

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

	//// по идее тут вставить цикл в котором будет go GET /comments?postId=(5)
	//// insertComments ?????
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go InsertComments(posts[i].Id, ctx, dbPool, wg)
	//}

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

	query := `INSERT INTO posts (id, userId, title, body) VALUES `

	for i := 0; i < 10; i++ {
		lastLetter := ", "
		if i == 9 {
			lastLetter = ";"
		}

		values := fmt.Sprintf("(%d, %d, '%s', '%s')", posts[i].Id, posts[i].UserId, posts[i].Title, posts[i].Body)
		query = query + values + lastLetter

	}

	_, err = dbPool.Query(ctx, query)
	if err != nil {
		log.Fatalln("Insert post query err:", err)
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

	query := "INSERT INTO comments (id, postId, name, email, body) VALUES "

	for i := 0; i < 5; i++ {
		lastLetter := ", "
		if i == 4 {
			lastLetter = ";"
		}

		values := fmt.Sprintf("(%d, %d, '%s', '%s', '%s')", comments[i].Id, comments[i].PostId, comments[i].Name, comments[i].Email, comments[i].Body)
		query = query + values + lastLetter

	}

	_, err = dbPool.Query(ctx, query)
	if err != nil {
		log.Fatalln("Insert comments query err:", err)
	}
}
