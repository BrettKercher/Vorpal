package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-martini/martini"
	golog "github.com/op/go-logging"
)

var log = golog.MustGetLogger("middleware")

func getRequestIp(request *http.Request) string {
	// prefer to use X-Real-Ip
	realIp := request.Header.Get("X-Real-Ip")
	if realIp != "" {
		return realIp
	}

	// use X-Forwared-For as a backup
	// https://en.wikipedia.org/wiki/X-Forwarded-For
	forwardedForIp := request.Header.Get("X-Forwarded-For")
	if forwardedForIp != "" {
		return strings.Split(forwardedForIp, ", ")[0]
	}

	return request.RemoteAddr
}

// Logger returns a middleware handler that logs the request as it goes in and
// the response as it goes out.
func Logger() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()

		log.Debug("Getting the real ip address of request")
		ipAddress := getRequestIp(request)

		log.Notice("Starting request with method=%q url=%q ip_address=%q",
			request.Method, request.URL.Path, ipAddress)
		martiniWriter := writer.(martini.ResponseWriter)
		log.Notice("Completed request with status=\"%v\" in elapsed_time=\"%v\"",
			martiniWriter.Status(), time.Since(startTime))
	}
}
