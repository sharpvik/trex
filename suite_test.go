package trex_test

import (
	"testing"

	"github.com/sharpvik/trex"
)

type Suite struct {
	trex.BasicSuite
	variable int
}

func (s *Suite) Setup() error {
	s.variable = 42
	return nil
}

func (s *Suite) TestFortyTwo(t *testing.T) {
	if s.variable != 42 {
		t.Fatalf("variable value mismatch: want 42; got %d", s.variable)
	}
}

func (s *Suite) TestIsEven(t *testing.T) {
	if s.variable%2 != 0 {
		t.Fatalf("variable value %d is odd; want even", s.variable)
	}
}

func TestSuite(t *testing.T) {
	trex.Run(t, &Suite{})
}
