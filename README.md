# Randport

A very simply library that returns a random TCP port. It does so by
asking the OS to listen on `:0`, which makes the OS choose a random
free port for you. Port will then immediately close the listener and
return the port chosen by the OS.

```go
port := randport.Get()
```

**Warning** If the OS fails to create the listener, the `Get()` will panic.

# License

MIT-LICENSE. See [LICENSE](http://olivere.mit-license.org/)
or the LICENSE file provided in the repository for details.
