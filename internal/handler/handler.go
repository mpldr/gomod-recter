package handler

import (
	"bytes"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"internal/data"
)

func FasthttpHandler(ctx *fasthttp.RequestCtx) {
	glog.Infof("%s requested %s using %s", ctx.RemoteIP(), ctx.Path(), ctx.UserAgent())
	ctx.SetContentTypeBytes([]byte("text/html"))

	path := bytes.Split(ctx.Path(), []byte("/"))

	glog.Debug(path)
	if len(path) < 2 || len(path[1]) == 0 {
		indexHandler(ctx)
		return
	}

	if name := viper.GetString("Projects." + string(path[1]) + ".Name"); name != "" {
		templateDir := viper.GetString("Directories.TemplateDir")
		if templateDir == "" {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.WriteString("sorry mate, internal server error")
			glog.Errorf("No template directory set! Please set one.")
			return
		}

		dataset, ok := data.GetProjectList()[string(path[1])]
		if !ok {
			glog.Tracef("found: %v", data.GetProjectList())
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.WriteString("sorry mate, internal server error")
			glog.Errorf("for some reason the dataset for the project '%s' was not found.", path[1])
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
	} else {
		glog.Warnf("project '%s' not found in config", path[1])
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.Redirect("/", fasthttp.StatusSeeOther)
	}
}
