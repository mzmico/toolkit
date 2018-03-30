package state

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/mzmico/mz"
	"github.com/mzmico/protobuf"
	"github.com/mzmico/toolkit/errors"
	"github.com/sirupsen/logrus"
)

type RpcState struct {
	State
	context context.Context
	err     *error
	logger  *logrus.Entry
}

func NewRpcState(
	service *mz.RpcService,
	ctx context.Context,
	session *protobuf.Session,
	err *error,
) *RpcState {

	ctxlogrus.AddFields(
		ctx,
		logrus.Fields{
			"platform":   session.Platform,
			"version":    session.Version,
			"describe":   session.Describe,
			"os":         session.Os,
			"trace_id":   session.TraceId,
			"time_stamp": session.TimeStamp,
			"token":      session.Token,
			"uid":        session.Uid,
			"account":    session.Account,
			"net":        session.Net,
		},
	)

	state := &RpcState{
		State: State{
			service: service,
			session: session,
		},
		logger:  ctxlogrus.Extract(ctx),
		context: ctx,
		err:     err,
	}

	return state
}

func (m *RpcState) LogInfo(args ...interface{}) {
	m.logger.Info(args...)
}

func (m *RpcState) LogInfof(format string, args ...interface{}) {
	m.logger.Infof(format, args...)
}

func (m *RpcState) LogInfoln(args ...interface{}) {
	m.logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func (m *RpcState) LogWarning(args ...interface{}) {
	m.logger.Warning(args...)
}

func (m *RpcState) LogWarningf(format string, args ...interface{}) {
	m.logger.Warningf(format, args...)
}

func (m *RpcState) LogWarningln(args ...interface{}) {
	m.logger.Warningln(args...)
}

func (m *RpcState) LogError(args ...interface{}) {
	m.logger.Error(args...)
}

func (m *RpcState) LogErrorf(format string, args ...interface{}) {
	m.logger.Errorf(format, args...)
}

func (m *RpcState) Errorf(format string, args ...interface{}) error {
	m.logger.Errorf(format, args...)

	return errors.New(format, args...)
}

func (m *RpcState) Error(err error) error {

	var (
		e = errors.By(err)
	)

	m.logger.Error(e)

	return err
}
func (m *RpcState) LogErrorln(args ...interface{}) {
	m.logger.Errorln(args...)
}
