# Playing with a Go wrapper of libui

You need to have GCC installed and in your path (compiled with SEH style exceptions). Old versions don't work. I'm not certain what the cutoff would be but on Windows I had issues with 5.10 TDM and [someone couldn't get GCC 5.4 via Cygwin to work](https://github.com/andlabs/ui/issues/196). 

```
go get github.com/andlabs/ui
go build -ldflags -H=windowsgui
```

The linking flag hides the terminal Window.

[&#x2190; Back to Project List](../README.md)