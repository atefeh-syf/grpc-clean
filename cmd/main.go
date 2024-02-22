package main

import (
	"fmt"
	"log"
	"net"

	dbConfig "github.com/atefeh-syf/grpc-clean/internal/db"
	"github.com/atefeh-syf/grpc-clean/internal/models"
	interfaces "github.com/atefeh-syf/grpc-clean/pkg/v1"
	repo "github.com/atefeh-syf/grpc-clean/pkg/v1/repository"
	handler "github.com/atefeh-syf/grpc-clean/pkg/v1/handler/grpc"
	usecase "github.com/atefeh-syf/grpc-clean/pkg/v1/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	db := dbConfig.DbConn()
	migrations(db)

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()

	userUseCase := initUserServer(db)
	handler.NewServer(grpcServer, userUseCase)
	log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	userRepo := repo.New(db)
	return usecase.New(userRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
