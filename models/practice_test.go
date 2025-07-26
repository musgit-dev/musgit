package models

import (
	"testing"
)

func TestPractice(t *testing.T) {

	p := NewPractice(0)
	if !p.Active() {
		t.Fatal("Practice should be active")
	}

	err := p.Complete(PracticeProgressNormal)
	if err != nil {
		t.Fatal(err)
	}

	if p.Active() {
		t.Fatal("Practice should not be active")
	}
	if !p.Completed() {
		t.Fatal("Practice should be completed")
	}

}
