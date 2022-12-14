package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/cccddd77/lock_service/lockservice"
)

var (
	port       = flag.Int("port", 50051, "The server port")
)

type lockServiceServer struct {
	pb.UnimplementedLockServiceServer
	lockedClient int64
	clientReqCache map[int64]string
	mutex sync.Mutex
	changed bool
}

func (s *lockServiceServer) getlock(id int64) string {
	var msg string
	s.mutex.Lock()
	if s.lockedClient == 0 {
		s.lockedClient = id
		s.changed = true;
		msg = "success"
	} else {
		msg = "failure"
		s.changed = false;
	}
	s.mutex.Unlock()
	return msg
}

func (s *lockServiceServer) DoLock(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	var rsp pb.Rsp
	rsp.CliID = req.CliID
	rsp.Operator = req.Operator
	tmp, ok := s.clientReqCache[rsp.CliID]
	if ok {
		if tmp == rsp.Operator {
			if s.changed {
				rsp.Msg = s.getlock(rsp.CliID)
				s.clientReqCache[rsp.CliID] = rsp.Operator
			} else {
				rsp.Msg = "repeat"
			}
		} else {
			rsp.Msg = s.getlock(rsp.CliID)
			s.clientReqCache[rsp.CliID] = rsp.Operator
		}
	} else {
		rsp.Msg = s.getlock(rsp.CliID)
		s.clientReqCache[rsp.CliID] = rsp.Operator
	}
	log.Print("Client ID: ", rsp.CliID, ", ", "Operator: ", rsp.Operator, " ", rsp.Msg, "\n")
	return &rsp, nil
}

func (s *lockServiceServer) droplock(id int64) string {
	var msg string
	s.mutex.Lock()
	if s.lockedClient == id {
		s.lockedClient = 0
		s.changed = true
		msg = "success"
	} else {
		s.changed = false
		msg = "failure"
	}
	s.mutex.Unlock()
	return msg
}

func (s *lockServiceServer) UnLock(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	var rsp pb.Rsp
	rsp.CliID = req.CliID
	rsp.Operator = req.Operator
	tmp, ok := s.clientReqCache[rsp.CliID]
	if ok {
		if tmp == rsp.Operator {
			if s.changed {
				rsp.Msg = s.droplock(rsp.CliID)
				s.clientReqCache[rsp.CliID] = rsp.Operator
			} else {
				rsp.Msg = "repeat"
			}
		} else {
			rsp.Msg = s.droplock(rsp.CliID)
			s.clientReqCache[rsp.CliID] = rsp.Operator
		}
	} else {
		rsp.Msg = s.droplock(rsp.CliID)
		s.clientReqCache[rsp.CliID] = rsp.Operator
	}
	log.Print("Client ID: ", rsp.CliID, ", ", "Operator: ", rsp.Operator, " ", rsp.Msg, "\n")
	return &rsp, nil
}


func newServer() *lockServiceServer {
	s := &lockServiceServer{clientReqCache: make(map[int64]string)}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Print("Listening at localhost: ", *port)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLockServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}