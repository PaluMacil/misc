# Git Crawl (gitc)

## Overview

Crawler for bulk git operations (in progress)

This library is a very simple tool for moving from TFS to Git en masse! It isn't full of features because there are only a few simple tasks I needed to execute, but hopefully it can assist others with the simple tasks it accomplished for me. While this isn't a very in depth project, I will fix issues and use it in production with a binary download link made available as well.

## Install

Git is not required. Run the following command:

```
go get github.com\PaluMacil\misc\gitc...
```

## Run

If installed, use `gitc -h` to see usage.

## Git ammending

This command is my experimentation with modifying cloned repos. The below parameters would clone a repo and create a commit for each day from the since date to present. This works on private repos.

```
go run main.go touch -email me@gmail.com -name "My Name" -password spaghetti -repo https://github.com/MrToast/important -since 170923 -user MrToast
```

[&#x2190; Back to Project List](../README.md)