package utils

import (
	"encoding/base64"
	"fmt"
)

type Token struct {
	AppID string
	Uid   string
	UUID  string
}

func NewToken(appid string, uid string) *Token {
	return &Token{
		AppID: appid,
		Uid:   uid,
		UUID:  NewShortUUID(),
	}
}

func (m *Token) Key() string {
	return fmt.Sprintf(
		"%s:%s",
		m.AppID,
		m.Uid,
	)
}

func (m *Token) String() string {

	base := fmt.Sprintf(
		"%s:%s:%s",
		m.AppID,
		m.Uid,
		m.UUID)

	return base64.StdEncoding.EncodeToString([]byte(base))
}
