package gonsen

import (
	"bytes"
	"html/template"
	"net/http"
	"path"
)

type Page[T interface{}] struct {
	template   *template.Template
	dataSource func(r *http.Request) (T, int)
}

func NewPage[T interface{}](source *Source, pageName string, dataSource func(r *http.Request) (T, int)) *Page[T] {
	fullPath := path.Join(PageSubdirectory, pageName)
	fileContents := mustReadFileString(source.templates, fullPath)

	clone := template.Must(source.base.Clone())

	return &Page[T]{
		template:   template.Must(clone.Parse(fileContents)),
		dataSource: dataSource,
	}
}

func (p *Page[T]) Render(data T) ([]byte, error) {
	var buf bytes.Buffer
	err := p.template.ExecuteTemplate(&buf, "base", data)

	return buf.Bytes(), err
}

func (p *Page[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data T
	if p.dataSource != nil {
		var statusCode int
		data, statusCode = p.dataSource(r)

		if statusCode != 200 {
			w.WriteHeader(statusCode)
			return
		}
	}

	body, err := p.Render(data)

	if err != nil {
		// TODO: This probably shouldn't happen... better way to handle/notify?
		w.WriteHeader(500)
		return
	}

	addHtmlHeaders(w)
	w.Write(body)
}
