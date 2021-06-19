package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	grabbing "github.com/ischenkx/innotech-backend/services/grabbing/implementation/grpc/pb/generated"
	"google.golang.org/grpc"
	"log"
)

func main() {
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	
	if err != nil {
		log.Println(err)
		return
	}
	
	client := grabbing.NewGrabbingClient(conn)
	
	fmt.Println("Welcome to interactive cli of our grabbing system!")
	fmt.Println("You can use following commands:")
	fmt.Println("	get +inn")
	fmt.Println()
	fmt.Println("+ means you must enter each argument on a new line")
	fmt.Println()
	
	for {
		var cmd string
		if _, err := fmt.Scan(&cmd); err != nil {
			log.Println("scan err:", err)
			continue
		}

		switch cmd {
		case "get":
			var inn string
			fmt.Print("inn:")
			if _, err := fmt.Scan(&inn); err != nil {
				log.Println(err)
				continue
			}

			response, err := client.Get(context.Background(), &grabbing.Product{
				Id:  uuid.New().String(),
				Inn: inn,
			})

			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}

			fmt.Printf("Response: %+v\n", response)
		}
		
	}
	
}

