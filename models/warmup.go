package models

import (
	"errors"
	"time"
)

type WarmupState int

const (
	WarmupScheduled WarmupState = iota
	WarmupActive
	WarmupCompleted
)

var (
	ErrWarmapStarted  = errors.New("Warmup has already been started")
	ErrWarmapCompeted = errors.New("Warmup has already been completed")
)

type Warmup struct {
	ID        int64       `json:"id"`
	LessonID  int64       `json:"lesson_id"`
	State     WarmupState `json:"status"`
	StartDate time.Time   `json:"start_date"`
	EndDate   time.Time   `json:"end_date"`
}

func NewWarmup(lessonId int64) *Warmup {
	return &Warmup{
		LessonID:  lessonId,
		StartDate: time.Now(),
		State:     WarmupActive,
	}
}

func (w *Warmup) Complete() error {
	if w.Completed() {
		return ErrWarmapCompeted
	}
	w.EndDate = time.Now()
	w.State = WarmupCompleted
	return nil
}

func (w *Warmup) Completed() bool {
	return w.EndDate.After(w.StartDate)
}
