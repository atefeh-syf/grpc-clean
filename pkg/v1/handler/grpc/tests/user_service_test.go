package grpc_test

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/atefeh-syf/grpc-clean/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreateUser(t *testing.T){
  conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    t.Fatal("the connection with the server cannot be established")
  }
  defer conn.Close()

  client := pb.NewUserServiceClient(conn)

  request := &pb.CreateUserRequest{
    Name: "test2",
    Email: "test2@test.com",
  }

  res, err := client.Create(context.Background(), request)
  if err != nil{
    t.Fatalf("CREATE FAILED: %v", err)
  }

  if res.Email != request.Email{
     t.Errorf("CREATE returned incorrect email, expected %s got %s", request.Email, res.Email)
  }

  if res.Name != request.Name{
     t.Errorf("CREATE returned incorrect Name, expected %s got %s", request.Name, res.Name)
  }

  if res.GetId() == ""{
    t.Error("CREATE function didnot returned id as the response")
  }
  fmt.Println("Eyvalllll Ati!!!!!!!!!!!!!!!!!!!")
}


func TestGetUser(t *testing.T){
  conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    t.Fatal("the connection with the server cannot be established")
  }
  defer conn.Close()

  client := pb.NewUserServiceClient(conn)

  request := &pb.SingleUserRequest{
    Id: "4",
  }

  res, err := client.Get(context.Background(), request)
  if err != nil{
    t.Fatalf("Get FAILED: %v", err)
  }
  fmt.Printf("\n User : %v", res)
  fmt.Println("\nEyvalllll Ati!!!!!!!!!!!!!!!!!!!")
}


func TestDeleteUser(t *testing.T){
  conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    t.Fatal("the connection with the server cannot be established")
  }
  defer conn.Close()

  client := pb.NewUserServiceClient(conn)

  request := &pb.SingleUserRequest{
    Id: "4",
  }

  res, err := client.Delete(context.Background(), request)
  if err != nil{
    t.Fatalf("delete FAILED: %v", err)
  }
  fmt.Printf("\n User : %v", res)
  fmt.Println("\n Eyvalllll Ati!!!!!!!!!!!!!!!!!!!")
}



func TestUpdateUser(t *testing.T){
  conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    t.Fatal("the connection with the server cannot be established")
  }
  defer conn.Close()

  client := pb.NewUserServiceClient(conn)

  request := &pb.UpdateUserRequest{
    Id: "5",
    Name: "test222",
    Email: "test222@test.com",
  }

  res, err := client.Update(context.Background(), request)
  if err != nil{
    t.Fatalf("update FAILED: %v", err)
  }
  fmt.Printf("\n User : %v", res)
  fmt.Println("\n Eyvalllll Ati!!!!!!!!!!!!!!!!!!!")
}