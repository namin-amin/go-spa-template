package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namin-amin/simplesend/pkg/loadenv"
	"github.com/namin-amin/simplesend/ui"
)

func main() {
	loadenv.LoadRequiredEnvFiles()

	buildEnv:=os.Getenv("RUNENV")
	PORT := os.Getenv("PORT")

	fmt.Printf("Running with %s Mode\n",buildEnv)

	server := echo.New()

	apiGroup := server.Group("api",
		middleware.AddTrailingSlash())

	apiGroup.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world iam namin")
	})

	apiGroup.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, fmt.Sprintf("[hello world iam namin %s]", id))
	})

	if os.Getenv("RUNENV") == "build" {
		ui.RegisterAllStaticFs(func(s string, f fs.FS) {
			server.StaticFS(s,f)
		})
	} else {
		server.GET("/*", echo.WrapHandler(ui.ProxyPass("http://localhost:4200")))
	}

	log.Println(server.Start(fmt.Sprintf(":%s",PORT)))
}
