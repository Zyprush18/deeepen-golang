package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/helper"
)

func AuthMiddleware(next http.Handler) http.Handler  {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gettoken := strings.Split(r.Header.Get("Authorization")," ")
		if len(gettoken) != 2 || gettoken[1] == "" ||  gettoken[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Unauthorized",
			})
			return
		}

		parsToken,err := helper.ParseToken(gettoken[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helper.Messages{
				Message: "Unauthorized: Invalid Token",
			})
			return 
		}

		token := parsToken.Claims.(*helper.CustomJwt)
		ctx := context.WithValue(r.Context(), helper.Email, token.Subject)
		ctx = context.WithValue(ctx, helper.UserId, token.ID)

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}


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

		ctx := context.WithValue(r.Context(), helper.Uuid, claims.Uuid)
		ctx = context.WithValue(ctx, helper.ToUserId, toUser)
		ctx = context.WithValue(ctx, helper.Name, claims.Name)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
