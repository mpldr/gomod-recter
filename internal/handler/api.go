package handler

import (
	"encoding/json"

	"mpldr.codes/recter/internal/data"

	"git.sr.ht/~poldi1405/glog"
	"github.com/valyala/fasthttp"
)

func apiHandler(ctx *fasthttp.RequestCtx, proj *data.Project, remainingPath []byte) {
	ctx.SetContentTypeBytes([]byte("application/json"))
	glog.Tracef("remainder: %s", remainingPath)

	switch string(remainingPath) {
	case "/versions":
		response, _ := json.Marshal(proj.Versions)
		ctx.Write(response)
	case "/versions/latest":
		response, _ := json.Marshal(getLatestVersion(proj))
		ctx.Write(response)
	case "/details":
		response, _ := json.Marshal(proj)
		ctx.Write(response)
	default:
		api404Handler(ctx)
	}
}

func api404Handler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)

	type response404 struct {
		Message            string   `json:"message"`
		AvailableEndpoints []string `json:"allowed_endpoints"`
	}
	response, _ := json.Marshal(response404{
		Message:            "requested API endpoint was not found",
		AvailableEndpoints: []string{"/versions", "/versions/latest", "/details"},
	})
	ctx.Write(response)
}

type latestVersion struct {
	V string `json:"latest_version"`
}

func getLatestVersion(proj *data.Project) *latestVersion {
	lv := "none"

	if len(proj.Versions) != 0 {
		lv = proj.Versions[0]
	}

	return &latestVersion{V: lv}
}
