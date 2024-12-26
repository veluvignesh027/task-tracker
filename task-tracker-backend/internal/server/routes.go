package server

import (
	"errors"
	"log"
	"net/http"
	"task-tracker/internal/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

    v1 := r.Group("/api/v1/story")
    {
        v1.GET("/get/:id", s.GetHandler)
        v1.POST("/create", s.PostHandler)
        v1.PUT("/update/:id", s.PutHandler)
        v1.DELETE("/delete/:id", s.DeleteHandler)
    }

    ui := r.Group("/frontend")
    {
        ui.GET("/index", s.IndexPageHandler)
        ui.GET("/home/:user", s.HomePageHandler)
    }
    return r
}

// Method: GET
// Returns: The Story details from the database with story_id as query param
func (s *Server) GetHandler(c *gin.Context){
   storyid := c.Param("id")
   if storyid == ""{
        log.Println("story id is empty in the URL param")
        c.JSON(http.StatusBadRequest, gin.H{"error":"story id is empty in the URL param"})
        return
   }

   var story models.Story
   dbinc := s.db.GetDBInstance()
   err := dbinc.Where("story_id = ?", string(storyid)).First(&story).Error
   if err != nil{
       if errors.Is(err, gorm.ErrRecordNotFound) {
           c.JSON(http.StatusNotFound, gin.H{"error":"record not found"})
            return
       }
       c.JSON(http.StatusBadRequest, gin.H{"error":err})
       return
   }

   c.JSON(http.StatusFound, story)
}

func (s *Server) PostHandler(c *gin.Context){
    var story models.StoryStruct
    if err := c.ShouldBind(&story); err != nil{
        log.Println("Error binding the body with struct.", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    dbStoryObj := models.ParseStoryObj(story)

    dbinc := s.db.GetDBInstance()
    err := dbinc.FirstOrCreate(&models.Story{}, dbStoryObj).Error
    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    c.JSON(http.StatusCreated, story)
}

func (s *Server) PutHandler(c *gin.Context){
    var story models.StoryStruct
    if err := c.ShouldBind(&story); err != nil{
        log.Println("Error binding the body with struct.", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    dbStoryObj := models.ParseStoryObj(story)
    dbinc := s.db.GetDBInstance()
    err := dbinc.FirstOrCreate(&models.Story{}, dbStoryObj).Error
    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
    c.JSON(http.StatusOK, story)
}

func (s *Server) DeleteHandler(c *gin.Context){
   storyid := c.Param("id")
   if storyid == ""{
       log.Println("story id is empty in the URL param")
       c.JSON(http.StatusBadRequest, gin.H{"error":"story id is empty in the URL param"})
       return
   }

   var story models.Story
   dbinc := s.db.GetDBInstance()
   err := dbinc.Where("story_id = ?", string(storyid)).First(&story).Error
   if err != nil{
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error":"record not found"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error":err})
        return
   } 

   err = dbinc.Delete(story).Error
   if err != nil{
       c.JSON(http.StatusInternalServerError,"story can not be deleted")
       return
   }
   c.JSON(http.StatusOK, story)
}
