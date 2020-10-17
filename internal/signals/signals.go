package signals

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var sigChannel = make(chan os.Signal)

func Handle(cancel context.CancelFunc) {
	go func() {
		// kill -9 sounds not like a graceful
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
