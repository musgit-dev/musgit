package models

import (
	"errors"
	"time"
)

type LessonState int

var ErrNotActiveLesson = errors.New("No active lesson")

const (
	LessonScheduled LessonState = iota
	LessonActive
	LessonPaused
	LessonCompleted
)

type Lesson struct {
	ID        int64
	State     LessonState
	StartDate time.Time
	EndDate   time.Time
	Comment   string
}

func NewLesson() *Lesson {
	lesson := Lesson{StartDate: time.Now(), State: LessonActive}
	return &lesson
}

func (l *Lesson) Pause() *Lesson {
	l.State = LessonPaused
	return l
}

func (l *Lesson) Resume() *Lesson {
	l.State = LessonActive
	return l
}

func (l *Lesson) Finish() *Lesson {
	l.EndDate = time.Now()
	l.State = LessonCompleted
	return l
}

func (l *Lesson) AddNote(comment string) *Lesson {
	l.Comment = comment
	return l
}
