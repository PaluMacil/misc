# gForm Exercise

This repo is currently in a broken state but can be worked around. The original author did a lot of work on documentation, showing MSDN links above functions in the underlying win 32 api wrapper lib. Hopefully this will be fixed, but since I think [this PR](https://github.com/AllenDang/w32/commit/a61dc27a068a1e324c26b5f5d84655cb4824eecb) broke it in April 2016, hope for this project seems dim. Using some pointers in [the issue that discusses the build problems](https://github.com/AllenDang/w32/issues/65), I might be able to compile with a specific commit mentioned in the comments.

You can at least get the application to run if you use the below commands in PowerShell, but it doesn't seem to run without a console, doesn't listen well when you try to get it to stop (use ctrl-c), and overall my assessment is that this is a great project that just never got completed.

```
go get github.com/AllenDang/w32
cd $env:GOPATH/src/github.com/AllenDang/w32/
git checkout 4255f5e69a7e10216b3b58db2129d72a7ac85906
cd $env:GOPATH/src/github.com/palumacil/misc/gformex/
go build
.\gformex
```

[&#x2190; Back to Project List](../README.md)