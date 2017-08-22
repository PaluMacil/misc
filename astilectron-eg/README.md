# Go Astilectron

## Summary

This is the example code from [asticode/astilectron-go](https://github.com/asticode/go-astilectron). I've copied it here with the intent of adding my own examples which might not necessarily be worth push back upstream if they aren't demonstrating anything unique. Any changes to the original code will be documented here. This library provides a lazy-loaded (or or embedded as an alternative if you use [jteeuwen/go-bindata](github.com/jteeuwen/go-bindata/)) Electron solution for Go. I was surprised how it all just worked. Once I tried the embedded variant as well, my surprise was bordering on shock. This is a fantastic project! The download was quick, and the output file was only 6.5MB. I didn't determine where the Electron download wound up, but I will attempt to determine this shortly.

## Install

Switch to the example folder desired.

```
cd <example folder>
go get ./...
go build
.\<binary name>.exe
```

For embedding Electron in your binary (example 5), you need to also do the following just before `go build` in the instructions above.

```
go get -u github.com/jteeuwen/go-bindata/...
go generate
```

## License

See the [original project's MIT license included here](./LICENSE.txt).

[&#x2190; Back to Project List](../README.md)