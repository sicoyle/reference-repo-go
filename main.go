package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/reference-repo-go/api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("\n%s", err))
		os.Exit(1)
	}
}

func run() error {
	s := server.New()
	s.Routes()
	log := logrus.New()
	port := "8080"
	log.Infof("reference-repo server starting on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), s)
	return nil
}
