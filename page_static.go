package gonsen

import "net/http"

type StaticPage struct {
	contents []byte
}

func NewStaticPage(source *Source, pageName string) *StaticPage {
	p := NewPage[interface{}](source, pageName, nil)

	contents, err := p.Render(nil)

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
