package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed all:dist
var f embed.FS

func main() {

	env := "dev2"
	server := echo.New()

	if env == "dev" {
		server.GET("/", echo.StaticFileHandler("dist/browser/index.html", f))
		staticdir, _ := fs.Sub(f, "dist/browser")
		server.StaticFS("/", staticdir)
	}else if env == "dev2"{
		server.GET("/*",echo.WrapHandler(http.HandlerFunc(proxyPass)))

	} else {
		server.Static("/", "dist/browser/")
		server.GET("/", func(c echo.Context) error {
			return c.File("dist/browser/index.html")
		})
	}

	apigroup := server.Group("api",
		middleware.AddTrailingSlash())

	apigroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tenant := c.Request().Header.Get("tenant")
			if tenant == "" {
				return echo.ErrBadRequest
			}
			return next(c)
		}
	})

	apigroup.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})

	apigroup.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		tenat := c.Request().Header.Get("tenant")
		return c.JSON(http.StatusOK, fmt.Sprintf("%s%s", id, tenat))
	})

	log.Println(server.Start(":3000"))
}

func proxyPass(res http.ResponseWriter, req  *http.Request) {
	// Encrypt Request here
	// ...
  
	url, _ := url.Parse("http://localhost:4200/")
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
  } 