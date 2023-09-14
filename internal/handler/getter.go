package handler

import (
	"NIX-Education/internal/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	getPostsURL    = "https://jsonplaceholder.typicode.com/posts?userId="
	getCommentsURL = "https://jsonplaceholder.typicode.com/comments?postId="
)

func GetPostsByUserID(id int) ([]model.Post, error) {
	res, err := http.Get(getPostsURL + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var posts []model.Post
	err = json.Unmarshal([]byte(body), &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func GetCommentsByPostID(id int) ([]model.Comment, error) {
	res, err := http.Get(getCommentsURL + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var comments []model.Comment
	err = json.Unmarshal([]byte(body), &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
