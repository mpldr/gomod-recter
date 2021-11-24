package handler

import (
	"io"
	"mime"
	"os"
	"path"
	"strings"

	"git.sr.ht/~poldi1405/glog"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func assetHandler(ctx *fasthttp.RequestCtx) {
	filepath := viper.GetString("Directories.AssetDir") + strings.TrimPrefix(string(ctx.Path()), "/assets")

	glog.Debugf("serving asset: %s", filepath)
	fh, err := os.Open(filepath)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	ctx.SetContentType(mime.TypeByExtension(path.Ext(filepath)))
	_, err = io.Copy(ctx.Response.BodyWriter(), fh)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	return
}
