package main

import (
  "context"
  "log"
  "net"

  "google.golang.org/grpc"
  pb "github.com/devincl18/go-rpc/chat"
)

type server struct {
  pb.UnimplementedChatServiceServer
}

func (s *server) SignIn(ctx context.Context, in *pb.User) (*pb.Message, error) {
  log.Printf("Received request sign in:\nUUID: %d\nName: %s", in.Uuid,in.Name)
  return &pb.Message{
    Message: "Sign In request access",
    Status:  true,
  }, nil
}

func (s *server) SignOut(ctx context.Context, in *pb.User) (*pb.Message, error) {
  log.Printf("Received request sign out:\nUUID: %d\nName: %s", in.Uuid,in.Name)
  return &pb.Message{
    Message: "Sign Out request access",
    Status:  true,
  }, nil
}

func (s *server) SignUp(ctx context.Context, in *pb.NewUser) (*pb.Message, error) {
  log.Printf("Received request sign out:\nUUID: %d\nName: %s\nAge: %d\nAddress: %s", in.Uuid, in.Name, in.Age, in.Address)
  return &pb.Message{
    Message: "Sign Up request access",
    Status:  true,
  }, nil
}

func main() {
  const PORT = ":9000"

  lis, err := net.Listen("tcp", PORT)
  if err != nil {
    log.Fatalf("Failed to listen on port 9000: %v", err)
  }

  grpcServer := grpc.NewServer()
  pb.RegisterChatServiceServer(grpcServer, &server{})

  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("Failed to serve gRPC server over port %d: %v", PORT, err)
  }
}
