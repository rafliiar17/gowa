package views

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
)

//go:embed index.html
var indexRaw embed.FS

//go:embed assets components
var viewsRaw embed.FS

type gowaFS struct {
	underlying fs.FS
}

func (g gowaFS) Open(name string) (fs.File, error) {
	cleaned := name
	if strings.HasPrefix(name, "views/") {
		cleaned = strings.TrimPrefix(name, "views/")
	} else if name == "views" {
		cleaned = "."
	}
	return g.underlying.Open(cleaned)
}

func (g gowaFS) ReadDir(name string) ([]fs.DirEntry, error) {
	cleaned := name
	if strings.HasPrefix(name, "views/") {
		cleaned = strings.TrimPrefix(name, "views/")
	} else if name == "views" {
		cleaned = "."
	}
	if rdf, ok := g.underlying.(fs.ReadDirFS); ok {
		return rdf.ReadDir(cleaned)
	}
	return nil, fmt.Errorf("read dir not supported")
}

var EmbedIndexRaw fs.FS = gowaFS{underlying: indexRaw}
var EmbedViewsRaw fs.FS = gowaFS{underlying: viewsRaw}
