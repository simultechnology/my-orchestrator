package task

import (
	"github.com/google/uuid"
	"my-orchestrator/job"
)

type Task struct {
	ID uuid.UUID
	Name string
	State job.State
}

func Show() {
	job.Job2()
}