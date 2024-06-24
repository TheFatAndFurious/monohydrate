package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Print("error geting the wd")
	}
	fmt.Printf("wd is %q", path)
	server := NewServer(":8080")
	server.Run()

}

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("coucou"))
	})
	server := http.Server {
		Addr: s.listenAddr,
		Handler: router,
	}
	log.Print("server has been started")
	return server.ListenAndServe()
}

