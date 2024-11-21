package app

import (
	"flag"
	"fmt"
	"log"
	"mms/internal/handler/gapi"
	"mms/internal/middleware"
	"mms/internal/repository"
	"mms/internal/service"

	"mms/pkg/database/postgres"
	"mms/pkg/dotenv"
	"net"

	pb "mms/internal/pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func RunServer() {
	fmt.Println("Grpc Server running ...")

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatal(err)
	}

	err = dotenv.Viper()

	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db)

	services := service.NewService(service.Deps{Repository: repository})

	userGapi := gapi.NewUserHandlerGrpcHandler(services.User)

	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryJWTInterceptor))

	pb.RegisterUserServiceServer(s, userGapi)

	fmt.Println("start gRPC server running on port : 50051")

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}
