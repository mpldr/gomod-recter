package handler

import (
	"bytes"
	"time"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func FasthttpHandler(ctx *fasthttp.RequestCtx) {
	t := time.Now()
	defer func(t time.Time) { glog.Debugf("request took %s", time.Since(t)) }(t)
	remoteAddr := ctx.RemoteIP().String()

	if string(ctx.Method()) != "GET" && string(ctx.Method()) != "HEAD" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if header := viper.GetString("IPHeaderField"); header != "" {
		remoteAddr = string(ctx.Request.Header.Peek(header))
	}

	glog.Infof("%s requested %s using %s", remoteAddr, ctx.Path(), ctx.UserAgent())
	glog.Tracef("Headers: %v", &ctx.Request.Header)
	ctx.SetContentTypeBytes([]byte("text/html; charset=utf-8"))

	path := bytes.Split(ctx.Path(), []byte("/"))

	glog.Debug(path)
	if len(path) < 2 || len(path[1]) == 0 {
		indexHandler(ctx)
		return
	}

	name := viper.GetString("Projects." + string(path[1]) + ".Name")
	switch {
	case bytes.Compare(path[1], []byte("assets")) == 0:
		assetHandler(ctx)
	case name != "":
		remainingPath := ctx.Path()[len(path[1])+1:]
		projectHandler(ctx, string(path[1]), remainingPath)
	default:
		glog.Warnf("project '%s' not found in config", path[1])
		if ctx.URI().QueryArgs().GetBool("go-get") {
			glog.Debug("detected go-get, sending 404")
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.WriteString(`<html><head><title>Project not found</title></head><body><h1>404 - Not found!</h1><p>This Project does not exist.</p></body></html>`)
			return
		}
		ctx.Redirect("/", fasthttp.StatusSeeOther)
	}
}
