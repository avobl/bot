package todoist

import "time"

// Task is a representation in Todoist
type Task struct {
	CreatorID    string    `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	AssigneeID   string    `json:"assignee_id"`
	AssignerID   string    `json:"assigner_id"`
	CommentCount int       `json:"comment_count"`
	IsCompleted  bool      `json:"is_completed"`
	Content      string    `json:"content"`
	Description  string    `json:"description"`
	Due          *Due      `json:"due"`
	Duration     *Duration `json:"duration"`
	ID           string    `json:"id"`
	Labels       []string  `json:"labels"`
	Order        int       `json:"order"`
	Priority     int       `json:"priority"`
	ProjectID    string    `json:"project_id"`
	SectionID    string    `json:"section_id"`
	ParentID     string    `json:"parent_id"`
	URL          string    `json:"url"`
}

// Due represents schedule configuration
type Due struct {
	Date        string    `json:"date"`
	IsRecurring bool      `json:"is_recurring"`
	Datetime    time.Time `json:"datetime"`
	String      string    `json:"string"`
	Timezone    string    `json:"timezone"`
}

// Duration of how much time dedicated to a task
type Duration struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}
