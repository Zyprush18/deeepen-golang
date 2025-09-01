package response


type MessageResponse struct {
	From string `json:"from"`
	Type string `json:"type"`
	To string `json:"to"`
	Message string `json:"message"`
}