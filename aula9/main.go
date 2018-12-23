package main

import (
	"log"
	"net/http"

	"./routers"
)

func main() {
	router := routers.GetRouters()
	log.Fatal(http.ListenAndServe(":8000", router))
}
