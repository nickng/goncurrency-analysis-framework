package main

import (
	"flag"
	"go/build"
	"log"
	"path"

	"github.com/nickng/demotool/webservice"
)

var (
	addr string // Listen interface.
	port string // Listen port.
)

const basePkg = "github.com/nickng/demotool"

func init() {
	flag.StringVar(&addr, "bind", "127.0.0.1", "Bind address")
	flag.StringVar(&port, "port", "6060", "Listen port")
	p, err := build.Default.Import(basePkg, "", build.FindOnly)
	if err != nil {
		log.Fatal("Could not find base path")
	}
	basePath := p.Dir
	flag.StringVar(&webservice.TemplateDir, "templates",
		path.Join(basePath, "templates"), "Templates directory")
	flag.StringVar(&webservice.StaticDir, "static",
		path.Join(basePath, "static"), "Static files directory")
	flag.StringVar(&webservice.ExamplesDir, "examples",
		path.Join(basePath, "examples", "srepls6"), "Examples directory")
	flag.Parse()
}

func main() {
	server := webservice.NewServer(addr, port)
	server.Start()
	server.Close()
}
