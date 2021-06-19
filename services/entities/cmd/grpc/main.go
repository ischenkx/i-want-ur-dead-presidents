package main

import (
	"context"
	"fmt"
	"github.com/ischenkx/innotech-backend/services/entities/implementation/grpc/pb/generated"
	grpcProducts "github.com/ischenkx/innotech-backend/services/entities/implementation/grpc/server"
	"github.com/ischenkx/innotech-backend/services/entities/implementation/mongodb"
	"github.com/ischenkx/innotech-backend/services/entities/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	ConfigPath = "./"
)

var Config struct {
	Port     int
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

	srv := service.New(db)

	s := grpc.NewServer()

	products.RegisterProductsServer(s, grpcProducts.New(srv))

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
