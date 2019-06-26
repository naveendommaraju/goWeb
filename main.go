package main

import (
	"fmt"
	"net/http"
	"os"

	"goWeb/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:9234", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
