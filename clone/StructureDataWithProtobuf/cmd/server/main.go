package main

import (
	"gostudy/clone/StructureDataWithProtobuf/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer("8080")
	log.Fatal(srv.ListenAndServe())
}
