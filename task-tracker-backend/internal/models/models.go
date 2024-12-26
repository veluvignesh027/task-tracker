package models

type StatusString string
type PriorityString string

const (
    New StatusString = "New"
    Assigned StatusString = "Assigned"
    Completed StatusString = "Completed"
    Closed StatusString = "Closed"
    Hold StatusString = "Hold"
)

const (
    Low PriorityString = "Low"
    Medium PriorityString = "Medium"
    Normal PriorityString = "Normal"
    High PriorityString = "High"
    Critical PriorityString = "Critical"
)

type StoryStruct struct{
    StoryName string `json:"story_name"`
    UserCreated string `json:"creator_name"`
    UserAssigned string `json:"assignee_name"`
    Description string `json:"description"`
    Status StatusString `json:"status"`
    Priority PriorityString `json:"priority"`
}
