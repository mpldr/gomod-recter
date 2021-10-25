package handler

import (
	"fmt"
	"html/template"
	"internal/data"
)

func getTemplateWithHelper(proj *data.Project) *template.Template {
	return template.New("").Funcs(template.FuncMap{"ModuleHeader": func() string {
		return fmt.Sprintf("%s %s %s", proj.RootPath, proj.VCS, proj.Repo)
	}})
}
