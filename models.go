package main

type Account interface {
}

type account struct {
}

func NewAccount() Account {
	return &account{}
}

type Post interface {
}

type post struct {
}

func NewPost() Post {
	return &post{}
}

type Reply interface {
}

type reply struct {
}

func NewReply() Post {
	return &reply{}
}
