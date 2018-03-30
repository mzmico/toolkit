package wechat

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/mzmico/toolkit/errors"
)

var (
	c = &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := &net.Dialer{
					KeepAlive: 10 * time.Minute,
					Timeout:   10 * time.Second,
				}

				return d.DialContext(ctx, network, addr)
			},
		},
		Timeout: 10 * time.Second,
	}
)

func GET(addr string, v interface{}) error {

	response, err := c.Get(addr)

	if err != nil {
		return errors.By(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return errors.By(err)
	}

	defer func() {
		response.Body.Close()
	}()

	err = json.Unmarshal(bytes, v)

	if err != nil {
		return errors.By(err)
	}

	value, ok := v.(IError)

	if ok {
		if value.GetErrorCode() != 0 {

			u, err := url.Parse(addr)

			if err != nil {
				return errors.New(
					"wechat call remote method fail, code=%d,message=%s",
					value.GetErrorCode(),
					value.GetErrorMessage(),
				)
			}

			return errors.New(
				"wechat call remote method %s fail, code=%d,message=%s",
				u.Path,
				value.GetErrorCode(),
				value.GetErrorMessage(),
			)
		}
	}

	return nil
}
