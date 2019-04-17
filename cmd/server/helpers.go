package main

import (
	"net/http"
	"strings"
)

type corsHandler struct {
	handler http.Handler
}

func (h corsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if req.Method == "OPTIONS" {
		return
	}
	h.handler.ServeHTTP(w, req)
}

// fileSystemHandler custom file system handler
type fileSystemHandler struct {
	fs http.FileSystem
}

// Open opens file
func (h fileSystemHandler) Open(path string) (http.File, error) {
	f, err := h.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := h.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
