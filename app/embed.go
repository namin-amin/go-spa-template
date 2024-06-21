package app

import (
	"embed"
	"io/fs"
	"log"

	"github.com/namin-amin/simplesend/pkg/customerros"
)

//go:embed all:dist
var FDir embed.FS

func GetAllRequiredPathsAndFS() map[string]fs.FS {
	errs := customerros.NewErrors()
	assets, err := fs.Sub(FDir, "dist/assets")
	errs.AddNewError(err)
	dist, err := fs.Sub(FDir, "dist/images")
	errs.AddNewError(err)

	if errs.DoesErrorExists() {
		log.Fatalln("could not register required file systems")
	}

	faMap := make(map[string]fs.FS)
	faMap["assets"] = assets
	faMap["images"] = dist

	return faMap
}
