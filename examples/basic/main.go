package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luxploit/red"
)

func AwaitInterrupt() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	if sig := <-c; sig != nil {
		print("test2")
		os.Exit(0)
	}

	return nil
}

func funny() error {
	print("test1")
	time.Sleep(20 * time.Second)
	print("test3")
	return nil
}

func main() {
	app := red.New()

	app.Use(
		red.Invoke(AwaitInterrupt),
		red.Invoke(funny),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
