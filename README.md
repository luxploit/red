# [luxploit](https://luxploit.net)/[red](https://github.com/luxploit/red)

> [!WARNING]
> This library is an experimental, but working, state and bugs are fully expected. _You could use this for production code now if you want_.
>
> Contributions for bugfixes and new features are welcome!

`red` is a Golang v1.24+ library is an _experimental_ DI/Dependency Injection framework, designed to help in the creation of microservices.

To install, just run `go get -u github.com/luxploit/red`

### Example Usage (DI Container)

```go
func main() {
    app := red.New()

    app.Use(
	    red.Provide(settings.New),
	    red.Provide(log.New),
	    red.Invoke(lifecycle.AwaitInterrupt),
	    red.Invoke(server.ListenHTTP), 
    )

    if err := app.Run(); err != nil {
	    panic(err)
    }
}
```

**NOTE:** `red` does is not responsible for keeping the program alive after calling `app.Run()`, unless you have exisiting code listening for an interrupt, you can use a function like this as an `red.Invoke` handler to handle it:

```go
func AwaitInterrupt(log *log.Logger) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		if sig := <-c; sig != nil {
			log.Info("Lifecycle", "Captured %v! Stopping Server...", sig)
			os.Exit(0)
		}
	}()

	return nil
}
```

### Design

The design is largely inspired by [uber](https://github.com/uber-go)'s [fx](https://github.com/uber-go/fx) and [dig](https://github.com/uber-go/dig) libraries, however since I have way too much time on my hands and reinventing the wheel is always fun for learning, I decided to write my own more minimal framework around Dependency Injection.

The big difference between most DI libraries and `red` is that `red` also ships functionality for usage as a SL/Service Locator (that also needs to be better tested and documented). Providers and Services can be used interchangeably, however its recommended to stick to one or the other.

`red` is still in an alright state. Bugfixes, feature requests and contributions are very welcome. If you prefer a cup of tea and a nice chat instead of an issue ticket, shoot me an E-Mail over at [laura@luxploit.net](mailto:laura@luxploit.net) or a DM on Discord `@luxploit.net`
