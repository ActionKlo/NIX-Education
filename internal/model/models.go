package model

type Post struct {
	Id, UserId  int
	Title, Body string
}

type Comment struct {
	Id, PostId        int
	Name, Email, Body string
}
