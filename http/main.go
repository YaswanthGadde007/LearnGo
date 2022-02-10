package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Printf("%+v", res)
	// body, err := io.ReadAll(res.Body)
	// fmt.Println(body, err)

	lw := logWriter{}

	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
