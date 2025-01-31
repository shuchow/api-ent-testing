package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shuchow/api-ent-testing/ent"
	"net/http"
)

func CreateUserHandler(entClient *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type UserRequest struct {
			UserName string `json:"userName,omitempty"`
			Email    string `json:"email,omitempty"`
		}

		decoder := json.NewDecoder(r.Body)
		userRequest := UserRequest{}
		err := decoder.Decode(&userRequest)
		if err != nil {
			fmt.Println("Error decoding userRequest:", err)
			w.WriteHeader(400)
		}

		userObj := entClient.User.Create()
		userObj.SetUserName(userRequest.UserName)
		userObj.SetEmail(userRequest.Email)
		_, err = userObj.Save(context.Background())

		if err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(201)
		}
	}
}
