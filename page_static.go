package gonsen

import "net/http"

type StaticPage struct {
	contents []byte
}

func NewStaticPageWithContext[C interface{}](source *Source, pageName string, ctxSource func(r *http.Request) (C, int)) *Page[interface{}, C] {
	return NewPageWithContext[interface{}](source, pageName, nil, ctxSource)
}

func NewStaticPage(source *Source, pageName string) *StaticPage {
	p := NewPage[interface{}](source, pageName, nil)

	contents, err := p.render(nil, nil)

	if err != nil {
		panic(err)
	}

	return &StaticPage{
		contents: contents,
	}
}

func (p *StaticPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addHtmlHeaders(w)
	w.Write(p.contents)
}
