package util

import (
    "net/http"
    "io"
    "fmt"
)

func HttpErrorResponse(w http.ResponseWriter, code int, format string, a ...interface{}) {
    w.WriteHeader(code);
    io.WriteString(w, fmt.Sprintf(format, a...));
}

func ParseRequest(w http.ResponseWriter, req *http.Request) {
}
