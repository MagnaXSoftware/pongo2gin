// Package pongo2gin is a template renderer that can provides a Gin-compatible
// template renderer using the Pongo2 template library https://github.com/flosch/pongo2
package pongo2gin

import (
	"fmt"
	"net/http"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// Pongo2 is a custom Gin template renderer using Pongo2.
type Pongo2 struct {
	Suffix string

	set *pongo2.TemplateSet
}

// New creates a new Pongo2 instance with custom Options.
func New(set *pongo2.TemplateSet, suffix string) *Pongo2 {
	return &Pongo2{set: set, Suffix: suffix}
}

// Default creates a Pongo2 instance with default options.
func Default() *Pongo2 {
	return New(pongo2.DefaultSet, ".html.twig")
}

// Instance should return a new Pongo2 struct per request and prepare
// the template by either loading it from disk or using pongo2's cache.
func (p Pongo2) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	path := fmt.Sprintf("%s%s", name, p.Suffix)

	if gin.Mode() == gin.DebugMode {
		template = pongo2.Must(p.set.FromFile(path))
	} else {
		template = pongo2.Must(p.set.FromCache(path))
	}

	return &Render{
		tpl:  template,
		data: data.(pongo2.Context),
	}
}

type Render struct {
	tpl  *pongo2.Template
	data pongo2.Context
}

// Render should render the template to the response.
func (r Render) Render(w http.ResponseWriter) error {
	if r.data == nil {
		r.data = make(pongo2.Context)
	}
	return r.tpl.ExecuteWriter(r.data, w)
}

// WriteContentType should add the Content-Type header to the response
// when not set yet.
func (r Render) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header.Set("content-type", "text/html; charset=utf8")
	}
}
