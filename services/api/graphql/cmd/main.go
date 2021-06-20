package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ischenkx/innotech-backend/services/api/graphql/api"
	"github.com/ischenkx/innotech-backend/services/api/graphql/graph/generated"
	authClient "github.com/ischenkx/innotech-backend/services/auth/implementation/grpc/client"
	billingClient "github.com/ischenkx/innotech-backend/services/billing/implementation/grpc/client"
	"github.com/ischenkx/innotech-backend/services/users/implementation/grpc/client"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	entitiesClient "github.com/ischenkx/innotech-backend/services/entities/implementation/grpc/client"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	time.Sleep(time.Second * 8)

	usersGrpcClient, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Println("failed to connect users service:", err)
	}
	authGrpcClient, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	entitiesGrpcClient, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	billingGrpcClient, err := grpc.Dial("localhost:3031", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	serv := api.New(authClient.New(authGrpcClient), client.New(usersGrpcClient), entitiesClient.New(entitiesGrpcClient), billingClient.New(billingGrpcClient))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: serv}))
	serv.Engine().GET("/", gin.WrapF(playground.Handler("GraphQL playground", "/query")))
	serv.Engine().POST("/query", gin.WrapH(srv))
	serv.Engine().Run(":"+port)
}
