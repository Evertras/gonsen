package gonsen

import (
	"html/template"
	"io"
	"io/fs"
	"net/http"
)

const (
	TemplateNameBase = "base"

	BaseFilename     = "site/base.html"
	PageSubdirectory = "site/pages"
)

type Source struct {
	templates          fs.FS
	base               *template.Template
	statusCodeBehavior map[int]func(w http.ResponseWriter, r *http.Request)
}

func NewSource(templates fs.FS) *Source {
	g := &Source{
		base:      template.Must(template.New(TemplateNameBase).Parse(mustReadFileString(templates, BaseFilename))),
		templates: templates,

		statusCodeBehavior: make(map[int]func(w http.ResponseWriter, r *http.Request)),
	}

	return g
}

func (s *Source) OnStatus(code int, do func(w http.ResponseWriter, r *http.Request)) {
	s.statusCodeBehavior[code] = do
}

func (s *Source) AssetsHandler() http.Handler {
	assetFS, err := fs.Sub(s.templates, "site")

	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(assetFS))
}

func mustReadFileString(fs fs.FS, file string) string {
	f, err := fs.Open(file)

	if err != nil {
		panic(err)
	}

	raw, err := io.ReadAll(f)

	return string(raw)
}
