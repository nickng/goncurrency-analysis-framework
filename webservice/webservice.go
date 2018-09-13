package webservice

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
)

// Handler is a convenient type to keep all handler initialision in one place.
type Handler struct {
	Name     string
	InitFunc func()
}

type Config struct {
	Synthesis bool // CFSMs synthesis.
	Gong      bool // Gong verifier.
	Godel     bool // Godel checker.

	ExamplesDir string // Demo example dir.
	StaticDir   string // Static files dir.
	TemplateDir string // HTML Template dir.

	Handlers []Handler
}

func (cfg *Config) LoadHandlers() {
	cfg.Handlers = append(cfg.Handlers, Handler{"SSA", initSSA})
	if cfg.Synthesis {
		cfg.Handlers = append(cfg.Handlers,
			Handler{"CFSM extraction", initCFSMExtract},
			Handler{"CFSM synthesis", initSynthesis})
	}
	if cfg.Gong {
		cfg.Handlers = append(cfg.Handlers,
			Handler{"MiGo inference v1", initMigoV1},
			Handler{"Gong verifier", initGong},
		)
	}
	if cfg.Godel {
		cfg.Handlers = append(cfg.Handlers,
			Handler{"MiGo inference v2", initMigoV2},
			Handler{"Godel checker", initGodel},
		)
	}
}

// Server encapsulates all parameters required to run the webservice.
type Server struct {
	listener net.Listener
	iface    string
	port     string

	cfg Config
	mu  sync.Mutex
}

// NewServer creates a new instance of server.
func NewServer(iface, port string, cfg Config) *Server {
	s := Server{
		iface: iface,
		port:  port,
		cfg:   cfg,
	}
	return &s
}

func (s *Server) Close() {
	s.Listener().Close()
}

// Start runs the server and starts listening for incoming connection.
func (s *Server) Start() {
	origin := &url.URL{Scheme: "http", Host: net.JoinHostPort(s.iface, s.port)}
	initPlayground(origin)

	WebEditor{s.cfg}.init()
	for _, m := range s.cfg.Handlers {
		m.InitFunc()
		log.Printf("Initialised %s handler", m.Name)
	}
	log.Printf("Listening at %s", s.URL())
	defer s.Close()
	(&http.Server{}).Serve(s.Listener())
}

// Listener creates or returns the existing listener associated to this server.
func (s *Server) Listener() net.Listener {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.listener != nil {
		return s.listener
	}
	// don't use
	ln, err := net.Listen("tcp4", net.JoinHostPort("", s.port))
	if err != nil {
		log.Fatal("cannot listen:", err)
	}
	s.listener = ln
	return s.listener
}

func (s *Server) URL() string {
	return fmt.Sprintf("http://%s", s.Listener().Addr())
}
