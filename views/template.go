package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/csrf"
	"github.com/sinasezza/go-web-dev/context"
	"github.com/sinasezza/go-web-dev/models"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFs(fs fs.FS, patterns ...string) (Template, error) {
	// Create a new template with the name of the first pattern
	tpl := template.New(filepath.Base(patterns[0]))

	// Add the csrfField function
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", fmt.Errorf("csrfField Not Implemented")
		},
		"currentUser": func() (template.HTML, error) {
			return "", fmt.Errorf("currentUser Not Implemented")
		},
	})

	// Parse all the files
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		HtmlTpl: tpl,
	}, nil
}

type Template struct {
	HtmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.HtmlTpl.Clone()
	if err != nil {
		log.Printf("Cloning Template: %v", err)
		http.Error(w, "There is an error cloning the template!", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
		"currentUser": func() *models.User {
			return context.User(r.Context())
		},
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("Executing Template: %v", err)
		http.Error(w, "There is an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
