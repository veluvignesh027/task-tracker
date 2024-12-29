package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	UpstreamIP   string = os.Getenv("UPSTREAM_SERVER")
	UpstreamPort string = os.Getenv("UPSTREAM_PORT")
	UrlPrefix    string = "http://" + UpstreamIP + ":" + UpstreamPort
)

func GetStoryByID(c *gin.Context) {
	userid := c.Param("id")
	url := UrlPrefix + "/api/v1/story/getall/" + userid
	log.Println("Requesting upstream server with the URI: ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var stories []Story
	err = json.Unmarshal(body, &stories)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "list.html", stories)
}

func GetTicketsByID(c *gin.Context) {
	userid := c.Param("id")
	url := UrlPrefix + "/api/v1/tickets/getall/" + userid
	log.Println("Requesting upstream server with the URI: ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Writer.Write(body)
}
