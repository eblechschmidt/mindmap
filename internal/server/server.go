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

// New creates a new server
func New(url string) (*Server, error) {
	s := &Server{http: &http.Server{Addr: url}}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.home)

	s.http.Handler = mux

	return s, nil
}

// Run starts the server
func (s *Server) Run() error {
	go func() {
		if err := s.http.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	log.Printf("Mindmap served at http://%s", s.http.Addr)
	return nil
}

// Stop gracefully shuts the server down
func (s *Server) Stop() error {
	ctx, release := context.WithTimeout(context.Background(), 10*time.Second)
	defer release()

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %w", err)
	}
	return nil
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	w.Write(web.Index)
}
