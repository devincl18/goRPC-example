package main

import (
  "context"
  "log"
  "time"

  "google.golang.org/grpc"
  pb "github.com/devincl18/go-rpc/chat"
)

func main() {
  const (
    address = "localhost:9000"
  )

  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
  if err != nil {
    log.Fatalf("Did not connect: %v", err)
  }
  defer conn.Close()

  c := pb.NewChatServiceClient(conn)

  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  var (
      newUser = pb.NewUser{
        Uuid:    0,
        Name:    "Kien",
        Age:     0,
        Address: "Ha Noi",
      }
      user = pb.User{
        Uuid: 0,
        Name: "Kien",
      }
  )
  r, err := c.SignIn(ctx, &user)
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  log.Printf("Message: %s\nStatus: %v", r.Message, r.Status)

  r1, err1 := c.SignUp(ctx, &newUser)
  if err1 != nil {
    log.Fatalf("could not greet: %v", err1)
  }
  log.Printf("Message: %s\nStatus: %v", r1.Message, r1.Status)

  r2, err2 := c.SignOut(ctx, &user)
  if err2 != nil {
    log.Fatalf("could not greet: %v", err2)
  }
  log.Printf("Message: %s\nStatus: %v", r2.Message, r2.Status)
}
