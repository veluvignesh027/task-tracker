package server

import (
	"errors"
    log "github.com/golang/glog"
	"net/http"
	"task-tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) GetAllUsers(c *gin.Context) {
	var users []models.User
	dbins := s.db.GetDBInstance()
	if err := dbins.Find(&users).Error; err != nil {
		log.Error(err)
		c.JSON(http.StatusNoContent, err)
		return
	}
        
    c.JSON(http.StatusOK, users)
}

func (s *Server) GetUserByEmail(c *gin.Context) {
    email := c.Query("email")
    if email == ""{
        log.Error("The email in the URL param is missing")
        c.JSON(http.StatusBadRequest, gin.H{
            "error":" The email in the URL param is missing",
        })
        return
    }

    var user models.User
    dbins := s.db.GetDBInstance()
    
    if err := dbins.Where("email = ?",email).Find(&user).Error; err != nil{
        log.Error(err)
        c.JSON(http.StatusNotFound, err)
        return
    }
    
    c.JSON(http.StatusOK, user)
}

func (s *Server) CreateUser(c *gin.Context) {
    var user models.UserStruct
    if err := c.ShouldBind(&user); err != nil {
        log.Error("Error binding the body with struct.", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userinfo := models.ParseUserObj(user)
    log.Info("The User Info for creating user is ", userinfo)
    dbinc := s.db.GetDBInstance()
    if err := dbinc.Create(&userinfo).Error; err != nil{
        if errors.Is(err, gorm.ErrRecordNotFound){
            log.Error("No user found in the table. Creating a new record")
        }else{
            log.Error(err)
            c.JSON(http.StatusBadRequest, gin.H{
                "error":err.Error(),
            })
            return
        }
    }

    c.JSON(http.StatusCreated, user) 
}

func (s *Server) DeleteUser(c *gin.Context) {
    userid := c.Param("id")
    if userid == "" {
        log.Error("user id is empty in the URL param")
        c.JSON(http.StatusBadRequest, gin.H{"error": "story id is empty in the URL param"})
        return
    }

    var user models.User
    dbinc := s.db.GetDBInstance()
    err := dbinc.Where("id = ?", string(userid)).Delete(&models.User{}).Error
    if err != nil {
        log.Error(err)
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    c.JSON(http.StatusOK, user)    
}

func (s *Server) GetHomeDataForUser(c *gin.Context){
    var homedata models.HomeData
    
    homedata.UserName = c.Query("username")
    dbins := s.db.GetDBInstance()
    
    uid, err := getUserIDFromUserName(dbins, homedata.UserName)
    if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }
    homedata.UserID = uid
    homedata.UserEmail = homedata.UserName
    
    err = dbins.Model(&models.Story{}).Where("user_assigned_id = ?", uid).Count(&homedata.StoryCount).Error
    if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }
    

    err = dbins.Model(&models.Ticket{}).Where("user_assigned_id = ?", uid).Count(&homedata.TicketCount).Error
    if err != nil{
        log.Error(err)
        c.JSON(http.StatusBadRequest, err)
        return
    }

    // Taking the count of completed tickets
    dbins.Model(&models.Ticket{}).Where("user_assigned_id = ? AND status = ? OR status = ?",uid, models.Completed, models.Closed).Count(&homedata.CompletedCount)

    // Taking the count of Pending tickets
    dbins.Model(&models.Ticket{}).Where("user_assigned_id = ? AND status = ? OR status = ?",uid, models.Assigned, models.New).Count(&homedata.CompletedCount)

    
    c.JSON(http.StatusOK, homedata)
}



