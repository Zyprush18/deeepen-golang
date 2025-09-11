package friend

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/friend/servicefriend"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/helper"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/request"
	"gorm.io/gorm"
)

type handleFriend struct {
	svc servicefriend.FriendRepo
}

func NewHandle(s servicefriend.FriendRepo) handleFriend {
	return handleFriend{svc: s}
}

func (h *handleFriend) AllFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Get Method Is Allowed",
		})
		return
	}

	id,err := strconv.Atoi(r.Context().Value(helper.UserId).(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid type id",
		})
		return
	}

	resp,err := h.svc.ListAllFriend(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "No friends",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
		Data: resp,
	})

}
func (h *handleFriend) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Post Method Is Allowed",
		})
		return
	}

	req := new(request.Friends)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
		})
		return
	}

	if err := h.svc.CreateFriend(req); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Friend Already Exists",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
	})

}

func (h *handleFriend) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Put Method Is Allowed",
		})
		return
	}

	role := r.URL.Query().Get("role")

	req := new(request.Friends)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Body Request Is Missing",
		})
		return
	}

	if err := h.svc.UpdateFriend(req,role); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Not Found",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
	})

}

func (h *handleFriend) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Delete Method Is Allowed",
		})
		return
	}

	from,err := strconv.Atoi(r.Context().Value(helper.UserId).(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Type id",
		})
		return
	}
	to,err := strconv.Atoi(r.PathValue("to"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Invalid Type Params",
		})
		return
	}


	if err := h.svc.DeleteFriend(from,to); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Not Found",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success",
	})

}
