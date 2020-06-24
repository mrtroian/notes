package signals

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var sigChannel chan os.Signal

func Handle(cancel context.CancelFunc) {
	go func() {
		signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

		for {
			sig := <-sigChannel
			switch sig {
			case os.Interrupt:
				cancel()
				return
			}
		}
	}()
}

func init() {
	sigChannel = make(chan os.Signal)
}
