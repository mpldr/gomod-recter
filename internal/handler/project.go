package handler

import (
	"internal/data"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func projectHandler(ctx *fasthttp.RequestCtx, project string) {
	dataset, ok := data.GetProjectList()[project]
	if !ok {
		glog.Tracef("found: %v", data.GetProjectList())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("for some reason the dataset for the project '%s' was not found.", project)
		return
	}
	glog.Tracef("Dataset for %s: %v", project, dataset)

	if dataset.Redirect {
		ctx.Redirect(dataset.Repo, fasthttp.StatusSeeOther)
	}

	templateDir := viper.GetString("Directories.TemplateDir")
	if templateDir == "" {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("No template directory set! Please set one.")
		return
	}

	tmpl, err := getTemplateWithHelper(&dataset).ParseFiles(templateDir + "project.tmpl")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("Template '%s' could not be parsed: %v", templateDir+"project.tmpl", err)
		return
	}

	err = tmpl.ExecuteTemplate(ctx.Response.BodyWriter(), "project.tmpl", dataset)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("could not execute template: %v", err)
		return
	}
}
