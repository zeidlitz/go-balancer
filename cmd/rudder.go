package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/zeidlitz/rudder/internal/config"
	"github.com/zeidlitz/rudder/internal/env"
	"github.com/zeidlitz/rudder/internal/server"
)

func main() {
	conf := &config.Config{}
	conf.Servers = env.GetStrings("SERVERS", []string{"http://192.168.1.103:8080", "http://192.168.1.103:8081"})
	conf.Hostname = env.GetString("HOSTNAME", "localhost")
	conf.Algorithm = env.GetString("ALGORITHM", "roundrobin")
	conf.Port = env.GetString("PORT", "8080")


	server := server.Server{}
	conf.Host = fmt.Sprint(conf.Hostname + ":" + conf.Port)
	err := server.Configure(conf)
	if err != nil {
		os.Exit(1)
	}

	slog.Info("Rudder is running!")
	slog.Info("Configuration:", "servers", conf.Servers, "hostname", conf.Hostname, "port", conf.Port, "algorithm", conf.Algorithm)
	server.Start(conf.Host)
}
