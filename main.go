package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/coding-yogi/grpcserver/employee"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterEmployeeServer(s, new(employeeService))
	s.Serve(lis)
}

type employeeService struct {
}

func (s *employeeService) GetEmployee(ctx context.Context, in *pb.EmployeeID) (*pb.EmployeeDetails, error) {
	return &pb.EmployeeDetails{FirstName: "Aniket", LastName: "Gadre"}, nil
}

func (s *employeeService) CreateEmployee(ctx context.Context, in *pb.EmployeeDetails) (*pb.EmployeeID, error) {

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("MD: %v\n", md)
	}

	fmt.Printf("\nCreating employee with FN: %v LN: %v\n", in.FirstName, in.LastName)
	return &pb.EmployeeID{Id: int32(randInt(10, 100))}, nil
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
