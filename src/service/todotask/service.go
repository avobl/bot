package todotask

import (
	"context"
	"fmt"
	"time"

	"github.com/avobl/bot/src/external/todoist"
)

const (
	todoistTimeLayout = "2006-01-02"
)

type TodoistClient interface {
	GetTasks(ctx context.Context, token string) ([]*todoist.Task, error)
}

type Service struct {
	todoistClient TodoistClient
}

func NewService(todoistClient TodoistClient) *Service {
	return &Service{
		todoistClient: todoistClient,
	}
}

func (s *Service) GetTasks(ctx context.Context, token string) ([]*todoist.Task, error) {
	tasks, err := s.todoistClient.GetTasks(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %v", err)
	}

	today := time.Now().Format(todoistTimeLayout)
	dueToday := make([]*todoist.Task, 0, len(tasks))
	for _, task := range tasks {
		if task.Due != nil && task.Due.Date == today {
			dueToday = append(dueToday, task)
		}
	}

	return dueToday, nil
}
