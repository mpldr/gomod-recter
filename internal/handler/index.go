package handler

import (
	"html/template"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"internal/data"
)

func indexHandler(ctx *fasthttp.RequestCtx) {
	templateDir := viper.GetString("Directories.TemplateDir")
	if templateDir == "" {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("No template directory set! Please set one.")
		return
	}

	tmpl, err := template.ParseFiles(templateDir + "index.tmpl")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("Template '%s' could not be parsed: %v", templateDir+"index.tmpl", err)
		return
	}

	dataset := data.GetProjectList()
	err = tmpl.Execute(ctx.Response.BodyWriter(), struct{ Data interface{} }{Data: dataset})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.WriteString("sorry mate, internal server error")
		glog.Errorf("could not execute template: %v", err)
		return
	}
}
