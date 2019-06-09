package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yasukotelin/go-openurl"
)

const (
	addr = ":8080"
	url  = "http://localhost" + addr
)

const (
	errorExitCode   = 0
	successExitCode = 1
)

func main() {
	fmt.Println("open the database")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	}
	if err = initTableIfNotExists(db); err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	}

	routing()
	if err := listen(); err != nil {
		fmt.Println(err)
		os.Exit(errorExitCode)
	}
	os.Exit(successExitCode)
}

func routing() {
	http.Handle("/", http.FileServer(http.Dir("view/build/")))
}

func listen() error {
	fmt.Println("server started.")
	fmt.Printf("listen on the %s\n", url)
	fmt.Println()
	fmt.Println("stop server pressed Ctrl+C.")
	if err := openurl.OpenWithBrowser(url); err != nil {
		return err
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}

func initTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS accounts (id INTEGER PRIMARY KEY, personal_id TEXT, password TEXT)`)
	return err
}
