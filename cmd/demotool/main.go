package main

import (
	"flag"
	"go/build"
	"log"
	"os"
	"path"

	"github.com/nickng/goncurrency-analysis-framework/webservice"
)

var (
	addr string // Listen interface.
	port string // Listen port.

	examplesDir     string
	staticDir       string
	templateDir     string
	enableSynthesis bool
	enableGong      bool
	enableGodel     bool

	SREPLS6    webservice.Config
	GolangUK17 webservice.Config
)

const basePkg = "github.com/nickng/goncurrency-analysis-framework"

func init() {
	envport := os.Getenv("PORT")
	if envport == "" {
		envport = "6060"
	}
	flag.StringVar(&addr, "bind", "127.0.0.1", "Bind address")
	flag.StringVar(&port, "port", envport, "Listen port")
	p, err := build.Default.Import(basePkg, "", build.FindOnly)
	if err != nil {
		log.Fatal("Could not find base path")
	}
	basePath := p.Dir
	flag.StringVar(&templateDir, "templates", path.Join(basePath, "templates"), "Templates directory")
	flag.StringVar(&staticDir, "static", path.Join(basePath, "static"), "Static files directory")
	flag.StringVar(&examplesDir, "examples", path.Join(basePath, "examples", "default"), "Examples directory")
	flag.BoolVar(&enableSynthesis, "synthesis", false, "Enable CFSM synthesis")
	flag.BoolVar(&enableGong, "gong", false, "Enable Gong verification")
	flag.BoolVar(&enableGodel, "godel", true, "Enable Godel checker")
}

func main() {
	flag.Parse()
	cfg := webservice.Config{
		ExamplesDir: examplesDir,
		StaticDir:   staticDir,
		TemplateDir: templateDir,
		Synthesis:   enableSynthesis,
		Gong:        enableGong,
		Godel:       enableGodel,
	}
	cfg.LoadHandlers()

	server := webservice.NewServer(addr, port, cfg)
	server.Start()
	server.Close()
}
