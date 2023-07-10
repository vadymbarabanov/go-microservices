package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	user "github.com/vadymbarabanov/go-microservices/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userClient := user.NewUserClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := userClient.Create(
			context.Background(),
			&user.CreateUserRequest{
				Email:    "hello@world",
				Password: "BLABLA",
			},
		)

		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(*response)
	})

	http.ListenAndServe(":3000", nil)
}
