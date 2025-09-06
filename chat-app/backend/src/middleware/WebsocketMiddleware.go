package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/helper"
)

func MiddlewareWs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		toUser := r.URL.Query().Get("toUser")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Unauthorized",
			})
			return
		}

		getTkn, err := helper.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Unauthorized: Invalid Token",
			})
			return
		}

		claims := getTkn.Claims.(*helper.CustomJwt)

		ctx := context.WithValue(r.Context(), helper.UserId, claims.ID)
		ctx = context.WithValue(ctx, helper.ToUserId, toUser)
		ctx = context.WithValue(ctx, helper.Name, claims.Name)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
