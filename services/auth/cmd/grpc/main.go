package main

import (
	"context"
	"fmt"
	"github.com/ischenkx/innotech-backend/services/auth/implementation/grpc/pb/generated"
	authGrpcServer "github.com/ischenkx/innotech-backend/services/auth/implementation/grpc/server"
	"github.com/ischenkx/innotech-backend/services/auth/implementation/mongodb"
	"github.com/ischenkx/innotech-backend/services/auth/service"
	"google.golang.org/grpc"
	"time"

	"github.com/spf13/viper"
	"log"
	"net"
)

const (
	ConfigPath = "./"
)

var Config struct {
	Port     int
	JWT struct {
		Key string
		Alg string
		ExpiresIn time.Duration
	}
	Database struct {
		Name       string
		Url        string
		Collection string
	}
}

func initConfig() error {
	viper.AddConfigPath(ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("failed to read config:", err)
		return
	}

	db, err := mongodb.Connect(Config.Database.Url, Config.Database.Name, Config.Database.Collection)
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	log.Println("expires in:", Config.JWT.ExpiresIn)

	srv := service.New(db, Config.JWT.Key, Config.JWT.Alg, Config.JWT.ExpiresIn)


	authServer := authGrpcServer.New(srv)

	s := grpc.NewServer()

	auth.RegisterAuthServer(s, authServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", Config.Port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
