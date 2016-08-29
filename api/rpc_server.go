package api

import (
	"net"

	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/Sirupsen/logrus"
)

// RPCServer for communicate with clients
type RPCServer struct {
	Advertise string
}

// HeartBeat handle HeartBeat from client
func (s *RPCServer) HeartBeat(ctx context.Context, req *HeartBeatRequest) (res *HeartBeatResponse, err error) {
	log.Infoln("ClientID", req.ClientID, "ClientIP", req.ClientIP)
	return &HeartBeatResponse{IsOk: true}, nil
}

// Stream for test
func (s *RPCServer) Stream(reqStream RPCServer_StreamServer) error {
	log.Infoln("============>>> Stream")
	for {
		log.Infoln("============>>> Stream 1")
		req, err := reqStream.Recv()
		if err != nil {
			if err == io.EOF {
				return reqStream.Send(&HeartBeatResponse{IsOk: false})
			}
			return err
		}
		log.Infoln(req.ClientID, req.ClientIP)

		err = reqStream.Send(&HeartBeatResponse{IsOk: true})
		if err != nil {
			return err
		}
	}
}

// Serve start RPCServer
func (s *RPCServer) Serve() (err error) {
	var lis net.Listener
	lis, err = net.Listen("tcp", s.Advertise)
	if err != nil {
		return
	}

	log.Infof("Start server at %s", s.Advertise)
	rpcS := grpc.NewServer()
	RegisterRPCServerServer(rpcS, s)
	rpcS.Serve(lis)
	return
}
