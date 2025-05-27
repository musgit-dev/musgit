package domain

import "time"

type PracticeProgressEvalutation int

const (
	Bad PracticeProgressEvalutation = iota
	Normal
	Good
)

type Practice struct {
	StartDate time.Time                   `json:"start_date"`
	EndDate   time.Time                   `json:"end_date"`
	Progress  PracticeProgressEvalutation `json:"progress"`
}
