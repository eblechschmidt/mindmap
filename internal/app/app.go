package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eblechschmidt/mindmap/internal/server"
	"github.com/pkg/browser"
)

func Serve(file string) error {
	url := "127.0.0.1:1337"
	s, err := server.New(url)
	if err != nil {
		return fmt.Errorf("could not start server: %w")
	}
	s.Run()

	browser.OpenURL(fmt.Sprintf("http://%s", url))

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	return s.Stop()
}
