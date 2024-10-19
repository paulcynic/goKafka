package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type Producer interface {
	Start() error
}

type HTTPProducer struct {
	listenAddr string
	server     *Server
	producech  chan<- Message
}

func NewHTTPProducer(listenAddr string, producech chan Message) *HTTPProducer {
	return &HTTPProducer{
		listenAddr: listenAddr,
		producech:  producech,
	}
}

func (p *HTTPProducer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// path := fmt.Sprintf(r.URL.Path)
	var (
		path  = strings.TrimPrefix(r.URL.Path, "/")
		parts = strings.Split(path, "/")
	)
	// commit
	if r.Method == "GET" {
	}
	if r.Method == "POST" {
		if len(parts) != 2 {
			fmt.Println("invalid action")
			return
		}
		p.producech <- Message{
			Data:  []byte("we dont know yet"),
			Topic: parts[1],
		}
	}
	fmt.Println(parts)
	// w.Write([]byte(path))
}

func (p *HTTPProducer) Start() error {
	slog.Info("HTTP transport started", "port", p.listenAddr)
	http.ListenAndServe(p.listenAddr, p)

	return nil
}
