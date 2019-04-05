package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/yasukotelin/go-openurl"
)

const (
	port = ":8080"
)

var (
	exitCode = 0
)

func main() {
	routing()
	listen()
	os.Exit(exitCode)
}

func getRunURL() string {
	return fmt.Sprintf("http://localhost%s", port)
}

func routing() {
	http.Handle("/", http.FileServer(http.Dir("public")))
}

func listen() {
	fmt.Println("server started.")
	fmt.Printf("listen on the localhost%s\n", port)
	fmt.Println()
	fmt.Println("stop server pressed Ctrl+C.")
	if err := openurl.OpenWithBrowser(getRunURL()); err != nil {
		fmt.Println(err)
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
		exitCode = 1
		return
	}
}
