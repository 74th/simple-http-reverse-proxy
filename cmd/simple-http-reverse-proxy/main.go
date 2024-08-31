package main

import (
	simplehttpreverseproxy "github.com/74th/simple-http-reverse-proxy"
	"github.com/alecthomas/kingpin/v2"
)

var (
	addr       = kingpin.Flag("addr", "addr").Required().String()
	configPath = kingpin.Flag("config", "config path").Required().String()
	static     = kingpin.Flag("static", "static file serve path").String()
)

func main() {

	kingpin.Parse()

	config := simplehttpreverseproxy.LoadHostsConfig(*configPath)

	server := simplehttpreverseproxy.NewServer(*addr, config, *static)
	server.Run()
}
