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

I copied the controls demo code to this repo ([see BSD license](.\controls\LICENSE)) because I wondered if the resource file might be necessary for some functionality since it references Microsoft.Windows.Common-Controls (bascially controls available on Windows from the days of VB 6.0). This can be built and run, but there is still no exit, and the common dialogs mentioned don't show. I looked into common controls and realized they might be 32-bit only, though I didn't know for sure until I checked [Stack Overflow](https://stackoverflow.com/questions/836368/is-there-a-64-bit-version-of-microsoft-common-controls-mscomctl-ocx) and looked back at the [Microsoft download page](https://www.microsoft.com/en-us/download/details.aspx?id=10019).

Even though 64-bit Windows can run a 32-bit application just fine, I don't have a lot of interest in compiling for 32-bit. It's possible this library still works great if I compile for 32-bit, and the code still serves as great win32 documentation, so I'll leave this active.

[&#x2190; Back to Project List](../README.md)