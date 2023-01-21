package trex_test

import (
	"testing"

	"github.com/sharpvik/trex"
)

type Suite struct {
	trex.BasicSuite
	variable int
}

func NewSuite(t *testing.T) *Suite {
	s := &Suite{
		BasicSuite: trex.Basic(t),
	}
	s.Schedule(s.TestFortyTwo)
	return s
}

func (s *Suite) Setup() error {
	s.variable = 42
	return nil
}

func (s *Suite) TestFortyTwo() {
	if s.variable != 42 {
		s.T().Fatalf("variable value mismatch: want 42; got %d", s.variable)
	}
}

func TestSuite(t *testing.T) {
	trex.Run(NewSuite(t))
}
