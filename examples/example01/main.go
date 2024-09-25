package main

import (
	ez "github.com/Rickiarty/EZLog_Go"
)

func main() {
	ez.EZLog("This is a detailed message")
	ez.EZLog("This is a detailed message", 1)
	ez.EZLog("This is a detailed message", 1, "DEBUG")
}
