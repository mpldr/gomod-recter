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
	defer func(t time.Time){ glog.Debugf("request took %s", time.Since(t)) }(t)
	glog.Infof("%s requested %s using %s", ctx.RemoteIP(), ctx.Path(), ctx.UserAgent())
	ctx.SetContentTypeBytes([]byte("text/html"))

	path := bytes.Split(ctx.Path(), []byte("/"))

	glog.Debug(path)
	if len(path) < 2 || len(path[1]) == 0 {
		indexHandler(ctx)
		return
	}

	if name := viper.GetString("Projects." + string(path[1]) + ".Name"); name != "" {
		remainingPath := ctx.Path()[len(path[1])+1:]
		projectHandler(ctx, string(path[1]), remainingPath)
	} else {
		glog.Warnf("project '%s' not found in config", path[1])
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.Redirect("/", fasthttp.StatusSeeOther)
	}
}
