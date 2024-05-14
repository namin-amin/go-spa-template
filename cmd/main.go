package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	app "github.com/namin-amin/simplesend/app"
)

func main() {

	isdev := os.Getenv("RUNENV")
	log.Println(isdev)
	server := echo.New()

	apigroup := server.Group("api",
		middleware.AddTrailingSlash())

	apigroup.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK,"hello world iam namin")
	})


	apigroup.GET("/:id",func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK,fmt.Sprintf("hello world iam namin %s",id))
	})

	apigroup.GET("/*", func(c echo.Context) error {
		tenat := c.Request().Header.Get("tenant")
		return c.JSON(http.StatusOK, tenat)
	})

	if isdev == "build" {
		RegisterAllStaticFs(app.GetAllRequiredPathsAndFS(), server)
		server.GET("/*", echo.StaticFileHandler("dist/index.html", app.Fdir))

	} else {
		server.GET("/*", echo.WrapHandler(http.HandlerFunc(proxyPass)))
	}

	log.Println(server.Start(":3000"))
}

func proxyPass(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse("http://localhost:5173/")
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
}

func RegisterAllStaticFs(fsMap map[string]fs.FS, server *echo.Echo) {
	for dir, fsdir := range fsMap {
		server.StaticFS(dir, fsdir)
	}
}
