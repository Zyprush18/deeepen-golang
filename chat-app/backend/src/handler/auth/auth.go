package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/auth/servicesauth"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/helper"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/request"

	"gorm.io/gorm"
)

type handleAuth struct {
	svc servicesauth.AuthService
}

func NewHandle(s servicesauth.AuthService) handleAuth {
	return handleAuth{svc: s}
}

func (h *handleAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Method Post Is Allowed",
		})
		return
	}

	req := new(request.Register)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed: Body Request is Missing",
		})
		return
	}

	if err := h.svc.Register(req); err != nil {
		if helper.CheckDuplicate(err) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Email Is Already Exists",
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Something Went Wrong",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(helper.Messages{
		Message: "Success Register",
	})
}

func (h *handleAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Only Method Post Is Allowed",
		})
		return
	}

	req := new(request.Login)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.Messages{
			Message: "Failed: Body Request is Missing",
		})
		return
	}

	tokens, err := h.svc.Login(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "pass incorrect" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Invalid Credential",
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
		Message: "Success Login",
		Token:   tokens,
	})
}
