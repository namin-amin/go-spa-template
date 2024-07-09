package ui

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/namin-amin/simplesend/pkg/customerros"
)

//go:embed all:dist
var FDir embed.FS

// route : path of the route you want to assign
//
// fileSystem : to be served for the given route
type RegisterStaticRouteHandler func(route string, fileSystem fs.FS)

func getAllRequiredPathsAndFS() map[string]fs.FS {
	errs := customerros.NewErrors()
	assets, err := fs.Sub(FDir, "dist/browser")
	errs.AddNewError(err)

	if errs.DoesErrorExists() {
		log.Fatalln("could not register required file systems")
	}

	faMap := make(map[string]fs.FS)
	faMap["/*"] = assets
	return faMap
}

func ProxyPass(uiEndpoint string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		uri, _ := url.Parse(uiEndpoint)
		proxy := httputil.NewSingleHostReverseProxy(uri)
		proxy.ServeHTTP(res, req)
	}
}

func RegisterAllStaticFs(registerStaticRoutes RegisterStaticRouteHandler) {
	fsMap := getAllRequiredPathsAndFS()
	for dir, fsDir := range fsMap {
		registerStaticRoutes(dir, fsDir)
	}
}
