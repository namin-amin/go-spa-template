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
	"github.com/namin-amin/simplesend/app"
	"github.com/namin-amin/simplesend/pkg/loadenv"
)

func main() {
	loadenv.LoadRequiredEnvFiles()

	fmt.Println(os.Getenv("RUNENV"))

	server := echo.New()

	apiGroup := server.Group("api",
		middleware.AddTrailingSlash())

	apiGroup.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "hello world iam namin")
	})

	apiGroup.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, fmt.Sprintf("hello world iam namin %s", id))
	})

	apiGroup.GET("/*", func(c echo.Context) error {
		tenant := c.Request().Header.Get("tenant")
		return c.JSON(http.StatusOK, tenant)
	})

	if os.Getenv("RUNENV") == "build" {
		RegisterAllStaticFs(app.GetAllRequiredPathsAndFS(), server)
		server.GET("/*", echo.StaticFileHandler("dist/index.html", app.FDir))

	} else {
		server.GET("/*", echo.WrapHandler(http.HandlerFunc(proxyPass)))
	}

	log.Println(server.Start(":3000"))
}

func proxyPass(res http.ResponseWriter, req *http.Request) {
	uri, _ := url.Parse("http://localhost:5173/")
	proxy := httputil.NewSingleHostReverseProxy(uri)
	proxy.ServeHTTP(res, req)
}

func RegisterAllStaticFs(fsMap map[string]fs.FS, server *echo.Echo) {
	for dir, fsDir := range fsMap {
		server.StaticFS(dir, fsDir)
	}
}
