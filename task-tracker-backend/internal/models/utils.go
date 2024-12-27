package models


func ParseStoryObj(obj StoryStruct, userid int, assigneeid int) Story {
	var story Story
	story.Name = obj.StoryName
	story.Status = string(obj.Status)
	story.UserAssignedID = uint(assigneeid)
	story.UserCreatedID = uint(userid)
	story.Description = obj.Description
	story.Priority = string(obj.Priority)

	return story
}

func ParseUserObj(obj UserStruct) User{
    var user User
    user.Name = obj.UserName
    user.Email = obj.UserEmail
    user.Password = obj.UserPass
    return user
}

