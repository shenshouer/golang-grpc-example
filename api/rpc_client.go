package api

import (
	"context"
	"fmt"
	"io"
	"time"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc"
)

// RPCClient for communicate with server
type RPCClient struct {
	conn       *grpc.ClientConn
	ServerAddr string
	rpcClient  RPCServerClient
}

// Init rpc client
func (c *RPCClient) init() error {
	// Set up a connection to the server.
	var err error
	c.conn, err = grpc.Dial(c.ServerAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("did not connect: %v", err)
	}
	// defer conn.Close()

	c.rpcClient = NewRPCServerClient(c.conn)
	return nil
}

// Close close connection
func (c *RPCClient) Close() {
	c.conn.Close()
}

// HeartBeat send HeartBeat message to server
func (c *RPCClient) HeartBeat() (err error) {
	if err = c.init(); err != nil {
		return
	}
	defer c.Close()

	for i := 0; i < 20; i++ {
		if res, err := c.rpcClient.HeartBeat(context.Background(), &HeartBeatRequest{
			ClientID: fmt.Sprintf("%d", i),
			ClientIP: fmt.Sprintf("192.168.1.%d", i),
		}); err != nil {
			log.Errorln(err)
		} else {
			log.Infoln(res.IsOk)
		}
		time.Sleep(1 * time.Second)
	}

	return
}

// Stream for test
func (c *RPCClient) Stream(dataSources []*HeartBeatRequest) error {
	if err := c.init(); err != nil {
		return err
	}
	defer c.Close()

	ctx := context.Background()
	stream, err := c.rpcClient.Stream(ctx)
	defer ctx.Done()

	if err != nil {
		return err
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Errorln(err)
				return
			}
			log.Infoln(in.IsOk, err)
		}
	}()
	for _, note := range dataSources {
		if err := stream.Send(note); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	return nil
}
