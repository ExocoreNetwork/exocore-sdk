package client

import (
	grpc1 "github.com/gogo/protobuf/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	clientConn grpc1.ClientConn
}

func NewGRPCClient(url string, options ...grpc.DialOption) GrpcClient {
	if options == nil {
		options = []grpc.DialOption{grpc.WithInsecure()}
	}

	clientConn, err := grpc.Dial(url, options...)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return GrpcClient{clientConn: clientConn}
}

func (g GrpcClient) GenConn() (grpc1.ClientConn, error) {
	return g.clientConn, nil
}
