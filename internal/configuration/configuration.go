package configuration

import "github.com/NodeFactoryIo/vedran/pkg/http-tunnel/server"

type Configuration struct {
	AuthSecret string
	Name       string
	Capacity   int64
	Whitelist  []string
	Fee        float32
	Selection  string
	Port       int32
	TunnelURL  string
	PortPool   *server.AddrPool
}

var Config Configuration
