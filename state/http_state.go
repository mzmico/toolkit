package state

import (
	"net/http"

	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"github.com/mzmico/protobuf"
	"github.com/mzmico/toolkit/errors"
)

type HttpState struct {
	c       *gin.Context
	session *protobuf.Session
	err     error
}

func (m *HttpState) parseSession() {

	s := m.session

	m.QueryValues(
		QueryValues{
			"_platform":   &s.Platform,
			"_version":    &s.Version,
			"_describe":   &s.Describe,
			"_os":         &s.Os,
			"_trace_id":   &s.TraceId,
			"_time_stamp": &s.TimeStamp,
			"_token":      &s.Token,
			"_uid":        &s.Uid,
			"_account":    &s.Account,
			"_net":        &s.Net,
		},
	)
}

func (m *HttpState) Param(name string, v interface{}) error {

	if m.err != nil {
		return m.err
	}

	value, ok := m.c.Params.Get(name)

	if !ok {
		m.err = errors.New("url param <%s> not found.", name)
		return m.err
	}

	err := conv.Infer(v, value)

	if err != nil {
		m.err = errors.New("url param <%s> parse fail, %s", name, err)
	}

	return m.err

}

type QueryValues map[string]interface{}

func (m *HttpState) QueryValues(ql QueryValues) {

	for name, value := range ql {
		m.Query(name, value)
	}
}

func (m *HttpState) Session() *protobuf.Session {
	return m.session
}

func (m *HttpState) Query(name string, v interface{}) error {

	if m.err != nil {
		return m.err
	}

	value := m.c.Query(name)

	if len(value) != 0 {
		err := conv.Infer(v, value)

		if err != nil {
			m.err = errors.New("url param <%s> parse fail, %s", name, err)
		}
	}

	return m.err
}

func (m *HttpState) GetLastError() error {
	return m.err
}

func NewHttpState(c *gin.Context) *HttpState {

	state := &HttpState{
		c:       c,
		session: &protobuf.Session{},
	}
	state.parseSession()
	return state
}

func (m *HttpState) Error(code int, message string, err error) {

	m.c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code":    code,
			"message": message,
			"data":    map[string]interface{}{},
		},
	)
}

func (m *HttpState) JSON(v interface{}) {

	m.c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"code":    0,
			"message": "success",
			"data":    v,
		},
	)
}

func GinHandler(cb func(state *HttpState)) func(context *gin.Context) {

	return func(context *gin.Context) {

		state := NewHttpState(context)

		if state.GetLastError() != nil {
			context.JSON(
				http.StatusOK,
				map[string]interface{}{
					"code":    int(protobuf.State_ArgumentError),
					"message": protobuf.State_ArgumentError,
				},
			)
		}

		cb(state)

		return

	}
}
