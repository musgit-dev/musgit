package domain

import (
	"errors"
	"time"
)

type PracticeProgressEvalutation int

const (
	PracticeProgressBad PracticeProgressEvalutation = iota
	PracticeProgressNormal
	PracticeProgressGood
)

type Practice struct {
	ID        int64                       `json:"id"`
	StartDate time.Time                   `json:"start_date"`
	EndDate   time.Time                   `json:"end_date"`
	Progress  PracticeProgressEvalutation `json:"progress"`
	LessonID  int64                       `json:"lesson_id"`
}

func NewPractice() *Practice {
	return &Practice{StartDate: time.Now()}
}

func (p *Practice) Complete(
	evaluation PracticeProgressEvalutation,
) error {
	if p.Completed() {
		return errors.New(
			"Practice has already been completed.",
		)
	}
	p.EndDate = time.Now()
	p.Progress = evaluation
	return nil
}

func (p *Practice) Active() bool {
	return !p.StartDate.IsZero() && p.EndDate.IsZero()
}

func (p *Practice) Completed() bool {
	return p.EndDate.After(p.StartDate)
}
