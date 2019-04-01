# GoLauncher
A Linux Application Launcher, based on [Ulauncher](https://ulauncher.io/) and written in Golang.

For now, it's basically just a reimplementation of [Ulauncher](https://ulauncher.io/) but aims to be faster and more flexible.

[App Search Mode Preview](https://streamable.com/umzzp)

(*The other modes are also available. Run the launcher to test them all out!*)

## So... What's different about GoLauncher?
Right now, the main advantages of **Golauncher** over [Ulauncher](https://ulauncher.io/) are:
* Scrollable results
* The number of results shown at a time is a user preference. (*And, because the results are scrollable, that doesn't limit your searches.*)
* It's written in Go, therefore, it should be faster (maybe).
* If an app already has an instance open, it switches focus instead of starting a new one. (This is useful for most apps, like Spotify, for example)
* The extension system (**TODO**) will use JSON to communicate with the launcher. This is an advantage because, 
unlike Ulauncher (that uses Python specific formats), pretty much every language can use JSON, and that means that,
provided that a library is made for that purpose, extensions can be written in any language.
