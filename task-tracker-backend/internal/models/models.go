package models

type StatusString string
type PriorityString string

const (
	New       StatusString = "New"
	Assigned  StatusString = "Assigned"
	Completed StatusString = "Completed"
	Closed    StatusString = "Closed"
	Hold      StatusString = "Hold"
)

const (
	Low      PriorityString = "Low"
	Medium   PriorityString = "Medium"
	Normal   PriorityString = "Normal"
	High     PriorityString = "High"
	Critical PriorityString = "Critical"
)

type StoryStruct struct {
	StoryName    string         `json:"story_name"`
	UserCreated  string         `json:"creator_name"`
	UserAssigned string         `json:"assignee_name"`
	Description  string         `json:"description"`
	Status       StatusString   `json:"status"`
	Priority     PriorityString `json:"priority"`
}

type UserStruct struct{
    UserName string `json:"username"`
    UserEmail string `json:"email"`
    UserPass string `json:"password"`
}

type HomeData struct{
    UserID int `json:"userid"`
    UserName string `json:"username"`
    UserEmail string `json:"useremail"`
    StoryCount int64 `json:"story_count"`
    TicketCount int64 `json:"ticket_count"`
    PendingCount int64 `json:"pending_count"`
    CompletedCount int64 `json:"completed_count"`
}


