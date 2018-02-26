# Package Sharpener

I created this commandline tool to remove duplicate files in the packages and bin child folders to the working directory. The publish command 
below also demonstrates how to compile a self contained exe from a .NET Core project. Note that this project doesn't actually deduplicate successfully because the packages folder is actually arranged by type of architecture, so I would need to add some additional logic. Luckily, I am no longer cleaning up after people using manual references instead of Nuget.

Build an exe into .\bin\Release\netcoreapp2.0\win10-x64\

```
dotnet publish -c Release --self-contained --runtime win10-x64
```

Add to path for this session

```
$env:Path = (Get-Item -Path ".\" -Verbose).FullName + "\bin\Release\netcoreapp2.0\win10-x64;$env:Path"
```

Or add `"cwd": "<path of a .net project>"` to launch.json

[&#x2190; Back to Project List](../README.md)