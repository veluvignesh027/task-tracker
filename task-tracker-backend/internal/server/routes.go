package server

import (
	"errors"
	log "github.com/golang/glog"
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

	user := r.Group("/users")
	{
		user.GET("/getall", s.GetAllUsers)
		user.GET("/get", s.GetUserByEmail)
		user.POST("/create", s.CreateUser)
		user.DELETE("/delete/:id", s.DeleteUser)
	    user.GET("/homedata/:id",s.GetHomeDataForUser)
    }
    log.Info("Registered all the routes for User & Story handles")
	return r
}

// Method: GET
// Returns: The Story details from the database with story_id as query param
func (s *Server) GetHandler(c *gin.Context) {
	storyid := c.Param("id")
	if storyid == "" {
		log.Error("story id is empty in the URL param")
		c.JSON(http.StatusBadRequest, gin.H{"error": "story id is empty in the URL param"})
		return
	}

	var story models.Story
	dbinc := s.db.GetDBInstance()
	err := dbinc.Where("story_id = ?", string(storyid)).First(&story).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Error("Can not get the story. ",err)
            c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusFound, story)
}

func (s *Server) PostHandler(c *gin.Context) {
	var story models.StoryStruct
	if err := c.ShouldBind(&story); err != nil {
		log.Error("Error binding the body with struct.", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
    }

    dbinc := s.db.GetDBInstance()
    creator, err  := getUserIDFromUserName(dbinc, story.UserCreated)
	if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, gin.H{"error":"not a valid user"})
    }
    assigner, err := getUserIDFromUserName(dbinc, story.UserAssigned)
    if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, gin.H{"error":"not a valid user"})
    }

    dbStoryObj := models.ParseStoryObj(story, creator, assigner)
	err = dbinc.FirstOrCreate(&models.Story{}, dbStoryObj).Error
	if err != nil {
        log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
    }
	c.JSON(http.StatusCreated, story)
}

func (s *Server) PutHandler(c *gin.Context) {
	var story models.StoryStruct
	if err := c.ShouldBind(&story); err != nil {
		log.Error("Error binding the body with struct.", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
    }

	dbinc := s.db.GetDBInstance()
    creator, err  := getUserIDFromUserName(dbinc, story.UserCreated)
	if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, gin.H{"error":"not a valid user"})
        return
    }

    assigner, err := getUserIDFromUserName(dbinc, story.UserAssigned)
    if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, gin.H{"error":"not a valid user"})
        return
    }

    dbStoryObj := models.ParseStoryObj(story,creator ,assigner)
	err = dbinc.FirstOrCreate(&models.Story{}, dbStoryObj).Error
	if err != nil {
        log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	    return
    }
	c.JSON(http.StatusOK, story)
}

func (s *Server) DeleteHandler(c *gin.Context) {
	storyid := c.Param("id")
	if storyid == "" {
		log.Error("story id is empty in the URL param")
		c.JSON(http.StatusBadRequest, gin.H{"error": "story id is empty in the URL param"})
		return
	}

	var story models.Story
	dbinc := s.db.GetDBInstance()
	err := dbinc.Where("story_id = ?", string(storyid)).Delete(&models.Story{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
            c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, story)
}

func getUserIDFromUserName(db *gorm.DB, name string)(int, error){
    var id int

    res := db.Model(&models.User{}).Select("user_id").Where("email = ?", name).Scan(&id)
    if res.Error != nil{
        log.Error(res.Error)
        return id, res.Error
    }
    
    return id, nil
}
