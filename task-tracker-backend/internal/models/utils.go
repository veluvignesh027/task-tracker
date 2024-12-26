package models

import(
    "time"
)

func ParseStoryObj(obj StoryStruct) Story {
    var story Story
    story.StoryID = getNewStoryID()
    story.Name = obj.StoryName
    story.Status = string(obj.Status)
    story.UserAssignedID = 123
    story.UserCreatedID = 000
    story.Description = obj.Description
    story.Priority = string(obj.Priority)

    return story
}

func getNewStoryID() int{
    return int(time.Now().UnixNano())% 100000
}
