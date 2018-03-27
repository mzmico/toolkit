package balance

import "google.golang.org/grpc"

type DnsBalance struct {
	conn map[string]*grpc.ClientConn
}
