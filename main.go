package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":1234"

func main() {
	http.HandleFunc("/food", handlerFood)

	fmt.Println("server started at port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handlerFood(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Query().Get("nomer")))
}
