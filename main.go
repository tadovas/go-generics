package main

import (
	"fmt"
	"github.comm/tadovas/go-generics/http_json"
	"net"
	"net/http"
	"os"
)

func runMe() error {
	l, err := net.Listen("tcp", ":6161")
	if err != nil {
		return fmt.Errorf("listener: %v", err)
	}
	var h http_json.JsonHandlerFunc = func(input http_json.SomeStruct) http_json.HttpResponse {
		return http_json.HttpResponse{Code: http.StatusOK}
	}
	fmt.Println("Will serve you now")
	return http.Serve(l, http_json.HandleJson[http_json.SomeStruct, http_json.HttpResponse](h))
}

func main() {
	if err := runMe(); err != nil {
		fmt.Printf("Main error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
