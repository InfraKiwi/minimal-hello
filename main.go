package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello world! OS: %s, arch: %s", runtime.GOOS, runtime.GOARCH)
}
