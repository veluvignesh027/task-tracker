package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.LoadHTMLGlob("internal/template/**/*.html")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", s.IndexPageHandler)

	app := r.Group("/app/frontend")
	{
		app.GET("/home", s.HomePageHandler)
	}

	upstream := r.Group("/upstream")
	{
		upstream.GET("/story/getall/:id", GetStoryByID)
		upstream.GET("/tickets/getall/:id", GetTicketsByID)
	}
	return r
}

func (s *Server) IndexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (s *Server) HomePageHandler(c *gin.Context) {
	var homedata HomeData

	url := "http://" + UpstreamIP + ":" + UpstreamPort + "/users/homedata?username=" + c.Query("username")
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	jsondata, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err = json.Unmarshal(jsondata, &homedata); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("username", homedata.UserEmail, 3600, "/", "localhost", false, true)

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Username":  homedata.UserEmail,
		"userid":    homedata.UserID,
		"Stories":   homedata.StoryCount,
		"Tickets":   homedata.TicketCount,
		"Completed": homedata.CompletedCount,
		"Pending":   homedata.PendingCount,
	})
}
