package webservice

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
)

type Server struct {
	listener net.Listener
	iface    string
	port     string

	listenerMtx sync.Mutex
}

func NewServer(iface string, port string) *Server {
	return &Server{
		iface: iface,
		port:  port,
	}
}

func (s *Server) Start() {
	origin := &url.URL{Scheme: "http", Host: net.JoinHostPort(s.iface, s.port)}
	initPlayground(origin)
	http.HandleFunc("/", indexHandler)
	fs := http.FileServer(http.Dir(StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	initSSA()
	initExampleLoader(ExamplesDir)
	initCFSMExtract()
	initMigo()
	initGong()
	initGodel()
	initMigoinfer()
	initSynthesis()

	log.Printf("Listening at %s", s.URL())
	(&http.Server{}).Serve(s.Listener())
}

func (s *Server) Close() {
	s.Listener().Close()
}

func (s *Server) URL() string {
	return fmt.Sprintf("http://%s/", s.Listener().Addr())
}

func (s *Server) Listener() net.Listener {
	s.listenerMtx.Lock()
	defer s.listenerMtx.Unlock()

	if s.listener != nil {
		return s.listener
	}

	ifaceAndPort := fmt.Sprintf("%v:%v", s.iface, s.port)
	listener, err := net.Listen("tcp4", ifaceAndPort)
	if err != nil {
		log.Fatal(err)
	}

	s.listener = listener
	return s.listener
}
