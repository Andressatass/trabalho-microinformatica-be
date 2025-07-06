package handlers

import (
	"log"
	"net/http"
)

func StartServer(port string) error {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		return err
	}

	log.Printf("start server on the port: %s", port)
	return nil
}
