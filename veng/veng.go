package veng

import (
	"net/http"
	"fmt"
	"html/template"
)

// Conf is a configuration struct for view engine
type Conf struct {
	BaseDir string
	PagesDir string
	LayoutsDir string
}

type viewEngine struct {
	layout *template.Template
	templates map[string]*template.Template
	config *Conf
}

//Package globals
var eng *viewEngine


//Initialize preapares view engine for use
func Initialize(conf *Conf) error {
	if conf == nil {
		conf = &Conf{"views", "pages", "layouts"}
	}

	eng = &viewEngine{
		templates: make(map[string]*template.Template, 0),
		config: conf,
	}

	if err := loadLayout(); err != nil {
		return err
	}

	if err := loadTemplates(); err != nil {
		return err
	}

	return nil
}

func Render(w http.ResponseWriter, templateName string, payload interface{}) {
	tmpl := eng.templates[templateName]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	tmpl.Execute(w, payload)
}

// Private functions
func badInitialization() bool {
	return eng == nil || eng.config == nil || eng.templates == nil
}

func loadLayout() error {
	files, err := readDirHtmlFiles(eng.config.BaseDir)
	if err != nil {
		return err
	}

	layout := template.New("layout")
	layout, _ = layout.ParseFiles(files[0])
	eng.layout = layout

	return nil
}

func loadTemplates() error {
	files, err := walkDirFiles(fmt.Sprintf("./%s/%s", eng.config.BaseDir, eng.config.PagesDir))
	if err != nil {
		return err
	}

	for k, file := range files {
		templ, _ := eng.layout.Clone()
		k = getFileNameWithoutExt(k)
		eng.templates[k] = template.Must(templ.ParseFiles(file))
	}

	return nil
}


