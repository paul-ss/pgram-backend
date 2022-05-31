package domain

/*
 * Password:
 * at least 7 letters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character
 * not similar to username
 * is not in common password list??
 */

// User is database model
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
