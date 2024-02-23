package main

import (
	"database/sql"
	"net"

	"github.com/devfullcycle/grpc/internal/database"
	"github.com/devfullcycle/grpc/internal/pb"
	"github.com/devfullcycle/grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
