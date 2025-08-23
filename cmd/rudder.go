package main

import (
	"fmt"
	"log/slog"

	"github.com/zeidlitz/rudder/internal/env"
	"github.com/zeidlitz/rudder/internal/server"
)

func main() {
	servers := env.GetStrings("SERVERS", []string{"http://localhost", "http://hostlocal"})
	hostname := env.GetString("HOSTNAME", "localhost")
	algorithm := env.GetString("ALGORITHM", "roundrobin")
	port := env.GetString("PORT", "8080")

	server := server.Server{}
	host := fmt.Sprint(hostname + ":" + port)
	err := server.Configure(algorithm, servers)
	if err != nil {
		slog.Error("Error during configuration")
		// TODO: What do we do here?
	}
	slog.Info("Rudder is running!")
	slog.Info("configuration", "servers", servers, "hostname", hostname, "port", port, "algorithm", algorithm)
	server.Start(host)
}
