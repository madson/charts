package main

import (
	"net/http"
)

func (s *server) configureRoutes() {
	s.router.Handle("/", handlerListFiles())
}

func handlerListFiles() http.Handler {
	dir := http.Dir("/Users/madson/workspace/logictake/pie_charts/output")
	fs := http.FileServer(dir)
	return fs
}
