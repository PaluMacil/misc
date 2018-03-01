# C Webview

## Summary

I haven't written any C or C++ since highschool a decade and a half ago. I'm using this webview example to try to remember how to use external non-standard library dependencies. The [zserge/webview](https://github.com/zserge/webview) looks pretty good, and I'll probably use the Go part of it in the future if I want a lighter alternative to adding Electron.

The [usage in Go is very straightforward](..\webview\).

## Notes

 - Still trying to decide how to use the c project. `gcc main.c webview.c -DWEBVIEW_WINAPI=1 -lole32 -lcomctl32 -loleaut32 -luuid -lgdi32 -o webview-example.exe` created a binary but it seems to panic without an error message.
 - Go project looks great! Examples just work (well, except the canvas one). I have mingw64 installed.
 - Compiling a DLL from the 

```
cmake .
```

## Links
 - C Compiler review from [caltech.edu...compiling\_c](http://courses.cms.caltech.edu/cs11/material/c/mike/misc/compiling_c.html)
 - [Summary of purpose of make vs cmake](https://stackoverflow.com/questions/39761924/understanding-roles-of-cmake-make-and-gcc?rq=1)
 - [MakeFile tutorial on TutorialsPoint](https://www.tutorialspoint.com/makefile/makefile_example.htm)
 - [Win32 API C programming Tutorial from Winprog.org](http://www.winprog.org/tutorial/start.html)
 - [Win32 API C programming Tutorial from Zetcode](http://zetcode.com/gui/winapi/)