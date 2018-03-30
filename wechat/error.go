package wechat

type IError interface {
	GetErrorCode() int64
	GetErrorMessage() string
}

type Error struct {
	ErrorCode    int64  `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

func (m *Error) GetErrorCode() int64 {
	return m.ErrorCode
}

func (m *Error) GetErrorMessage() string {
	return m.ErrorMessage
}
