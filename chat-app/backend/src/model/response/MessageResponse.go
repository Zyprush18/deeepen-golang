package response


type MessageResponse struct {
	Name string `json:"name"`
	From string `json:"from"`
	Type string `json:"type"`
	To string `json:"to"`
	Message string `json:"message"`
}