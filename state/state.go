package state

import (
	"github.com/mzmico/mz"
	"github.com/mzmico/protobuf"
	"google.golang.org/grpc"
)

type State struct {
	service mz.IService
	session *protobuf.Session
}

func (m *State) Rpc(endpoint string) *grpc.ClientConn {
	return nil
}
