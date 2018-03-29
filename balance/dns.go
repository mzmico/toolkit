package balance

import (
	"github.com/mzmico/toolkit/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type DNSBalance struct {
	conn map[string]*grpc.ClientConn
}

func (m *DNSBalance) GetConn(name string) *grpc.ClientConn {
	return m.conn[name]
}

func NewDNSBalance(v *viper.Viper) (*DNSBalance, error) {

	address := v.GetStringMapString("balance.dns")

	var (
		b = &DNSBalance{
			conn: make(map[string]*grpc.ClientConn),
		}
	)

	if address == nil {
		return b, nil
	}

	for name, addr := range address {

		conn, err := grpc.Dial(addr, grpc.WithInsecure())

		if err != nil {
			return nil, errors.By(err)
		}

		b.conn[name] = conn

	}

	return b, nil

}
