package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/mmuazam98/demo-grpc/service"
	"github.com/mmuazam98/demo-grpc/user"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serverCmd represents the gRPC server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on port 50051: %v", err)
		}

		grpcServer := grpc.NewServer()
		userService := &service.CustomUserService{
			Users:  make(map[int32]*user.User), // In-memory storage for users
			NextID: 0,
		}

		user.RegisterUserServiceServer(grpcServer, userService)
		reflection.Register(grpcServer)

		fmt.Println("gRPC server running on port 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	},
}
