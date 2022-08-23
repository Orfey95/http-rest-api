package main

import (
	"log"

	"github.com/Orfey95/http-rest-api/cmd/apiserver/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
