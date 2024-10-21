package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmuazam98/demo-grpc/user"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// clientCmd represents the gRPC client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the gRPC client",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		defer conn.Close()

		client := user.NewUserServiceClient(conn)

		// Example usage of the client
		createUser(client)
		getUser(client, 3)
	},
}

func createUser(client user.UserServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &user.CreateUserRequest{
		Name:    "Alice",
		Email:   "alice@example.com",
		Phone:   "1234567890",
		Address: "123 Main St",
	}

	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	fmt.Printf("Created User: %+v\n", resp.User)
}

func getUser(client user.UserServiceClient, userID int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &user.GetUserRequest{Id: userID}

	resp, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("Error getting user: %v", err)
	}

	fmt.Printf("Retrieved User: %+v\n", resp.User)
}
