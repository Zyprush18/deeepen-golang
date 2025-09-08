package friend

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/friend/servicefriend"
// 	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/helper"
// )

// type handleFriend struct {
// 	svc servicefriend.FriendRepo
// }

// func NewHandle(s servicefriend.FriendRepo) handleFriend  {
// 	return handleFriend{svc: s}
// }

// func (h *handleFriend) GetAll(w http.ResponseWriter, r *http.Request)  {
// 	w.Header().Set("Content-Type","application/json")
// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		json.NewEncoder(w).Encode(helper.Messages{
// 			Message: "Only Get Method Is Allowed",
// 		})
// 		return
// 	}

// 	id,err := strconv.Atoi(r.PathValue("id"))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(helper.Messages{
// 			Message: "Invalid Type Params",
// 		})
// 		return
// 	}

// 	resp,err := h.svc.GetAll(id)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(helper.Messages{
// 			Message: "Something Went Wrong",
// 		})
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(helper.Messages{
// 		Message: "Success",
// 		Data: resp,
// 	})

// }