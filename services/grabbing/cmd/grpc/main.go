package main

import (
	"context"
	"fmt"
	"github.com/ischenkx/innotech-backend/common"
	"github.com/ischenkx/innotech-backend/services/grabbing/implementation/grabber"
	"github.com/ischenkx/innotech-backend/services/grabbing/implementation/grpc/pb/generated"
	grpcUsers "github.com/ischenkx/innotech-backend/services/grabbing/implementation/grpc/server"
	"github.com/ischenkx/innotech-backend/services/grabbing/implementation/mongodb"
	"github.com/ischenkx/innotech-backend/services/grabbing/service"
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
	Damia struct {
		Fns     string `yaml:"fns"`
		Scoring string `yaml:"scoring"`
		Arbitr  string `yaml:"arbitr"`
	}
}

//TODO put api-keys to safer place
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
	fmt.Println(Config)

	return nil
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("failed to read config:", err)
		return
	}

	Config.Database.Url = common.LoadMongoFromEnv()

	db, err := mongodb.Connect(Config.Database.Url, Config.Database.Name, Config.Database.Collection)
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Println(err)
		}
	}()

	graber := grabber.Grabber{
		FnsKey:     Config.Damia.Fns,
		ScoringKey: Config.Damia.Scoring,
		ArbitrKey:  Config.Damia.Arbitr,
	}

	fmt.Println(Config.Damia)

	srv := service.New(db, &graber)

	grabbingServer := grpcUsers.New(srv)

	s := grpc.NewServer()

	grabbing.RegisterGrabbingServer(s, grabbingServer)

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
