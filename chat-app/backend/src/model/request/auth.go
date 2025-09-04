package request

type Register struct {
	Username string	`json:"username"`
	Email string	`json:"email"`
	Password string	`json:"password"`
	Uuid string	`json:"uuid"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}