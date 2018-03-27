package state

import (
	"context"

	"github.com/mzmico/protobuf"
)

type rpcState struct {
	context context.Context
	session *protobuf.Session
	err     *error
}

func NewRpcState(
	ctx context.Context,
	session *protobuf.Session,
	err *error,
) *rpcState {
	return &rpcState{
		context: ctx,
		session: session,
		err:     err,
	}
}
