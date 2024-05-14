package app

import (
	"embed"
	"io/fs"
	"log"

	"github.com/namin-amin/simplesend/pkg/customerros"
)

//go:embed all:dist
var Fdir embed.FS

func GetAllRequiredPathsAndFS() map[string]fs.FS {
	errs:= customerros.NewErrors()
	assets, err := fs.Sub(Fdir, "dist/assets")
	errs.AddNewError(err)
	dist, err := fs.Sub(Fdir, "dist/images")
	errs.AddNewError(err)
	
	if errs.DoErrorExists() {
		log.Fatalln("could not register required file systems")
	}

	famap := make(map[string]fs.FS)
	famap["assets"] = assets
	famap["images"] = dist

	return famap
}
