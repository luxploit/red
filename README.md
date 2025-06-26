# [luxploit](https://luxploit.net)/[red](https://github.com/luxploit/red)

> [!WARNING]
> This library is an experimental, but working, state and bugs are fully expected. _You could use this for production code now if you want_
>
> Contributions for bugfixes and new features are welcome!

`red` is a Golang v1.24+ library is an _experimental_ DI/Dependency Injection framework, designed to help in the creation of microservices.

The design is largely inspired by [uber](https://github.com/uber-go)'s [fx](https://github.com/uber-go/fx) and [dig](https://github.com/uber-go/dig) libraries, however since I have way too much time on my hands and reinventing the wheel is always fun for learning, I decided to write my own more minimal framework around Dependency Injection.

The big difference between most DI libraries and `red` is that `red` also ships functionality for usage as a SL/Service Locator. Providers and Services can be used interchangeably, however its recommended to stick to one or the other.

`red` is still in an alright state. Bugfixes, feature requests and contributions are very welcome. If you prefer a cup of tea and a nice chat instead of an issue ticket, shoot me an E-Mail over at [laura@luxploit.net](mailto:laura@luxploit.net) or a DM on Discord `@luxploit.net`
