package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eblechschmidt/mindmap/web"
)

// Server serves a web page showing the current state of the mind map rendered
// in svg
type Server struct {
	http *http.Server
}

func New() (*Server, error) {
	s := &Server{http: &http.Server{Addr: "127.0.0.1:1337"}}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.home)

	s.http.Handler = mux

	return s, nil
}
func (s *Server) Run() {
	go func() {
		if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	log.Printf("Server listening at %s", s.http.Addr)
}

func (s *Server) Stop() error {
	ctx, release := context.WithTimeout(context.Background(), 10*time.Second)
	defer release()

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %w")
	}
	return nil
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	w.Write(web.Index)
}
