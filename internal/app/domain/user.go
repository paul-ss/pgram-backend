package domain

type User struct {
	Id        int
	Nickname  string
	FirstName string
	LastName  string
	About     string
	Email     string
	Password  string
	Image     string
}

type UserDelivery interface {
}

type UserUsecase interface {
}

type UserRepository interface {
}
