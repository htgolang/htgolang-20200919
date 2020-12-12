package logger

import (
	"io"
	"log"
	"net"
	"net/http"
)

var logger *log.Logger

func InitLogger(writer io.Writer) {
	logger = log.New(writer, "", 0)
}

func LoggerWrapper(action http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action(w, r)
		if logger != nil {
			addr, _, _ := net.SplitHostPort(r.RemoteAddr)
			logger.Printf(
				`%s %s %s "%s"`,
				addr,
				r.Method,
				r.URL, r.Header.Get("User-Agent"),
			)
		}
	}
}
