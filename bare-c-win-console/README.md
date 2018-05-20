# bare-c-win-console

This is a bare Windows console app in c. The result is a 5 or 6 kb file.

## Compile

Sourceforge.net provides a build of the [Mingw-w64 web installer](https://sourceforge.net/projects/mingw-w64/) which 
lets you select version, exception style, thread style, etc. I recommend installing GCC 7.3.0 x86_64-posix-seh-rev0.

This assumes a bin of a mingw-w64 is in the `PATH`.

Run `.\compile.ps1` to compile.

## References

 - [Chris Wellons' 01/31/16 blog article at nullprogramming.com](http://nullprogram.com/blog/2016/01/31/) discusses not
   linking against a standard library, instead using windows.h. This is to avoid the buggy Windows msvcrt.dll, which is
   an ancient C runtime. Instead he manually links to kernel32.dll which is essentially a standard library on its own. 
   One could also link to another clib, but many have GPL licensing and a light library can often get away without it.
 - [Win-builds.org](http://win-builds.org/doku.php/start) software packaging for Windows. 92 libraries and tools: Qt,
   Cairo, Curl, EFL, expat (XML parsing), ffmpeg, freetype, GTK+, jansson, json-c, libgcrypt, libsoup, lua, nettle, 
   openssl, pcre, SDL2, sqlite, wget, winpthreads, xz, zlib, and more.
   More libs: [cJSON](https://github.com/DaveGamble/cJSON) (simple json lib), a Python requests-like 
   [requests](https://github.com/mossberg/librequests) lib (Windows support might not exist--it will at least need 
   curl; worth a try), [stretchy buffer](https://github.com/nothings/stb/blob/master/stretchy_buffer.h), a C `<vector>`.
 - [Simple Dynamic Strings](https://github.com/antirez/sds) is a string library for C designed to augment the limited
   libc string handling functionalities by adding heap allocated strings that are:
   - Simpler to use.
   - Binary safe.
   - Computationally more efficient.
   - But yet... Compatible with normal C string functions.
 - Types... [Wiki: Handles and Data Types](https://en.wikibooks.org/wiki/Windows_Programming/Handles_and_Data_Types)
   - *HANDLE*: a reference. Usually an integer or pointer value that refers to some object, but you as the consumer 
     shouldn't care about the actual handle value.
     - A file, socket, bitmap, etc.
     - IntPtr in C#.
     - Most methods that deal with handles are in the Marshal class.  The .NET framework does its best to hide them from
       you by wrapping them in a class.  FileStream, Socket, Bitmap, etc.  If a class has a Dispose() method, there's 
       likely to be a handle stored in a private member.
   - *HINSTANCE*: handles to a program instance.
   - *DWORD*: typedef for an `unsigned long`
   - *TCHAR*: typedef for a standard `char` or `wchar_t` (Unicode), hold either standard 1-byte ASCII characters, or 
     wide 2-byte Unicode characters. 
   - *LPCTSTR*: typedef for a `const char *` or `const wchar_t *` (Unicode). `#define TCHAR * LPTSTR`
   - *CString*: typedef based on `CStringT`, an alternative to `std::string`
     - can grow as a result of concatenation operations
     - follow "value semantics." Think of a `CString` object as an actual string, not as a pointer to a string.
     - can freely substitute `CString` objects for `const char*` and `LPCTSTR` function arguments.
     - conversion operator gives direct access to the string's characters as a read-only array of characters (a C-style
       string).
 - [Get ip address](https://stackoverflow.com/questions/3989446/get-an-ip-address-string-from-getadaptersaddresses) with
   [GetAdaptersAddresses](https://msdn.microsoft.com/en-us/library/aa365915%28v=VS.85%29.aspx)
 - USER32.DLL implements the Windows USER component that creates and manipulates the standard elements of the Windows user interface, such as the desktop, windows, and menus. It thus enables programs to implement a graphical user interface (GUI) that matches the Windows look and feel.
 - kernel32.dll is the 32-bit dynamic link library found in the Windows operating system kernel. It handles memory management, input/output operations, and interrupts. When Windows boots up, kernel32.dll is loaded into a protected memory space so other applications do not take that space over.
 - [Symantec description of system32 files](https://www.symantec.com/connect/blogs/cwindowssystem32-files-explained)
 - [How to specify lib path](http://www.mingw.org/wiki/HOWTO_Specify_the_Location_of_Libraries_for_use_with_MinGW)
 - [VS Code Intellisense configuration](https://github.com/Microsoft/vscode-cpptools/blob/master/Documentation/Getting%20started%20with%20IntelliSense%20configuration.md)

[&#x2190; Back to Project List](../README.md)