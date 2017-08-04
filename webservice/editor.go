package webservice

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type WebEditor struct {
	Config
}

func (p WebEditor) init() {
	http.HandleFunc("/", p.indexHandler)
	fs := http.FileServer(http.Dir(p.StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/load", p.loadHandler)
	log.Print("Initialised web editor,example loader")
}

func (p WebEditor) indexHandler(w http.ResponseWriter, req *http.Request) {
	var examples []string
	t, err := template.ParseFiles(path.Join(p.TemplateDir, "index.tmpl"))
	if err != nil {
		NewErrInternal(err, "Cannot load template").Report(w)
		return
	}
	d, err := ioutil.ReadDir(p.ExamplesDir)
	if err != nil {
		NewErrInternal(err, "Cannot read examples").Report(w)
		return
	}
	for _, f := range d {
		if f.IsDir() {
			examples = append(examples, f.Name())
		}
	}

	data := struct {
		Title     string
		Examples  []string
		Synthesis bool
		Gong      bool
		Godel     bool
	}{
		Title:     "Go verification framework",
		Examples:  examples,
		Synthesis: p.Synthesis,
		Gong:      p.Gong,
		Godel:     p.Godel,
	}
	err = t.Execute(w, data)
	if err != nil {
		NewErrInternal(err, "Template execute failed").Report(w)
		return
	}
}

func (p WebEditor) loadHandler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		NewErrInternal(err, "Cannot read input").Report(w)
		return
	}
	if err := req.Body.Close(); err != nil {
		NewErrInternal(err, "Cannot close request").Report(w)
		return
	}
	log.Println("Load example:", string(b))
	file, err := os.Open(path.Join(p.ExamplesDir, string(b), "main.go"))
	if err != nil {
		NewErrInternal(err, "Cannot open file").Report(w)
		return
	}
	io.Copy(w, file)
}
