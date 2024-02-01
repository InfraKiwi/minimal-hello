package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	argsWithoutProg := os.Args[1:]
	command := ""
	if len(argsWithoutProg) > 0 {
		command = argsWithoutProg[0]
	}

	switch command {
	case "serve":
		serve()
	default:
		fmt.Printf("Hello world! OS: %s, arch: %s", runtime.GOOS, runtime.GOARCH)
	}

}

func serve() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Got request!")
		writer.Write([]byte("hello!"))
	})
	http.ListenAndServe(":8090", nil)
}
