package server

type HomeData struct {
	UserID         int    `json:"userid"`
	UserName       string `json:"username"`
	UserEmail      string `json:"useremail"`
	StoryCount     int64  `json:"story_count"`
	TicketCount    int64  `json:"ticket_count"`
	PendingCount   int64  `json:"pending_count"`
	CompletedCount int64  `json:"completed_count"`
}

type Story struct {
	StoryID        int
	Name           string
	UserCreatedID  uint
	UserAssignedID uint
	Description    string
	Status         string
	Priority       string
}
