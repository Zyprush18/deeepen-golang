package request

type Register struct {
	Username string
	Email string
	Password string
	Uuid string
}

type Login struct {
	Email string
	Password string
}