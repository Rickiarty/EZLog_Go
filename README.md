# EZLog_Go
an easy logger without setting via Go Modules on GitHub platform

```
go get -u github.com/Rickiarty/EZLog_Go@latest
```

// go.mod
```
module github.com/yourusername/GoTEST


go 1.22


require (
    github.com/Rickiarty/EZLog_Go v1.0.0
)
```

// main.go
```
package main


// go.mod
module github.com/yourusername/GoTEST

go 1.22

require (
    github.com/Rickiarty/EZLog_Go v1.0.1
)


// main.go
package main

import (

    ...
<<<<<<< HEAD
    ez "github.com/Rickiarty/EZLog_Go/ezlog"
    ...
)

func main() {
	ez.EZLog("This is a detailed message")
	ez.EZLog("This is a detailed message", 1)
	ez.EZLog("This is a detailed message", 1, "DEBUG")
}
=======
    
    ez "github.com/Rickiarty/EZLog_Go/ezlog"
    
    ...
    
)

func main() {

	ez.EZLog("This is a detailed message")
 
	ez.EZLog("This is a detailed message", 1)
 
	ez.EZLog("This is a detailed message", 1, "DEBUG")
 
}
```
>>>>>>> 4429169328fea3d7890c3c5a0fa11b4e59e9bcb6
