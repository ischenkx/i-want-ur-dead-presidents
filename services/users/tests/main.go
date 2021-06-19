package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	users "github.com/ischenkx/innotech-backend/services/users/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Println(err)
		return
	}

	client := users.NewUsersClient(conn)

	fmt.Println("Welcome to interactive cli of our user-managing system!")
	fmt.Println("You can use following commands:")
	fmt.Println("	get +id")
	fmt.Println("	register +email +password")
	fmt.Println("	login +email +password")
	fmt.Println()
	fmt.Println("+ means you must enter each argument on a new line")
	fmt.Println()

	for {
		var cmd string
		if _, err := fmt.Scanln(&cmd); err != nil {
			log.Println(err)
			continue
		}

		switch cmd {
		case "get":
			var id string
			fmt.Print("userid:")
			if _, err := fmt.Scanln(&id); err != nil {
				log.Println(err)
				continue
			}

			user, err := client.Get(context.Background(), &users.GetUserRequest{Id: id})

			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}

			fmt.Printf("USER: %+v\n", user)
		case "register":
			var password string
			fmt.Print("password:")
			if _, err := fmt.Scanln(&password); err != nil {
				log.Println(err)
				continue
			}
			model := &users.RegisterRequest{
				Username: uuid.New().String(),
				Password: password,
			}

			start := time.Now()

			user, err := client.Register(context.Background(), model)

			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}

			fmt.Println(time.Since(start))

			fmt.Printf("USER: %+v\n", user)
		case "login":
			var un, password string
			fmt.Print("username:")
			if _, err := fmt.Scanln(&un); err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("password:")
			if _, err := fmt.Scanln(&password); err != nil {
				log.Println(err)
				continue
			}
			user, err := client.Login(context.Background(), &users.LoginRequest{
				Username: un,
				Password: password,
			})

			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}

			fmt.Printf("USER: %+v\n", user)
		}

	}

}
