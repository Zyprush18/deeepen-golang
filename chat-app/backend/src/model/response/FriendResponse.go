package response

import "github.com/Zyprush18/deeepen-golang/chat-app/backend/src/database"

type Friend struct {
	Id 	int `json:"id"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
	Status string `json:"status"`

}

func ParseFriend(data []database.Friend) (resp []Friend)   {
	for _, v := range data {
		resp = append(resp, Friend{
			Id: int(v.Users.ID),
			Name: v.Name,
			Uuid: v.Users.UUID,
			Status: v.Status,
		})
	}
	return resp
}