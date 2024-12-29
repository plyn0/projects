package main

import (
	"api-htmx/domain"
	"api-htmx/route"
	"api-htmx/template"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var writers = []domain.Writer{{Id: uuid.NewString(), Name: "hugo"}}

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
	// to be imported by the HTML code, static files (js, css) must be served
	router.Static("/public", "./public")
}

func setupRoutes(router *gin.Engine) {
	router.GET("/", GetPage)
	router.GET("/info", GetInfo)
	router.POST("/note", PostNote)
	router.GET("/api/v1/product", route.GetProducts)
	router.DELETE("/items/:id", DeleteItem)
}

func DeleteItem(c *gin.Context) {
	idToDelete := c.Param("id")
	// if err != nil {
	// 	template.Error().Render(c.Request.Context(), c.Writer)
	// }
	newWriters := []domain.Writer{}
	for _, item := range writers {
		if item.Id != idToDelete {
			newWriters = append(newWriters, item)
		}
	}
	writers = newWriters
}

func GetPage(c *gin.Context) {
	template.Root(writers).Render(c.Request.Context(), c.Writer)
}

func GetInfo(c *gin.Context) {
	template.Info().Render(c.Request.Context(), c.Writer)
}

/*
More efficient: return the HTML for the single item that was added to the list, an used hx-swap: beforeend
*/
func PostNote(c *gin.Context) {
	note := c.PostForm("note")
	item := domain.Writer{
		Id:   uuid.NewString(),
		Name: note,
	}
	// template.Hello(note).Render(c.Request.Context(), c.Writer)
	writers = append(writers, item)
	template.RenderList(writers).Render(c.Request.Context(), c.Writer)
}
