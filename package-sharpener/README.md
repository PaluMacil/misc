Build an exe into .\bin\Release\netcoreapp2.0\win10-x64\

```
dotnet publish -c Release --self-contained --runtime win10-x64
```

Add to path for this session

```
$env:Path = (Get-Item -Path ".\" -Verbose).FullName + "\bin\Release\netcoreapp2.0\win10-x64;$env:Path"
```

Or add `"cwd": "<path of a .net project>"` to launch.json