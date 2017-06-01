Download Mingw (32 bit only) installer: https://sourceforge.net/projects/mingw/files/latest/download?source=files

Once the download completes, run mingw-get-setup.exe (run as administrator if you need to install for all users).

MinGW-w64: https://sourceforge.net/projects/mingw-w64/ [Options explained](https://stackoverflow.com/questions/29947302/meaning-of-options-in-mingw-w64-installer)

Select a location to install. I used the default. For these settings it recommended `C:\Program Files\mingw-w64\x86_64-7.1.0-posix-sjlj-rt_v5-rev0`

Add the bin folder to your system path. This will confuse any applications requiring 32 bit GCC, so if this applies to you, don't add it to path and instead modify your gcc environmental variable to point here. My bin folder was right here: `C:\Program Files\mingw-w64\x86_64-7.1.0-posix-sjlj-rt_v5-rev0\mingw64\bin`

In the Windows start bar, type `system var` and select to edit the environmental variables. The "Advanced" tab of the "System Properties" dialog should appear.

Click the button near the bottom titled "Environmental Variables".

To make it available to all users, click the "Edit..." button after selecting Path under the "System variables" control group. (If you only want GCC added to path for yourself, edit the Path under "User variables" instead.)

If your terminal (Powershell, CMD, or an integrated terminal in your IDE) was already open, you'll need to close and re-open it to get your new environmental variables loaded.

Fetch and install Pixel and all dependencies:

```
go get -u github.com/go-gl/glfw/v3.2/glfw
go get -u github.com/palumacil/misc/pixgame...
```