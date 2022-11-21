package routers

import (
	"fmt"
	"html/template"
)

// parseTemplateFile
func parseTemplateFile(filesname ...string) (tmpl *template.Template){
	tmpl = template.New("layout")
	var files []string
	for _, file := range filesname {
		files = append(files, fmt.Sprintf("template/%s.html", file))
	}
	tmpl = template.Must(tmpl.ParseFiles(files...))
	return
}