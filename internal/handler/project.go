package handler

import (
	"bytes"
	"fmt"

	"mpldr.codes/recter/internal/data"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func projectHandler(ctx *fasthttp.RequestCtx, project string, remainingPath []byte) {
	glog.Tracef("remainder: %s", remainingPath)

	dataset, ok := data.GetProjectList()[project]
	if !ok {
		glog.Tracef("found: %v", data.GetProjectList())
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("for some reason the dataset for the project '%s' was not found.", project)
		return
	}
	glog.Tracef("Dataset for %s: %v", project, dataset)

	if ctx.URI().QueryArgs().GetBool("go-get") {
		glog.Debug("detected go-get, sending project details")
		ctx.WriteString(fmt.Sprintf(`<html><head><meta name="go-import" content="%s %s %s">%s</head></html>`, dataset.RootPath, dataset.VCS, dataset.Repo, dataset.GetGoSource()))
		return
	}

	if bytes.HasPrefix(remainingPath, []byte("/api")) {
		apiHandler(ctx, &dataset, remainingPath[4:])
		return
	}

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
