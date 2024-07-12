package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	app := http.NewServeMux()

	app.HandleFunc("POST /create", handleCreate)
	app.HandleFunc("UPDATE /:id", handleUpdate)

	slog.Info("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", app))
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create"))
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	// use curl -X POST http://localhost:3000 to test this
	slog.Info("post request received")

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("post request received\n"))
}
