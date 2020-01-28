package system

import (
	"html/template"
)

const (
	endPoint = "views\\"
)

type view struct {
	basePath string
}

func (v *view) renderView(_template string, data interface{}) {

	fullPath := GetApplication().view.basePath + _template + ".gohtml"
	tmpl := template.Must(template.ParseFiles(fullPath))
	tmpl.Execute(GetApplication().response.rw, data)
	return
}

func View(template string, data interface{}) {
	GetApplication().view.renderView(template, data)
}
