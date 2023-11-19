package main

import (
	"accommodation-booking-system/user-service/pkg/database"
	"accommodation-booking-system/user-service/pkg/handlers"
	"accommodation-booking-system/user-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func main() {
	var cfg Config
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+cfg.Server.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Database.User, cfg.Database.Password,
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.Name)
	dbManager, err := database.NewDatabaseManager(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database manager: %v", err)
	}
	userServiceServer := handlers.NewUserServiceServer(dbManager)
	proto.RegisterUserServiceServer(s, userServiceServer)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
