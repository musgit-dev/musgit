package models

import (
	"errors"
	"time"
)

type WarmupState int

const (
	WarmupActive WarmupState = iota
	WarmupCompleted
)

type Warmup struct {
	ID        int64       `json:"id"`
	LessonID  int64       `json:"lesson_id"`
	Status    WarmupState `json:"status"`
	StartDate time.Time   `json:"start_date"`
	EndDate   time.Time   `json:"end_date"`
}

func NewWarmup(lessonId int64) *Warmup {
	return &Warmup{
		LessonID:  lessonId,
		StartDate: time.Now(),
		Status:    WarmupActive,
	}
}

func (w *Warmup) Complete() error {
	if w.Completed() {
		return errors.New(
			"Warmup has already been completed.",
		)
	}
	w.EndDate = time.Now()
	w.Status = WarmupCompleted
	return nil
}

func (w *Warmup) Completed() bool {
	return w.EndDate.After(w.StartDate)
}
