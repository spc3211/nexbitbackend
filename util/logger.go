package util

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func Init() {
	fmt.Println("Initiating logger")
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	Log.SetLevel(logrus.InfoLevel)
	Log.SetLevel(logrus.DebugLevel)

}

// WithRequest logs only request details and user
// logger.WithRequest(r).Info()  Debug() Error()
func WithRequest(r *http.Request) *logrus.Entry {
	userID, ok := r.Context().Value("userID").(string)
	fields := logrus.Fields{
		"http_method": r.Method,
		"remote_addr": r.RemoteAddr,
		"uri":         fmt.Sprintf("%s%s", r.Host, r.RequestURI),
	}
	if ok && userID != "" {
		fields["userID"] = userID
	}
	return Log.WithFields(fields)
}

// WithUser logs with only userID
// logger.WithUser(r).Info()  Debug() Error()
func WithUser(userID string) *logrus.Entry {
	fields := logrus.Fields{}
	if userID != "" {
		fields["userID"] = userID
	}
	return Log.WithFields(fields)
}

// WithContext logs with generic fields
func WithContext(context context.Context) *logrus.Entry {
	if logger, ok := context.Value("logger").(*logrus.Entry); ok {
		return logger
	}

	// Return a default logger if not found in context
	return logrus.NewEntry(logrus.StandardLogger())
}

// DebugWithUser logs with only userID If ENV is non prod
func DebugWithUser(userID string, message ...any) {
	if os.Getenv("STAGE") == "prod" {
		return
	}
	fields := logrus.Fields{}
	if userID != "" {
		fields["userID"] = userID
	}
	Log.WithFields(fields).Debugln(message...)
}

func WithResponse(r *http.Response) *logrus.Entry {
	fields := logrus.Fields{
		"Status":     r.Status,
		"StatusCode": r.StatusCode,
		"Body":       r.Body,
	}
	return Log.WithFields(fields)
}
