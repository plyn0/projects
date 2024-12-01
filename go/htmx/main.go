package main

import (
	"api-htmx/route"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	slog.Info("here is a structured log")
	// logs: add date, time and filename with line number (https://pkg.go.dev/log#pkg-constants)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// err := db.IngestData()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

func main() {
	router := gin.Default()
	setupStaticFiles(router)
	setupRoutes(router)
	router.Run(":8080") // serve and listen on port 0.0.0.0:8080
}

// setupRouter creates and returns a router with the associated routes.
func setupStaticFiles(router *gin.Engine) {
	// load all the templates (must be called before LoadHTMLFiles)
	router.LoadHTMLGlob("assets/*")
	// load the starting point of the app
	// router.LoadHTMLFiles("public/index.html", "templates/info.html")
	fmt.Println(router)
	// to be imported by the HTML code, static files (js, css) must be served
	router.Static("/assets", "./assets")
}

func setupRoutes(router *gin.Engine) {
	router.GET("/", GetPage)
	router.GET("/info", GetInfo)
	router.POST("/note", PostNote)
	router.GET("/api/v1/product", route.GetProducts)
}

func GetPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}

func GetInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "info.html", gin.H{
		"title": "Posts",
	})
}

func PostNote(c *gin.Context) {

}
