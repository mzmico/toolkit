package utils

import "github.com/renstrom/shortuuid"

func NewShortUUID() string {
	return shortuuid.New()
}
