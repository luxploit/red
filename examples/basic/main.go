package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/luxploit/red"
)

func AwaitInterrupt() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	if sig := <-c; sig != nil {
		os.Exit(0)
	}

	return nil
}

func main() {
	app := red.New()

	app.Use(
		red.Invoke(AwaitInterrupt),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
