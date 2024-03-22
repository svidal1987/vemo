package main

import (
	"fmt"
	"net/http"

	"github.com/vemo/handler"
)

func main() {

	mux := http.NewServeMux()
	handler.HandlerToDo(mux)

	fmt.Println("Listen :8080")
	http.ListenAndServe(":8080", mux)
}
