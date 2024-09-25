# EZLog_Go
an easy logger without setting via Go Modules on GitHub platform

to clean cache:
```
go clean -modcache
```
to download modules and generate a "go.sum" file:
```
go mod tidy
```
to download the module:
```
go get -u github.com/Rickiarty/EZLog_Go@latest
```

// go.mod
```
module github.com/yourusername/GoTEST

go 1.22

require (
    github.com/Rickiarty/EZLog_Go v1.0.2
)
```

// main.go
```
package main

import (

    ...

    ez "github.com/Rickiarty/EZLog_Go/ezlog"

    ...

)

func main() {
	ez.EZLog("This is a detailed message")
	ez.EZLog("This is a detailed message", 1)
	ez.EZLog("This is a detailed message", 1, "DEBUG")
}
```
