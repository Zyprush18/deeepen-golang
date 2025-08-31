package request

type Register struct {
	Username string
	Email string
	Password string
}

type Login struct {
	Email string
	Password string
}