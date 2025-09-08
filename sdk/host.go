package sdk

import "net"

type HostOptions struct {
    Addr string
}

func Listen(addr string) (net.Listener, error) {
    return net.Listen("tcp", addr)
}
