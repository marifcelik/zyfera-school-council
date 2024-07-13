package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"

	"school_council/config"
	"school_council/handler"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("GET /health_check", handler.HealthCheck)
	app.HandleFunc("POST /create", handler.Create)
	app.HandleFunc("PATCH /{stdNumber}", handler.Update)

	addr := net.JoinHostPort(config.C.Host, config.C.Port)
	slog.Info(fmt.Sprintf("server is running on %s", addr))
	log.Fatal(http.ListenAndServe(addr, app))
}
