package gonsen

import (
	"bytes"
	"html/template"
	"net/http"
	"path"
)

type Page[T interface{}, C interface{}] struct {
	template   *template.Template
	dataSource func(r *http.Request) (T, int)
	ctxSource  func(r *http.Request) (C, int)
}

func NewPage[T interface{}](
	source *Source,
	pageName string,
	dataSource func(r *http.Request) (T, int),
) *Page[T, interface{}] {

	return NewPageWithContext[T, interface{}](
		source,
		pageName,
		dataSource,
		nil,
	)
}

func NewPageWithContext[T interface{}, C interface{}](
	source *Source,
	pageName string,
	dataSource func(r *http.Request) (T, int),
	ctxSource func(r *http.Request) (C, int),
) *Page[T, C] {
	fullPath := path.Join(PageSubdirectory, pageName)
	fileContents := mustReadFileString(source.templates, fullPath)

	clone := template.Must(source.base.Clone())

	return &Page[T, C]{
		template:   template.Must(clone.Parse(fileContents)),
		dataSource: dataSource,
		ctxSource:  ctxSource,
	}
}

func (p *Page[T, C]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data T
	if p.dataSource != nil {
		var statusCode int
		data, statusCode = p.dataSource(r)

		if statusCode != 200 {
			w.WriteHeader(statusCode)
			return
		}
	}

	var ctx C

	if p.ctxSource != nil {
		var statusCode int
		ctx, statusCode = p.ctxSource(r)

		if statusCode != 200 {
			w.WriteHeader(statusCode)
			return
		}
	}

	body, err := p.render(data, ctx)

	if err != nil {
		// TODO: This probably shouldn't happen... better way to handle/notify?
		w.WriteHeader(500)
		return
	}

	addHtmlHeaders(w)
	w.Write(body)
}

type pageData struct {
	Data    interface{}
	Context interface{}
}

func (p *Page[T, C]) render(data T, ctx C) ([]byte, error) {

	var buf bytes.Buffer
	err := p.template.ExecuteTemplate(&buf, "base", pageData{
		Data:    data,
		Context: ctx,
	})

	return buf.Bytes(), err
}
