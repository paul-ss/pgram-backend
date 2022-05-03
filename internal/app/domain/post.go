package domain

import "time"

type Post struct {
	Id      int64
	UserId  int
	GroupId int
	Content string
	Created time.Time
	Image   string
}

type PostDelivery interface {
}

type PostUsecase interface {
}

type PostRepository interface {
}
