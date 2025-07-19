package domain

import (
	"testing"
)

func TestPractice(t *testing.T) {

	p := NewPractice()
	if !p.Active() {
		t.Fatal("Practice should be active")
	}

	err := p.Complete(Normal)
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
