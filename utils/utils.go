package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func JsonResp(w http.ResponseWriter, data any, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	if len(status) > 0 {
		w.WriteHeader(status[0])
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		slog.Error("json encode", "err", err)
		InternalErrResp(w, err)
	}
}

func ErrResp(w http.ResponseWriter, status int, err ...any) {
	var text string
	if len(err) > 0 && err[0] != nil {
		switch t := err[0].(type) {
		case string:
			text = t
		case error:
			text = t.Error()
		default:
			slog.Warn("unknown error type: %T", t)
			text = http.StatusText(status)
		}
	} else {
		text = http.StatusText(status)
	}

	http.Error(w, text, status)
}

func InternalErrResp(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
