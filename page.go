package gonsen

import (
	"bytes"
	"html/template"
	"net/http"
	"path"
)

type Page[T interface{}] struct {
	template *template.Template
}

func NewPage[T interface{}](source *Source, pageName string) *Page[T] {
	fullPath := path.Join(PageSubdirectory, pageName)
	fileContents := mustReadFileString(source.templates, fullPath)

	clone := template.Must(source.base.Clone())

	return &Page[T]{
		template: template.Must(clone.Parse(fileContents)),
	}
}

func (p *Page[T]) Render(data T) ([]byte, error) {
	var buf bytes.Buffer
	err := p.template.ExecuteTemplate(&buf, "base", data)

	return buf.Bytes(), err
}

func (p *Page[T]) HandlerWithSource(dataSource func(r *http.Request) (T, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, statusCode := dataSource(r)

		if statusCode != 200 {
			w.WriteHeader(statusCode)
			return
		}

		body, err := p.Render(data)

		if err != nil {
			// TODO: This probably shouldn't happen... better way to handle/notify?
			w.WriteHeader(500)
			return
		}

		w.Header().Add("Content-Type", "text/html")
		w.Write(body)
	}
}

func MustRenderStaticPage(source *Source, pageName string) []byte {
	p := NewPage[interface{}](source, pageName)

	contents, err := p.Render(nil)

	if err != nil {
		panic(err)
	}

	return contents
}
