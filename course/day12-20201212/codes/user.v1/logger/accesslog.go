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

// 代理模式
type Response struct {
	writer     http.ResponseWriter
	statusCode int
}

func (r *Response) Header() http.Header {
	return r.writer.Header()
}

func (r *Response) Write(b []byte) (int, error) {
	return r.writer.Write(b)
}

func (r *Response) WriteHeader(code int) {
	r.writer.WriteHeader(code)
	r.statusCode = code
}

func LoggerWrapper(action http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := &Response{writer: w, statusCode: 200}
		action(response, r)
		// w无法获取响应状态码
		if logger != nil {
			addr, _, _ := net.SplitHostPort(r.RemoteAddr)
			logger.Printf(
				`%s %s %s "%s" %d`,
				addr,
				r.Method,
				r.URL, r.Header.Get("User-Agent"),
				response.statusCode,
			)
		}
	}
}
