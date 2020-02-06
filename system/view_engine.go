package system

import (
	"fmt"
	"html/template"
)

const (
	endPoint = "views\\"
)

type view struct {
	Errors []string
	//basePath string
}

type ViewValues map[string]interface{}

func (v view) renderView(_template string, data ViewValues) {

	if data == nil {
		data = make(ViewValues)
	}

	data["Errors"] = v.Errors

	fullPath := GetApplication().config.viewBasePath + _template + ".gohtml"

	tmpl := template.Must(template.New(_template + ".gohtml").Funcs(template.FuncMap{
		"hasError": func() bool {
			return len(v.Errors) > 0
		},
	}).ParseFiles(fullPath))
	tmpl.Execute(GetApplication().response.rw, data)
	return
}

func View(template string, data ViewValues) {
	fmt.Println("Request: ", GetRequest().All())
	fmt.Println("wqe: ", GetRequest().Validate())
	GetApplication().view.renderView(template, data)
}
