package domain

import "time"

type PracticeProgressEvalutation int

const (
	Bad PracticeProgressEvalutation = iota
	Normal
	Good
)

type Practice struct {
	StartDate time.Time
	EndDate   time.Time
	Progress  PracticeProgressEvalutation
}
