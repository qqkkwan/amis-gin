package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed *.html
var tmpl embed.FS

//go:embed resources
var resources embed.FS

//go:embed pages
var pages embed.FS

func main() {
	r := gin.Default()

	t, _ := template.ParseFS(tmpl, "./*.html")
	r.SetHTMLTemplate(t)

	resources, _ := fs.Sub(resources, "resources")
	r.StaticFS("/resources", http.FS(resources))

	pages, _ := fs.Sub(pages, "pages")
	r.StaticFS("/pages", http.FS(pages))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
