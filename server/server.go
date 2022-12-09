package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/cccddd77/lock_service/lockservice"
)

var (
	port       = flag.Int("port", 50051, "The server port")
)

type lockServiceServer struct {
	pb.UnimplementedLockServiceServer
}

func (s *lockServiceServer) DoLock(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	var rsp pb.Rsp
	rsp.CliID = req.CliID
	rsp.Operator = req.Operator
	rsp.Msg = "success"
	log.Print("Client ID: ", rsp.CliID, ", ", "Operator: ", rsp.Operator, "\n")
	return &rsp, nil
}

func (s *lockServiceServer) UnLock(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	var rsp pb.Rsp
	rsp.CliID = req.CliID
	rsp.Operator = req.Operator
	rsp.Msg = "success"
	log.Print("Client ID: ", rsp.CliID, ", ", "Operator: ", rsp.Operator, "\n")
	return &rsp, nil
}


func newServer() *lockServiceServer {
	s := &lockServiceServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLockServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}