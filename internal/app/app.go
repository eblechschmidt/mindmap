package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eblechschmidt/mindmap/internal/server"
)

func Serve(file string) error {
	s, err := server.New()
	if err != nil {
		return fmt.Errorf("could not start server: %w")
	}
	s.Run()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	return s.Stop()
}
