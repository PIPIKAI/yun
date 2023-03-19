package errs

import "google.golang.org/grpc/status"

var (
	EmptyErr = status.Error(200, "Not Record")
	TimeOut  = status.Error(200, "Time Out")
)
